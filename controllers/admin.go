/*
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /gpt-zmide-server/controllers/admin.go
 */
package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (ctl *Admin) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}

func (ctl *Admin) SignOut(c *gin.Context) {
	auth := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", -1)
	if auth != "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Redirect(http.StatusTemporaryRedirect, "/admin")
}
