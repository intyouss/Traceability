<script setup>
import FTagList from '~/layouts/components/FTagList.vue';
import VCard from '~/layouts/components/video/VCard.vue';
import {onBeforeMount} from 'vue';
import {useVideoByPage} from '~/composables/VideoManager.js';
const {
  Videos,
  getIndexVideos,
} = useVideoByPage();
onBeforeMount(() => {
  getIndexVideos();
});
</script>
<style scoped>
.tag-list {
  @apply h-[40px];
}
:deep(.el-scrollbar__bar.is-horizontal){
  height: 0 !important;
}
</style>
<template>
  <div class="tag-list">
  <f-tag-list class="tag-list"/>
  </div>
  <el-scrollbar>
    <div style="padding: 5px 20px 5px 20px;">
      <el-row :gutter="15">
        <el-col :span="6" v-for="item in Videos" :key="item.id">
          <v-card
              :title="item.title"
              :author="item.author.username"
              :cover-url="item.cover_url"
              :created-at="item.created_at"
          />
        </el-col>
      </el-row>
    </div>
  </el-scrollbar>
</template>

<style scoped>
</style>
