/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/pages/install/components/Header.tsx
 */
import React from 'react'
import { Divider, Layout } from '@arco-design/web-react';

import { LogoView } from '@/components';
import { axios } from '@/apis';

const LayoutHeader = Layout.Header;

interface HeaderProps {
}

export default function Header(props: HeaderProps) {
    return (
        <div className='app_header'>
            <LayoutHeader style={{
                height: 55,
                display: 'flex',
                flexDirection: 'row',
            }}>
                <LogoView />
            </LayoutHeader>
            <Divider style={{ margin: 0 }} />
        </div>
    )
}