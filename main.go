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
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"gpt-zmide-server/routers"
)

//go:embed static
var FSStatic embed.FS

//go:embed views
var FSViews embed.FS

const AppName = "gpt-zmide-server"

func main() {
	r := gin.Default()

	// 静态资源嵌入
	templ := template.Must(template.New("").ParseFS(FSViews, "views/*"))
	r.SetHTMLTemplate(templ)

	// 注册路由
	r = routers.BuildRouter(r)

	// 配置静态文件路由
	if !IsRelease() {
		r.StaticFS("/static", http.Dir("./static")) // 前端调试模式
	} else {
		staticDir, _ := fs.Sub(FSStatic, "static")
		r.StaticFS("/static", http.FS(staticDir))
	}

	// Listen and Server in 0.0.0.0:8091
	r.Run("0.0.0.0:8091")
}

func IsRelease() bool {
	arg1 := strings.ToLower(os.Args[0])
	name := filepath.Base(arg1)
	return strings.Index(name, AppName) == 0 && strings.Index(arg1, "go-build") < 0
}
