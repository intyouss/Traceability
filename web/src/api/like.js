import {
  authAPI,
} from '~/axios';

/**
 * 获取点赞列表
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getLikeList(params) {
  return authAPI.get('/like/list', {params: params});
}

/**
 * 点赞/取消点赞操作
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function likeAction(data) {
  return authAPI.post('/like/action', data);
}

