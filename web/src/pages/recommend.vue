<script setup>
import VPlayer from '~/layouts/components/video/VPlayer.vue';
import {useVideoByPage} from '~/composables/videoManager.js';
import {onBeforeMount} from 'vue';

const {
  Videos,
  getRecommendVideos,
} = useVideoByPage();

onBeforeMount(() => {
  getRecommendVideos();
});
</script>

<template>
  <div>
    <v-player :videos="Videos" v-if="Videos.length > 0"/>
    <el-empty
        v-else
        description="暂无推荐视频"/>
  </div>
</template>

<style scoped>

</style>
