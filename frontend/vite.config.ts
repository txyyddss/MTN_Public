import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'node:path'
import { fileURLToPath } from 'node:url'

import viteCompression from 'vite-plugin-compression'
import { ViteImageOptimizer } from 'vite-plugin-image-optimizer'

const _dirname = typeof __dirname !== 'undefined'
    ? __dirname
    : path.dirname(fileURLToPath(import.meta.url))

export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd(), '')
    const backendUrl = env.VITE_BACKEND_URL

    return {
        plugins: [
            vue(),
            viteCompression(),
            ViteImageOptimizer({
                webp: {
                    quality: 80,
                },
                png: {
                    quality: 80,
                },
                jpeg: {
                    quality: 80,
                },
                jpg: {
                    quality: 80,
                },
                tiff: {
                    quality: 80,
                },
                gif: {},
                svg: {
                    multipass: true,
                },
            }),
        ],
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
