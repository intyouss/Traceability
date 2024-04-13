<template>
  <el-config-provider :locale="locale" :size="size">
    <router-view></router-view>
  </el-config-provider>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useStore } from 'vuex'

const store = useStore()
const i18n = useI18n()
const size = computed(() => store.state.app.elementSize)
const messages = i18n.messages.value
const locale = computed(() => {
  return {
    name: i18n.locale.value,
    el: messages[i18n.locale.value].el
  }
})
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 100%;
  height: 100vh;
}
.el-tooltip__trigger {
  outline: none
}

</style>
