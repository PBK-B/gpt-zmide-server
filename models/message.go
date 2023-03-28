/*
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /gpt-zmide-server/models/message.go
 */
package models

type Message struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	ChatID  uint   `json:"chat_id"`
	Role    string `json:"role"`
	Content string `json:"content"`
	Raw     string `json:"-"`
	BaseModel
}
