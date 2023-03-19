/*
 * @Author: Bin
 * @Date: 2023-03-19
 * @FilePath: /gpt-zmide-server/src/pages/admin/screens/empty/index.tsx
 */
import React from 'react'
import { Button, Empty } from '@arco-design/web-react'
import { IconExclamation } from '@arco-design/web-react/icon'

export default function index() {
    return (
        <div style={{ minHeight: 350, display: 'flex', flexDirection: 'column', justifyContent: 'center' }}>
            <Empty
                icon={
                    <div
                        style={{
                            background: '#f2994b',
                            display: 'inline-flex',
                            borderRadius: '50%',
                            width: 50,
                            height: 50,
                            marginBottom: 20,
                            fontSize: 30,
                            alignItems: 'center',
                            color: 'white',
                            justifyContent: 'center',
                        }}
                    >
                        <IconExclamation />
                    </div>
                }
                description='大师兄，这个页面被师傅抓走啦…'
            />
            <div style={{ display: 'flex', justifyContent: 'center', marginTop: 10 }}>
                <Button type="primary" onClick={() => {
                    window.location.reload()
                }}>刷新一下</Button>
            </div>
        </div>
    )
}
