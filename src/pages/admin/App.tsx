import React from 'react'
import { useLocation, useNavigate } from 'react-router-dom';
import { Layout, Menu } from '@arco-design/web-react';

import { Footer, LogoView } from '@/components'
import Routers from './routers'
import { Header } from './components';

const { Sider, Content } = Layout;
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
                <LogoView />
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
                <Header />
                <Layout>
                    <Content style={{ padding: '0 24px' }}>
                        <Routers />
                    </Content>
                    <Footer />
                </Layout>
            </Layout>
        </Layout >
    )
}

export default App
