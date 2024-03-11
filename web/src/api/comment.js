import {
  publicAPI,
  authAPI,
} from '~/axios';

/**
 * 获取评论列表
 * @param {Number} videoId 视频ID
 * @param {Number} page 0表示不限制 页码
 * @param {Number} limit 0表示不限制 每页数量
 * @return {Promise<AxiosResponse<any>>}
 */
export function getCommentList(videoId, page=0, limit=0) {
  return publicAPI.get('/comment/list', {params: {
    'video_id': videoId,
    'page': page,
    'limit': limit,
  }});
}

/**
 * 添加评论
 * @param {Number} videoId 视频ID
 * @param {String} content 评论内容
 * @return {Promise<AxiosResponse<any>>}
 */
export function addComment(videoId, content) {
  return authAPI.post('/comment/add', {
    'video_id': videoId,
    'content': content,
  });
}

/**
 * 删除评论
 * @param {Number} commentId 评论ID
 * @return {Promise<AxiosResponse<any>>}
 */
export function deleteComment(commentId) {
  return authAPI.post('/comment/delete', {
    'comment_id': commentId,
  });
}
