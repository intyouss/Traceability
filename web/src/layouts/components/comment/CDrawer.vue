<script setup>
import {ref, watch} from 'vue';
import CCard from '~/layouts/components/comment/CCard.vue';

const props = defineProps({
  openComment: Boolean,
  commentCount: Number,
});

const open = ref(props.openComment);

watch(() => props.openComment, (val) => {
  open.value = val;
});
const emits = defineEmits(['close']);

const handleClose = () => {
  emits('close', 'close');
};

</script>

<template>
    <el-drawer
        v-model="open"
        title="评论"
        :size="336"
        style="background-color: rgb(40, 40, 40); text-shadow: 0 1px 1px rgba(222,222,222,.2);"
        :before-close="handleClose"
        class="kk"
    >
      <template #header="{titleId, titleClass}">
        <div>
          <div class="flex">
            <h3 :id="titleId" :class="titleClass">评论</h3>
          </div>
          <div class="feafaw"></div>
        </div>
      </template>
      <div class="egwafw">
        <span class="afwafwa">全部评论({{props.commentCount}})</span>
      </div>
      <c-card />
    </el-drawer>
</template>

<style scoped>
:deep(.el-overlay) {
  position: initial;
}
.feafaw {
  width: calc(100% + 32px);
  height: 1px;
  min-height: 1px;
  opacity: .06;
  z-index: 12;
  background-color: #fff;
  display: block;
  position: relative;
  margin-top: 10px;
  left: 0;
}
.egwafw {
  height: unset;
  margin: 12px 0 0;
  padding: 0;
  text-shadow: 0 1px 1px rgba(222,222,222,.2);
  color: rgba(222,222,222);
  z-index: 1;
  flex-grow: 0;
  flex-shrink: 0;
  justify-content: space-between;
  align-items: center;
  font-family: PingFang SC, DFPKingGothicGB-Medium, sans-serif;
  font-size: 16px;
  font-weight: 500;
  line-height: 24px;
  display: flex;
  position: relative;
  top: 0;
}
.afwafwa {
  z-index: 2;
  display: inline-block;
  position: relative;
  font-size: 12px;
}
</style>

<style>
.kk .el-drawer__header {
  margin-bottom: 0 !important;
}
.kk .el-drawer__body {
  padding-top: 0 !important;
}
</style>
