/*
 * @Author: Bin
 * @Date: 2023-03-09
 * @FilePath: /gpt-zmide-server/controllers/apis/open.go
 */
package apis

import (
	"gpt-zmide-server/helper"

	"github.com/gin-gonic/gin"
)

type Open struct {
	Controller
}

func (ctl *Open) Index(c *gin.Context) {
	app := c.MustGet(helper.MiddlewareAuthAppKey)
	ctl.Success(c, app)
}
