<script setup>
import {reactive, ref, watch} from 'vue';
import CCard from '~/layouts/components/comment/CCard.vue';
import {useComment} from '~/composables/commentManager.js';

const props = defineProps({
  openComment: Boolean,
  commentCount: Number,
  videoId: Number,
});

const {
  Loading,
  More,
  Comments,
  loadComments,
  createComment,
} = useComment(props.videoId);


const open = ref(props.openComment);

watch(() => props.openComment, (val) => {
  open.value = val;
});
const emits = defineEmits(['close']);

const handleClose = () => {
  emits('close', 'close');
};

const form = reactive({
  text: '',
});

const handleClick = () => {
  createComment(form.text);
  form.text = '';
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
    <div class="grefd">
      <div class="egwafw" @mousedown.stop>
        <span class="afwafwa">全部评论({{ props.commentCount }})</span>
      </div>
      <el-scrollbar @wheel.stop style="margin-top: 0">
        <div>
          <ul
              v-infinite-scroll="loadComments"
              style="list-style: none;padding: 0;margin-top: 0"
          >
            <!--            <li-->
            <!--                v-for="item in 4"-->
            <!--                :key="item"-->
            <!--                style="padding-right: 13px"-->
            <!--            >-->
            <!--              <c-card-->
            <!--                  style="margin-bottom: 5px"-->
            <!--              />-->
            <!--            </li>-->
            <li
                v-for="item in Comments"
                :key="item.id"
                style="padding-right: 13px"
            >
              <c-card
                  :avatar="item.user.avatar"
                  :user-id="item.user.id"
                  :username="item.user.username"
                  :content="item.content"
                  :time="item.created_at"
                  style="margin-bottom: 5px"
              />
            </li>
          </ul>
          <p
              v-if="Loading"
              class="xguewi"
          >
            加载中...
          </p>
          <p
              v-if="More"
              class="xguewi"
          >
            暂时没有更多评论了
          </p>
        </div>
      </el-scrollbar>
      <div class="gL8GFAmM">
        <el-input
            v-model="form.text"
            placeholder="写下你的评论吧"
            type="textarea"
            maxlength="100"
            autosize
            show-word-limit
        />
        <template v-if="form.text">
          <el-button text class="w-[10px] ml-2" @click="handleClick">
            <font-awesome-icon :icon="['fas', 'paper-plane']" />
          </el-button>
        </template>
      </div>
    </div>
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
  text-shadow: 0 1px 1px rgba(222, 222, 222, .2);
  color: rgba(222, 222, 222);
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

.grefd {
  height: 100%;
  flex-direction: column;
  display: flex;
}
.xguewi {
  font-size: 12px;
  text-align: center;
  margin: 24px 0;
  padding-right: 20px;
  color: #FFFFFFE6;
  line-height: 20px;
}
:deep(.el-textarea) {
  height: 100%;
  border-radius: 12px;
}

:deep(.el-textarea .el-input__count) {
  background: none;
  line-height: initial;
  right: 20px;
}

:deep(.el-textarea__inner) {
  width: 100%;
  height: 100%;
  overflow: hidden;
  padding: 5px 80px 5px 10px;
  min-height: 44px;
  resize: none;
  background-color: rgba(255, 255, 255, .2);
  border: unset !important;
  border-radius: 12px;
  justify-content: center;
  align-items: center;
  font-family: PingFang SC, DFPKingGothicGB-Medium, sans-serif;
  display: flex;
  position: relative;
  box-shadow: none;
  font-size: 14px;
  font-weight: 400;
  line-height: 22px;
  color: #FFFFFFE6;
  text-shadow: 0 1px 1px rgba(0, 0, 0, .2);
}

.gL8GFAmM {
  @apply flex;
  width: 100%;
  max-height: calc(100% - 74px);
  padding-right: 13px;
  z-index: 10;
  flex-grow: 0;
  flex-shrink: 0;
  margin: 10px 0 0;
  position: relative;
}
</style>

<style>
.kk .el-drawer__header {
  margin-bottom: 0 !important;
}

.kk .el-drawer__body {
  padding-top: 0 !important;
  height: calc(100% - 46px);
  flex: 1;
}
</style>
