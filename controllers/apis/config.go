/*
 * @Author: Bin
 * @Date: 2023-03-19
 * @FilePath: /gpt-zmide-server/controllers/apis/config.go
 */
package apis

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"gpt-zmide-server/helper"
	"gpt-zmide-server/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
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

	model := helper.Config.OpenAI.Model
	secret_key := helper.Config.OpenAI.SecretKey
	if model != "" && secret_key != "" {
		proxy_host := helper.Config.OpenAI.HttpProxyHost
		proxy_port := helper.Config.OpenAI.HttpProxyPort

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

	ctl.Success(c, gin.H{
		"status":   status,
		"callback": callback,
	})
}
