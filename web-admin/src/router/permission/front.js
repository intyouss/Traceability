/**
 * 前端路由管理
 **/

/** 引入需要权限的Modules */
import dashboard from '../modules/dashboard'
import userManage from '../modules/userManage'
import pages from '../modules/pages'

/** 登录后需要动态加入的本地路由 */
const FrontRoutes = [
  ...dashboard,
  ...userManage,
  ...pages
]

export default FrontRoutes
