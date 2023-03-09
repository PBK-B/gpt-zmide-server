/*
 * @Author: Bin
 * @Date: 2023-03-09
 * @FilePath: /gpt-zmide-server/middleware/open.go
 */
package middleware

import (
	"encoding/base64"
	"errors"
	"gpt-zmide-server/helper"
	"gpt-zmide-server/models"
	"net/http"
	"strings"

	"gpt-zmide-server/controllers/apis"

	"github.com/gin-gonic/gin"
)

func applicationCredential(authorization string) (*models.Application, error) {
	if authorization == "" {
		return nil, errors.New("authorization 为空")
	}

	if app_key, err := base64.StdEncoding.DecodeString(strings.Replace(authorization, "Bearer ", "", -1)); err == nil && app_key != nil {
		app := &models.Application{AppKey: string(app_key)}
		if err := models.DB.Where("app_key = ?", app_key).First(app).Error; err != nil {
			return nil, err
		}
		if app != nil && app.Status != 0 {
			return app, nil
		}
	}

	return nil, errors.New("authorization 异常")
}

func BasicAuthOpen() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Search user in the slice of allowed credentials
		app, err := applicationCredential(c.Request.Header.Get("Authorization"))
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
