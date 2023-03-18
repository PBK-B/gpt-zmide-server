/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/pages/admin/components/Header.tsx
 */
import React from 'react'
import { Avatar, Divider, Layout, Menu, Dropdown } from '@arco-design/web-react';
import { IconUser } from '@arco-design/web-react/icon';

const LayoutHeader = Layout.Header;

interface HeaderProps {
}

export default function Header(props: HeaderProps) {
    return (
        <div>
            <LayoutHeader style={{
                height: 55,
                display: 'flex',
                flexDirection: 'row',
            }}>
                <div style={{ flex: 1 }} />
                <div style={{ height: '100%', marginRight: 15, display: 'flex', alignItems: 'center' }}>
                    <Dropdown
                        position="br"
                        droplist={
                            <Menu>
                                <Menu.Item key='1'>修改密码</Menu.Item>
                                <Menu.Item key='2'>退出登录</Menu.Item>
                            </Menu>
                        }
                    >
                        <Avatar size={35} style={{ backgroundColor: '#3370ff' }}>
                            <IconUser />
                        </Avatar>
                    </Dropdown>
                </div>
            </LayoutHeader>
            <Divider style={{ margin: 0 }} />
        </div>

    )
}