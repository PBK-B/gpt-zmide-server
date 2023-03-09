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
	"gpt-zmide-server/middleware"
)

func BuildRouter(r *gin.Engine) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	r.GET("/", new(controllers.Index).Index)

	r.GET("/admin", middleware.BasicAuth(), new(controllers.Admin).Index)

	// r.GET("/test", new(controllers.InstallController).Test) // 测试路由

	api := r.Group("/api")
	{

		apisCtl := new(apis.Controller)
		apisApp := new(apis.Application)

		notDefault := func(ctx *gin.Context) {
			apisCtl.Fail(ctx, "404 route not found.")
		}

		api.GET("/", notDefault)
		api.Any("/:route/*no", notDefault)

		adminApis := api.Group("/admin")

		// 后台管理应用接口
		adminApp := adminApis.Group("/application")
		adminApp.GET("/index", apisApp.Index)
		adminApp.POST("/create", apisApp.Create)
		adminApp.POST("/:id/update", apisApp.Update)

	}

	return r
}
