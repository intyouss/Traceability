<template>
  <div class="tags-view-item" :class="props.active? 'active' : ''">
    <router-link :to="{ path: props.menu.path, query: props.menu.query }" v-if="props.menu.meta.title">
      {{ isBackMenu ? props.menu.meta.title : $t(props.menu.meta.title) }}
    </router-link>
    <el-icon @click.stop="reload" v-if="props.active"><refresh-right /></el-icon>
    <el-icon @click.stop="closeTab" v-if="!props.menu.meta.hideClose" :alt="$t('message.common.del')"><close /></el-icon>
  </div>
</template>

<script setup>
import { RefreshRight, Close } from '@element-plus/icons'
import { isBackMenu } from '@/config'

const props = defineProps({
  menu: {
    type: Object,
    default: () => {
      return {
        path: '',
        meta: {
          title: '',
          hideClose: false
        }
      }
    }
  },
  active: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['close', 'reload'])

const closeTab = () => {
  emit('close')
}

// 刷新按钮
const reload = () => {
  emit('reload')
}
</script>

<style lang="scss" scoped>
  .tags-view-item {
    display: inline-flex;
    align-items: center;
    position: relative;
    cursor: pointer;
    height: 26px;
    line-height: 26px;
    border: 1px solid var(--system-header-border-color);
    color: var(--system-header-text-color);
    background: var(--system-header-tab-background);
    padding: 0 8px;
    font-size: 12px;
    margin-left: 5px;
    margin-top: 4px;
    border-radius: 2px;
    a {
      color: var(--system-header-text-color);
      height: 26px;
      display: inline-block;
      padding-left: 8px;
      padding-right: 8px;
    }
    .el-icon-refresh-right {
      display: inline-block;
      margin-right: 5px;
    }
    .el-icon-close {
      display: inline-block;
      height: 26px;
    }
    &:first-of-type {
      margin-left: 15px;
    }
    &:last-of-type {
      margin-right: 15px;
    }
    &.active {
      background: var(--system-primary-color);
      border-color: var(--system-primary-color);
      color: var(--system-primary-text-color);
      a {
        color: var(--system-primary-text-color);
      }
      &:hover {
        background: var(--system-primary-color);
      }
    }
    &:hover {
      background-color: var(--system-header-item-hover-color);
    }
  }
</style>
