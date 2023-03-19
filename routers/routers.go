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
	r.GET("/admin/signout", new(controllers.Admin).SignOut)

	// r.GET("/test", new(controllers.InstallController).Test) // 测试路由

	api := r.Group("/api")
	{

		apisCtlApp := new(apis.Application)
		apisCtlOpen := new(apis.Open)
		apisCtlConfig := new(apis.Config)

		notDefault := func(ctx *gin.Context) {
			apis.APIDefaultController.Fail(ctx, "404 route not found.")
		}

		api.GET("/", notDefault)
		api.Any("/:route/*no", notDefault)

		// 开放接口
		openApis := api.Group("/open", middleware.BasicAuthOpen())
		openApis.POST("/", apisCtlOpen.Index)
		openApis.POST("/query", apisCtlOpen.Query)

		adminApis := api.Group("/admin", middleware.BasicAuthAdmin())

		// 系统配置
		adminConfig := adminApis.Group("/config")
		adminConfig.POST("/update/password", apisCtlConfig.UpdatePassword)
		adminConfig.GET("/system/info", apisCtlConfig.SystemInfo)
		adminConfig.GET("/ping/openai", apisCtlConfig.PingOpenAI)

		// 后台管理应用接口
		adminApp := adminApis.Group("/application")
		adminApp.GET("/", apisCtlApp.Index)
		adminApp.POST("/create", apisCtlApp.Create)
		adminApp.POST("/:id/update", apisCtlApp.Update)

	}

	return r
}
