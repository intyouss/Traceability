<script setup>
import {onBeforeMount, ref, watch} from 'vue';
import VUserCard from '~/layouts/components/video/VUserCard.vue';
import {useStore} from 'vuex';
import {useVideoByOwner} from '~/composables/videoManager.js';
import DefaultAvatar from '~/assets/icon/default_avatar.jpg';
import RRelationDialog from '~/layouts/components/relation/RRelationDialog.vue';
import {useRelationList} from '~/composables/relationManager.js';
import {useUserByOwner} from '~/composables/userManager.js';

const {
  getUserInfo,
} = useUserByOwner();

const {
  RelationList,
  getRelationList,
  OpenRelationList,
  OpenRelationType,
  openRelationList,
  closeRelationList,
} = useRelationList();

const store = useStore();
const {
  Videos,
  getVideos,
} = useVideoByOwner();
const handleClick = (tab) => {
  getVideos(tab.props.name);
};
onBeforeMount(() => {
  getVideos('作品');
  getUserInfo();
});
const count = (name) => {
  if (name === '作品') {
    return store.state.user.video_count;
  } else if (name === '喜爱') {
    return store.state.user.like_count;
  }
  return store.state.user.collect_count;
};
const tagList = ['作品', '喜爱', '收藏'];

const isDefaultAvatar = () => {
  return store.state.user.avatar === '' ? DefaultAvatar : store.state.user.avatar;
};

const handleRelation = (type) => {
  getRelationList(store.state.user.id, type);
  getUserInfo();
};

const User = ref(store.state.user);

watch(() => store.state.user, (value) => {
  User.value = value;
});
</script>

<template>
  <div>
    <div
          class="ml-2 mr-2"
          :style="{ background: 'linear-gradient(40deg,white,transparent),url('+ isDefaultAvatar() +') center center'}"
      >
        <div class="p-6" style="backdrop-filter: blur(5px);">
          <el-row :gutter="20">
            <el-col :span="4">
              <el-avatar
                  class="avatar"
                  :size="25"
                  :src="isDefaultAvatar()"
              />
            </el-col>
            <el-col :span="20">
              <div class="introduce">
                <div class="username">
                  <h1 class="text-xl m-0">
                  <span class="nameSpan">
                    {{ $store.state.user.username }}
                  </span>
                  </h1>
                </div>
                <div class="count">
                  <div
                      class="option other"
                      @click="openRelationList($store.state.user.id, 1)"
                  >
                    <div class="title">关注</div>
                    <div class="number">{{$store.state.user.focus_count}}</div>
                  </div>
                  <div
                      class="option other"
                      @click="openRelationList($store.state.user.id,2)"
                  >
                    <div class="title">粉丝</div>
                    <div class="number">{{$store.state.user.fans_count}}</div>
                  </div>
                  <div class="option">
                    <div class="title">获赞</div>
                    <div class="number">{{$store.state.user.liked_count}}</div>
                  </div>
                </div>
                <p class="info">
                <span class="age">
                  <font-awesome-icon
                      :icon="['fas', 'mars']"
                      style="color: #005eff;"
                  />
                  <span>
                    22岁
                  </span>
                </span>
                  <span class="city">
                  衡阳
                </span>
                </p>
                <div class="introduction" >
                  <div class="flex relative">
                    <div class="signature">
                      <div class="signature1">
                        <span style="
                          max-width: 300px;
                          overflow: hidden;
                          text-overflow: ellipsis;
                          white-space: nowrap;"
                        >
                          {{ $store.state.user.signature}}
                        </span>
                        <el-tooltip
                            v-if="$store.state.user.signature.length > 25"
                            class="box-item"
                            effect="light"
                            placement="bottom-end"
                        >
                          <template #content>
                            {{$store.state.user.signature}}
                          </template>
                          <div class="flex ml-4px">
                            <span class="more">更多</span>
                          </div>
                        </el-tooltip>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </el-col>
          </el-row>
        </div>
      </div>
    <el-tabs model-value="作品" class="tab p-2" @tab-click="handleClick">
        <el-scrollbar>
          <template v-for="item in tagList" :key="item">
            <el-tab-pane
                :name="item"
                style="z-index: -1;"
            >
              <template #label>
            <span class="text-lg p-2 font-medium">
              <span class="mr-2">{{item}}</span>
              <span>{{count(item)}}</span>
            </span>
              </template>
              <el-row :gutter="10">
                <template v-for="item in Videos" :key="item.id">
                  <el-col :span="6" >
                    <v-user-card
                        :cover-url="item.cover_url"
                        :title="item.title"
                    />
                  </el-col>
                </template>
              </el-row>
            </el-tab-pane>
          </template>
        </el-scrollbar>
    </el-tabs>
    <r-relation-dialog
        :open="OpenRelationList"
        :type="OpenRelationType"
        :close="closeRelationList"
        :user="User"
        :list="RelationList"
        @click="handleRelation"
    />
  </div>
</template>

<style scoped>
.avatar {
  @apply h-[112px] w-[112px] rounded-1/2 relative;
  border: 1px solid rgba(22,24,35,.06)!important;
  box-sizing: content-box;
  flex-grow: 0;
  flex-shrink: 0;
  overflow: hidden;
}

.nameSpan {
  color: #000000;
  display: block;
  flex: none;
  font-size: 20px;
  font-weight: 550;
  line-height: 28px;
  max-width: 300px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.username {
  display: flex;
  position: relative;
  width: 100%;
}
.introduce {
  align-content: center;
  align-items: center;
  display: flex;
  flex: 1 1;
  flex-wrap: wrap;
  margin-left: 32px;
  min-height: 120px;
}
.count {
  display: flex;
  margin-top: 4px;
  width: 100%;
}
.option {
  align-items: center;
  display: flex;
}
.other {
  cursor: pointer;
}
.other::after {
  border-left: 1px solid #f2f2f4;
  content: "";
  height: 16px;
  margin: 0 16px;
  width: 0;
}
.title {
  color: #000000;
  font-size: 14px;
  line-height: 22px;
  margin-right: 6px;
}
.number {
  color: #000000;
  font-size: 16px;
  line-height: 24px;
}
.info {
  align-items: center;
  display: flex;
  height: 20px;
  margin-top: 12px;
  width: 100%;
}
.age, .city {
  align-items: center;
  background: #f2f2f4;
  border-radius: 4px;
  color: rgba(22,24,35,.75);
  display: flex;
  font-size: 12px;
  height: 20px;
  line-height: 20px;
  margin-right: 4px;
  padding: 0 8px;
}
.introduction {
  display: flex;
  height: 20px;
  margin-top: 4px;
  position: relative;
  width: 100%;
}
.introduction span {
  font-size: 12px;
  line-height: 20px;
}
.signature {
  width: 100%;
  height: 20px;
  color: #000000;
  margin-top: 4px;
  display: flex;
  position: relative;
}
.signature1 {
  display: flex;
  position: relative;
}
.signature1 span {
  font-size: 12px;
  line-height: 20px;
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
}
.more {
  color: #000000;
  cursor: default;
  margin-left: 4px;
  position: relative;
}
:deep(.el-tabs__active-bar){
  display: none;
}

:deep(.el-tabs__content) {
  height: 540px;
  overflow-y: auto;
}
:deep(.el-tabs__nav) {
  z-index: 0;
}

:deep(.el-scrollbar__bar.is-horizontal){
  height: 0 !important;
}
:deep(.el-tabs__nav-wrap::after) {
  z-index: initial;
}
</style>
