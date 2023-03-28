import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

import { resolve } from 'path'
import fs from 'fs'

const VIEWS_TMP_DIR = "views/" // html 模版文件路径

function getAllInputOption(dirPath: string): any {
    const files = fs.readdirSync(dirPath)
    const options = {}
    for (const key in files) {
        const item = files[key]
        if (!item) {
            continue
        }
        options[item] = resolve(__dirname, dirPath + item)
    }
    return options
}

const allInputOption = getAllInputOption(VIEWS_TMP_DIR)

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        react()
    ],
    publicDir: "static",
    build: {
        chunkSizeWarningLimit: 2048,
        copyPublicDir: true,
        rollupOptions: {
            input: {
                ...allInputOption,
            },
        }
    },
    resolve: {
        alias: {
            "@": resolve(__dirname, "src")
        }
    },
    server: {
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:8091',
                changeOrigin: true,
            },
        }
    }
})
