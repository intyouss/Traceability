<script setup>
const asideMenus = [{
  'name': 'intyou',
  'avatar': 'https://img.zcool.cn/community/01e79c554be56e000001bf72f7cc20.jpg?x-oss-process=image/auto-orient,1/resize,m_lfit,w_1280,limit_1/sharpen,100/quality,q_100/format,webp',
}, {
  'name': 'intyou专用测试bot',
  'avatar': 'https://img.zcool.cn/community/0121595ae3530ca801214a6172f04a.jpg?x-oss-process=image/auto-orient,1/resize,m_lfit,w_1280,limit_1/sharpen,100/quality,q_100/format,webp',
}, {
  'name': '奇异生物逮捕者',
  'avatar': 'https://img.zcool.cn/community/014fa45f033f5aa801215aa0494bed.jpg?x-oss-process=image/auto-orient,1/resize,m_lfit,w_1280,limit_1/sharpen,100/quality,q_100/format,webp',
}];
</script>

<template>
  <div class="f-message-menu">
    <el-menu
        class="el-menu-vertical-demo border-0"
    >
      <template v-for="(item, index) in asideMenus" :key="index">
        <el-menu-item class="m-1 rounded-lg">
          <el-avatar class="mr-2" :size="40" :src="item.avatar" />
          <span style="text-overflow:ellipsis; width: 8em;white-space: nowrap;overflow: hidden;">{{ item.name }}</span>
        </el-menu-item>
      </template>
    </el-menu>
  </div>
</template>

<style>
.f-message-menu{
  transition: 0.4s;
  overflow-y: auto;
  overflow-x: hidden;
}
.f-message-menu::-webkit-scrollbar{
  width: 0;
}
</style>
