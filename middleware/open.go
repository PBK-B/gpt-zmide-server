/*
 * @Author: Bin
 * @Date: 2023-03-09
 * @FilePath: /gpt-zmide-server/middleware/open.go
 */
package middleware

import (
	"encoding/json"
	"errors"
	"gpt-zmide-server/helper"
	"gpt-zmide-server/models"
	"net/http"
	"strings"

	"gpt-zmide-server/controllers/apis"

	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt/aes"
)

func applicationCredential(token string, appKey string) (*models.Application, error) {
	if token == "" {
		return nil, errors.New("authorization 为空")
	}

	app := &models.Application{AppKey: appKey}
	if err := models.DB.Where("app_key = ?", appKey).First(app).Error; err != nil {
		return nil, err
	}

	// 校验 token 是否有效
	plaintext, err := aes.AesCbcDecryptByBase64(token, []byte(strings.Replace(app.AppSecret, "-", "", -1)), make([]byte, 16))
	if err != nil || string(plaintext) != appKey {
		return nil, errors.New("authorization 异常")
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

		appKey := c.Request.Header.Get("Applicationkey")
		if appKey == "" {
			appKey = c.Query("app_key")
		}

		app, err := applicationCredential(auth, appKey)
		if err != nil || app == nil {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			apis.APIDefaultController.Fail(c, "应用认证失败。")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 如启用加密则需进行数据解密
		encryptBody := c.Request.Header.Get("EncryptBody")
		if encryptBody == "1" {
			body, _ := c.GetRawData()
			plaintext, _ := aes.AesCbcDecryptByBase64(string(body), []byte(strings.Replace(app.AppSecret, "-", "", -1)), make([]byte, 16))

			var bodyMap map[string]string
			if err := json.Unmarshal(plaintext, &bodyMap); err != nil {
				apis.APIDefaultController.Fail(c, "body 校验失败。")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set(helper.PostBodyKey, bodyMap)
		}

		// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
		// c.MustGet(gin.AuthUserKey).
		c.Set(helper.MiddlewareAuthAppKey, app)
	}
}
