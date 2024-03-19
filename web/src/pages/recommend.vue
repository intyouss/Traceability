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
    <v-player :videos="Videos"/>
  </div>
</template>

<style scoped>

</style>
