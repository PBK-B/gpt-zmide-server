<?php
/*
 * @Author: Bin
 * @Date: 2023-03-11
 * @FilePath: /gpt-zmide-server/docs/examples/php/src/example.php
 */

namespace src;

use Exception;
use \GuzzleHttp\Client;
use \GuzzleHttp\Psr7\Request;

class ZmideChatgpt
{
    private static $host  = "https://example.zmide.com";
    private static $token = "GDAiQmNHhoHdqoFqxrGObTSObYSySast";

    /**
     * @description: 查询 Chatgpt 消息
     * @param {string} $content 查询消息内容
     * @param {int} $chat_id 会话 ID 可为空，为空时表示创建新会话
     * @return {*}
     */
    public static function query(string $content, int $chat_id = 0)
    {
        if ($content == null || $content == "") {
            throw new Exception("查询内容不得为空。");
        }

        $multipart = [
            [
                'name'     => 'content',
                'contents' => $content,
            ],
        ];
        if ($chat_id != null && $chat_id != 0) {
            array_push($multipart, [
                'name'     => 'chat_id',
                'contents' => $chat_id,
            ]);
        }

        $options = [
            'multipart'   => $multipart,
            'synchronous' => true,
            'verify'      => false,
        ];
        $headers = [
            'Authorization' => 'Bearer ' . ZmideChatgpt::$token,
        ];
        $client = new Client([
            'base_uri' => ZmideChatgpt::$host,
        ]);
        $request = new Request('POST', '/api/open/query', $headers);
        $res     = $client->sendAsync($request, $options)->wait();
        $body    = $res->getBody();
        $resObj  = json_decode($body->__toString(), true);

        // print_r($resObj);
        // print_r(PHP_EOL);

        return $resObj;
    }
}

require 'vendor/autoload.php';

try {
    $callback = ZmideChatgpt::query("计算机是什么？");
    $code     = $callback['code']; // 是否请求成功 code 200 = 成功
    if ($code != 200) {
        throw new Exception("请求失败", $code, $callback);
    }
    $message         = $callback['data']; // 消息对象
    $message_content = $message['content']; // 消息的内容
    $message_chat_id = $message['chat_id']; // 消息的会话 ID 通过该 id 可以连续对话
    print_r($callback);
} catch (\Throwable$th) {
    //throw $th;
    print_r($th->getMessage());
}

print_r(PHP_EOL);
