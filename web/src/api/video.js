import {
  publicAPI,
  authAPI,
} from '~/axios.js';

/**
 * 获取首页视频
 * @param {Number} type
 * @param {String} latestTime
 * @return {Promise<AxiosResponse<any>>}
 */
export function getIndexVideo(type, latestTime=0) {
  return publicAPI.get('/video/feed', {params: {
    'type': type,
    'latest_time': latestTime,
  }});
}

/**
 * 获取推荐,关注,朋友视频
 * @param {Number} type
 * @param {String} latestTime
 * @return {Promise<AxiosResponse<any>>}
 */
export function getAuthVideo(type, latestTime=0) {
  return authAPI.get('/video/feed', {params: {
    'type': type,
    'latest_time': latestTime,
  }});
}


/**
 * 上传视频
 * @param {String} title
 * @param {Data} coverImageData
 * @param {Data} Data
 * @return {Promise<AxiosResponse<any>>}
 */
export function uploadVideo(title, coverImageData, Data) {
  return authAPI.post(
      '/video/upload',
      {
        'title': title,
        'cover_image_data': coverImageData,
        'data': Data,
      }, {headers: {'content-type': 'multipart/form-data'}},
  );
}

/**
 * 删除视频
 * @param {Number} videoId
 * @return {Promise<AxiosResponse<any>>}
 */
export function deleteVideo(videoId) {
  return authAPI.post('/video/delete', {params: {
    'video_id': videoId,
  }});
}

/**
 * 获取其他用户发布视频列表
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getOtherUserVideoList(userId) {
  return publicAPI.get('/video/list', {params: {
    'user_id': userId,
  }});
}

/**
 * 获取用户本人发布视频列表
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getUserVideoList(userId) {
  return authAPI.get('/video/list', {params: {
    'user_id': userId,
  }});
}

/**
 * 获取视频搜索结果
 * @param {String} key
 * @param {Number} type
 * @return {Promise<AxiosResponse<any>>}
 */
export function getVideoSearch(key, type) {
  return publicAPI.get('/video/search', {params: {
    'key': key,
    'type': type,
  }});
}

/**
 * 获取视频详情
 * @param {Number} id
 * @return {Promise<AxiosResponse<any>>}
 */
export function getVideoInfo(id) {
  return authAPI.get('/video/info', {params: {
    'id': id,
  }});
}


