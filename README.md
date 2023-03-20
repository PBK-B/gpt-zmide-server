<!--
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/README.md
-->
# gpt-zmide-server 💡

> zmide ChatGPT 应用服务，用于管理应用程序对接和集成 ChatGPT API 的服务应用，提供简单易用的 API 服务。

## 开始 🎀

1. 创建 `app.conf` 配置文件

2. 启动服务 `go run .`

3. 修改 `app.conf` 配置文件，重启服务
    ```
    site_name: 站点名称
    domain_name: https://demo.zmide.com
    host: 0.0.0.0
    port: 8091
    admin_user:
        user: admin
        password: 
    mysql:
        host: localhost
        port: 3306
        user: root
        password:
        database:
    openai:
        secret_key:
        model: gpt-3.5-turbo
        http_proxy_host:
        http_proxy_port:
    ```

4. 访问 `http://127.0.0.1:8091/admin`

## 文档 📜

API 文档请参考: [docs/README.md](/docs/README.md)

## 截图 🔦

<img src="docs/images/screenshot_1001.png" width="560">

<img src="docs/images/screenshot_1002.png" width="560">

## 谁在使用

- [全能搜题](https://github.com/zmide/study.zmide.com) 全能搜题项目是一个基于开源社区公开贡献的永久免费搜题系统。

## 感谢支持 😋

- [OpenAI](https://openai.com/) Creating safe artificial general intelligence that benefits all of humanity

- [gin](https://gin-gonic.com/) Gin Web Framework

- [gorm.io/gorm](https://gorm.io/) The fantastic ORM library for Golang


