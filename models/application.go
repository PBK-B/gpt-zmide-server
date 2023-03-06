/*
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /gpt-zmide-server/models/application.go
 */
package models

type Application struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"unique;index" json:"name"`
	AppSecret string `gorm:"unique" json:"app_secret"`
	AppKey    string `gorm:"unique" json:"app_key"`
	Status    uint   `json:"status"`
	Model
}
