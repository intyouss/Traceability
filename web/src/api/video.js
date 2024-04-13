import {
  publicAPI,
  authAPI
} from '~/axios.js'

/**
 * 获取首页视频
 * @param {Number} type
 * @param {String} latestTime
 * @return {Promise<AxiosResponse<any>>}
 */
export function getIndexVideo (type, latestTime = 0) {
  return publicAPI.get('/video/feed', {
    params: {
      type,
      latest_time: latestTime
    }
  })
}

/**
 * 获取推荐,关注,朋友视频
 * @param {Number} type
 * @param {Number} userId
 * @param {String} latestTime
 * @return {Promise<AxiosResponse<any>>}
 */
export function getAuthVideo (type, userId = 0, latestTime = 0) {
  return authAPI.get('/video/feed', {
    params: {
      type,
      user_id: userId,
      latest_time: latestTime
    }
  })
}

/**
 * 上传视频
 * @param {String} title
 * @param {Data} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function uploadVideo (title, data) {
  return authAPI.post(
    '/video/upload/video',
    {
      title,
      data
    }, { headers: { 'content-type': 'multipart/form-data' } }
  )
}

/**
 * 上传视频封面
 * @param {String} title
 * @param {Data} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function uploadImage (title, data) {
  return authAPI.post(
    '/video/upload/image',
    {
      title,
      cover_image_data: data
    }, { headers: { 'content-type': 'multipart/form-data' } }
  )
}

export function publishVideo (title, playUrl, coverUrl) {
  return authAPI.post('/video/publish', {
    title,
    video_url: playUrl,
    cover_image_url: coverUrl
  })
}

/**
 * 删除视频
 * @Param {String} title
 * @Param {Number} type
 * @return {Promise<AxiosResponse<any>>}
 */
export function abolishVideoUpload (title, type = 1) {
  return authAPI.post('/video/upload/abolish', {
    title,
    type
  })
}

/**
 * 获取其他用户发布视频列表
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getOtherUserVideoList (userId) {
  return publicAPI.get('/video/list', {
    params: {
      user_id: userId
    }
  })
}

/**
 * 获取用户本人发布视频列表
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getUserVideoList (userId) {
  return authAPI.get('/video/list', {
    params: {
      user_id: userId
    }
  })
}

/**
 * 获取视频搜索结果
 * @param {String} key
 * @param {Number} type
 * @return {Promise<AxiosResponse<any>>}
 */
export function getVideoSearch (key, type) {
  return publicAPI.get('/video/search', {
    params: {
      key,
      type
    }
  })
}

/**
 * 获取视频详情
 * @param {Number} id
 * @return {Promise<AxiosResponse<any>>}
 */
export function getVideoInfo (id) {
  return authAPI.get('/video/info', {
    params: {
      id
    }
  })
}
