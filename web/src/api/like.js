import {
  authAPI,
} from '~/axios';

/**
 * 获取点赞列表
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getLikeList(userId) {
  return authAPI.get('/like/list', {params: {
    'user_id': userId,
  }});
}

/**
 * 点赞/取消点赞操作
 * @param {Number} videoId
 * @param {Number} actionType
 * @return {Promise<AxiosResponse<any>>}
 */
export function likeAction(videoId, actionType) {
  return authAPI.post('/like/action', {
    'video_id': videoId,
    'action_type': actionType,
  });
}

