/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/controllers/apis/api.go
 */
package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
}

// 统一请求成功回调数据结构
func (ctl *ApiController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "code": 200, "data": data})
}

// 统一请求失败回调数据结构
func (ctl *ApiController) Fail(c *gin.Context, massage string) {
	c.JSON(http.StatusOK, gin.H{"status": "fail", "code": 400, "data": nil, "msg": massage})
}
