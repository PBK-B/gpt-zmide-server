/*
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /gpt-zmide-server/models/chat.go
 */
package models

type Chat struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	AppId  uint   `json:"app_id"`
	Remark string `json:"remark"`
	Model
}
