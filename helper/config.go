/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/helper/config.go
 */
package helper

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

var Config *DefaultConfig

const configPath = "./app.conf"

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

func InitConfig() *DefaultConfig {
	c := DefaultConfig{}
	c.AppKey = RandStr(32)
	c.SiteName = "gpt-zmide-server"
	c.DomainName = "https://demo.zmide.com"
	c.Host = "0.0.0.0"
	c.Port = 8091
	c.AdminUser.User = "admin"
	pwd := fmt.Sprintf("%x", md5.Sum([]byte("admin")))
	c.AdminUser.Password = pwd
	c.Mysql.Host = "127.0.0.1"
	c.Mysql.Port = 3306
	c.Mysql.User = "root"
	c.Mysql.Database = "gpt_zmide_server"
	c.OpenAI.Model = "gpt-3.5-turbo"
	return &c
}

func ReadConfig() (*DefaultConfig, error) {
	_, err := os.Stat(configPath)
	if err == nil {
		// 文件存在，读取配置文件
		content, err := ioutil.ReadFile(configPath)

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
	yaml.Unmarshal([]byte(configStr), config)
	return config, nil
}

// 保存配置文件
func (c *DefaultConfig) SaveConfig() error {
	if content, err := yaml.Marshal(c); err == nil {
		err = ioutil.WriteFile(configPath, content, 0766)
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
	u, err := url.Parse("http://" + c.Mysql.Host + ":" + strconv.Itoa(c.Mysql.Port))
	if err != nil {
		return nil, err
	}
	return u, nil
}
