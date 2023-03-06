/*
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /gpt-zmide-server/models/message.go
 */
package models

type Message struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	ChatId uint `json:"chat_id"`
	Data   string
	Model
}
