<template>
  <el-row :gutter="20" class="card-list">
    <el-col :lg="12" :md="24" v-for="row in list">
      <Row :key="row.id" :row="row"/>
    </el-col>
  </el-row>
</template>

<script setup>
import Row from './row.vue'
import { onBeforeMount, onBeforeUnmount, reactive } from 'vue'
import { getUserTotal } from '@/api/user'
import { getVideoTotal } from '@/api/video'
const list = reactive([
  { id: 1, name: '注册用户总数', data: '8053', color: '#4e73df', icon: 'fa-solid fa-user' },
  { id: 2, name: '发布视频总数', data: '7463', color: '#1cc88a', icon: 'fa-solid fa-video' }
])

onBeforeMount(async () => {
  await getUserTotal().then(res => {
    list[0].data = res.data.total
  })
  await getVideoTotal().then(res => {
    list[1].data = res.data.total
  })
})

let timer = null
timer = setInterval(async () => {
  await getUserTotal().then(res => {
    list[0].data = res.data.total
  })
  await getVideoTotal().then(res => {
    list[1].data = res.data.total
  })
}, 60000)

onBeforeUnmount(() => {
  clearInterval(timer)
  timer = null
})
</script>

<style lang="scss" scoped>
.card-list {
  margin-left: -10px;
  padding-right: 26px;
  display: flex;
  flex-wrap: wrap;
}
</style>
