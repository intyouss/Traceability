import {
  authAPI
} from '~/axios'

/**
 * 获取收藏列表
 * @param {Number} userId 用户ID
 * @return {Promise<AxiosResponse<any>>}
 */
export function getCollectList (userId) {
  return authAPI.get('/collect/list', { params: { user_id: userId } })
}

/**
 * 收藏/取消收藏操作
 * @param {Number} videoId 视频ID
 * @param {Number} actionType 操作类型 1:收藏 2:取消收藏
 * @return {Promise<AxiosResponse<any>>}
 */
export function collectAction (videoId, actionType) {
  return authAPI.post('/collect/action', {
    video_id: videoId,
    action_type: actionType
  })
}
