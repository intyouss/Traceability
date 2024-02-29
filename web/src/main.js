import {createApp} from 'vue';

import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import App from './App.vue';
import router from './router';
import store from '~/store';
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import 'animate.css';
import {fas} from '@fortawesome/free-solid-svg-icons';
import {far} from '@fortawesome/free-regular-svg-icons';
import {fab} from '@fortawesome/free-brands-svg-icons';
import {library} from '@fortawesome/fontawesome-svg-core';
import VuePlyr from 'vue-plyr';
import 'vue-plyr/dist/vue-plyr.css';

import {
  FontAwesomeIcon, FontAwesomeLayers,
  FontAwesomeLayersText,
} from '@fortawesome/vue-fontawesome';
const app = createApp(App);
app.use(store);
app.use(router);
app.use(ElementPlus);
app.use(VuePlyr, {
  plyr: {
    i18n: {
      speed: '速度',
      normal: '正常',
    },
  },
});
library.add(fas, far, fab);

app.component('font-awesome-icon', FontAwesomeIcon);

app.component('font-awesome-layers', FontAwesomeLayers);

app.component('font-awesome-layers-text', FontAwesomeLayersText);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}
import 'virtual:windi.css';

import './permission.js';

import 'nprogress/nprogress.css';

app.mount('#app');
