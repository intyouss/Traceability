import axios from 'axios'
import store from '@/store'
import { ElMessage } from 'element-plus'
import { notify } from '@/composables/util'

const baseURL = import.meta.env.VITE_BASE_URL

export const service = axios.create({
  baseURL,
  timeout: 5000
})

export const publicAPI = axios.create({
  baseURL: '/api/v1/public'
})

export const authAPI = axios.create({
  baseURL: '/api/v1/admin'
})

// 请求前的统一处理
service.interceptors.request.use(
  (config) => {
    // JWT鉴权处理
    if (store.getters['user/token']) {
      config.headers.token = store.state.user.token
    }
    return config
  },
  (error) => {
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

service.interceptors.response.use(
  (response) => {
    const res = response.data
    if (res.code === 200) {
      return res
    } else {
      showError(res)
      return Promise.reject(res)
    }
  },
  (error) => {
    console.log(error) // for debug
    const badMessage = error.message || error
    const code = parseInt(badMessage.toString().replace('Error: Request failed with status code ', ''))
    showError({ code, message: badMessage })
    return Promise.reject(error)
  }
)

// 添加请求拦截器
authAPI.interceptors.request.use(
  function (config) {
    if (store.getters['user/token']) {
      config.headers.Authorization = 'Bearer ' + store.state.user.token
    }
    return config
  }, function (error) {
    return Promise.reject(error)
  })

// 添加响应拦截器
authAPI.interceptors.response.use(
  function (response) {
    if (response.data.code !== 0) {
      const msg = response.data.msg || '请求失败'
      notify(msg, 'error')
      return Promise.reject(new Error(response.data.msg))
    }
    if (response.headers.token) {
      store.commit('user/tokenChange', response.headers.token)
    }
    return response.data
  }, (error) => {
    return Promise.reject(error)
  })

// 添加响应拦截器
publicAPI.interceptors.response.use(
  function (response) {
    if (response.data.code !== 0) {
      const msg = response.data.msg || '请求失败'
      notify(msg, 'error')
      return Promise.reject(new Error(response.data.msg))
    }
    return response.data
  }, (error) => {
    return Promise.reject(error)
  })

// 错误处理
function showError (error) {
  // token过期，清除本地数据，并跳转至登录页面
  if (error.code === 403) {
    // to re-login
    store.dispatch('user/loginOut')
  } else {
    ElMessage({
      message: error.msg || error.message || '服务异常',
      type: 'error',
      duration: 3 * 1000
    })
  }
}
