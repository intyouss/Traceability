import vue from '@vitejs/plugin-vue'
import { viteMockServe } from 'vite-plugin-mock'
import {vitePluginSvg} from "@webxrd/vite-plugin-svg"
import { resolve } from 'path'
import {defineConfig, loadEnv} from "vite";

const pathResolve = (dir) => {
  return resolve(__dirname, ".", dir)
}

const alias = {
  '@': pathResolve("src")
}

export default defineConfig(({ mode, command }) => {
  const prodMock = true;
  const env = loadEnv(mode, process.cwd());
  const { VITE_BASE_ENDPOINT, VITE_BASE_PORT } = env;
  const BaseUrl = VITE_BASE_ENDPOINT + ':' + VITE_BASE_PORT;
  return {
    base: './',
    resolve: {
      alias
    },
    server: {
      proxy: {
        '/api/v1': {
          target: BaseUrl,
          changeOrigin: true,
        },
      },
    },
    build: {
      rollupOptions: {
        output: {
          manualChunks: {
            'echarts': ['echarts']
          }
        }
      }
    },
    plugins: [
      vue(),
      viteMockServe({
        mockPath: 'mock',
        localEnabled: command === 'serve',
        prodEnabled: command !== 'serve' && prodMock,
        watchFiles: true,
        injectCode: `
          import { setupProdMockServer } from '../mockProdServer';
          setupProdMockServer();
        `,
        logger: true,
      }),
      vitePluginSvg({
        // 必要的。必须是绝对路径组成的数组。
        iconDirs: [
            resolve(__dirname, 'src/assets/svg'),
        ],
        // 必要的。入口script
        main: resolve(__dirname, 'src/main.js'),
        symbolIdFormat: 'icon-[name]'
      }),
    ],
    css: {
      postcss: {
        plugins: [
            {
              postcssPlugin: 'internal:charset-removal',
              AtRule: {
                charset: (atRule) => {
                  if (atRule.name === 'charset') {
                    atRule.remove();
                  }
                }
              }
            }
        ],
      },
    }
  };
});
