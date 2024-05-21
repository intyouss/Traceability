import {authAPI} from '@/utils/system/request'

// 删除用户
export function deleteUser (data) {
  return authAPI.post('/user/delete', data)
}

// 添加用户
export function addUser (data) {
  return authAPI.post('/user/add', data)
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
