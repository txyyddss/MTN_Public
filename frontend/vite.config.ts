import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'node:path'
import { fileURLToPath } from 'node:url'

const _dirname = typeof __dirname !== 'undefined'
    ? __dirname
    : path.dirname(fileURLToPath(import.meta.url))

export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd(), '')
    const backendUrl = env.VITE_BACKEND_URL || 'https://mtnapi.1391399.xyz'

    return {
        plugins: [vue()],
        resolve: {
            alias: {
                '@': path.resolve(_dirname, './src')
            }
        },
        server: {
            proxy: {
                '/api': {
                    target: backendUrl.endsWith('/api') ? backendUrl.slice(0, -4) : backendUrl,
                    changeOrigin: true
                }
            }
        }
    }
})
