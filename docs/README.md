<!--
 * @Author: Bin
 * @Date: 2023-03-11
 * @FilePath: /gpt-zmide-server/docs/README.md
-->
# API

## 准备
以下文档中的 `example.zmide.com` API 地址请替换为您的线上服务器地址。

### API 调用方式

调用服务端接口时，需要使用 HTTPS 协议、JSON 数据格式、UTF8 编码。请求需要把访问凭证 app_key 放到 Header 的 Authorization 中。

Authorization 格式为 `Bearer ` + app_key 例如: `Bearer GDAiQmNHhoHdqoFqxrGObTSObYSySast`

示例:
```bash
curl -X POST 'https://example.zmide.com/api/open/'
-H 'Authorization: Bearer GDAiQmNHhoHdqoFqxrGObTSObYSySast'
-d '{
	"content": "Hello World"
}'
```

## 接口列表

### 查询

---

**请求**
| 基本 ||
| --- | --- |
| HTTP Path | /api/open/query |
| HTTP Method | POST |

**请求头**
| 名称 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 应用密钥 app_key<br> **示例值：**"Bearer GDAiQmNHhoHdqoFqxrGObTSObYSySast" |

**请求体**
| 名称 | 类型 | 必填 | 描述 |
| --- | ---| --- | --- |
| content | string | 是 | 消息内容<br>**示例值:**"在 CPU 中配置高速缓冲器（Cache）是为了解决啥？" |
| chat_id | string | 否 | 会话 ID，默认 0 为创建新会话<br>**示例值:**"1" |
| remark | string | 否 | 用户标记，用于标识用户<br>**示例值:** "user001"|

**请求体示例**

```
curl --location 'https://example.zmide.com/api/open/query' \
--header 'Authorization: Bearer GDAiQmNHhoHdqoFqxrGObTSObYSySast' \
--form 'content="在 CPU 中配置高速缓冲器（Cache）是为了解决啥？"'
```