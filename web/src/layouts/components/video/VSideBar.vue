<script setup>
import {ref} from 'vue';
import DefaultAvatar from '~/assets/icon/default_avatar.jpg';

const isFavorite = ref(false);
const handleFavorite = ()=>{
  if (isFavorite.value) {
    isFavorite.value = false;
    return;
  }
  isFavorite.value = true;
};

const isCollect = ref(false);
const handleCollect = ()=>{
  if (isCollect.value) {
    isCollect.value = false;
    return;
  }
  isCollect.value = true;
};
const props = defineProps({
  likeCount: Number,
  commentCount: Number,
  collectCount: Number,
});
</script>
<template>
  <div class="k" style="height: auto; bottom: 80px;">
    <div class="k1">
      <div class="icon">
        <el-avatar
            class="avatar"
            :size="40"
            :src="$store.state.user.avatar === ''?DefaultAvatar:$store.state.user.avatar"
        />
      </div>
      <div class="icon" @click="handleFavorite">
        <font-awesome-icon
            class="f-icon"
            :icon="['fas', 'heart']"
            size="2xl"
            style="color: #ffffff;"
            v-if="!isFavorite"
        />
        <font-awesome-icon
            class="f-icon"
            :icon="['fas', 'heart']"
            size="2xl"
            style="color: #f70258;"
            v-else
        />
        <div class="count">
          {{ props.likeCount }}
        </div>
      </div>
      <div class="icon">
        <font-awesome-icon
            class="f-icon"
            :icon="['fas', 'comment-dots']"
            size="2xl"
            style="color: #ffffff;"
        />
        <div class="count">
          {{ props.commentCount }}
        </div>
      </div>
      <div class="icon" @click="handleCollect">
        <font-awesome-icon
            class="f-icon"
            :icon="['fas', 'star']"
            size="2xl"
            style="color: #ffffff;"
            v-if="!isCollect"
        />
        <font-awesome-icon
            class="f-icon"
            :icon="['fas', 'star']"
            size="2xl"
            style="color: #fbc709;"
            v-else
        />
        <div class="count">
          {{ props.collectCount }}
        </div>
      </div>

    </div>
  </div>

</template>

<style>
  .f-icon{
    cursor: pointer;
    height: 30px;
    justify-content: center;
    align-items: center;
    display: flex;
  }
  .k {
    padding-right: 38px;
    z-index: 11;
    flex-direction: column;
    justify-content: flex-end;
    align-items: center;
    display: flex;
    position: absolute;
    right: 0;
  }
  .k1 {
    margin-bottom: 0;
    filter: drop-shadow(0 0 3px rgba(0, 0, 0, .3));
    flex-direction: column;
    flex-shrink: 0;
    justify-content: center;
    align-items: center;
    display: flex;
    position: relative;
  }
  .icon {
    position: relative;
    vertical-align: bottom;
    padding-bottom: 15px;
  }
  .icon .avatar {
    margin-top: 24px;
    margin-bottom: 23px;
    position: relative;
    vertical-align: bottom;
  }
  .icon .count {
    font-weight: 400;
    color: #fff;
    opacity: .9;
    justify-content: center;
    align-items: center;
    font-family: PingFang SC, DFPKingGothicGB-Medium, sans-serif;
    font-size: 15px;
    line-height: 23px;
    display: flex;
    position: relative;
  }
</style>
