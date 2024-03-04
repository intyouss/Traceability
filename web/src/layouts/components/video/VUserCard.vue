<script setup>
const props = defineProps({
  coverUrl: String,
  title: String,
});
</script>

<template>
  <div class="v-card">
    <el-card class="el-card">
      <img class="img" :src="props.coverUrl" alt="">
      <template #footer>
        <div style="position: relative;">
          <span class="tmp1">{{props.title}}</span>
        </div>
      </template>
    </el-card>
  </div>
</template>

<style scoped>
.el-card {
  @apply rounded-2xl mb-4;
}
.img {
  height: 200px;
  width: 100%;
}
:deep(.el-card__body) {
  padding: 0;
}
.tmp1 {
  word-break: break-all;
  text-overflow: ellipsis;
  -webkit-line-clamp: 2;
  cursor: pointer;
  color: #161823;
  -webkit-box-orient: vertical;
  font-size: 15px;
  line-height: 23px;
  display: -webkit-box;
  overflow: hidden;
}
</style>
