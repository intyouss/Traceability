import {
  authAPI,
} from '~/axios';

/**
 * 关注/取关
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function relationAction(data) {
  return authAPI.post('/relation/action', data);
}

/**
 * 获取关注列表
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getFocusList(params) {
  return authAPI.get('/relation/focus/list', {params: params});
}

/**
 * 获取粉丝列表
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getFansList(params) {
  return authAPI.get('/relation/fans/list', {params: params});
}
