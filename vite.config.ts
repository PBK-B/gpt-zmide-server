import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import {resolve} from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        react()
    ],
    publicDir: "static",
    build: {
        copyPublicDir: true,
        rollupOptions: {
            input: {
                index: resolve(__dirname, 'views/index.html'),
                admin: resolve(__dirname, 'views/admin.html'),
            },
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
