/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/pages/admin/screens/application/index.tsx
 */
import React from 'react'
import { Breadcrumb, Button, Input, Message, Modal, Result, Statistic, Table, TableColumnProps, Tag, Form } from '@arco-design/web-react'
import useAxios from 'axios-hooks';
import { axios } from '@/apis';

type createAppConfigType = {
    visible: boolean,
    id?: number,
    name?: string
}

export default function index() {

    const [{ data, error, loading }, refresh] = useAxios({
        url: "/api/admin/application/"
    })

    const columns: TableColumnProps[] = [
        {
            title: '应用名称',
            dataIndex: 'name',
        },
        {
            title: '密钥',
            dataIndex: 'app_key',
        },
        {
            title: '状态',
            dataIndex: 'status',
            render: (status) => {
                return status === 1 ? <Tag color='green'>已启用</Tag> : <Tag color='red'>已禁用</Tag>
            }
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
        },
        {
            title: '操作',
            dataIndex: 'id',
            align: 'center',
            render: (id, item) => {
                return item ? <>
                    <Button
                        type='text'
                        disabled={!item?.id}
                        onClick={() => {
                            updateAppStatus(item.id, item.status)
                        }}
                    >
                        {item?.status === 1 ? '禁用' : '启用'}
                    </Button>
                    <Button
                        type='text'
                        disabled={!item?.id}
                        onClick={() => {
                            setCreateAppConfig({
                                ...createAppConfig,
                                visible: true,
                                id: item?.id,
                                name: item?.name,
                            })
                        }}
                    >
                        修改
                    </Button>
                </> : undefined
            }
        },
    ];

    const [createAppConfig, setCreateAppConfig] = React.useState<createAppConfigType>({
        visible: false,
        id: undefined,
        name: undefined,
    })

    // 创建应用
    const createApp = (name: string) => {
        if (!name || name == "") {
            Message.warning('应用名不得为空。')
            return
        }

        const formData = new FormData();
        formData.append("name", name)

        axios.post("/api/admin/application/create", formData).then((response) => {
            const { code, msg, data } = response.data
            if (code !== 200) {
                Message.info(`请求失败，${msg || code}`)
                return
            }
            // 成功
            refresh() // 刷新数据 
            setCreateAppConfig({
                ...createAppConfig,
                visible: false,
                id: undefined,
                name: undefined,
            }) // 关闭弹窗
        }).catch(err => {
            Message.info(`请求失败，${err.message || '请稍后重试'}`)
        })
    }

    // 更新应用
    const updateApp = (id: number, name: string) => {
        if (!name || name == "") {
            Message.warning('应用名不得为空。')
            return
        }

        const formData = new FormData();
        formData.append("name", name)

        axios.post(`/api/admin/application/${id}/update`, formData).then((response) => {
            const { code, msg, data } = response.data
            if (code !== 200) {
                Message.info(`请求失败，${msg || code}`)
                return
            }
            // 成功
            refresh() // 刷新数据 
            setCreateAppConfig({
                ...createAppConfig,
                visible: false,
                id: undefined,
                name: undefined,
            }) // 关闭弹窗
        }).catch(err => {
            Message.info(`请求失败，${err.message || '请稍后重试'}`)
        })
    }

    // 更新应用状态
    const updateAppStatus = (id: number, status: number) => {
        if (!id || id < 1) {
            Message.warning('应用异常。')
            return
        }

        const formData = new FormData();
        formData.append("status", status === 1 ? '2' : '1')

        axios.post(`/api/admin/application/${id}/update`, formData).then((response) => {
            const { code, msg, data } = response.data
            if (code !== 200) {
                Message.info(`请求失败，${msg || code}`)
                return
            }
            // 成功
            refresh() // 刷新数据 
            Message.success(`${status === 1 ? '禁用' : '启用'}成功`)
        }).catch(err => {
            Message.info(`请求失败，${err.message || '请稍后重试'}`)
        })
    }

    return (
        <div style={{ marginTop: 20 }}>
            <div style={{ display: 'flex', flexDirection: 'row' }}>
                <div style={{ flex: 1 }}></div>
                <Button
                    style={{ marginBottom: 10, }}
                    type='primary'
                    onClick={() => setCreateAppConfig({
                        ...createAppConfig,
                        visible: true,
                        id: undefined,
                        name: undefined,
                    })}
                >
                    创建应用
                </Button>
            </div>
            {loading || (!error && data?.data) ? (
                <Table rowKey={(item) => item.id} loading={loading} columns={columns} data={data?.data} />
            ) : (
                <Result
                    status='warning'
                    title={error ? `出错啦，${error.message}` : '应用列表为空，请先创建一个应用。'}
                    extra={<Button type='primary' onClick={refresh}>刷新</Button>}
                />
            )}

            <Modal
                title={createAppConfig.id ? '修改应用' : '创建新应用'}
                visible={createAppConfig.visible}
                onOk={() => {
                    // console.log("创建应用", createAppConfig?.name);
                    if (createAppConfig?.id) {
                        // 修改应用
                        updateApp(createAppConfig.id, createAppConfig.name || '')
                    } else {
                        // 创建应用
                        createApp(createAppConfig?.name || '')
                    }
                }}
                onCancel={() => setCreateAppConfig({
                    ...createAppConfig,
                    visible: false,
                })}
                autoFocus={false}
                focusLock={true}
                okText={createAppConfig.id ? '修改' : '创建'}
            >
                {createAppConfig.visible && (
                    <Form autoComplete='off' layout="vertical" >
                        <Form.Item label='应用名称'>
                            <Input defaultValue={createAppConfig.name} onChange={(value) => {
                                setCreateAppConfig({
                                    ...createAppConfig,
                                    name: value,
                                })
                            }} placeholder='请输入应用名称' />
                        </Form.Item>
                    </Form>
                )}
            </Modal>
        </div>
    )
}
