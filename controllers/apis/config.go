/*
 * @Author: Bin
 * @Date: 2023-03-19
 * @FilePath: /gpt-zmide-server/controllers/apis/config.go
 */
package apis

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"gpt-zmide-server/helper"
	"gpt-zmide-server/models"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Controller
}

// 修改密码
func (ctl *Config) UpdatePassword(c *gin.Context) {

	old_password := c.PostForm("old_password")
	new_password := c.PostForm("new_password")
	if old_password == "" || new_password == "" || len(new_password) < 3 {
		ctl.Fail(c, "参数异常")
		return
	}

	old_pwd := fmt.Sprintf("%x", md5.Sum([]byte(old_password)))
	if old_pwd != helper.Config.AdminUser.Password {
		ctl.Fail(c, "旧密码错误")
		return
	}

	if old_password == new_password {
		ctl.Fail(c, "两次输入密码相同")
		return
	}

	new_pwd := fmt.Sprintf("%x", md5.Sum([]byte(new_password)))
	helper.Config.AdminUser.Password = new_pwd
	helper.Config.SaveConfig()

	ctl.Success(c, "ok")
}

func (ctl *Config) SystemInfo(c *gin.Context) {
	// 获取应用程序数量
	var applicationCount int64 = 0
	models.DB.Model(&models.Application{}).Count(&applicationCount)

	// 获取会话数量
	var chatCount int64 = 0
	models.DB.Model(&models.Chat{}).Count(&chatCount)

	// 统计接口调用次数
	var messageCount int64 = 0
	models.DB.Model(&models.Message{}).Count(&messageCount)

	// 计算预计扣费
	var assistantMessages []models.Message
	models.DB.Model(&models.Message{}).Where("role = ?", "assistant").Find(&assistantMessages)
	assistantMessageWordCount := 0
	for _, item := range assistantMessages {
		assistantMessageWordCount = assistantMessageWordCount + len(item.Content)
	}

	estimatedCost := float64(assistantMessageWordCount) / 1000 * 0.002

	ctl.Success(c, gin.H{
		"app_count":      applicationCount,
		"chat_count":     chatCount,
		"use_api_count":  messageCount,
		"estimated_cost": estimatedCost,
	})
}

func (ctl *Config) PingOpenAI(c *gin.Context) {

	status := false
	callback := ""

	secret_key := helper.Config.OpenAI.SecretKey
	if secret_key != "" {
		proxy_host := helper.Config.OpenAI.HttpProxyHost
		proxy_port := helper.Config.OpenAI.HttpProxyPort
		status, callback = helper.PingOpenAI(secret_key, proxy_host, proxy_port)
	}

	ctl.Success(c, gin.H{
		"status":   status,
		"callback": callback,
	})
}

func (ctl *Config) ConfigInfo(c *gin.Context) {
	systemConfig := helper.Config
	openAiConfig := systemConfig.OpenAI

	ctl.Success(c, gin.H{
		"site_config": gin.H{
			"site_name":   systemConfig.SiteName,
			"domain_name": systemConfig.DomainName,
			"port":        strconv.Itoa(systemConfig.Port),
		},
		"openai_config": gin.H{
			"openai_secret_key": openAiConfig.SecretKey,
			"openai_model":      openAiConfig.Model,
			"openai_proxy_host": openAiConfig.HttpProxyHost,
			"openai_proxy_port": openAiConfig.HttpProxyPort,
		},
	})
}

func (ctl *Config) ConfigInfoSave(c *gin.Context) {
	name := c.PostForm("name")
	data := c.PostForm("data")
	if name == "" || data == "" {
		ctl.Fail(c, "参数异常")
		return
	}

	var err error
	switch name {
	case "site":
		err = ctl.siteConfig(data)
	case "openai":
		err = ctl.openaiConfig(data)
	}
	if err != nil {
		ctl.Fail(c, err.Error())
		return
	}

	ctl.Success(c, "ok!")
}

// 配置站点
func (ctl *Config) siteConfig(data string) error {
	type SiteConfig struct {
		SiteName   string `json:"site_name"`
		DomainName string `json:"domain_name"`
		Port       string `json:"port"`
	}

	var config SiteConfig
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return errors.New("参数填写错误")
	}

	if config.SiteName == "" {
		return errors.New("站点名称不得为空")
	}

	domainName, err := url.Parse(config.DomainName)
	if config.DomainName == "" || err != nil || domainName == nil {
		return errors.New("站点域名错误")
	}

	port, err := strconv.Atoi(config.Port)
	if err != nil || port == 0 {
		return errors.New("应用端口错误")
	}

	helper.Config.SiteName = config.SiteName
	helper.Config.DomainName = config.DomainName
	helper.Config.Port = port

	if err := helper.Config.SaveConfig(); err != nil {
		return err
	}

	return nil
}

// 配置 OpenAI
func (ctl *Config) openaiConfig(data string) error {
	type OpenAIConfig struct {
		SecretKey string `json:"openai_secret_key"`
		ProxyHost string `json:"openai_proxy_host"`
		ProxyPort string `json:"openai_proxy_port"`
		Model     string `json:"openai_model"`
	}

	var config OpenAIConfig
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return errors.New("参数填写错误")
	}

	if config.SecretKey == "" {
		return errors.New("OpenAI SecretKey 填写错误")
	}

	status, callback := helper.PingOpenAI(config.SecretKey, config.ProxyHost, config.ProxyPort)
	if !status {
		return errors.New("OpenAI 服务器连接失败，" + callback)
	}

	helper.Config.OpenAI.SecretKey = config.SecretKey
	helper.Config.OpenAI.HttpProxyHost = config.ProxyHost
	helper.Config.OpenAI.HttpProxyPort = config.ProxyPort
	helper.Config.OpenAI.Model = config.Model

	if err := helper.Config.SaveConfig(); err != nil {
		return err
	}

	return nil
}
