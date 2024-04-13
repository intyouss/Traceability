import {
  createRouter,
  createWebHashHistory
} from 'vue-router'

import Layout from '~/layouts/layout.vue'
import Index from '~/pages/index.vue'
import Login from '~/pages/login.vue'
import NotFound from '~/pages/404.vue'
import Focus from '~/pages/focus.vue'
import Mine from '~/pages/mine.vue'
import Friend from '~/pages/friend.vue'
import Recommend from '~/pages/recommend.vue'
import Search from '~/pages/search.vue'
import userMine from '~/pages/userMine.vue'

const routes = [{
  path: '/',
  component: Layout,
  children: [{
    path: '/',
    component: Index,
    name: 'index',
    meta: {
      title: '首页'
    }
  }, {
    path: '/focus',
    component: Focus,
    name: 'focus',
    meta: {
      title: '关注'
    }
  }, {
    path: '/mine',
    component: Mine,
    name: 'mine',
    meta: {
      title: '我的'
    }
  }, {
    path: '/user/:id',
    component: userMine,
    name: 'userMine',
    meta: {
      title: '用户详情'
    }
  }, {
    path: '/friend',
    component: Friend,
    name: 'friend',
    meta: {
      title: '朋友'
    }
  }, {
    path: '/recommend',
    component: Recommend,
    name: 'recommend',
    meta: {
      title: '推荐'
    }
  }, {
    path: '/search',
    component: Search,
    name: 'search',
    meta: {
      title: '搜索'
    }
  }]
},
{
  path: '/:pathMatch(.*)*',
  name: 'NotFound',
  component: NotFound
}, {
  path: '/login',
  component: Login,
  meta: {
    title: '登录'
  }
}]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
