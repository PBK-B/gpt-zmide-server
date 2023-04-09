/*
 * @Author: Bin
 * @Date: 2023-04-09
 * @FilePath: /gpt-zmide-server/src/pages/admin/screens/system/Logs.tsx
 */
import React, { useRef, useEffect } from 'react'

import { MonacoLanguageLog } from '@/pages/admin/utils';
import Editor, { Monaco, OnMount } from '@monaco-editor/react';
import { editor as MonacoEditor } from 'monaco-editor';
import useAxios from 'axios-hooks';

export default function Logs(props: any) {
    const [{ data, error, loading }, refresh] = useAxios("/api/admin/config/system/log")

    const refreshRuntime = useRef<any>(null)
    useEffect(() => {
        if (refreshRuntime.current === null) {
            refreshRuntime.current = setInterval(() => {
                refresh()
            }, 2000)
        }
        return () => {
            clearInterval(refreshRuntime.current)
            refreshRuntime.current = undefined
        }
    }, [])

    const monacoEditor = useRef<MonacoEditor.IStandaloneCodeEditor>()
    const handleEditorWillMount = (monaco: Monaco) => {
        MonacoLanguageLog(monaco)
    }

    const handleEditorDidMount: OnMount = (editor, monaco) => {
        monacoEditor.current = editor
        listenScrollRefreshEvent(editor)
    }

    // 滚动刷新
    const enableScrollingRefresh = useRef(true)
    const autoScrollRefresh = useRef(false)
    const scrollToRefresh = () => {
        if (!monacoEditor.current || !enableScrollingRefresh) { return }
        const editor = monacoEditor.current
        const lineCount = editor.getModel()?.getLineCount() || 0
        autoScrollRefresh.current = true
        editor.revealLine(lineCount)
    }
    const listenScrollRefreshEvent = (editor: MonacoEditor.IStandaloneCodeEditor) => {
        if (!editor) { return }
        editor.onDidScrollChange((e) => {
            const eViewHeight = editor.getDomNode()?.clientHeight || 0
            if (e.scrollTop > 0) {
                if (!autoScrollRefresh.current) {
                    if (eViewHeight + e.scrollTop === e.scrollHeight) {
                        enableScrollingRefresh.current = true
                    } else {
                        enableScrollingRefresh.current = false
                    }
                }
                autoScrollRefresh.current = false
                // console.log("enableScrollingRefresh", enableScrollingRefresh.current);
            }
        })
    }

    useEffect(() => {
        if (data?.data && enableScrollingRefresh.current) {
            scrollToRefresh()
        }
    }, [data, enableScrollingRefresh.current])

    return <Editor
        width="70vw"
        height="80vh"
        theme="logview"
        defaultLanguage="log"
        value={data?.data || ''}
        options={{
            scrollBeyondLastLine: false,
            automaticLayout: true,
            readOnly: true
        }}
        beforeMount={handleEditorWillMount}
        onMount={handleEditorDidMount}
    />
}
