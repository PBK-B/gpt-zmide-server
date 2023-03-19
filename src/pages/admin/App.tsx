import React from 'react'
import { useLocation, useNavigate } from 'react-router-dom';
import { Layout, Menu } from '@arco-design/web-react';
import { IconGithub } from '@arco-design/web-react/icon';

import Routers from './routers'
import { Header } from './components';

const { Sider, Footer, Content } = Layout;
const MenuItem = Menu.Item;

const menus = [
    {
        name: "系统状态",
        router: "/"
    },
    {
        name: "应用管理",
        router: "/app"
    },
    {
        name: "会话查询",
        router: "/chat"
    },
    {
        name: "系统设置",
        router: "/system"
    },
]

function App() {
    const navigate = useNavigate();
    const location = useLocation();

    return (
        <Layout className='layout-collapse'>
            <Sider
                breakpoint='xl'
            >
                <div>
                    <h3 style={{ padding: 15 }}>ChatGPT API 网关</h3>
                </div>
                <Menu
                    selectedKeys={[location.pathname]}
                    onClickMenuItem={(key) => {
                        navigate(key, {
                            replace: true,
                        });
                    }}
                    style={{ flex: 1, width: '100%' }}
                >
                    {menus.map((item) => (
                        <MenuItem key={item.router}>
                            {item.name}
                        </MenuItem>
                    ))}
                </Menu>
            </Sider>
            <Layout>
                <Header></Header>
                <Layout>
                    <Content style={{ padding: '0 24px' }}>
                        <Routers />
                    </Content>
                    <Footer style={{ display: 'flex', background: '#00000006', padding: '15px 24px' }}>
                        <p style={{ flex: 1 }}>Copyright © {new Date().getFullYear()} zmide studio All rights reserved.</p>
                        <a style={{
                            display: 'flex', textAlign: 'center', justifyContent: 'center', alignItems: 'center', color: '#333', textDecoration: 'none'
                        }}
                            href="https://github.com/PBK-B/gpt-zmide-server"
                            target="_blank"
                        >
                            <IconGithub fontSize={16} /><span style={{ fontWeight: 300, fontSize: 13, marginLeft: 6 }}>开源地址</span>
                        </a>
                    </Footer>
                </Layout>
            </Layout>
        </Layout >
    )
}

export default App
