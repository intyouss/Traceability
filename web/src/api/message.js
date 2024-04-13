import {
  authAPI
} from '~/axios'

/**
 * 发送消息
 * @param {Number} toUserId
 * @param {String} content
 * @return {Promise<AxiosResponse<any>>}
 */
export function sendMessage (toUserId, content) {
  return authAPI.post('/message/send', {
    to_user_id: Number(toUserId),
    content
  })
}

/**
 * 获取消息列表
 * @param {Number} toUserId
 * @param {String} preMsgTime
 * @return {Promise<AxiosResponse<any>>}
 */
export function getMessageList (toUserId, preMsgTime = '0') {
  return authAPI.get('/message/chat', {
    params: {
      to_user_id: Number(toUserId),
      pre_msg_time: preMsgTime
    }
  })
}

/**
 * 获取开放联系人列表
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getOpenUsers (userId) {
  return authAPI.get('/message/open', {
    params: {
      user_id: Number(userId)
    }
  })
}

/**
 * 添加开放联系人
 * @param {Number} openUserId
 * @return {Promise<AxiosResponse<any>>}
 */
export function addOpenUser (openUserId) {
  return authAPI.post('/message/open/add', {
    open_user_id: Number(openUserId)
  })
}

/**
 * 删除开放联系人
 * @param {Number} openUserId
 * @return {Promise<AxiosResponse<any>>}
 */
export function deleteOpenUser (openUserId) {
  return authAPI.post('/message/open/delete', {
    open_user_id: Number(openUserId)
  })
}
