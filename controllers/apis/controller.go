/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/controllers/apis/controller.go
 */
package apis

import (
	"gpt-zmide-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var APIDefaultController = new(Controller)

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

// 统一请求成功分页列表数据回调数据结构
func (ctl *Controller) SuccessList(c *gin.Context, list interface{}, pageForm *models.PaginateForm, pageTotal int) {
	ctl.Success(c, gin.H{
		"list":       list,
		"page_index": pageForm.Index,
		"page_limit": pageForm.Limit,
		"page_total": pageTotal,
	})
}
