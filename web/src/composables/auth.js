import { useCookies } from '@vueuse/integrations/useCookies'
const TokenKey = 'token'
const cookie = useCookies()

/**
 * 获取token
 * @return {any}
 */
export function getToken () {
  return cookie.get(TokenKey)
}

/**
 * 设置token
 * @param {string} token
 * @return {void}
 */
export function setToken (token) {
  return cookie.set(TokenKey, token)
}

/**
 * 移除token
 * @return {void}
 */
export function removeToken () {
  return cookie.remove(TokenKey)
}
