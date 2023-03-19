/*
 * @Author: Bin
 * @Date: 2023-03-09
 * @FilePath: /gpt-zmide-server/controllers/apis/open.go
 */
package apis

import (
	"gpt-zmide-server/helper"
	"gpt-zmide-server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Open struct {
	Controller
}

func (ctl *Open) Index(c *gin.Context) {
	app := c.MustGet(helper.MiddlewareAuthAppKey)
	ctl.Success(c, app)
}

func (ctl *Open) Query(c *gin.Context) {
	appTmp := c.MustGet(helper.MiddlewareAuthAppKey)

	if appTmp == nil {
		ctl.Fail(c, "应用异常")
		return
	}

	var app *models.Application
	var ok bool
	if app, ok = appTmp.(*models.Application); !ok || app == nil || app.Status != 1 {
		ctl.Fail(c, "应用异常")
		return
	}

	content := c.PostForm("content")
	p_chat_id := c.PostForm("chat_id")
	p_remark := c.PostForm("remark")

	// content 参数为必传
	if content == "" {
		ctl.Fail(c, "参数异常")
		return
	}

	chat := &models.Chat{}

	if p_chat_id != "" {
		if id, err := strconv.Atoi(p_chat_id); err == nil && id != 0 {
			// 当 chat_id 合法时，去数据库查找 chat
			chat.ID = uint(id)
			if err := models.DB.First(chat).Error; err != nil || chat.AppID != app.ID {
				ctl.Fail(c, "chat_id 不合法")
				return
			}
			chat.AppID = app.ID
		}
	}

	if chat.AppID == 0 {
		chat.AppID = app.ID
		if err := models.DB.Create(chat).Error; err != nil {
			ctl.Fail(c, "chat 处理异常")
			return
		}
	}

	// 当 remark 参数存在时更新 chat remark
	if p_remark != "" {
		chat.Remark = p_remark
		models.DB.Updates(chat)
	}

	message := &models.Message{
		ChatID:  chat.ID,
		Role:    "user",
		Content: content,
		Raw:     "",
	}

	if err := models.DB.Create(message).Error; err != nil {
		ctl.Fail(c, "消息处理失败")
		return
	}

	// 刷新 chat Messages
	models.DB.Preload("Messages").Find(chat)

	callback, err := chat.QueryChatGPT()
	if err != nil {
		ctl.Fail(c, err.Error())
		return
	}

	ctl.Success(c, callback)
}
