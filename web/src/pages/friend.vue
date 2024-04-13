<script setup>
import VPlayer from '~/layouts/components/video/VPlayer.vue'
import { useVideoByPage } from '~/composables/videoManager.js'
import { onBeforeMount } from 'vue'

const {
  Videos,
  getFriendVideos
} = useVideoByPage()

onBeforeMount(() => {
  getFriendVideos()
})
</script>

<template>
  <div>
    <el-empty
        v-if="Videos.length === 0"
        description="你还没有和其它用户成为朋友哦"
    />
    <v-player v-else :videos="Videos"/>
  </div>
</template>

<style scoped>

</style>
