import {defineConfig, loadEnv} from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')

  return {
    plugins: [react()],
    build: {
      outDir: 'dist',
    },
    server:{
      host: true,
      port: 5173,
      strictPort: true,
      watch: {
        usePolling: true,
        interval: 1000
      },
      proxy:{
        '/api': {
          // target: 'http://backend_dev:8080', // para local
          // target: process.env.VITE_API_URL, // para dev/prod
          target: env.VITE_API_URL, // para dev/prod
          // target: process.env.VITE_API_URL || 'http://hoooooosttttt:8080',
          changeOrigin: true,
        },
        '/ws': {
          target: env.VITE_WS_URL?.replace(/^ws:/, 'http:')?.replace(/^wss:/, 'https:') || 'http://localhost:8081',
          ws: true,
          changeOrigin: true
        }
      },
    }
  }
})