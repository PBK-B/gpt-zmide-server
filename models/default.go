/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/models/default.go
 */
package models

import (
	"database/sql/driver"
	"fmt"
	"gpt-zmide-server/helper"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Model struct {
	CreatedAt LocalTime `json:"created_at"`
	UpdatedAt LocalTime `json:"updated_at"`
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

type LocalTime struct {
	time.Time
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = LocalTime{t1}
	return err
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
