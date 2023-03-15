/*
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /gpt-zmide-server/controllers/admin.go
 */
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (ctl *Admin) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}
