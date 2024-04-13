<script setup>
import { computed, ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useStore } from 'vuex'
import { Sunny } from '@element-plus/icons-vue'
import { getToken } from '~/composables/auth.js'
import { notify } from '~/composables/util.js'
import { useUserByOwner } from '~/composables/userManager.js'

const router = useRouter()
const store = useStore()
const route = useRoute()
const { getUserInfo } = useUserByOwner()

const defaultActive = ref(route.path)
watch(() => route.path, (val) => {
  defaultActive.value = val
})

const isCollapse = computed(() => {
  return store.state.asideWidth === '64px'
})

const asideMenus = [{
  name: '首页',
  icon: 'HomeFilled',
  path: '/'
}, {
  name: '推荐',
  icon: 'StarFilled',
  path: '/recommend'
}, {
  name: '关注',
  icon: 'Flag',
  path: '/focus'
}, {
  name: '朋友',
  icon: 'UserFilled',
  path: '/friend'
}, {
  name: '我的',
  icon: 'User',
  path: '/mine'
}]

const handleSelect = (e) => {
  if (e === '/') {
    return router.push('/')
  } else if (!getToken()) {
    notify('请先登录', 'warning')
    return router.push('/login')
  } else if (e === '/mine') {
    getUserInfo()
    return router.replace('/mine')
  }
  return router.push(e)
}
</script>

<template>
  <div class="f-menu" :style="{ width:$store.state.asideWidth }">
      <el-menu
          :default-active="defaultActive"
          class="el-menu-vertical-demo border-0"
          :collapse="isCollapse"
          @select="handleSelect"
      >
        <template v-for="(item, index) in asideMenus" :key="index">
          <el-menu-item :index="item.path" :route="{path: item.path}">
            <el-icon><component :is="item.icon"></component></el-icon>
            <span>{{ item.name }}</span>
          </el-menu-item>
        </template>
        <el-menu-item>
          <el-icon><Sunny /></el-icon>
          <span>切换</span>
        </el-menu-item>
      </el-menu>
  </div>
</template>

<style>
  .f-menu{
    transition: 0.4s;
    overflow-y: auto;
    overflow-x: hidden;
    top: 61px;
    bottom: 0;
    left: 0;
    @apply fixed bg-light-50;
  }
  .f-menu::-webkit-scrollbar{
    width: 0;
  }
</style>
