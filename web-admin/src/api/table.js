import { service } from '@/utils/system/request'

// 获取数据api
export function getData (data) {
  return service({
    url: '/table/list',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 获取分类数据
export function getCategory (data) {
  return service({
    url: '/table/category',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 获取树组织数据
export function getTree (data) {
  return service({
    url: '/table/tree',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 新增
export function add (data) {
  return service({
    url: '/table/add',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 编辑
export function update (data) {
  return service({
    url: '/table/update',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 删除
export function del (data) {
  return service({
    url: '/table/del',
    method: 'post',
    baseURL: '/mock',
    data
  })
}
