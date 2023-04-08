<!--
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/README.md
-->
# gpt-zmide-server 💡

[![GitHub Repo stars](https://img.shields.io/github/stars/pbk-b/gpt-zmide-server?style=social)](https://github.com/PBK-B/gpt-zmide-server)
[![Docker Image Version (latest by date)](https://img.shields.io/docker/v/pbkbin/zmide-gpt-started?label=Docker%20Image%20Version)
](https://hub.docker.com/repository/docker/pbkbin/zmide-gpt-started/general)


> zmide ChatGPT 应用服务，用于管理应用程序对接和集成 ChatGPT API 的服务应用，提供简单易用的 API 服务。

## 开始 🎀

### Build Run

1. 创建 `app.conf` 配置文件

2. 启动服务 `go run .`

3. 访问 `http://127.0.0.1:8091/install` 开始安装

4. 访问 `http://127.0.0.1:8091/admin` 登录管理后台

### Docker Install

```
docker pull pbkbin/zmide-gpt-started:v1.1
```

## 文档 📜

API 文档请参考: [docs/README.md](/docs/README.md)

## 截图 🔦

<img src="docs/images/screenshot_1003.png" width="760">

<img src="docs/images/screenshot_1001.png" width="760">

<img src="docs/images/screenshot_1002.png" width="760">

## 计划

- [x] ~~安装引导页面~~

- [ ] 后台会话查询

- [x] ~~后台系统设置~~

- [ ] 敏感词过滤设置

- [ ] 应用请求限速设置

- [ ] 应用单独配置模型

## 开发 🔨

```shell
# 启动前端
yarn && yarn dev

# 启动后端
DEBUG=1 go run .

# 编译项目 (跨平台交叉编译可以修改 Makefile go build 相关参数)
make all
```

## 谁在使用

- [全能搜题](https://github.com/zmide/study.zmide.com) 全能搜题项目是一个基于开源社区公开贡献的永久免费搜题系统。

## 感谢支持 😋

- [OpenAI](https://openai.com/) Creating safe artificial general intelligence that benefits all of humanity

- [gin](https://gin-gonic.com/) Gin Web Framework

- [gorm.io/gorm](https://gorm.io/) The fantastic ORM library for Golang

## 许可 📝

该项目基于 `MIT` 协议自由分发

```
The MIT License

Copyright (c) 2023 zmide Studio Development Team

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
