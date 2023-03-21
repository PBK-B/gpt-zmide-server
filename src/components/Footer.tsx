/*
 * @Author: Bin
 * @Date: 2023-03-21
 * @FilePath: /gpt-zmide-server/src/components/Footer.tsx
 */
import React from 'react'
import { Layout } from '@arco-design/web-react';
import { IconGithub } from '@arco-design/web-react/icon';

export default function Footer() {
    return (
        <Layout.Footer style={{ display: 'flex', background: '#00000006', padding: '15px 24px' }}>
            <p style={{ flex: 1 }}>Copyright © {new Date().getFullYear()} zmide studio All rights reserved.</p>
            <a style={{
                display: 'flex', textAlign: 'center', justifyContent: 'center', alignItems: 'center', color: '#333', textDecoration: 'none'
            }}
                href="https://github.com/PBK-B/gpt-zmide-server"
                target="_blank"
            >
                <IconGithub fontSize={16} /><span style={{ fontWeight: 300, fontSize: 13, marginLeft: 6 }}>开源地址</span>
            </a>
        </Layout.Footer>
    )
}
