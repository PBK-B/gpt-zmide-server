/*
 * @Author: Bin
 * @Date: 2023-03-21
 * @FilePath: /gpt-zmide-server/controllers/install.go
 */
package controllers

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"gpt-zmide-server/controllers/apis"
	"gpt-zmide-server/helper"
	"gpt-zmide-server/models"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Install struct {
	apis.Controller
}

func (ctl *Install) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "install.html", nil)
}

// 提交配置
func (ctl *Install) Config(c *gin.Context) {
	if !helper.IsInitialize() {
		ctl.Fail(c, "出错啦")
		return
	}

	step := c.PostForm("step")
	data := c.PostForm("data")
	if step == "" || data == "" {
		ctl.Fail(c, "参数异常")
		return
	}

	var err error
	switch step {
	case "site":
		err = ctl.siteConfig(data)
	case "openai":
		err = ctl.openaiConfig(data)
	case "database":
		err = ctl.databaseConfig(data)
	}
	if err != nil {
		ctl.Fail(c, err.Error())
		return
	}

	ctl.Success(c, "ok!")
}

// 配置站点
func (ctl *Install) siteConfig(data string) error {
	type SiteConfig struct {
		SiteName      string `json:"site_name"`
		DomainName    string `json:"domain_name"`
		Port          string `json:"port"`
		AdminUser     string `json:"admin_user"`
		AdminPassword string `json:"admin_password"`
	}

	var config SiteConfig
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return errors.New("参数填写错误")
	}

	if config.SiteName == "" {
		return errors.New("站点名称不得为空")
	}

	domainName, err := url.Parse(config.DomainName)
	if config.DomainName == "" || err != nil || domainName == nil {
		return errors.New("站点域名错误")
	}

	port, err := strconv.Atoi(config.Port)
	if err != nil || port == 0 {
		return errors.New("应用端口错误")
	}

	if config.AdminUser == "" || len(config.AdminUser) < 3 {
		return errors.New("管理员用户名为空或小于 3 位")
	}

	if config.AdminPassword == "" || len(config.AdminPassword) < 6 {
		return errors.New("管理员密码为空或小于 6 位")
	}

	helper.Config.SiteName = config.SiteName
	helper.Config.DomainName = config.DomainName
	helper.Config.Port = port
	helper.Config.AdminUser.User = config.AdminUser
	helper.Config.AdminUser.Password = fmt.Sprintf("%x", md5.Sum([]byte(config.AdminPassword)))

	if err := helper.Config.SaveConfig(); err != nil {
		return err
	}

	return nil
}

// 配置 OpenAI
func (ctl *Install) openaiConfig(data string) error {
	type OpenAIConfig struct {
		SecretKey string `json:"openai_secret_key"`
		ProxyHost string `json:"openai_proxy_host"`
		ProxyPort string `json:"openai_proxy_port"`
		Model     string `json:"openai_model"`
	}

	var config OpenAIConfig
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return errors.New("参数填写错误")
	}

	if config.SecretKey == "" {
		return errors.New("OpenAI SecretKey 填写错误")
	}

	status, callback := helper.PingOpenAI(config.SecretKey, config.ProxyHost, config.ProxyPort)
	if !status {
		return errors.New("OpenAI 服务器连接失败，" + callback)
	}

	helper.Config.OpenAI.SecretKey = config.SecretKey
	helper.Config.OpenAI.HttpProxyHost = config.ProxyHost
	helper.Config.OpenAI.HttpProxyPort = config.ProxyPort
	helper.Config.OpenAI.Model = config.Model

	if err := helper.Config.SaveConfig(); err != nil {
		return err
	}

	return nil
}

// 配置数据库
func (ctl *Install) databaseConfig(data string) error {
	type DatabaseConfig struct {
		MysqlHost     string `json:"mysql_host"`
		MysqlPort     string `json:"mysql_port"`
		MysqlUser     string `json:"mysql_user"`
		MysqlPassword string `json:"mysql_password"`
		MysqlDatabase string `json:"mysql_database"`
	}

	var config DatabaseConfig
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return errors.New("参数填写错误")
	}

	port, err := strconv.Atoi(config.MysqlPort)
	if err != nil || port == 0 {
		return errors.New("数据库端口错误")
	}

	if config.MysqlHost == "" {
		return errors.New("数据库地址错误")
	}

	dbURL, err := helper.GetMysqlUrl(config.MysqlHost, port)
	if err != nil || dbURL == nil {
		return errors.New("数据库地址或端口错误")
	}

	if config.MysqlUser == "" {
		return errors.New("数据库用户名错误")
	}

	if config.MysqlDatabase == "" {
		return errors.New("数据库名称错误")
	}

	u_p := config.MysqlUser
	if config.MysqlPassword != "" {
		u_p = u_p + ":" + config.MysqlPassword
	}

	// 检查数据库是否连接成功
	dsn := u_p + "@tcp(" + dbURL.Host + ")/" + config.MysqlDatabase + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("数据库连接失败" + err.Error())
	}
	if db == nil {
		return errors.New("数据库连接失败")
	}

	helper.Config.Mysql.Host = config.MysqlHost
	helper.Config.Mysql.Port = port
	helper.Config.Mysql.User = config.MysqlUser
	helper.Config.Mysql.Password = config.MysqlPassword
	helper.Config.Mysql.Database = config.MysqlDatabase

	if err := helper.Config.SaveConfig(); err != nil {
		return err
	}

	// 初始化数据库
	if err := models.InitDB(); err != nil {
		return errors.New("数据迁移失败")
	}

	return nil
}
