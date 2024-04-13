import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route = [
  {
    path: '/pages',
    component: Layout,
    redirect: '/pages/jump',
    meta: { title: 'message.menu.pages.name', icon: 'sfont system-shuliang_mianxing' },
    children: [
      {
        path: 'jump',
        component: createNameComponent(() => import('@/views/main/pages/index.vue')),
        meta: { title: 'message.menu.pages.jump' }
      }
    ]
  }
]

export default route
