/*
 * @Author: Bin
 * @Date: 2023-03-27
 * @FilePath: /gpt-zmide-server/controllers/apis/chat.go
 */
package apis

import (
	"gpt-zmide-server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Chat struct {
	Controller
}

func (ctl *Chat) Index(c *gin.Context) {

	pageForm := &models.PaginateForm{
		Limit: 10,
		Index: 1,
	}

	if c.ShouldBindQuery(&pageForm) != nil {
		c.ShouldBindJSON(&pageForm)
	}

	var chats []models.Chat

	// 获取分页数据
	pageForm, pageOffset, pageTotal := models.ModelPaginate(&chats, pageForm)

	err := models.DB.Limit(pageForm.Limit).
		Offset(pageOffset).
		Preload("Application",
			func(query *gorm.DB) *gorm.DB {
				return query.Model(models.Application{})
			}).
		Find(&chats).Error
	if err != nil {
		ctl.Fail(c, err.Error())
		return
	}

	type ChatItem struct {
		models.Chat
		MessagesCount int64 `json:"massages_count"`
	}
	chatsList := []ChatItem{}
	for _, chat := range chats {
		newChat := &ChatItem{}
		newChat.ID = chat.ID
		newChat.AppID = chat.AppID
		newChat.Remark = chat.Remark
		newChat.Messages = chat.Messages
		newChat.Application = chat.Application
		newChat.CreatedAt = chat.CreatedAt
		newChat.UpdatedAt = chat.UpdatedAt

		models.DB.Model(&models.Message{}).Where("chat_id = ?", chat.ID).Count(&newChat.MessagesCount)
		chatsList = append(chatsList, *newChat)
	}

	ctl.SuccessList(c, chatsList, pageForm, pageTotal)
}
