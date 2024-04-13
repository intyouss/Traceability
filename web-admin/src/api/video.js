import {
  authAPI
} from '@/utils/system/request'

/** 获取月总日视频发布数增长列表 */
export function getVideoIncreaseList (data) {
  return authAPI.get('/video/increase', {
    params: data
  })
}

// 获取发布视频总数
export function getVideoTotal () {
  return authAPI.get('/video/total')
}
