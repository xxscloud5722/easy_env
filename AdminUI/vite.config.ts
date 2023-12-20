import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react-swc';
import eslintPlugin from 'vite-plugin-eslint';
import viteCompression from 'vite-plugin-compression';
import { resolve } from 'path';

const config = loadEnv(process.argv[2] === 'build' ? 'production' : '*', process.cwd());

// https://vitejs.dev/config/
export default defineConfig({
  base: config.VITE_BASE,
  server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8080',
        rewrite: (path: { replace: (arg0: RegExp, arg1: string) => any; }) => path.replace(/^\/api/, ''),
        changeOrigin: true
      }
    }
  },
  build: {
    assetsDir: './assets',
    terserOptions: {
      compress: {
        drop_console: true,
        drop_debugger: true
      }
    },
    rollupOptions: {
      output: {
        entryFileNames: 'assets/[hash].js',
        chunkFileNames: 'assets/[hash].js',
        assetFileNames: 'assets/[hash].[ext]',
        globals: {
          global: 'window'
        }
      }
    }
  },
  plugins: [
    react(),
    eslintPlugin(),
    viteCompression()
  ],
  resolve: {
    alias: {
      '@': resolve(process.cwd(), './src'),
      '@core': resolve(process.cwd(), './src/core'),
      '@components': resolve(process.cwd(), './src/components'),
      '@assets': resolve(process.cwd(), './src/assets'),
      '@data': resolve(process.cwd(), './src/data.ts')
    },
    extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json']
  },
  css: {
    preprocessorOptions: {
      less: {
        javascriptEnabled: true
      }
    }
  }
});
