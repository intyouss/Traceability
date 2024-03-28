<script setup>
import VVideo from '~/layouts/components/video/VVideo.vue';
import {ref} from 'vue';

const props = defineProps({
  video: Object,
});

const Open = ref(false);

const OpenScreen = () => {
  Open.value = true;
};

</script>

<template>
  <div class="v-card" @click="OpenScreen">
    <el-card class="el-card">
      <img class="img" :src="props.video.cover_url" alt="">
      <template #footer>
        <div style="position: relative;">
          <span class="tmp1">{{props.video.title}}</span>
          <div class="all">
          <span class="tmp2">
            <span style="vertical-align: top;">@</span>
            <span class="author">{{ props.video.author.username }}</span>
          </span>
            <span class="time"> · {{props.video.created_at}}</span>
          </div>
        </div>
      </template>
    </el-card>
    <el-dialog
      v-model="Open"
      :close-on-click-modal="false"
      style="pointer-events: auto;"
      :fullscreen="true"
    >
      <v-video :video="props.video"/>
    </el-dialog>
  </div>
</template>

<style scoped>
.v-card {
  z-index: 0;
  cursor: pointer;
}
.v-card .el-card {
  @apply rounded-2xl mb-4;
}
.v-card .el-card .img {
  height: 180px;
  width: 100%;
}
:deep(.el-card__body) {
  padding: 0;
}
.tmp1 {
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
.tmp2{
  text-overflow: ellipsis;
  white-space: nowrap;
  backface-visibility: hidden;
  display: inline-block;
  overflow: hidden;
}
.time {
  color: rgba(22,24,35,.6);
  margin-left: 5px;
  font-family: PingFang SC,DFPKingGothicGB-Regular,sans-serif;
  font-size: 12px;
  font-weight: 400;
  line-height: 23px;
}
.author {
  margin-left: 4px;
}
.all {
  word-break: break-all;
  width: 100%;
  white-space: nowrap;
  text-overflow: ellipsis;
  color: rgba(22,24,35,.6);
  cursor: pointer;
  margin-top: 8px;
  font-size: 14px;
  line-height: 22px;
  display: inline-flex;
  overflow: hidden;
}
</style>
