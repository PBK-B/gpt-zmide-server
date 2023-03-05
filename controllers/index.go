/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/controllers/index.go
 */
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Index struct {
}

func (ctl *Index) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tpl", nil)
}
