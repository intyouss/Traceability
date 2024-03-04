<script setup>
import {useVideos} from '~/composables/useManager.js';
import {
  onBeforeMount,
  reactive,
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

const {
  Videos,
  getVideoList,
} = useVideos();

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
  getVideoList();
  console.log(Videos);
});
const props = defineProps({
  sideBarDisplay: String,
});
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
              <div class="flex justify-center items-center relative">
                <div class="videoPlayer">
                  <vue-plyr class="plyr">
                      <video
                          controls
                          crossorigin
                          playsinline
                      >
                        <source
                            :src="item.play_url"
                            type="video/mp4"
                        />
<!--                        <track-->
<!--                            default-->
<!--                            kind="captions"-->
<!--                            label="English captions"-->
<!--                            src="/path/to/english.vtt"-->
<!--                            srclang="en"-->
<!--                        />-->
                      </video>
                  </vue-plyr>
                </div>
                <v-side-bar
                    :comment-count="item.comment_count"
                    :like-count="item.like_count"
                    :collect-count="item.collect_count"
                />
                <v-footer
                    :title="item.title"
                    :author="item.author.username"
                    :created-at="item.created_at"
                />
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
.videoPlayer {
  width: 100%;
  height: 600px;
  border-radius: 20px;
  @apply shadow-md shadow-gray-500;
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
}
.plyr--full-ui {
  border-radius: 20px;
}
</style>
