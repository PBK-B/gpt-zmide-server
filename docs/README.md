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
```shell
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

```shell
curl --location 'https://example.zmide.com/api/open/query' \
--header 'Authorization: Bearer GDAiQmNHhoHdqoFqxrGObTSObYSySast' \
--form 'content="在 CPU 中配置高速缓冲器（Cache）是为了解决啥？"'
```

**响应体**
| 名称 | 类型 | 描述 |
| --- | --- | --- |
| code | int | 响应代码: 200 为请求成功  |
| status | string | 请求响应状态 |
| msg | string | 请求失败消息 |
| data | object | 数据对象 |
| - id | int | 消息 ID |
| - chat_id | int | 会话 ID |
| - role | int | 会话角色 |
| - content | int | 消息内容 |
| - created_at | int | 创建时间 |

**响应示例**
```json
{
    "code": 200,
    "data": {
        "id": 667,
        "chat_id": 142,
        "role": "assistant",
        "content": "高速缓存器的主要目的是解决CPU和内存的速度不匹配问题。CPU的处理速度比内存快得多，但是当CPU需要从内存中读取数据时，传输速度却显得很慢，这就会影响CPU的性能。高速缓存器被设计成一个小而快速的存储介质，它可以存储已经读取过的数据，并提供快速访问，这样CPU就可以在不用等待内存响应的情况下，快速地访问缓存中的数据，提高了整个计算机系统的性能。",
        "created_at": "2023-03-21 15:28:32",
        "updated_at": "2023-03-21 15:28:32"
    },
    "status": "ok"
}
```
