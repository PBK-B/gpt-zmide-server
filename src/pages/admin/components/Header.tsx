/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/pages/admin/components/Header.tsx
 */
import React from 'react'
import { Avatar, Divider, Layout, Menu, Popover, Modal, Input, Form, Message } from '@arco-design/web-react';
import { IconUser } from '@arco-design/web-react/icon';
import { axios } from '@/apis';

const LayoutHeader = Layout.Header;

interface HeaderProps {
}

type updatePasswordConfigType = {
    visible: boolean,
    oldPassword?: string,
    newPassword?: string
}

export default function Header(props: HeaderProps) {

    const [updatePasswordConfig, setUpdatePasswordConfig] = React.useState<updatePasswordConfigType>({
        visible: false,
    })

    // 修改管理员密码
    const updatePassword = (oldPassword: string, newPassword: string) => {
        if (!oldPassword || !newPassword) {
            Message.warning('密码不得为空。')
            return
        }

        if (oldPassword === newPassword) {
            Message.warning('新密码不得和旧密码相同。')
            return
        }

        const formData = new FormData()
        formData.append('old_password', oldPassword)
        formData.append('new_password', newPassword)

        axios.post("/api/admin/config/update/password", formData)
            .then(response => {
                const { code, msg } = response.data
                if (code !== 200) {
                    Message.error(msg || '修改失败，请稍后重试')
                    return
                }

                setUpdatePasswordConfig({
                    ...updatePasswordConfig,
                    visible: false,
                    oldPassword: undefined,
                    newPassword: undefined
                })

                Message.success({
                    content: '密码修改成功',
                    onClose: () => {
                        window.location.replace("/admin")
                    }
                })

            })
            .catch((err) => {
                Message.error(err.message || '服务器响应失败，请稍后重试')
            })

    }

    return (
        <div className='app_header'>
            <LayoutHeader style={{
                height: 55,
                display: 'flex',
                flexDirection: 'row',
            }}>
                <div style={{ flex: 1 }} />
                <div style={{ height: '100%', marginRight: 15, display: 'flex', alignItems: 'center' }}>
                    <Popover
                        position="br"
                        trigger="click"
                        className="app_header_popover"
                        style={{ padding: 0 }}
                        content={
                            <Menu>
                                <Menu.Item key='1' onClick={() => {
                                    setUpdatePasswordConfig({
                                        ...updatePasswordConfig,
                                        visible: true,
                                        oldPassword: undefined,
                                        newPassword: undefined
                                    })
                                }} >修改密码</Menu.Item>
                                <Menu.Item key='2' onClick={() => {
                                    window.location.replace("/admin/signout")
                                }}>退出登录</Menu.Item>
                            </Menu>
                        }
                    > <Avatar size={35} style={{ backgroundColor: '#3370ff' }}>
                            <IconUser />
                        </Avatar>

                    </Popover>
                </div>
            </LayoutHeader>
            <Divider style={{ margin: 0 }} />

            <Modal
                title='修改密码'
                visible={updatePasswordConfig.visible}
                onOk={() => {
                    const { oldPassword = '', newPassword = '' } = updatePasswordConfig
                    updatePassword(oldPassword, newPassword)
                }}
                onCancel={() => setUpdatePasswordConfig({
                    ...updatePasswordConfig,
                    visible: false,
                })}
                autoFocus={false}
                focusLock={true}
                okText='修改'
            >
                {updatePasswordConfig.visible && (
                    <Form autoComplete='off' layout="vertical" >
                        <Form.Item label='旧密码'>
                            <Input defaultValue={updatePasswordConfig.oldPassword} onChange={(value) => {
                                setUpdatePasswordConfig({
                                    ...updatePasswordConfig,
                                    oldPassword: value,
                                })
                            }} placeholder='请输入旧密码' />
                        </Form.Item>
                        <Form.Item label='新密码'>
                            <Input defaultValue={updatePasswordConfig.newPassword} onChange={(value) => {
                                setUpdatePasswordConfig({
                                    ...updatePasswordConfig,
                                    newPassword: value,
                                })
                            }} placeholder='请输入新密码' />
                        </Form.Item>
                    </Form>
                )}
            </Modal>
        </div>

    )
}