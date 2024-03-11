import {
  authAPI,
} from '~/axios';

/**
 * 关注/取关
 * @param {Number} userId
 * @param {Number} actionType
 * @return {Promise<AxiosResponse<any>>}
 */
export function relationAction(userId, actionType) {
  return authAPI.post('/relation/action', {
    'user_id': userId,
    'action_type': actionType,
  });
}

/**
 * 获取关注列表
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getFocusList(userId) {
  return authAPI.get('/relation/focus/list', {params: {
    'user_id': userId,
  }});
}

/**
 * 获取粉丝列表
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getFansList(userId) {
  return authAPI.get('/relation/fans/list', {params: {
    'user_id': userId,
  }});
}
