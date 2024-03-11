<script setup>
import {useVideoByPage} from '~/composables/VideoManager.js';
import {
  onBeforeMount,
  reactive, ref,
} from 'vue';
import SwiperCore, {
  Navigation,
  Mousewheel,
} from 'swiper';
import {Swiper, SwiperSlide} from 'swiper/vue';
SwiperCore.use([Mousewheel, Navigation]);
import 'swiper/swiper-bundle.css';
import VFooter from '~/layouts/components/video/VFooter.vue';
import VSideBar from '~/layouts/components/video/VSideBar.vue';
import CDrawer from '~/layouts/components/comment/CDrawer.vue';

const {
  Videos,
  getVideos,
} = useVideoByPage();

const swiperOption = reactive({
  slidesPerView: 1,
  direction: 'vertical',
  mousewheel: true,
  thresholdTime: 700,
  navigation: {
    nextEl: '.swiper-button-next',
    prevEl: '.swiper-button-prev',
  },
});

onBeforeMount(() => {
  getVideos();
  console.log(Videos);
});
const props = defineProps({
  sideBarDisplay: String,
});
const openComment = ref(false);
const handleClick = (val) => {
  if (val === 'true') {
    openComment.value = !openComment.value;
  }
};
const handleClose = (val) => {
  if (val === 'close') {
    openComment.value = false;
  }
};

const setWidth = () => {
  return openComment.value ? 'calc(100% - 336px)' : '100%';
};
</script>

<template>
    <el-row :gutter="20">
      <el-col :span="23">
          <swiper
              :slidesPerView="swiperOption.slidesPerView"
              :direction="swiperOption.direction"
              :mousewheel="swiperOption.mousewheel"
              :thresholdTime="swiperOption.thresholdTime"
              :navigation="swiperOption.navigation"
              class="swiper"
          >
            <!--          <div class="swiper-wrapper" style="z-index: 0">-->
            <swiper-slide v-for="item in Videos" :key="item" >
<!--              class="flex justify-center items-center relative"-->
              <div class="ribvri">
                <div class="grgwh">
                  <div class="breoi">
                    <div class="fegew">
                      <div class="breinrb" :style="{width: setWidth()}">
                        <vue-plyr
                            class="plyr"
                            :data-poster="item.cover_url"
                        >
                          <div :style="{ background: 'linear-gradient(40deg,gray,transparent),url(' + item.cover_url + ') center center'}">
                              <video
                                  controls
                                  crossorigin
                                  playsinline
                                  style="backdrop-filter: blur(10px);"
                              >
                                <source
                                    :src="item.play_url"
                                    type="video/mp4"
                                />
                              </video>

                          </div>

                        </vue-plyr>
                        <v-side-bar
                            :comment-count="item.comment_count"
                            :collect-count="item.collect_count"
                            :like-count="item.like_count"
                            :is-like="item.is_like"
                            :is-collect="item.is_collect"
                            :video-id="item.id"
                            :avatar="item.author.avatar"
                            :user-id="item.author.id"
                            @click="handleClick"
                        />
                        <v-footer
                            :title="item.title"
                            :author="item.author.username"
                            :created-at="item.created_at"
                        />
                      </div>
                      <div class="dadwdwad">
                        <c-drawer :open-comment="openComment" :comment-count="item.comment_count" @close="handleClose"/>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </swiper-slide>
            <!--          </div>-->
          </swiper>
      </el-col>
      <el-col :span="1">
        <div class="swiper-navigation" :style="{display: props.sideBarDisplay}">
            <font-awesome-icon
                :icon="['fas', 'chevron-up']"
                class="swiper-button-prev" size="2xl"
            />
            <font-awesome-icon
                :icon="['fas', 'chevron-down']"
                class="swiper-button-next" size="2xl"
            />
        </div>
      </el-col>
    </el-row>
</template>

<style>
.dadwdwad {
  width: 336px;
  white-space: normal;
  scrollbar-width: none;
  height: 100%;
  z-index: 1000;
  display: inline-block;
  position: relative;
  transform: translateZ(0);
}
.breinrb {
  white-space: normal;
  background-color: transparent;
  position: relative;
  height: 100%;
  z-index: 2;
  vertical-align: top;
  display: inline-block;
  transition: 0.4s;
}
.videoPlayer video{
  border-radius: 20px;
}
.swiper-wrapper {
  transition: 0.3s !important;
}
.swiper-button-prev, .swiper-button-next {
  position: initial;
  --swiper-navigation-color: "";
  --swiper-theme-color: "";
  margin: auto;
  top: initial;
}
.swiper-navigation {
  border: 1px solid var(--el-border-color);
  @apply bg-gray-200;
  margin-top: 240px;
  border-radius: 22px;
  height: 88px;
  width: 44px;
}

.swiper {
  height: 640px !important;
  z-index: 0;
}
.plyr {
  --plyr-color-main: write;
  --plyr-video-background: blur(20px);
}
.plyr--full-ui {
  border-radius: 20px;
}
.ribvri {
  position: absolute;
  left: 0;
  top: calc(0% + 0px);
  width: 100%;
  height: calc(100% - 12px);
  overflow: visible;
  padding: 0 30px;
  padding-left: 8px !important;
  padding-right: 60px !important;
}
.grgwh {
  transform: translate3d(0px, 0px, 0px);
  transition-duration: 0ms;
  flex-direction: column;
  width: 100%;
  height: 100%;
  display: flex;
}
.breoi {
  height: 600px;
  margin-bottom: 12px;
  user-select: none;
  flex-shrink: 0;
  flex-direction: column;
  display: flex;
  position: relative;
}
.fegew {
  border-radius: 16px;
  width: 100%;
  height: 100%;
  opacity: 1;
  background-color: #000;
  white-space: nowrap;
  background-clip: content-box;
  flex-grow: 1;
  transition: all .15s linear;
  position: relative;
  overflow: hidden;
}
</style>
