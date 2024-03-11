<script setup>
import UAvatar from '~/layouts/components/user/UAvatar.vue';

const props = defineProps({
  avatar: String,
  userId: Number,
  username: String,
  title: String,
  createdAt: String,
  playUrl: String,
  loading: Boolean,
});
</script>

<template>
  <el-skeleton
      :loading="loading"
      animated
      style="--el-skeleton-circle-size: 55px"
  >
    <template #template>
      <el-card class="mb-4">
        <div>
          <div class="a" >
            <el-skeleton-item variant="circle" />
            <div class="c" style="white-space: normal">
              <div class="d">
                <el-skeleton-item
                    variant="h1"
                    style="width: 6%;margin-right: 5px"
                />
                <el-skeleton-item variant="h1" style="width: 5%" />
              </div>
            </div>
          </div>
          <div class="h">
            <el-skeleton-item variant="h1" style="width: 100%" />
          </div>
            <el-skeleton-item
                variant="caption"
                style="width: 100%;height: 600px;border-radius: 20px;"
            />
          </div>
      </el-card>
    </template>
    <template #default>
      <el-card class="mb-4">
        <div>
          <div class="a" >
            <div class="b">
              <u-avatar :user-id="props.userId" :avatar="props.avatar"/>
            </div>
            <div class="c">
              <div class="d">
                <a href="" class="e">
                  <p class="f">
                    {{props.username}}
                  </p>
                </a>
                <p class="g">
            <span style="margin: 0 4px;">
                · {{props.createdAt}}
              </span>
                </p>
              </div>

            </div>
          </div>
          <div class="h">
      <span class="i">
        {{props.title}}
      </span>
          </div>
          <div class="flex justify-center items-center relative">
            <div class="videoPlayer">
              <vue-plyr class="plyr">
                <video
                    controls
                    crossorigin
                    playsinline
                >
                  <source
                      :src="props.playUrl"
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
          </div>
        </div>
      </el-card>
    </template>
  </el-skeleton>
</template>

<style scoped>
.a {
  @apply w-[100%];
  align-items: center;
  display: flex;
}
.videoPlayer {
  width: 100%;
  height: 600px;
  border-radius: 20px;
  @apply shadow-md shadow-gray-500;
}
.videoPlayer video{
  border-radius: 20px;
}
.b {
  align-items: center;
  display: flex;
}
.c {
  flex: 1 1;
  margin-left: 12px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  justify-content:center;
  align-items: center;
}
.d {
  align-items: center;
  display: flex;
}
.e {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  position: relative;
  text-decoration: none;
}
.f {
  color: #161823;
  font-size: 16px;
  line-height: 24px;
  font-family: PingFang SC,DFPKingGothicGB-Medium,sans-serif;
  font-weight: 500;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.g {
  color: #161823;
  flex-shrink: 0;
  font-family: PingFang SC,DFPKingGothicGB-Regular,sans-serif;
  font-size: 12px;
  font-weight: 400;
  letter-spacing: .6px;
  line-height: 20px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.h {
  margin: 16px 0;
  display: flex;
  position: relative;
  width: 100%;
  font-family: PingFang SC,DFPKingGothicGB-Regular,sans-serif;
  font-weight: 500;
  overflow: hidden;
  max-height: 48px;
  font-size: 18px;
}
.plyr {
  --plyr-color-main: write;
}
.plyr--full-ui {
  border-radius: 20px;
}
</style>
