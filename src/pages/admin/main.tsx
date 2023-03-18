/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/pages/admin/main.tsx
 */
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './scss/index.scss'
import 'vite/modulepreload-polyfill'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <App />
    </React.StrictMode>,
)
