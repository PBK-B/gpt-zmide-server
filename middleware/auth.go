/*
 * @Author: Bin
 * @Date: 2023-03-08
 * @FilePath: /gpt-zmide-server/middleware/auth.go
 */
package middleware

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"gpt-zmide-server/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthUserKey is the cookie name for user credential in basic auth.
const AuthUserKey = "user"

func searchCredential(authValue string) (string, bool) {
	if decoded, err := base64.StdEncoding.DecodeString(strings.Replace(authValue, "Basic ", "", -1)); err == nil && decoded != nil {
		if input := strings.Split(string(decoded), ":"); len(input) == 2 {
			userObj := helper.Config.AdminUser
			user, password := input[0], fmt.Sprintf("%x", md5.Sum([]byte(input[1])))
			if user == userObj.User && password == userObj.Password {
				return userObj.User, true
			}
		}
	}
	return "", false
}

func BasicAuth() gin.HandlerFunc {
	realm := "Basic realm=" + strconv.Quote("Authorization Required")
	return func(c *gin.Context) {

		// 判断程序未初始化，跳转安装部署页面
		if helper.IsInitialize() {
			c.Redirect(http.StatusTemporaryRedirect, "/install")
			return
		}

		// Search user in the slice of allowed credentials
		user, found := searchCredential(c.Request.Header.Get("Authorization"))
		if !found {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
		// c.MustGet(gin.AuthUserKey).
		c.Set(AuthUserKey, user)
	}
}
