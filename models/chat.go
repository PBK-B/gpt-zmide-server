/*
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /gpt-zmide-server/models/chat.go
 */
package models

import (
	"encoding/json"
	"errors"
	"gpt-zmide-server/helper"
	"gpt-zmide-server/helper/logger"
)

type Chat struct {
	ID              uint             `gorm:"primaryKey" json:"id"`
	AppID           uint             `json:"-"`
	Remark          string           `json:"remark"`
	Messages        []*Message       `gorm:"foreignKey:ChatID" json:"messages"`
	App             *Application     `json:"-"`
	ChatApplication *ChatApplication `gorm:"-" json:"app"`
	Model
}

type ChatApplication struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
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
	var msgsTmp = []*MsgTmp{}
	msgCount := 0
	// 倒序遍历消息记录
	for i := len(chat.Messages) - 1; i >= 0; i-- {
		item := chat.Messages[i]
		contextCount := msgCount + len(item.Content)
		// 避免消息上下文超过 4600 字数限制
		if contextCount > 4500 {
			// 判断应用是否需要修复长消息
			DB.Preload("App").Find(chat)
			if chat.App != nil && chat.App.EnableFixLongMsg != 1 {
				continue
			} else {
				return nil, errors.New("消息上下文超过 4600 字数限制")
			}
		}
		msgCount = contextCount
		msgsTmp = append(msgsTmp, &MsgTmp{
			Role:    item.Role,
			Content: item.Content,
		})
	}
	// 修正消息顺序
	var msgs = []*MsgTmp{}
	for i := len(msgsTmp) - 1; i >= 0; i-- {
		msgs = append(msgsTmp, msgsTmp[i])
	}

	msgsStr, err := json.Marshal(msgs)
	if err != nil {
		return nil, err
	}

	client, err := helper.Config.GetOpenAIHttpClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.R().
		SetBody(`{
			"model": "` + model + `",
			"user": "` + helper.Config.SiteName + `",
			"max_tokens": 1800,
			"messages": ` + string(msgsStr) + `
		}`).
		Post("/v1/chat/completions")

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
		logger.Warn("OpenAI CallBack Data unusual " + resp.String())
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
		// fmt.Println("message create error " + err.Error())
		logger.Error("message create error " + err.Error())
	}

	return msg, nil
}
