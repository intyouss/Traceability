import {
  publicAPI,
  authAPI,
} from '~/axios';


/**
 * 登录
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function login(data) {
  return publicAPI.post('/user/login', data);
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
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function register(data) {
  return publicAPI.post('/user/register', data);
}

/**
 * 获取用户信息（公共）
 * @param {object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getPublicInfo(params) {
  return publicAPI.get('/public/user/', {
    params: params,
  });
}

/**
 * 获取用户信息 (私有)
 * @param {object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getInfo(params) {
  return authAPI.get('/user/', {
    params: params,
  });
}

/**
 * 获取用户列表
 * @param {object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getUserList(params) {
  return authAPI.get('/user/list', {
    params: params,
  });
}

/**
 * 更新用户信息
 * @param {object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function updateUser(data) {
  return authAPI.post('/user/update', data);
}

/**
 * 删除用户
 * @param {Object} data
 * @return {Promise<AxiosResponse<any>>}
 */
export function deleteUser(data) {
  return authAPI.post('/user/delete', data);
}

/**
 * 获取用户搜索结果
 * @param {Object} params
 * @return {Promise<AxiosResponse<any>>}
 */
export function getUserSearch(params) {
  return publicAPI.get('/user/search', {
    params: params,
  });
}
