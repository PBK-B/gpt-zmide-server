/*
 * @Author: Bin
 * @Date: 2023-03-09
 * @FilePath: /gpt-zmide-server/middleware/open.go
 */
package middleware

import (
	"errors"
	"gpt-zmide-server/helper"
	"gpt-zmide-server/models"
	"net/http"
	"strings"

	"gpt-zmide-server/controllers/apis"

	"github.com/gin-gonic/gin"
)

func applicationCredential(token string) (*models.Application, error) {
	if token == "" {
		return nil, errors.New("authorization 为空")
	}

	app := &models.Application{AppKey: token}
	if err := models.DB.Where("app_key = ?", token).First(app).Error; err != nil {
		return nil, err
	}
	if app != nil && app.Status == 1 {
		return app, nil
	}

	return nil, errors.New("authorization 异常")
}

func BasicAuthOpen() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Search user in the slice of allowed credentials
		auth := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", -1)
		if auth == "" {
			auth = c.Query("token")
		}

		app, err := applicationCredential(auth)
		if err != nil || app == nil {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			apis.APIDefaultController.Fail(c, "应用认证失败。")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
		// c.MustGet(gin.AuthUserKey).
		c.Set(helper.MiddlewareAuthAppKey, app)
	}
}
