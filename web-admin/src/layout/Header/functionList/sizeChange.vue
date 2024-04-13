<template>
  <el-dropdown @command="handleCommand" size="default">
    <span class="el-dropdown-link">
      <i class="sfont system-zuixiaohua"></i>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item
          v-for="d in list"
          :key="d.size"
          :command="d.size"
          :disabled=" elementSize === d.size "
        >
          {{ $t(d.name) }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup>
import { computed } from 'vue'
import { useStore } from 'vuex'

const store = useStore()
const elementSize = computed(() => store.state.app.elementSize)
const list = [
  { size: 'large', name: 'message.system.size.large' },
  { size: 'default', name: 'message.system.size.default' },
  { size: 'small', name: 'message.system.size.small' }
]

const handleCommand = (command) => {
  store.commit('app/stateChange', {
    name: 'elementSize',
    value: command
  })
}
</script>

<style lang="scss" scoped>

</style>
