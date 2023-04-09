/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/main.go
 */
package main

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"gpt-zmide-server/helper"
	"gpt-zmide-server/helper/logger"
	_ "gpt-zmide-server/models"
	"gpt-zmide-server/routers"
)

//go:embed dist/assets
var FSStatic embed.FS

//go:embed dist/views
var FSViews embed.FS

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func newScriptDevNode() html.Node {
	// scriptDev := `
	// <script type="module">
	// 	import RefreshRuntime from 'http://localhost:5173/@react-refresh'
	// 	RefreshRuntime.injectIntoGlobalHook(window)
	// 	window.$RefreshReg$ = () => {}
	// 	window.$RefreshSig$ = () => (type) => type
	// 	window.__vite_plugin_react_preamble_installed__ = true
	// </script>
	// `
	scriptDev := html.Node{
		Type:     html.ElementNode,
		Data:     "script",
		DataAtom: atom.Body,
		Attr: []html.Attribute{
			{
				Key: "type",
				Val: "module",
			},
		},
	}
	scriptDev.AppendChild(&html.Node{
		Type: html.TextNode,
		Data: `
		import RefreshRuntime from '/@react-refresh'
		RefreshRuntime.injectIntoGlobalHook(window)
		window.$RefreshReg$ = () => {}
		window.$RefreshSig$ = () => (type) => type
		window.__vite_plugin_react_preamble_installed__ = true
		`,
	})
	return scriptDev
}

func main() {
	// 判断是部署环境
	if helper.IsRelease() {
		gin.SetMode(gin.ReleaseMode)
	}

	logger.InitLogger()

	r := gin.Default()

	// 配置静态文件路由
	if gin.Mode() == "debug" {
		// 前端调试模式
		// 注入前端调试代码
		var templ = template.New("")
		templDir, err := template.ParseGlob("./views/*")
		if err == nil {
			for _, item := range templDir.Templates() {
				document, err := goquery.NewDocumentFromReader(strings.NewReader(item.Tree.Root.String()))
				if err == nil {
					headStr := document.Find("head")
					nodeItem := newScriptDevNode()
					headStr.AddBack().Get(0).AppendChild(&nodeItem)
				}

				templItem := template.Must(templ.Parse(renderNode(document.AddBack().Get(0))))
				// templ := template.Must(template.New(templ.Tree.Root.String()).Parse())
				logger.Debug("解析 html " + item.Name())
				templ.AddParseTree(item.Name(), templItem.Tree)
			}
		}
		// goquery.NewDocument();

		r.SetHTMLTemplate(templ)

		r.StaticFS("/assets", http.Dir("./dist/assets"))

		// 反向代理代码目录
		proxyHandle := func(c *gin.Context) {
			remote, err := url.Parse("http://localhost:5173")
			if err != nil {
				panic(err)
			}
			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.ServeHTTP(c.Writer, c.Request)
		}
		// r.Any("/src/:path/*name", proxyHandle)
		r.Any("/src/*name", proxyHandle)
		r.Any("/@id/*name", proxyHandle)
		r.Any("/node_modules/*name", proxyHandle)
		r.Any("/@vite/*name", proxyHandle)
		r.Any("/@react-refresh", proxyHandle)

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
	logger.Info("gpt-zmide-server start up, listening and serving HTTP on " + serverHost.String())
	r.Run(serverHost.Host)
}
