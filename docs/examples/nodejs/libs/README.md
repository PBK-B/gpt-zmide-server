<!--
 * @Author: Bin
 * @Date: 2023-04-01
 * @FilePath: /gpt-zmide-server/docs/examples/typescript/libs/README.md
-->
# zmide-gpt-libs

> the javascript library using gpt-zmide-server api

## install 
you can use `yarn` or `npm` to install.
```shell
yarn add zmide-gpt-libs
# or 
npm install zmide-gpt-libs
```

## use
```
import GPTServer from 'zmide-gpt-libs';

// 配置密钥
const server = new GPTServer("xxxxxx")

// 查询
server.Query("计算机是什么？").then((reply) => {
    const { data: { code, data, msg } } = reply
    if (code != 200) {
        console.error("查询失败", code, msg)
        return
    }
    console.log("reply", data)
}).catch((err) => {
    console.error("error", err)
})
```
