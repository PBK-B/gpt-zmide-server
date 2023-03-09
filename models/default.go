/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/models/default.go
 */
package models

import (
	"gpt-zmide-server/helper"
	"time"

	"gorm.io/driver/mysql"
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

	dbUrl, err := helper.Config.GetMysqlUrl()
	if err != nil || dbUrl == nil {
		panic("the database is not configured, please modify the app.conf file to configure the database")
	}

	dsn := helper.Config.Mysql.User + ":" + helper.Config.Mysql.Password + "@tcp(" + dbUrl.Host + ")/" + helper.Config.Mysql.Database + "?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if db, _ := DB.DB(); db != nil {
		db.SetMaxOpenConns(0)
	}

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err == nil && DB != nil {
		// 执行数据库迁移
		DB.AutoMigrate(
			&Application{},
			&Chat{},
			&Message{},
		)
	}
}
