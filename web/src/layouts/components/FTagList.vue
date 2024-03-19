<script setup>

</script>

<template>
  <div class="f-tag-list" :style="{left:$store.state.tagListWidth}">
    <el-tag
        v-for="item in 1"
        :key="item"
        :type="'info'"
        effect="light"
        class="f-tag-list__tag"
    >
      全部
    </el-tag>
  </div>
</template>

<style scoped>
.f-tag-list {
  @apply flex items-center fixed right-0 bg-light-100;
  transition: 0.4s;
}
.f-tag-list__tag {
  @apply rounded-lg text-sm text-black m-2 border-none bg-gray-150;
  width: 50px;
  height: 33px;
}
</style>
