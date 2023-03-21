/*
 * @Author: Bin
 * @Date: 2023-03-21
 * @FilePath: /gpt-zmide-server/src/pages/install/App.tsx
 */
import React, { ReactElement } from 'react'
import { Layout, Steps, Divider, Button, Form, Input, Result, Message } from '@arco-design/web-react';

import { Footer } from '@/components'
import { Header } from './components'
import { axios } from '@/apis';

type StepViewProps = {
    fromData?: any,
    setFromData?: (data: any) => void,
    onStepPrevious: () => void,
    onStepNext: () => void,
}

export default function App() {

    const [stepFromData, setStepFromData] = React.useState<any[]>([])
    const [stepIndex, setStepIndex] = React.useState(1)

    // 获取步进器表单数据
    const stepFromDataConfig = React.useMemo<any>(() => {
        return stepFromData[stepIndex - 1]
    }, [stepFromData, stepIndex])

    // 设置步进器表单数据
    const setStepFromDataConfig = React.useCallback(
        (config: any) => {
            const index = stepIndex - 1
            const datas = stepFromData

            datas[index] = {
                ...(stepFromData[index]),
                ...config
            }
            setStepFromData([
                ...datas,
            ])
        },
        [stepIndex],
    )

    // 提交配置数据
    const [requestConfig, setRequestConfig] = React.useState({
        loading: false
    })
    const sendConfig = (step: "site" | "openai" | "database", data: string) => {
        if (requestConfig.loading) {
            return
        }
        if (!step || !data) {
            Message.warning("请正确填写数据")
            return
        }

        setRequestConfig({
            ...requestConfig,
            loading: true,
        })

        const formData = new FormData();
        formData.append("step", step);
        formData.append("data", data);
        axios.post("/install/config", formData).then((response) => {
            const { code, msg, data } = response.data
            setRequestConfig({
                ...requestConfig,
                loading: false,
            })
            if (code !== 200) {
                Message.info(msg || `请求失败，${code || '稍后重试'}`)
                return
            }
            onStepNext() // 配置下一项
        }).catch((error) => {
            Message.error(`请求失败，请稍后重试 ${error}`);
            setRequestConfig({
                ...requestConfig,
                loading: false,
            })
        })
    }

    const views: { title: string, view: ((props: StepViewProps) => JSX.Element) | ReactElement<any, any>, }[] = [
        {
            title: "设置站点信息",
            view: (props: StepViewProps) => {
                const { fromData = {}, setFromData = () => { }, onStepPrevious = () => { }, onStepNext = () => { } } = props;
                const forms: { label: string, field: string, example?: string }[] = [
                    {
                        label: '站点名称',
                        field: 'site_name',
                        example: '天真的 ChatGPT 服务'
                    },
                    {
                        label: '站点域名',
                        field: 'domain_name',
                        example: 'https://example.zmide.com'
                    },
                    {
                        label: '应用端口',
                        field: 'port',
                        example: '80'
                    },
                    {
                        label: '管理员用户名',
                        field: 'admin_user'
                    },
                    {
                        label: '管理员密码',
                        field: 'admin_password'
                    }
                ]

                return <Form autoComplete='off' layout="vertical" >
                    {
                        forms.map((item, index) =>
                            <Form.Item key={`${item.field}_form_item_${index}`} label={item.label}>
                                <Input
                                    defaultValue={fromData[item.field]}
                                    placeholder={`请输入你的${item.label}${item.example ? ' ，例如: ' + item.example : ''}`}
                                    onChange={(value) => {
                                        const data = {
                                            ...fromData,
                                        }
                                        data[item.field] = value
                                        setFromData(data)
                                    }} />
                            </Form.Item>
                        )
                    }
                </Form>
            }
        },
        {
            title: "配置 OpenAI",
            view: (props: StepViewProps) => {
                const { fromData, setFromData = () => { }, onStepPrevious = () => { }, onStepNext = () => { } } = props;
                const forms: { label: string, field: string, example?: string }[] = [
                    {
                        label: 'OpenAI Secret Key',
                        field: 'openai_secret_key'
                    },
                    {
                        label: 'HTTP 代理地址',
                        field: 'openai_proxy_host',
                        example: '192.168.1.1'
                    },
                    {
                        label: 'HTTP 代理端口',
                        field: 'openai_proxy_port',
                        example: '8081'
                    }
                ]
                return fromData && <Form autoComplete='off' layout="vertical" >
                    {
                        forms.map((item, index) =>
                            <Form.Item key={`${item.field}_form_item_${index}`} label={item.label}>
                                <Input
                                    defaultValue={fromData[item.field]}
                                    placeholder={`请输入你的${item.label}${item.example ? ' ，例如: ' + item.example : ''}`}
                                    onChange={(value) => {
                                        const data = {
                                            ...fromData,
                                        }
                                        data[item.field] = value
                                        setFromData(data)
                                    }} />
                            </Form.Item>
                        )
                    }
                </Form>
            }
        },
        {
            title: "配置数据库",
            view: (props: StepViewProps) => {
                const { fromData = {}, setFromData = () => { }, onStepPrevious = () => { }, onStepNext = () => { } } = props;
                const forms: { label: string, field: string, example?: string }[] = [
                    {
                        label: 'MySql 数据库地址',
                        field: 'mysql_host',
                        example: '127.0.0.1'
                    },
                    {
                        label: 'MySql 数据库端口',
                        field: 'mysql_port',
                        example: '3306'
                    },
                    {
                        label: 'MySql 数据库用户名',
                        field: 'mysql_user',
                        example: 'root'
                    },
                    {
                        label: 'MySql 数据库密码',
                        field: 'mysql_password',
                    },
                    {
                        label: 'MySql 数据库名称',
                        field: 'mysql_database',
                        example: 'gpt_zmide_server'
                    }
                ]

                return <Form autoComplete='off' layout="vertical" >
                    {
                        forms.map((item, index) =>
                            <Form.Item key={`${item.field}_form_item_${index}`} label={item.label}>
                                <Input
                                    defaultValue={fromData[item.field]}
                                    placeholder={`请输入你的${item.label}${item.example ? ' ，例如: ' + item.example : ''}`}
                                    onChange={(value) => {
                                        const data = {
                                            ...fromData,
                                        }
                                        data[item.field] = value
                                        setFromData(data)
                                    }} />
                            </Form.Item>
                        )
                    }
                </Form>
            }
        },
        {
            title: "应用部署完成",
            view: <Result
                status='success'
                title='你的应用已经完成配置啦!'
                subTitle='现在点击完成按钮立即跳转后台管理页面创建应用。'
            ></Result>
        }
    ];

    const selectStepView = React.useMemo(() => views[((stepIndex - 1))].view, [stepIndex])

    // 是否为最后一步
    const isLastStep = React.useMemo(() => stepIndex >= views.length, [stepIndex])
    // 是否为第一步
    const isFirstStep = React.useMemo(() => stepIndex <= 1, [stepIndex])

    // 返回上一步
    const onStepPrevious = React.useCallback(
        () => {
            if (isFirstStep) {
                return
            }
            setStepIndex(stepIndex - 1)
        },
        [stepIndex, isFirstStep],
    )

    // 继续下一步
    const onStepNext = React.useCallback(
        () => {
            if (isLastStep) {
                return
            }
            setStepIndex(stepIndex + 1)

        },
        [stepIndex, isLastStep],
    )

    // 提交数据事件
    const onClickSendEvent = React.useCallback(
        (index: number, stepFromData: Array<any>) => {
            if (isLastStep) {
                // 最后一步，应该是完成按钮(跳转管理页面)
                window.location.href = "/admin"
                return
            }

            if (stepFromData.length < index) {
                Message.warning("请填写数据")
                return
            }

            const fromData = stepFromData[index - 1]

            if (!fromData || Object.keys(fromData).length < 1) {
                Message.clear()
                Message.warning("好像还没填写数据哎？")
                return
            }

            for (const key in Object.values(fromData)) {
                if (!Object.values(fromData)[key]) {
                    Message.clear()
                    Message.warning("部分数据未填写完整")
                    return
                }
            }

            const dataStr = JSON.stringify(fromData)
            switch (index) {
                case 1:
                    sendConfig("site", dataStr)
                    break;
                case 2:
                    sendConfig("openai", dataStr)
                    break;
                case 3:
                    sendConfig("database", dataStr)
                    break;
            }
        },
        [stepIndex, stepFromData, isLastStep],
    )

    return (
        <div className='layout_install'>
            <Header />
            <Layout style={{ flex: 1 }}>
                <Layout.Content>
                    <Steps current={stepIndex} style={{ maxWidth: 780, margin: '25px auto' }}>
                        {views.map((item, index) =>
                            <Steps.Step key={index} title={item.title} />
                        )}
                    </Steps>
                    <Divider />
                    <div className='step_container'>
                        <div style={{ minHeight: 410, padding: '30px 0' }}>
                            {(typeof selectStepView === "function" ? selectStepView({
                                fromData: { ...stepFromDataConfig },
                                setFromData: setStepFromDataConfig,
                                onStepPrevious,
                                onStepNext
                            }) : selectStepView)}
                            {

                            }
                        </div>
                        <div style={{ display: 'flex', flexDirection: 'row', marginBottom: 20 }}>
                            <div style={{ flex: 1 }} />
                            {isFirstStep ? null : (
                                <Button type="outline" style={{ marginRight: 20 }} disabled={requestConfig.loading} onClick={onStepPrevious}>上一步</Button>
                            )}
                            <Button
                                type="primary"
                                onClick={() => onClickSendEvent(stepIndex, stepFromData)}
                                loading={requestConfig.loading}
                                disabled={requestConfig.loading}
                            >
                                {isLastStep ? '完成' : '下一步'}
                            </Button>
                        </div>
                    </div>
                </Layout.Content>
            </Layout>
            <Footer />
        </div>
    )
}
