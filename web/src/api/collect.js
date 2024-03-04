import {
  authAPI,
} from '~/axios';

/**
 * 获取点赞列表
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getCollectList(params) {
  return authAPI.get('/collect/list', {params: params});
}

/**
 * 点赞/取消点赞操作
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function collectAction(data) {
  return authAPI.post('/collect/action', data);
}

