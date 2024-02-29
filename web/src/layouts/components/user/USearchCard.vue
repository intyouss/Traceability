<script setup>
const props = defineProps({
  avatar: String,
  username: String,
  signature: String,
  fansCount: Number,
  likedCount: Number,
  loading: Boolean,
});
</script>
<template>
  <el-skeleton
      :loading="loading"
      animated
      style="--el-skeleton-circle-size: 65px"
  >
    <template #template>
      <el-card class="mb-4">
        <div class="header">
          <el-skeleton-item variant="circle" />
          <div class="username">
            <el-skeleton-item variant="h1" style="width: 40%" />
          </div>
        </div>
        <div class="flex tag">
          <el-skeleton-item variant="h1" style="width: 12%;margin-right: 4px"/>
          <el-skeleton-item variant="h1" style="width: 12%" />
        </div>
        <el-skeleton-item variant="text" style="width: 100%" />
        <el-skeleton-item variant="text" style="width: 100%" />
      </el-card>
      </template>
      <template #default>
        <el-card class="mb-4">
          <div class="header">
            <el-avatar class="avatar" :size="50" :src="props.avatar" />
            <div class="username">
              <p class="a">
                {{ props.username }}
              </p>
            </div>
            <el-button
                class="button"
                type="primary"
                size="default"
                round
            >
              关注
            </el-button>
          </div>
          <div class="flex tag">
            <span class="count">{{props.likedCount}}获赞</span>
            <span class="count">{{props.fansCount}}粉丝</span>
          </div>
          <p class="t">
            {{ props.signature }}
          </p>
        </el-card>
      </template>
    </el-skeleton>
</template>


    <style scoped>
      .header {
        align-items: center;
        display: flex;
        height: 70px;
      }
      .avatar {
        @apply h-[65px] w-[65px] rounded-1/2 relative shadow-lg;
        border: 1px solid rgba(22,24,35,.06)!important;
        box-sizing: content-box;
        flex-grow: 0;
        flex-shrink: 0;
        overflow: hidden;
        border-radius: 50%;
        position: relative;
      }
      .username {
        flex: 1 1;
        margin-left: 12px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
      .a {
        font-family: PingFang SC, DFPKingGothicGB-Medium, sans-serif;
        font-weight: 500;
        height: 24px;
        max-width: 100%;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        color: #161823;
        font-size: 16px;
        line-height: 24px;
      }
      .button {
        align-items: center;
        border: 0;
        cursor: pointer;
        display: inline-flex;
        justify-content: center;
        margin: 0 8px;
        outline: none;
        padding: 0 16px;
        position: relative;
      }
      .tag {
        color: #161823;
        font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
        font-size: 12px;
        font-weight: 400;
        letter-spacing: .6px;
        line-height: 20px;
        margin-top: 7px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
      .t {
        color: rgba(22, 24, 35, .6);
        font-size: 12px;
        line-height: 20px;
        font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
        font-weight: 400;
        height: 20px;
        letter-spacing: .6px;
        margin-top: 4px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
      .count {
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
    </style>

