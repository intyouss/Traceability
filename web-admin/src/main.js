import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import { baidu } from './utils/system/statistics'
import 'element-plus/theme-chalk/display.css' // 引入基于断点的隐藏类
import 'element-plus/dist/index.css'
import 'normalize.css' // css初始化
import './assets/style/common.scss' // 公共css
import './theme/modules/chinese/index.scss'
import 'font-awesome/css/font-awesome.min.css'
import App from './App.vue'
import store from './store'
import router from './router'
import { getAuthRoutes } from './router/permission'
import i18n from './locale'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { library } from '@fortawesome/fontawesome-svg-core'

import {
  FontAwesomeIcon, FontAwesomeLayers,
  FontAwesomeLayersText
} from '@fortawesome/vue-fontawesome'

library.add(fas, far, fab)

if (import.meta.env.MODE !== 'development') { // 非开发环境调用百度统计
  baidu()
}

/** 权限路由处理主方法 */
getAuthRoutes().then(() => {
  const app = createApp(App)
  app.use(ElementPlus, { size: store.state.app.elementSize })
  app.use(store)
  app.use(router)
  app.use(i18n)
  app.component('font-awesome-icon', FontAwesomeIcon)

  app.component('font-awesome-layers', FontAwesomeLayers)

  app.component('font-awesome-layers-text', FontAwesomeLayersText)
  // app.config.performance = true
  app.mount('#app')
})
