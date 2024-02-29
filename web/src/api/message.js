import {
  authAPI,
} from '~/axios';

/**
 * 发送消息
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function sendMessage(data) {
  return authAPI.post('/message/send', data);
}

/**
 * 获取消息列表
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getMessageList(params) {
  return authAPI.get('/message/chat', {params: params});
}

