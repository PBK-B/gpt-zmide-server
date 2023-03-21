/*
 * @Author: Bin
 * @Date: 2023-03-21
 * @FilePath: /gpt-zmide-server/src/pages/install/main.tsx
 */
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import "@arco-design/web-react/dist/css/arco.css";
import './scss/index.scss'
import 'vite/modulepreload-polyfill'
import { HashRouter } from 'react-router-dom';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <HashRouter>
        <App />
    </HashRouter>
)
