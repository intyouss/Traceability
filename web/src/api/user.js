import {
  publicAPI,
  authAPI,
} from '~/axios';


/**
 * 登录
 * @param {string} username 用户名
 * @param {string} password 密码
 * @return {Promise<AxiosResponse<any>>}
 */
export function login(username, password) {
  return publicAPI.post('/user/login', {
    'username': username,
    'password': password,
  });
}

/**
 * 退出登录
 * @return {Promise<AxiosResponse<any>>}
 */
export function logout() {
  return authAPI.post('/user/logout');
}


/**
 * 注册
 * @param {string} username 用户名
 * @param {string} password 密码
 * @param {string} email 邮箱
 * @param {string} mobile 手机号
 * @return {Promise<AxiosResponse<any>>}
 */
export function register(username, password, email = '', mobile = '') {
  return publicAPI.post('/user/register', {
    'username': username,
    'password': password,
    'email': email,
    'mobile': mobile,
  });
}

/**
 * 获取用户信息（公共）
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getPublicInfo(userId) {
  return publicAPI.get('/public/user/', {
    params: {
      'user_id': userId,
    },
  });
}

/**
 * 获取用户信息 (私有)
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function getInfo(userId) {
  return authAPI.get('/user/', {
    params: {
      'id': userId,
    },
  });
}

/**
 * 更新用户信息
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function updateUser(data) {
  return authAPI.post('/user/update', {
    'password': data.password,
    'new_password': data.newPassword,
    'signature': data.signature,
    'email': data.email,
    'mobile': data.mobile,
    'avatar_url': data.avatarUrl,
  });
}

/**
 * 上传头像
 * @param {Data} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function uploadAvatar(data) {
  return authAPI.post('/user/upload/avatar', {
    'avatar_data': data,
  }, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
}

/**
 * 删除远程头像
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function abolishAvatar(userId) {
  return authAPI.post('/user/upload/avatar/abolish', {
    'user_id': userId,
  });
}

/**
 * 删除用户
 * @param {Number} userId
 * @return {Promise<AxiosResponse<any>>}
 */
export function deleteUser(userId) {
  return authAPI.post('/user/delete', {
    'user_id': userId,
  });
}

/**
 * 获取未登录用户搜索结果
 * @param {string} key
 * @param {Number} page
 * @param {Number} limit
 * @return {Promise<AxiosResponse<any>>}
 */
export function getPublicUserSearch(key, page = 0, limit = 0) {
  return publicAPI.get('/user/list', {
    params: {
      'key': key,
      'page': page,
      'limit': limit,
    },
  });
}

/**
 * 获取登录用户搜索结果
 * @param {string} key
 * @param {Number} page
 * @param {Number} limit
 * @return {Promise<AxiosResponse<any>>}
 */
export function getAuthUserSearch(key, page = 0, limit = 0) {
  return authAPI.get('/user/list', {
    params: {
      'key': key,
      'page': page,
      'limit': limit,
    },
  });
}
