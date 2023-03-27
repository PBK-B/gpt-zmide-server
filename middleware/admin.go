/*
 * @Author: Bin
 * @Date: 2023-03-10
 * @FilePath: /gpt-zmide-server/middleware/admin.go
 */
package middleware

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"gpt-zmide-server/controllers/apis"
	"gpt-zmide-server/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func adminCredential(token string) bool {
	if token == "" {
		return false
	}

	if decoded, err := base64.StdEncoding.DecodeString(token); err == nil && decoded != nil {
		if input := strings.Split(string(decoded), ":"); len(input) == 2 {
			userObj := helper.Config.AdminUser
			user, password := input[0], fmt.Sprintf("%x", md5.Sum([]byte(input[1])))
			if user == userObj.User && password == userObj.Password {
				return true
			}
		}
	}

	return false
}

func BasicAuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Search user in the slice of allowed credentials
		auth := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", -1)
		if auth == "" {
			auth = c.Query("token")
		}

		if !adminCredential(auth) {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			apis.APIDefaultController.Fail(c, "请登录管理员账号")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}

func InstallMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果系统已安装，直接禁止访问安装页面
		if !helper.IsInitialize() {
			c.AbortWithStatus(http.StatusNotFound)
			c.Writer.Write([]byte("404 page not found"))
			return
		}
	}
}
