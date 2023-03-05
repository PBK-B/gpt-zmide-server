/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/routers/routers.go
 */
package routers

import (
	"github.com/gin-gonic/gin"

	"gpt-zmide-server/controllers"
	"gpt-zmide-server/controllers/apis"
)

func BuildRouter(r *gin.Engine) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	r.GET("/", new(controllers.Index).Index)

	// r.GET("/test", new(controllers.InstallController).Test) // 测试路由

	api := r.Group("/api")
	{

		apisCtl := new(apis.ApiController)
		apisApp := new(apis.App)

		notDefault := func(ctx *gin.Context) {
			apisCtl.Fail(ctx, "404 route not found.")
		}

		api.Any("/", notDefault)
		api.Any("/:NoRoute", notDefault)

		// 应用层接口
		api.GET("/index", apisApp.Index)
		api.GET("/test", apisApp.Test)
	}

	return r
}
