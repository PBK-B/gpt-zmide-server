/*
 * @Author: Bin
 * @Date: 2023-04-01
 * @FilePath: /gpt-zmide-server/docs/examples/nodejs/libs/src/index.ts
 */
import axios from 'axios';
import FormData from 'form-data';

const DEFAULT_API_SERVER = "http://127.0.0.1:8091"

interface GPTServerConfig {
    server?: {
        URL?: string
    }
}

class GPTServer {
    private __token: string
    private __config?: GPTServerConfig

    constructor(token: string, config?: GPTServerConfig) {
        this.__token = token
        this.__config = config
    }

    /**
     * 查询
     * @param {string} query 内容
     * @param {number} chatId 会话 ID (利用会话 ID 连续对话), 默认 0 创建新会话
     * @param {string} remark 用户标记 用于应用标识用户
     * @return {any}
     * @link https://github.com/PBK-B/gpt-zmide-server/tree/master/docs#%E6%9F%A5%E8%AF%A2
     */
    public async Query(query: string, chatId?: number, remark?: string): Promise<any> {
        const url = this.__config?.server?.URL || DEFAULT_API_SERVER

        const data = new FormData();
        data.append("content", query);
        if (chatId) {
            data.append("chat_id", `${chatId}`);
        }
        if (remark) {
            data.append("remark", `${remark}`);
        }

        let config = {
            method: 'post',
            maxBodyLength: Infinity,
            url: `${url}/api/open/query`,
            headers: {
                "Authorization": `Bearer ${this.__token}`
            },
            data: data
        };

        return axios.request(config)
    }

}

export default GPTServer
export type { GPTServerConfig }
