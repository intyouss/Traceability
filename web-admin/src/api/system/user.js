import {
  service,
  authAPI
} from '@/utils/system/request'

// 获取数据api
export function getData (data) {
  return service({
    url: '/system/user/list',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 新增
export function add (data) {
  return service({
    url: '/system/user/add',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 编辑
export function update (data) {
  return service({
    url: '/system/user/update',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 状态变更
export function updateStatus (data) {
  return service({
    url: '/system/user/updateStatus',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 删除
export function del (data) {
  return service({
    url: '/system/user/del',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 获取角色列表
export function getRoles (data) {
  return authAPI.get('/user/role/list', {
    params: data
  })
}

// 添加角色
export function addRole (data) {
  return authAPI.post('/user/role/add', data)
}

// 删除角色
export function deleteRole (data) {
  return authAPI.post('/user/role/delete', data)
}

// 更新角色
export function updateRole (data) {
  return authAPI.post('/user/role/update', data)
}
