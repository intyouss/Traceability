<script setup>
const props = defineProps({
  coverUrl: String,
  title: String,
  author: String,
  createdAt: String,
});
</script>

<template>
  <div class="v-card">
    <el-card class="el-card">
      <img class="img" :src="props.coverUrl" alt="">
      <template #footer>
        <div style="position: relative;">
          <span class="tmp1">{{props.title}}</span>
          <div class="all">
          <span class="tmp2">
            <span style="vertical-align: top;">@</span>
            <span class="author">{{ props.author }}</span>
          </span>
            <span class="time"> · {{props.createdAt}}</span>
          </div>
        </div>

      </template>
    </el-card>
  </div>
</template>

<style scoped>
.v-card {
  z-index: 0;
  cursor: pointer;
}
.v-card .el-card {
  @apply rounded-2xl mb-4;
}
.v-card .el-card .img {
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
.tmp2{
  text-overflow: ellipsis;
  white-space: nowrap;
  backface-visibility: hidden;
  display: inline-block;
  overflow: hidden;
}
.time {
  color: rgba(22,24,35,.6);
  margin-left: 5px;
  font-family: PingFang SC,DFPKingGothicGB-Regular,sans-serif;
  font-size: 12px;
  font-weight: 400;
  line-height: 23px;
}
.author {
  margin-left: 4px;
}
.all {
  word-break: break-all;
  width: 100%;
  white-space: nowrap;
  text-overflow: ellipsis;
  color: rgba(22,24,35,.6);
  cursor: pointer;
  margin-top: 8px;
  font-size: 14px;
  line-height: 22px;
  display: inline-flex;
  overflow: hidden;
}
</style>
