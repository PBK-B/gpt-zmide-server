/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/controllers/apis/controller.go
 */
package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

// 统一请求成功回调数据结构
func (ctl *Controller) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "code": 200, "data": data})
}

// 统一请求失败回调数据结构
func (ctl *Controller) Fail(c *gin.Context, massage string) {
	c.JSON(http.StatusOK, gin.H{"status": "fail", "code": 400, "data": nil, "msg": massage})
}
