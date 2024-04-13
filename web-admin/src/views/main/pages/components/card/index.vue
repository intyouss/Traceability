<template>
  <el-row :gutter="20">
    <el-col :lg="12" :md="24"  v-for="row in list">
      <Row :key="row.id" :row="row"/>
    </el-col>
  </el-row>
</template>

<script setup>
import Row from './row.vue'

const MinioUrl = import.meta.env.VITE_BASE_ENDPOINT + ':' + import.meta.env.VITE_MINIO_PORT
const list = [
  { id: 1, name: 'minio对象存储', desc: '视频、图片等多媒体文件存储管理', url: MinioUrl },
  { id: 2, name: 'kafka消息队列', desc: '消息队列管理', url: 'https://kafka.apache.org/' },
  { id: 3, name: 'redis缓存', desc: '缓存管理', url: 'https://redis.io/' },
  { id: 4, name: 'mysql数据库', desc: '数据库管理', url: 'https://www.mysql.com/' }
]
</script>

<style lang="scss" scoped>
</style>
