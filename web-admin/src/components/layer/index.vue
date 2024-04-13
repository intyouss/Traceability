<template>
  <div v-drag="props.layer.show">
    <el-dialog
      v-model="props.layer.show"
      :title="props.layer.title"
      :width="props.layer.width"
      center
    >
      <slot></slot>
      <template #footer v-if="props.layer.showButton">
        <div>
          <el-button type="primary" @click="confirm">确认</el-button>
          <el-button @click="close">取消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import vDrag from '@/directive/drag/index'

const props = defineProps({
  layer: {
    type: Object,
    default: () => {
      return {
        show: false,
        title: '',
        showButton: false
      }
    },
    required: true
  }
})

const emits = defineEmits(['confirm'])

const confirm = () => {
  emits('confirm')
}
const close = () => {
  props.layer.show = false
}
</script>

<style lang="scss" scoped>
:deep(.el-form-item__label) {
  padding: 0 1px 0 0 !important;
}
</style>
