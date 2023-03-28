/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/controllers/apis/app.go
 */
package apis

import (
	"gpt-zmide-server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Controller
}

func (ctl *Application) Index(c *gin.Context) {
	var apps []models.Application
	if err := models.DB.Find(&apps).Error; err != nil {
		ctl.Fail(c, err.Error())
		return
	}
	ctl.Success(c, apps)
}

func (ctl *Application) Create(c *gin.Context) {

	name := c.PostForm("name")
	if name == "" {
		ctl.Fail(c, "参数异常")
		return
	}

	app, err := models.CreateApplication(name)
	if err != nil {
		// 创建失败
		ctl.Fail(c, err.Error())
		return
	}
	ctl.Success(c, app)
}

func (ctl *Application) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		ctl.Fail(c, err.Error())
		return
	}

	name, p_status, fix_long_msg := c.PostForm("name"),
		c.PostForm("status"),
		c.PostForm("fix_long_msg")
	if name == "" && p_status == "" && fix_long_msg == "" {
		ctl.Fail(c, "参数异常")
		return
	}

	app := models.Application{ID: uint(id)}
	if err = models.DB.First(&app).Error; err != nil {
		ctl.Fail(c, err.Error())
		return
	}

	if p_status != "" {
		status, err := strconv.Atoi(p_status)
		if err == nil {
			app.Status = uint(status)
		}
	}

	if fix_long_msg != "" {
		fixLongMsg, err := strconv.Atoi(fix_long_msg)
		if err == nil {
			app.EnableFixLongMsg = uint(fixLongMsg)
		}
	}

	if name != "" {
		app.Name = name
	}

	if err = models.DB.Updates(app).Error; err != nil {
		ctl.Fail(c, err.Error())
		return
	}

	ctl.Success(c, app)
}
