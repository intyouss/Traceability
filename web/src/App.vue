<script>
</script>
<script setup>
import { provide, nextTick, ref, onMounted } from 'vue'
import { useStore } from 'vuex'

const store = useStore()

onMounted(() => {
  window.addEventListener('beforeunload', () => {
    sessionStorage.setItem('state', JSON.stringify(store.state))
  })
})

const isRouterActive = ref(true)

provide('reload', () => {
  isRouterActive.value = false
  nextTick(() => {
    isRouterActive.value = true
  })
})

</script>

<template>
  <router-view v-if="isRouterActive"></router-view>
</template>

<style>

body {
  @apply bg-light-50;
}

#nprogress .bar{
  background-color: #a09ccf !important;
  height: 3px !important;
}
</style>
