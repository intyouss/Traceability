import { ElMessageBox, ElNotification } from 'element-plus'
import NProgress from 'nprogress'

/**
 * 弹出提示框
 * @param {String} message 提示内容
 * @param {String} type 提示类型
 * @param {Boolean} dangerouslyUseHTMLString 是否将 message 作为 HTML 片段处理
 */
export function notify (
  message,
  type = 'success',
  dangerouslyUseHTMLString = false
) {
  // eslint-disable-next-line new-cap
  ElNotification({
    message,
    type,
    dangerouslyUseHTMLString,
    duration: 3000
  })
}

/**
 * 弹出确认框
 * @param {String} content 提示内容
 * @param {String} type 提示类型
 * @param {String} title 提示标题
 * @return {Promise} Promise对象
 */
export function confirm (
  content = '提示内容',
  type = 'warning',
  title = '提示'
) {
  return ElMessageBox.confirm(
    content,
    title,
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type
    })
}

/**
 * 显示全屏loading
 */
export function showLoading () {
  NProgress.start()
}

/**
 * 隐藏全屏loading
 */
export function hideLoading () {
  NProgress.done()
}
