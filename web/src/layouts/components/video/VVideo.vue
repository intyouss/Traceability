<script setup>

import CDrawer from '~/layouts/components/comment/CDrawer.vue';
import VFooter from '~/layouts/components/video/VFooter.vue';
import VSideBar from '~/layouts/components/video/VSideBar.vue';
import {ref} from 'vue';

const enableComment = ref(false);
const handleClick = (val) => {
  if (val === 'true') {
    enableComment.value = !enableComment.value;
  }
};
const handleClose = (val) => {
  if (val === 'close') {
    enableComment.value = false;
  }
};

const width = () => {
  return enableComment.value ? 'calc(100% - 336px)' : '100%';
};

const props = defineProps({
  video: Object,
});
</script>

<template>
    <div class="nytytn">
      <div class="bgnffntn">
        <div class="tyjtyj">
          <div class="nrtheh" :style="{width: width()}">
            <vue-plyr
                class="plyr"
                :data-poster="props.video.cover_url"
            >
              <div :style="{ background: 'linear-gradient(40deg,gray,transparent),url(' + props.video.cover_url + ') center center'}">
                <video
                    class="ghrwfew"
                    controls
                    crossorigin
                    playsinline
                    style="backdrop-filter: blur(10px);"
                >
                  <source
                      :src="props.video.play_url"
                      type="video/mp4"
                  />
                </video>
              </div>
            </vue-plyr>
            <v-side-bar
                :comment-count="props.video.comment_count"
                :collect-count="props.video.collect_count"
                :like-count="props.video.like_count"
                :is-like="props.video.is_like"
                :is-collect="props.video.is_collect"
                :video-id="props.video.id"
                :avatar="props.video.author.avatar"
                :user-id="props.video.author.id"
                @click="handleClick"
            />
            <v-footer
                :title="props.video.title"
                :author="props.video.author.username"
                :created-at="props.video.created_at"
            />
          </div>
          <div class="nytrn">
            <c-drawer
                :video-id="props.video.id"
                :open-comment="enableComment"
                :comment-count="props.video.comment_count"
                @close="handleClose"
            />
          </div>
        </div>
      </div>
    </div>
</template>

<style scoped>
.nytrn {
  width: 336px;
  white-space: normal;
  scrollbar-width: none;
  height: 100%;
  z-index: 1000;
  display: inline-block;
  position: relative;
  transform: translateZ(0);
}
.nrtheh {
  white-space: normal;
  background-color: transparent;
  position: relative;
  height: 100%;
  z-index: 2;
  vertical-align: top;
  display: inline-block;
  transition: 0.3s;
}
.nrtheh .plyr {
  --plyr-color-main: write;
}
:deep(.plyr--full-ui) {
  border-radius: 20px;
}
.nytytn {
  transform: translate3d(0px, 0px, 0px);
  transition-duration: 0ms;
  flex-direction: column;
  width: 100%;
  height: 100%;
  display: flex;
  padding: 0 30px 0 30px;
}
.bgnffntn {
  height: 600px;
  user-select: none;
  flex-shrink: 0;
  flex-direction: column;
  display: flex;
  position: relative;
}
.tyjtyj {
  @apply border border-gray-500 shadow-xl;
  border-radius: 20px;
  width: 100%;
  height: 100%;
  opacity: 1;
  white-space: nowrap;
  background-clip: content-box;
  flex-grow: 1;
  transition: all .15s linear;
  position: relative;
  overflow: hidden;
}
.ghrwfew {
  border-radius: 20px;
}
</style>
