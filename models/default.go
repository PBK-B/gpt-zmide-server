/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/models/default.go
 */
package models

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Model struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {

	if DB != nil {
		return
	}

	var err error
	// DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{}) // 连接数据库
	// if db, _ := DB.DB(); db != nil {
	// 	db.SetMaxOpenConns(1)
	// }

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err == nil && DB != nil {
		// 执行数据库迁移
		DB.AutoMigrate()
	}
}
