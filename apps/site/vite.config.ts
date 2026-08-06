import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import ui from '@nuxt/ui/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
    ui({
      ui: {
        colors: {
          primary: 'zinc',
          secondary: 'zinc',
          success: 'emerald',
          info: 'sky',
          warning: 'amber',
          error: 'red',
          neutral: 'zinc',
        },
      },
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    // Vite в docker-контейнере: inotify не всегда видит изменения,
    // сделанные с хоста (напр., `task gen` переписывает src/types/api).
    watch: {
      usePolling: true,
      interval: 1000,
    },
    proxy: {
      // API ходит same-origin: CORS не нужен. Пути в OpenAPI уже с префиксом /api.
      // В dev-стеке (docker compose) цель переопределяется: API_PROXY_TARGET=http://api:8080.
      '/api': {
        target: process.env.API_PROXY_TARGET ?? 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
