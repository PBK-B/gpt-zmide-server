/*
 * @Author: Bin
 * @Date: 2023-03-28
 * @FilePath: /gpt-zmide-server/src/pages/admin/screens/chat/index.tsx
 */
import React from 'react'

import {
    Table,
    TableColumnProps,
    Button
} from '@arco-design/web-react'
import useAxios from 'axios-hooks';
import { axios } from '@/apis';

export default function index() {

    const [{ data, error, loading }, refresh] = useAxios({
        url: "/api/admin/chat/"
    })

    const columns: TableColumnProps[] = [
        {
            title: 'ID',
            dataIndex: 'id',
        },
        {
            title: '所属应用',
            dataIndex: 'app.name',
        },
        {
            title: '消息数',
            dataIndex: 'massages_count',
        },
        {
            title: '备注',
            dataIndex: 'remark',
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
                        disabled={true}
                    >
                        消息记录
                    </Button>
                </> : undefined
            }
        },
    ];

    return (
        <div>
            <div style={{ display: 'flex', flexDirection: 'row' }}>
                <div style={{ flex: 1 }}></div>
            </div>

            <Table
                loading={loading}
                columns={columns}
                style={{ margin: "20px 0" }}
                data={data?.data?.list}
                pagination={{
                    current: data?.data?.page_index || 1,
                    pageSize: data?.data?.page_limit || 10,
                    total: data?.data?.page_total * data?.data?.page_limit,
                    onChange(pageNumber, pageSize) {
                        refresh({
                            params: {
                                page_limit: pageSize,
                                page_index: pageNumber
                            }
                        })
                    },
                }}
            />
        </div>
    )
}
