<script setup>
import {ref} from 'vue';

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

</script>
<template>
  <div class="grid justify-items-center">
    <el-avatar class="my-4 cursor-pointer" :size="40" :src="$store.state.user.avatar" />
    <div @click="handleFavorite">
      <font-awesome-icon class="f-icon" :icon="['fas', 'heart']" size="2xl" style="color: #d1d7e0;" v-if="!isFavorite"/>
      <font-awesome-icon class="f-icon" :icon="['fas', 'heart']" size="2xl" style="color: #f70258;" v-else/>
    </div>
    <font-awesome-icon class="f-icon" :icon="['fas', 'comment-dots']" size="2xl" style="color: #d1d7e0;" />
    <div @click="handleCollect">
      <font-awesome-icon class="f-icon" :icon="['fas', 'star']" size="2xl" style="color: #d1d7e0;" v-if="!isCollect"/>
      <font-awesome-icon class="f-icon" :icon="['fas', 'star']" size="2xl" style="color: #fbc709;" v-else/>
    </div>

  </div>
</template>

<style>
  .f-icon{
    cursor: pointer;
    @apply my-3;
  }
</style>
