/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/main.go
 */
package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"

	"gpt-zmide-server/helper"
	_ "gpt-zmide-server/models"
	"gpt-zmide-server/routers"
)

//go:embed dist/assets
var FSStatic embed.FS

//go:embed dist/views
var FSViews embed.FS

func main() {
	// 判断是部署环境
	if helper.IsRelease() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// 配置静态文件路由
	if gin.Mode() == "debug" {
		// 前端调试模式
		templ := template.Must(template.New("").ParseGlob("./dist/views/*"))
		r.SetHTMLTemplate(templ)

		r.StaticFS("/assets", http.Dir("./dist/assets"))
	} else {
		templ := template.Must(template.New("").ParseFS(FSViews, "dist/views/*"))
		r.SetHTMLTemplate(templ)

		staticDir, _ := fs.Sub(FSStatic, "dist/assets")
		r.StaticFS("/assets", http.FS(staticDir))
	}

	// 注册路由
	r = routers.BuildRouter(r)

	// Listen and Server in 0.0.0.0:8091
	host := "http://0.0.0.0:8091"
	if helper.Config.Host != "" {
		host = "http://" + helper.Config.Host + ":" + strconv.Itoa(helper.Config.Port)
	}

	serverHost, err := url.Parse(host)
	if err != nil {
		serverHost = &url.URL{Host: "0.0.0.0:8091"}
	}
	r.Run(serverHost.Host)
}
