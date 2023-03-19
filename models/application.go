/*
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /gpt-zmide-server/models/application.go
 */
package models

import (
	"errors"
	"gpt-zmide-server/helper"

	"github.com/google/uuid"
)

type Application struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"unique" json:"name"`
	AppSecret string `gorm:"unique" json:"app_secret"`
	AppKey    string `gorm:"unique;index" json:"app_key"`
	Status    uint   `json:"status"`
	Model
}

// 创建应用
func CreateApplication(name string) (app *Application, err error) {
	if name == "" {
		return nil, errors.New("应用名不得为空")
	}
	app = &Application{
		Name: name,
	}

	if err = DB.Where(app).First(app).Error; err == nil {
		app, err = nil, errors.New("应用已经存在，请更换应用名")
		return
	}

	key := helper.RandomStr(32)
	app.AppSecret = uuid.NewString()
	app.AppKey = key
	app.Status = 1

	if err = DB.Create(app).Error; err != nil {
		app = nil
		return
	}
	return
}
