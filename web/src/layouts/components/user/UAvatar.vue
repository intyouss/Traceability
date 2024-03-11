<script setup>
import DefaultAvatar from '~/assets/icon/default_avatar.jpg';
import {useRouter} from 'vue-router';

const props = defineProps({
  userId: Number,
  avatar: String,
  mine: Boolean,
});

const router = useRouter();

const handleClick = (id=0) => {
  if (props.mine) {
    return router.replace('/mine');
  }
  router.replace('/user/' + id);
};

</script>

<template>
  <el-avatar
      class="user-avatar-component"
      :size="50"
      :src="props.avatar ===''?DefaultAvatar:props.avatar"
      @click="handleClick(props.userId)"
  />
</template>

<style scoped>
.user-avatar-component {
  @apply rounded-1/2 relative;
  border: 1px solid rgba(22,24,35,.06)!important;
  box-sizing: content-box;
  cursor: pointer;
  flex-grow: 0;
  flex-shrink: 0;
  overflow: hidden;
  outline: none;
  border-radius: 50%;
}
</style>
