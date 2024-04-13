<template>
  <el-dropdown @command="handleCommand" size="default">
    <span class="el-dropdown-link">
      <i class="sfont system-wenzi"></i>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item
          v-for="(locale, key) in $i18n.messages"
          :key="`locale-${locale.message.language}`"
          :command="key"
          :disabled="name === key"
        >
          {{ locale.message.language }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script lang="js" setup>
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import { useI18n } from 'vue-i18n'
import { changeTitle } from '@/utils/system/title'

const { locale } = useI18n()
const route = useRoute()
const store = useStore()
// 国际化语言切换
const handleCommand = (command) => {
  locale.value = command
  store.commit('app/stateChange', { name: 'lang', value: command })
  changeTitle(route.meta.title)
  document.querySelector('html') && document.querySelector('html').setAttribute('lang', command)
}
</script>

<style lang="scss" scoped>

</style>
