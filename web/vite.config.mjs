import {defineConfig} from 'vite';
import vue from '@vitejs/plugin-vue';
import WindiCSS from 'vite-plugin-windicss';

import {resolve} from 'path';

const pathResolve = (dir) => {
  return resolve(__dirname, ".", dir)
}

const alias = {
  '~': pathResolve("src")
}

// https://vitejs.dev/config/
export default defineConfig(() => {
  return {
    resolve: {
      alias,
    },
    server: {
      proxy: {
        '/api/v1': {
          target: 'http://127.0.0.1:8090',
          changeOrigin: true,
        },
      },
    },
    plugins: [vue(), WindiCSS()],
  }

});
