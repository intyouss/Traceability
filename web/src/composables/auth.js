import { useCookies } from '@vueuse/integrations/useCookies'
const TokenKey = 'token'
const UserKey = 'user'
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

/**
 * 获取用户信息
 * @return {any}
 */
export function getUser () {
  return cookie.get(UserKey)
}

/**
 * 设置用户信息
 * @param {any} user
 * @return {void}
 */
export function setUser (user) {
  return cookie.set(UserKey, user)
}

/**
 * 移除用户信息
 * @return {void}
 */
export function removeUser () {
  return cookie.remove(UserKey)
}
