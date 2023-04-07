/*
 * @Author: Bin
 * @Date: 2023-04-07
 * @FilePath: /gpt-zmide-server/docs/examples/nodejs/main.js
 */

import GPTServer from 'zmide-gpt-libs';

// 配置密钥
const server = new GPTServer("xxxxxx")

// 查询
server.Query("计算机是什么？").then((reply) => {
    const { data: { code, data, msg } } = reply
    if (code != 200) {
        console.error("查询失败", code, msg);
        return
    }
    console.log("reply", data);
}).catch((err) => {
    console.error("error", err);
})
