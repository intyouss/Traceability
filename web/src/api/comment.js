import {
  publicAPI,
  authAPI,
} from '~/axios';

/**
 * 获取评论列表
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getCommentList(params) {
  return publicAPI.get('/comment/list', {params: params});
}

/**
 * 添加评论
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function addComment(data) {
  return authAPI.post('/comment/add', data);
}

/**
 * 删除评论
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function deleteComment(data) {
  return authAPI.post('/comment/delete', data);
}
