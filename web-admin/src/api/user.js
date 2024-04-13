import {
  service,
  authAPI,
  publicAPI
} from '@/utils/system/request'

/** 登录api */
export function loginApi (data) {
  return publicAPI.post('/user/login', data)
}

/** 获取用户信息Api */
export function getInfoApi (data) {
  return authAPI.get('/user/', {
    params: data
  })
}

/** 获取用户列表Api */
export function getUserListApi (data) {
  return authAPI.get('/user/list', {
    params: data
  })
}

/** 获取月总日用户增长数列表Api */
export function getUserIncreaseList (data) {
  return authAPI.get('/user/increase', {
    params: data
  })
}

// 更新用户信息
export function updateUserApi (data) {
  return authAPI.post('/user/update', data)
}

// 获取用户总数
export function getUserTotal () {
  return authAPI.get('/user/total')
}

/** 退出登录Api */
export function loginOutApi () {
  return service({
    url: '/user/out',
    method: 'post',
    baseURL: '/mock'
  })
}

/** 获取用户信息Api */
export function passwordChange (data) {
  return service({
    url: '/user/passwordChange',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

/** 获取登录后需要展示的菜单 */
export function getMenuApi () {
  return service({
    url: '/menu/list',
    method: 'post',
    baseURL: '/mock'
  })
}
