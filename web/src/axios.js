import axios from 'axios';
import {getToken} from '~/composables/auth.js';
import {notify} from '~/composables/util.js';

export const publicAPI = axios.create({
  baseURL: '/api/v1/public',
});

export const authAPI = axios.create({
  baseURL: '/api/v1',
});

// 添加请求拦截器
authAPI.interceptors.request.use(
    function(config) {
      const token = getToken();
      if (token) {
        config.headers['Authorization'] = 'Bearer ' + token;
      }
      return config;
    }, function(error) {
      return Promise.reject(error);
    });

// 添加响应拦截器
authAPI.interceptors.response.use(
    function(response) {
      if (response.data.code !== 0) {
        const msg = response.data.msg || '请求失败';
        notify(msg, 'error');
        return Promise.reject(new Error(response.data.msg));
      }
      return response.data;
    }, (error) =>{
      return Promise.reject(error);
    });

// 添加响应拦截器
publicAPI.interceptors.response.use(
    function(response) {
      if (response.data.code !== 0) {
        notify(response.data.msg, 'error');
        return Promise.reject(new Error(response.data.msg));
      }
      return response.data;
    }, (error)=> {
      return Promise.reject(error);
    });
