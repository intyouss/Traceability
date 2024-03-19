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
  coverUrl: String,
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
          <div class="a">
            <el-skeleton-item variant="circle"/>
            <div class="c" style="white-space: normal">
              <div class="d">
                <el-skeleton-item
                    variant="h1"
                    style="width: 6%;margin-right: 5px"
                />
                <el-skeleton-item variant="h1" style="width: 5%"/>
              </div>
            </div>
          </div>
          <div class="h">
            <el-skeleton-item variant="h1" style="width: 100%"/>
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
        <div class="gwewf">
          <div class="a">
            <div class="b">
              <u-avatar :user-id="props.userId" :avatar="props.avatar"/>
            </div>
            <div class="c">
              <div class="gewvfuyf">
                <a href="" class="ntr">
                  <p class="f">
                    {{ props.username }}
                  </p>
                </a>
                <p class="g">
            <span style="margin: 0 4px;">
                · {{ props.createdAt }}
              </span>
                </p>
              </div>

            </div>
          </div>
          <div class="h">
            <span class="i">
              {{ props.title }}
            </span>
          </div>
          <div class="herwge">
            <div class="videoPlayer">
              <vue-plyr
                  class="plyr"
                  :data-poster="props.coverUrl"
              >
                <div :style="{ background: 'linear-gradient(40deg,gray,transparent),url(' + props.coverUrl + ') center center'}">
                  <video
                      controls
                      crossorigin
                      playsinline
                      style="backdrop-filter: blur(10px);"
                  >
                    <source
                        :src="props.playUrl"
                        type="video/mp4"
                    />
                  </video>
                </div>
              </vue-plyr>
            </div>
          </div>
        </div>
      </el-card>
    </template>
  </el-skeleton>
</template>

<style scoped>
.gwewf .a {
  @apply w-[100%];
  align-items: center;
  display: flex;
}

.herwge .videoPlayer {
  width: 100%;
  height: 600px;
  margin-bottom: 12px;
  user-select: none;
  flex-shrink: 0;
  flex-direction: column;
  display: flex;
  border-radius: 20px;
  opacity: 1;
  white-space: nowrap;
  background-clip: content-box;
  flex-grow: 1;
  transition: all .15s linear;
  position: relative;
  overflow: hidden;
  @apply shadow-lg border border-gray-500;
}

.herwge .videoPlayer video {
  border-radius: 20px;
}

.gwewf .a .b {
  align-items: center;
  display: flex;
}

.gwewf .a .c {
  flex: 1 1;
  margin-left: 12px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  justify-content: center;
  align-items: center;
}

.gewvfuyf {
  align-items: center;
  display: flex;
}

.ntr {
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
  font-family: PingFang SC, DFPKingGothicGB-Medium, sans-serif;
  font-weight: 500;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.g {
  color: #161823;
  flex-shrink: 0;
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
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
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
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
