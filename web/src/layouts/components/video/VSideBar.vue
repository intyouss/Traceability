<script setup>
import {ref, watch} from 'vue';
import {useLike} from '~/composables/likeManager.js';
import {useCollect} from '~/composables/collectManager.js';
import {getVideoInfo} from '~/api/video.js';
import UAvatar from '~/layouts/components/user/UAvatar.vue';


const props = defineProps({
  videoId: Number,
  userId: Number,
  avatar: String,
  isLike: Boolean,
  isCollect: Boolean,
  likeCount: Number,
  commentCount: Number,
  collectCount: Number,
  clickFunc: Function,
});

const emits = defineEmits(['click']);
const handleClick = () => {
  emits('click', 'true');
};
const {
  IsLiked,
  handleLikeAction,
} = useLike(props.isLike);

const {
  IsCollected,
  handleCollectAction,
} = useCollect(props.isCollect);


watch([() => IsLiked.value, () => IsCollected.value], () => {
  getVideoInfo(props.videoId).then((res) => {
    likeCount.value = res.data.video.like_count;
    commentCount.value = res.data.video.comment_count;
    collectCount.value = res.data.video.collect_count;
  });
});

const likeCount = ref(props.likeCount);
const commentCount = ref(props.commentCount);
const collectCount = ref(props.collectCount);
</script>
<template>
  <div class="k" style="height: auto; bottom: 80px;">
    <div class="k1">
      <div class="icon">
        <u-avatar :user-id="props.userId" :avatar="props.avatar"/>
      </div>
      <div class="icon" @click="handleLikeAction(props.videoId)">
        <font-awesome-icon
            class="f-icon"
            :icon="['fas', 'heart']"
            size="2xl"
            style="color: #ffffff;"
            v-if="!IsLiked"
        />
        <font-awesome-icon
            class="f-icon"
            :icon="['fas', 'heart']"
            size="2xl"
            style="color: #f70258;"
            v-else
        />
        <div class="count">
          {{ likeCount}}
        </div>
      </div>
      <div class="icon" @click="handleClick">
        <font-awesome-icon
            class="f-icon"
            :icon="['fas', 'comment-dots']"
            size="2xl"
            style="color: #ffffff;"
        />
        <div class="count">
          {{ commentCount }}
        </div>
      </div>
      <div class="icon" @click="handleCollectAction(props.videoId)">
        <font-awesome-icon
            class="f-icon"
            :icon="['fas', 'star']"
            size="2xl"
            style="color: #ffffff;"
            v-if="!IsCollected"
        />
        <font-awesome-icon
            class="f-icon"
            :icon="['fas', 'star']"
            size="2xl"
            style="color: #fbc709;"
            v-else
        />
        <div class="count">
          {{ collectCount }}
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
