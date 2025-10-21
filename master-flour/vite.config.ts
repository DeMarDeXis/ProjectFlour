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
          target: env.VITE_API_URL, // para dev/prod
          changeOrigin: true,
        },
        '/ws': {
          target: env.VITE_WS_URL?.replace(/^ws:/, 'http:')?.replace(/^wss:/, 'https:') || 'http://localhost:8081',
          ws: true,
          changeOrigin: true
        }
      },
      // proxy:{ //DEV Divide
      //   '/api': {
      //     target: 'http://flour_backend:8080',
      //     changeOrigin: true,
      //     rewrite: (path) => path.replace(/^\/api/, ''),
      //   },
      //   '/ws': {
      //     target: 'http://flour_backend:8081',
      //     ws: true,
      //     changeOrigin: true
      //   }
      // },
    }
  }
})