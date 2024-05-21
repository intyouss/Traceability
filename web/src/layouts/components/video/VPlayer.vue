<script setup>
import {
  reactive, ref, watch
} from 'vue'
import {
  Navigation,
  Mousewheel
} from 'swiper/modules'
import { Swiper, SwiperSlide } from 'swiper/vue'
import 'swiper/swiper-bundle.css'
import VFooter from '~/layouts/components/video/VFooter.vue'
import VSideBar from '~/layouts/components/video/VSideBar.vue'
import CDrawer from '~/layouts/components/comment/CDrawer.vue'

const swiperOption = reactive({
  slidesPerView: 1,
  direction: 'vertical',
  mousewheel: true,
  thresholdTime: 700,
  modules: [Navigation, Mousewheel],
  navigation: {
    nextEl: '.swiper-button-next',
    prevEl: '.swiper-button-prev'
  }
})
const props = defineProps({
  sideBarDisplay: String,
  videos: ref([])
})

const Videos = ref(props.videos)

watch(() => props.videos, (val) => {
  Videos.value = val
})

const openComment = ref([])
const handleClick = (val, index) => {
  if (val === 'true') {
    openComment.value[index] = !openComment.value[index]
  }
}
const handleClose = (val, index) => {
  if (val === 'close') {
    openComment.value[index] = false
  }
}

const setWidth = (index) => {
  return openComment.value[index] ? 'calc(100% - 336px)' : '100%'
}

</script>

<template>
    <el-row :gutter="20">
      <el-col :span="Videos.length > 1?23:24">
          <swiper
              :slidesPerView="swiperOption.slidesPerView"
              :direction="swiperOption.direction"
              :mousewheel="swiperOption.mousewheel"
              :thresholdTime="swiperOption.thresholdTime"
              :navigation="swiperOption.navigation"
              :modules="swiperOption.modules"
              class="swiper swiper-no-swiping"
              ref="swiperRef"
          >
            <swiper-slide
                v-for="(item,index) in Videos"
                :key="item.id"
            >
              <div class="ribvri">
                <div class="grgwh">
                  <div class="breoi">
                    <div class="fegew">
                      <div class="breinrb" :style="{width: setWidth(index)}">
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
                            @click="handleClick($event, index)"
                        />
                        <v-footer
                            :title="item.title"
                            :author="item.author.username"
                            :created-at="item.created_at"
                        />
                      </div>
                      <div class="dadwdwad">
                        <c-drawer
                            :video-id="item.id"
                            :open-comment="openComment[index]"
                            :comment-count="item.comment_count"
                            @close="handleClose($event, index)"
                        />
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </swiper-slide>
          </swiper>
      </el-col>
      <el-col :span="Videos.length > 1?1:0" :style="{display: props.sideBarDisplay}">
        <div class="swiper-navigation">
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
  transition: 0.3s;
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
.breinrb .plyr {
  --plyr-color-main: write;
}
:deep(.plyr--full-ui) {
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
</style>
