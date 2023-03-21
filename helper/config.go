/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/helper/config.go
 */
package helper

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"
)

var Config *DefaultConfig

type DefaultConfig struct {
	AppKey     string `yaml:"app_key"`
	SiteName   string `yaml:"site_name"`
	DomainName string `yaml:"domain_name"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	AdminUser  struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"admin_user"`
	Mysql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
	OpenAI struct {
		SecretKey     string `yaml:"secret_key"`
		Model         string `yaml:"model"`
		HttpProxyHost string `yaml:"http_proxy_host"`
		HttpProxyPort string `yaml:"http_proxy_port"`
	}
}

func init() {
	var err error
	Config, err = ReadConfig()
	if err != nil {
		// 读取配置失败
		fmt.Println("读取配置文件失败。" + err.Error())
	}
}

// 是否未完成初始化
func IsInitialize() bool {
	return Config == nil ||
		Config.AdminUser.User == "" ||
		Config.AdminUser.Password == "" ||
		Config.Mysql.Host == "" ||
		Config.Mysql.User == ""
}

// 获取配置目录
func getConfigPath() string {
	if IsRelease() {
		appPath, err := os.Executable()
		if err == nil {
			appPath = filepath.Dir(appPath)
			if appPath != "" {
				return appPath + "/app.conf"
			}
		}
	}
	return "./app.conf"
}

func InitConfig() *DefaultConfig {
	c := DefaultConfig{}
	c.AppKey = RandomStr(32)
	c.SiteName = "gpt-zmide-server"
	c.DomainName = "https://demo.zmide.com"
	c.Host = "0.0.0.0"
	c.Port = 8091
	c.AdminUser.User = "admin"
	pwd := fmt.Sprintf("%x", md5.Sum([]byte("admin")))
	c.AdminUser.Password = pwd
	c.Mysql.Host = "127.0.0.1"
	c.Mysql.Port = 3306
	c.Mysql.Database = "gpt_zmide_server"
	c.OpenAI.Model = "gpt-3.5-turbo"
	return &c
}

func ReadConfig() (*DefaultConfig, error) {
	_, err := os.Stat(getConfigPath())
	if err == nil {
		// 文件存在，读取配置文件
		content, err := ioutil.ReadFile(getConfigPath())

		if err != nil {
			// 配置文件读取失败
			return nil, err
		}
		return LoadConfig(string(content))
	}

	conf := InitConfig()

	if err = conf.SaveConfig(); err != nil {
		// 保存配置文件失败
		return nil, err
	}
	return conf, nil
}

// 重新加载配置文件
func LoadConfig(configStr string) (*DefaultConfig, error) {
	var config = &DefaultConfig{}
	err := yaml.Unmarshal([]byte(configStr), config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// 保存配置文件
func (c *DefaultConfig) SaveConfig() error {
	if content, err := yaml.Marshal(c); err == nil {
		err = ioutil.WriteFile(getConfigPath(), content, 0766)
		if err != nil {
			// 写入配置失败
			return err
		}
	} else {
		// 格式化配置失败
		return err
	}
	return nil
}

// 获取数据库地址
func (c *DefaultConfig) GetMysqlUrl() (*url.URL, error) {
	if c.Mysql.Host == "" || c.Mysql.Port == 0 {
		return nil, errors.New("database misconfiguration error")
	}
	return GetMysqlUrl(c.Mysql.Host, c.Mysql.Port)
}

func GetMysqlUrl(host string, port int) (*url.URL, error) {
	if host == "" || port == 0 {
		return nil, errors.New("database misconfiguration error")
	}
	u, err := url.Parse("http://" + host + ":" + strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	return u, nil
}

func PingOpenAI(secret_key string, proxy_host string, proxy_port string) (status bool, callback string) {
	model := Config.OpenAI.Model

	if model != "" && secret_key != "" {
		client := resty.New()
		if proxy_host != "" && proxy_port != "" {
			client.SetProxy("http://" + proxy_host + ":" + proxy_port)
		}
		client.SetTimeout(2 * time.Minute)
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Authorization", "Bearer "+secret_key).
			Get("https://api.openai.com/v1/models")
		if err == nil && resp.StatusCode() > 190 && resp.StatusCode() < 300 {
			type Model struct {
				Id         string        `json:"id"`
				Object     string        `json:"object"`
				OwnedBy    string        `json:"owned_by"`
				Permission []interface{} `json:"permission"`
			}
			var data struct {
				Data []Model `json:"data"`
			}
			callback = string(resp.Body())
			if err := json.Unmarshal(resp.Body(), &data); err == nil {
				status = true
				callback = "200"
			}
		} else {
			callback = resp.Status()
			if err != nil {
				callback = err.Error()
			}
		}
	}

	return
}
