/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/pages/admin/screens/home/index.tsx
 */
import React from 'react'
import { Statistic, Grid, Card, Badge } from '@arco-design/web-react';
import { IconArrowRise } from '@arco-design/web-react/icon';
import { axios } from '@/apis';
import useAxios from 'axios-hooks';

const Row = Grid.Row;
const Col = Grid.Col;

export default function index() {

    const [{ data: infoData, loading: infoLoading, error: infoError }] = useAxios({
        url: "/api/admin/config/system/info"
    })
    const [{ data: pingData, loading: pingLoading, error: pingError }] = useAxios({
        url: "/api/admin/config/ping/openai"
    })

    return (
        <div className='home_container' style={{ margin: '20px 0' }}>

            <Row gutter={24} style={{ alignItems: 'unset' }}>
                <Col span={4} >
                    <Card>
                        <Statistic
                            title='应用程序'
                            value={infoData?.data?.app_count || 0}
                            loading={infoLoading}
                        />
                    </Card>
                </Col>
                <Col span={4}>
                    <Card>
                        <Statistic
                            title='会话总计'
                            value={infoData?.data?.chat_count || 0}
                            loading={infoLoading}
                            prefix={<IconArrowRise />}
                            suffix='次'
                            countUp
                        />
                    </Card>
                </Col>
                <Col span={4}><Card>
                    <Statistic
                        title='接口调用次数'
                        value={infoData?.data?.use_api_count || 0}
                        loading={infoLoading}
                        prefix={<IconArrowRise />}
                        suffix='次'
                        countUp
                        styleValue={{ color: '#0fbf60' }}
                    />
                </Card>
                </Col>
                <Col span={4}>
                    <Card>
                        <Statistic
                            title='预计扣费'
                            value={infoData?.data?.estimated_cost || 0}
                            precision={2}
                            loading={infoLoading}
                            suffix='$'
                        />
                    </Card>
                </Col>
                <Col span={4} style={{ display: 'flex' }}>
                    <Card style={{ flex: 1 }}>
                        <p style={{ marginBottom: 8 }}>OpenAI 服务状态</p>
                        <Badge count={1} dot offset={[8, 8]} color={pingData?.data.status ? 'green' : 'red'}>
                            <span style={{ lineHeight: 1.9, fontSize: 22, fontWeight: 500, color: '#000' }}>
                                {pingData?.data.status ? '连接正常' : '连接失败'}
                            </span>
                        </Badge>
                    </Card>
                </Col>
            </Row>

        </div >
    )
}
