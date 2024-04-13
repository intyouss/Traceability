import router from '~/router'
import { getToken } from '~/composables/auth.js'
import { hideLoading, notify, showLoading } from '~/composables/util.js'
import { useStore } from 'vuex'

// 全局前置守卫
router.beforeEach(async (to, from, next) => {
  showLoading()

  const store = useStore()
  const token = getToken()
  const user = store.state.user

  // if (!token && to.path !== '/login') {
  //   notify('请先登录', 'error');
  //   return next({path: '/login'});
  // }
  if (token && to.path === '/login') {
    notify('请勿重复登陆', 'error')
    return next({ path: from.path ? from.path : '/' })
  }

  if (token && !user) {
    await store.dispatch('getUserInfo')
  }

  document.title = (to.meta.title ? to.meta.title : '') + '-IntWeb'

  next()
})

// 全局后置钩子
router.afterEach((to, from) => hideLoading())
