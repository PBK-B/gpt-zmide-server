/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/helper/config.go
 */
package helper

import (
	"fmt"
	"io/ioutil"
	"os"

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
