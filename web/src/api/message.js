import {
  authAPI,
} from '~/axios';

/**
 * 发送消息
 * @param {Number} toUserId
 * @param {String} content
 * @return {Promise<AxiosResponse<any>>}
 */
export function sendMessage(toUserId, content) {
  return authAPI.post('/message/send', {
    'to_user_id': toUserId,
    'content': content,
  });
}

/**
 * 获取消息列表
 * @param {Number} toUserId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getMessageList(toUserId) {
  return authAPI.get('/message/chat', {params: {
    'to_user_id': toUserId,
  }});
}

