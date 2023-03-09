/*
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /gpt-zmide-server/models/chat.go
 */
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"gpt-zmide-server/helper"

	"github.com/go-resty/resty/v2"
)

type Chat struct {
	ID       uint       `gorm:"primaryKey" json:"id"`
	AppID    uint       `json:"app_id"`
	Remark   string     `json:"remark"`
	Messages []*Message `gorm:"foreignKey:ChatID" json:"messages"`
	Model
}

func (chat *Chat) QueryChatGPT() (msg *Message, err error) {

	model := helper.Config.OpenAI.Model
	secret_key := helper.Config.OpenAI.SecretKey
	if model == "" || secret_key == "" {
		return nil, errors.New("OpenAI model 未设置或 secret_key 未设置")
	}

	if len(chat.Messages) < 1 {
		return nil, errors.New("chat messages 处理异常")
	}

	type MsgTmp struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}

	// 处理会话数据
	var msgs = []*MsgTmp{}
	for _, item := range chat.Messages {
		msgs = append(msgs, &MsgTmp{
			Role:    item.Role,
			Content: item.Content,
		})
	}
	msgsStr, err := json.Marshal(msgs)
	if err != nil {
		return nil, err
	}

	proxy_host := helper.Config.OpenAI.HttpProxyHost
	proxy_port := helper.Config.OpenAI.HttpProxyPort

	client := resty.New()
	if proxy_host != "" && proxy_port != "" {
		client.SetProxy("http://" + proxy_host + ":" + proxy_port)
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+secret_key).
		SetBody(`{
			"model": "` + model + `",
			"user": "` + helper.Config.SiteName + `",
			"max_tokens": 1800,
			"messages": ` + string(msgsStr) + `
		}`).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		return nil, err
	}

	// fmt.Println("数据" + resp.String())

	type Choice struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	}

	var data struct {
		Choices []Choice `json:"choices"`
	}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if len(data.Choices) < 1 {
		return nil, errors.New("openai api callback choices data error")
	}

	choiceFirst := data.Choices[0]
	msg = &Message{
		ChatID:  chat.ID,
		Raw:     resp.String(),
		Role:    choiceFirst.Message.Role,
		Content: choiceFirst.Message.Content,
	}

	if err = DB.Create(msg).Error; err != nil {
		fmt.Println("message create error " + err.Error())
	}

	return msg, nil
}
