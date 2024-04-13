import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route = [
  {
    path: '/userManage',
    component: Layout,
    redirect: '/userManage/menu',
    meta: { title: 'message.menu.userManage.name', icon: 'sfont system-xingmingyonghumingnicheng' },
    alwayShow: true,
    children: [
      {
        path: 'role',
        component: createNameComponent(() => import('@/views/main/userManage/role/index.vue')),
        meta: { title: 'message.menu.userManage.role' }
      },
      {
        path: 'user',
        component: createNameComponent(() => import('@/views/main/userManage/users/index.vue')),
        meta: { title: 'message.menu.userManage.user' }
      }
    ]
  }
]

export default route
