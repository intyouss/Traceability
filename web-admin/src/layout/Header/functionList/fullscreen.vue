<template>
  <div :title="isFullscreen ? $t('message.system.fullScreenBack') : $t('message.system.fullScreen')">
    <i class="sfont" :class="isFullscreen ? 'system-quanping':'system-quanping1'" @click="toggle"></i>
  </div>
</template>

<script setup>
import { useFullscreen } from '@vueuse/core'
const { isFullscreen, toggle } = useFullscreen()
</script>

<style lang="scss" scoped>
  i {
    cursor: pointer;
    font-size: 18px;
    &:focus {
      outline: none;
    }
  }
</style>
