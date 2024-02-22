import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from "path"
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"), // @当作根目录
    },
  },
  // css预处理器
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@import "@/styles/constant.scss";`, // 全局使用的scss变量
      },
    },
  }, 
})
