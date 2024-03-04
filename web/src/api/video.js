import {
  publicAPI,
  authAPI,
} from '~/axios.js';

/**
 * 获取首页视频
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getIndexVideo(params) {
  return publicAPI.get('/video/feed', {params: params});
}

/**
 * 获取推荐,关注,朋友视频
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getAuthVideo(params) {
  return authAPI.get('/video/feed', {params: params});
}


/**
 * 上传视频
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function uploadVideo(data) {
  return authAPI.post(
      '/video/upload',
      data, {headers: {'content-type': 'multipart/form-data'}},
  );
}

/**
 * 删除视频
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function deleteVideo(data) {
  return authAPI.post('/video/delete', {params: data});
}

/**
 * 获取其他用户发布视频列表
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getOtherUserVideoList(params) {
  return publicAPI.get('/video/list', {params: params});
}

/**
 * 获取用户本人发布视频列表
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getUserVideoList(params) {
  return authAPI.get('/video/list', {params: params});
}

/**
 * 获取视频搜索结果
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getVideoSearch(params) {
  return publicAPI.get('/video/search', {params: params});
}


