<script setup>
import VVideo from '~/layouts/components/video/VVideo.vue'
import { ref } from 'vue'

const props = defineProps({
  video: Object
})

const Open = ref(false)

const OpenScreen = () => {
  Open.value = true
}
</script>

<template>
  <div class="v-u-card" @click="OpenScreen">
    <el-card class="el-card">
      <img class="img" :src="props.video.cover_url" alt="">
      <template #footer>
        <div style="position: relative;">
          <span class="hwegwe">{{props.video.title}}</span>
        </div>
      </template>
    </el-card>
  </div>
  <el-dialog
      v-model="Open"
      :close-on-click-modal="false"
      style="pointer-events: auto;"
      :fullscreen="true"
  >
    <v-video :video="props.video"/>
  </el-dialog>
</template>

<style scoped>
.v-u-card {
  z-index: 0;
  cursor: pointer;
}

.v-u-card .el-card {
  @apply rounded-2xl mb-4;
}
.v-u-card .el-card .img {
  height: 180px;
  width: 100%;
}
:deep(.el-card__body) {
  padding: 0;
}
.hwegwe {
  word-break: break-all;
  text-overflow: ellipsis;
  -webkit-line-clamp: 2;
  cursor: pointer;
  color: #161823;
  -webkit-box-orient: vertical;
  font-size: 15px;
  line-height: 23px;
  display: -webkit-box;
  overflow: hidden;
}
</style>
