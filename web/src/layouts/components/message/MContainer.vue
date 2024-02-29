<script setup lang="ts">

</script>

<template>
  <div >
    <div class="h-[330px] pl-4 pt-4 pr-4">
      <el-scrollbar>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div><div>
        消息主体
      </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div>
        <div>
          消息主体
        </div><div>
        消息主体
      </div>
      </el-scrollbar>
    </div>
    <el-input placeholder="发送消息" maxlength="1000" class=" w-[100%] h-[70px] p-4">
      <template #suffix>
        <el-button text class="w-[10px]">
          <font-awesome-icon :icon="['fas', 'paper-plane']" />
        </el-button>
      </template>
    </el-input>
  </div>
</template>

<style scoped>
.el-input__wrapper {
  @apply bg-gray-200 shadow-sm;
}
:deep(.el-input__wrapper.is-focus) {
  @apply  border broder-gray-500/50;
}
.el-input {
  --el-input-focus-border-color: rgba(185,185,185) !important;
}
</style>
