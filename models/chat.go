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
)

type Chat struct {
	ID          uint             `gorm:"primaryKey" json:"id"`
	AppID       uint             `json:"-"`
	Remark      string           `json:"remark"`
	Messages    []*Message       `gorm:"foreignKey:ChatID" json:"messages"`
	Application *ChatApplication `gorm:"foreignKey:AppID" json:"app"`
	Model       string           `json:"model"`
	BaseModel
}

type ChatApplication struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	AppSecret string `json:"-"`
	AppKey    string `json:"-"`
	Status    uint   `json:"-"`
}

func (chat *Chat) QueryChatGPT() (msg *Message, err error) {

	model := chat.Model
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
		if contextCount > 4500 {
			// 避免消息上下文超过 4600 字数限制
			continue
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
		return nil, errors.New("openai api callback choices data error:" + string(resp.Body()))
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
