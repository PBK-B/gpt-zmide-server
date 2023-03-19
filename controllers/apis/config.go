/*
 * @Author: Bin
 * @Date: 2023-03-19
 * @FilePath: /gpt-zmide-server/controllers/apis/config.go
 */
package apis

import (
	"crypto/md5"
	"fmt"
	"gpt-zmide-server/helper"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Controller
}

// 修改密码
func (ctl *Config) UpdatePassword(c *gin.Context) {

	old_password := c.PostForm("old_password")
	new_password := c.PostForm("new_password")
	if old_password == "" || new_password == "" || len(new_password) < 3 {
		ctl.Fail(c, "参数异常")
		return
	}

	old_pwd := fmt.Sprintf("%x", md5.Sum([]byte(old_password)))
	if old_pwd != helper.Config.AdminUser.Password {
		ctl.Fail(c, "旧密码错误")
		return
	}

	if old_password == new_password {
		ctl.Fail(c, "两次输入密码相同")
		return
	}

	new_pwd := fmt.Sprintf("%x", md5.Sum([]byte(new_password)))
	helper.Config.AdminUser.Password = new_pwd
	helper.Config.SaveConfig()

	ctl.Success(c, "ok")
}
