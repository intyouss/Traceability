<script setup>
import {useStore} from 'vuex';
const store = useStore();


</script>

<template>
  <el-scrollbar>
  <div class="shadow-xl rounded-2xl">
      <div class="background ml-2 mr-2">
        <div class="p-6" style="backdrop-filter: blur(5px);">
          <el-row :gutter="20">
            <el-col :span="4">
              <el-avatar
                  class="avatar"
                  :size="25"
                  src="http://www.quazero.com/uploads/allimg/140303/1-140303215045.jpg"
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
                  <div class="option other">
                    <div class="title">关注</div>
                    <div class="number">{{$store.state.user.follow_count}}</div>
                  </div>
                  <div class="option other">
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
                      <span
                          style="
                          text-overflow:ellipsis;
                          width: 20em;
                          white-space: nowrap;
                          overflow: hidden;"
                      >
                        {{ $store.state.user.signature}}
                      </span>
                    <el-tooltip
                        class="box-item"
                        effect="dark"
                        placement="bottom-end"
                    >
                      <template #content>
                        {{ $store.state.user.signature}}
                      </template>
                      <div class="flex ml-4px">
                        <span class="more">更多</span>
                      </div>
                    </el-tooltip>
                  </div>
                </div>
              </div>
            </el-col>
          </el-row>
        </div>
      </div>
          <el-tabs  model-value="first" class="p-2">
            <el-scrollbar>
            <el-tab-pane name="first" style="z-index: -1;">
              <template #label>

                <span class="text-lg p-2 font-medium">
                  <span class="mr-2">作品</span>
                  <span>{{$store.state.user.video_count}}</span>
                </span>

              </template>
              <el-row
                  :gutter="20"
                  v-for="item in 4"
                  :key="item"
              >
                <el-col
                    v-for="item in 4"
                    :key="item"
                    :span="6"
                >
                  <el-card>
                    <img
                        src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                        class="image"
                        alt=""
                    />
                    <div style="padding: 14px">
                      <span>Yummy hamburger</span>
                    </div>
                  </el-card>
                </el-col>
              </el-row>
            </el-tab-pane>
            <el-tab-pane name="second">
              <template #label>
                <span class="text-lg  p-2 font-medium">
                  <span class="mr-2">喜爱</span>
                  <span>{{$store.state.user.like_count}}</span>
                </span>
              </template>
              <el-row :gutter="20">
                <el-col span="6">1</el-col>
                <el-col span="6">2</el-col>
                <el-col span="6">3</el-col>
                <el-col span="6">4</el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col span="6">1</el-col>
                <el-col span="6">2</el-col>
                <el-col span="6">3</el-col>
                <el-col span="6">4</el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col span="6">1</el-col>
                <el-col span="6">2</el-col>
                <el-col span="6">3</el-col>
                <el-col span="6">4</el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col span="6">1</el-col>
                <el-col span="6">2</el-col>
                <el-col span="6">3</el-col>
                <el-col span="6">4</el-col>
              </el-row>
            </el-tab-pane>
              <!--TODO: 后端需要写一个收藏API-->
            <el-tab-pane label="收藏" name="third">
              <template #label>
                <span class="text-lg p-2 font-medium">
                  <span class="mr-2">收藏</span>
                  <span>32</span>
                </span>
              </template>
              <el-row :gutter="20">
                <el-col span="6">1</el-col>
                <el-col span="6">2</el-col>
                <el-col span="6">3</el-col>
                <el-col span="6">4</el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col span="6">1</el-col>
                <el-col span="6">2</el-col>
                <el-col span="6">3</el-col>
                <el-col span="6">4</el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col span="6">1</el-col>
                <el-col span="6">2</el-col>
                <el-col span="6">3</el-col>
                <el-col span="6">4</el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col span="6">1</el-col>
                <el-col span="6">2</el-col>
                <el-col span="6">3</el-col>
                <el-col span="6">4</el-col>
              </el-row>
            </el-tab-pane>
              </el-scrollbar>
          </el-tabs>

  </div>
  </el-scrollbar>
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
  color: #161823;
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
  color: rgba(22,24,35,.6);
  font-size: 14px;
  line-height: 22px;
  margin-right: 6px;
}
.number {
  color: #161823;
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
.more {
  color: rgba(22,24,35,.34);
  cursor: default;
  margin-left: 4px;
  position: relative;
}
.background {
  background: url('http://www.quazero.com/uploads/allimg/140303/1-140303215045.jpg') no-repeat center center;
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
