/*
 * @Author: Wzq
 * @Date: 2023-03-27
 * @FilePath: /gpt-zmide-server/src/pages/admin/screens/system/index.tsx
 */
import React, { ReactElement, useEffect } from 'react'
import { Tabs, Form, Input, Button, Spin, Message } from '@arco-design/web-react';
import useAxios from 'axios-hooks';
import { axios } from '@/apis';

const TabPane = Tabs.TabPane;
const FormItem = Form.Item;

export default function index() {

    const [{ data: configData, loading: configLoading, error: infoError }] = useAxios({
        url: "/api/admin/config/system/config"
    })

    const [tabsIndex, setTabsIndex] = React.useState(1)
    const [tabsFromData, setTabsFromData] = React.useState<any[]>([])

    type FormViewProps = {
        fromData?: any,
        setFromData?: (data: any) => void,
    }

    // 获取表单数据
    const tabsFromDataConfig = React.useMemo<any>(() => {
        return tabsFromData[tabsIndex]
    }, [tabsFromData, tabsIndex])

    // 设置表单数据
    const setTabsFromDataConfig = React.useCallback(
        (config: any) => {
            const index = tabsIndex
            const datas = tabsFromData

            datas[index] = {
                ...(tabsFromData[index]),
                ...config
            }
            setTabsFromData([
                ...datas,
            ])
        },
        [tabsIndex],
    )

    const tabsView: { label: string, key: string, view: ((props: FormViewProps) => JSX.Element) | ReactElement<any, any> }[] = [
        {
            label: '站点信息配置',
            key: 'site_config',
            view: (props: FormViewProps) => {
                const { fromData, setFromData = () => { } } = props;
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
                ]

                return <Form autoComplete='off' layout="vertical" style={{ width: 600 }}>
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
            label: 'OpenAI配置',
            key: 'openai_config',
            view: (props: FormViewProps) => {
                const { fromData, setFromData = () => { } } = props;
                const forms: { label: string, field: string, example?: string, required?: boolean }[] = [
                    {
                        label: 'OpenAI Secret Key',
                        field: 'openai_secret_key',
                        required: true
                    },
                    {
                        label: 'OpenAI Model',
                        field: 'openai_model',
                        required: true
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

                return <Form autoComplete='off' layout="vertical" style={{ width: 600 }}>
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
    ]

    const selectTabsView = React.useMemo(() => tabsView[tabsIndex].view, [tabsIndex])

    useEffect(() => {
        if (configData) {
            const datas = tabsFromData
            tabsView.forEach((tab, i) => {
                datas[i] = {
                    ...configData.data[tab.key]
                }
            })
            setTabsFromData([
                ...datas,
            ])

            setTabsIndex(0)
        }
    }, [configData]);

    const sendConfig = (name: "site" | "openai", data: string) => {
        if (requestConfig.loading) {
            return
        }
        if (!name || !data) {
            Message.warning("请正确填写数据")
            return
        }

        setRequestConfig({
            ...requestConfig,
            loading: true,
        })

        const formData = new FormData();
        formData.append("name", name);
        formData.append("data", data);
        axios.post("/api/admin/config/system/config", formData).then((response) => {
            const { code, msg, data } = response.data
            setRequestConfig({
                ...requestConfig,
                loading: false,
            })
            if (code !== 200) {
                Message.info(msg || `请求失败，${code || '稍后重试'}`)
                return
            }

            Message.success('保存成功')
            window.location.reload()
        }).catch((error) => {
            Message.error(`请求失败，请稍后重试 ${error}`);
            setRequestConfig({
                ...requestConfig,
                loading: false,
            })
        })
    }


    // 提交配置数据
    const [requestConfig, setRequestConfig] = React.useState({
        loading: false
    })

    // 提交数据事件
    const onClickSubmitConfig = React.useCallback(
        (index: number, tabsFromData: Array<any>) => {
            if (tabsFromData.length < index) {
                Message.warning("请填写数据")
                return
            }

            const fromData = tabsFromData[index]

            if (!fromData || Object.keys(fromData).length < 1) {
                Message.clear()
                Message.warning("好像还没填写数据哎？")
                return
            }

            setRequestConfig({
                ...requestConfig,
                loading: true,
            })

            const dataStr = JSON.stringify(fromData)
            switch (index) {
                case 0:
                    sendConfig("site", dataStr)
                    break;
                case 1:
                    sendConfig("openai", dataStr)
                    break;
            }
        },
        [tabsIndex, tabsFromData],
    )

    return (
        <div className='system_container' style={{ margin: '20px 0' }}>
            <Spin dot loading={configLoading}>
                {!configLoading ? (
                    <Tabs defaultActiveTab={tabsView[0].key} tabPosition='left' onChange={(key) => {
                        setTabsIndex(tabsView.findIndex((item) => item.key == key))
                    }}>
                        {tabsView.map((item, index) =>
                            <TabPane key={item.key} title={item.label}>
                                {(typeof selectTabsView === "function" ? selectTabsView({
                                    fromData: { ...tabsFromDataConfig },
                                    setFromData: setTabsFromDataConfig,
                                }) : selectTabsView)}
                                <Button
                                    type="primary"
                                    onClick={() => onClickSubmitConfig(tabsIndex, tabsFromData)}
                                    loading={requestConfig.loading}
                                    disabled={requestConfig.loading}
                                >
                                    保存
                                </Button>
                            </TabPane>
                        )}
                    </Tabs>
                ) : ""}
            </Spin>
        </div >
    )
}
