import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'node:path'
import { fileURLToPath } from 'node:url'

const _dirname = typeof __dirname !== 'undefined'
    ? __dirname
    : path.dirname(fileURLToPath(import.meta.url))

export default defineConfig({
    plugins: [vue()],
    resolve: {
        alias: {
            '@': path.resolve(_dirname, './src')
        }
    },
    server: {
        proxy: {
            '/api': 'http://localhost:8080'
        }
    }
})
