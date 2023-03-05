/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/controllers/apis/app.go
 */
package apis

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	ApiController
}

func (ctl *App) Index(c *gin.Context) {
	ctl.Fail(c, "请求失败。")
}

func (ctl *App) Test(c *gin.Context) {
	ctl.Success(c, gin.H{"msg": "ok!!!"})
}
