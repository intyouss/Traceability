<script setup>
import {onBeforeMount} from 'vue';
import VPlayer from '~/layouts/components/video/VPlayer.vue';
import {useUserByOwner} from '~/composables/userManager.js';
import DefaultAvatar from '~/assets/icon/default_avatar.jpg';

const {
  Users,
  getUserFocusList,
} = useUserByOwner();

onBeforeMount(() => {
  getUserFocusList();
});
</script>

<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="3">
        <div class="menu">
          <el-menu class="border-0">
            <template v-for="(item, index) in Users" :key="index">
              <el-menu-item class="dawdaff">
                  <el-avatar class="adw broder-2 border-gray-500"
                             :src="item.avatar === ''?DefaultAvatar:item.avatar"
                  />
                <div class="dwad">
                  <span class="awdfw">{{ item.username }}</span>
                </div>

              </el-menu-item>
            </template>
          </el-menu>
        </div>
      </el-col>
      <el-col :span="21">
        <v-player />
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
.menu{
  overflow-y: auto;
  overflow-x: hidden;
  border-bottom-left-radius: 22px;
  border-top-left-radius: 22px;
  @apply shadow-2xl bg-light-100 w-[181px] h-[600px];
}
.menu::-webkit-scrollbar{
  width: 0;
}
.dawdaff {
  @apply rounded-2xl ml-2 mr-3;
  height: 48px;
  margin: 10px;
  -webkit-user-select: none;
  -ms-user-select: none;
  user-select: none;
  align-items: center;
  padding-left: 4px;
  display: flex;
  position: relative;
  list-style: none;
}
.adw {
  flex-shrink: 0;
  position: relative;
  width: 32px;
  height: 32px;
  box-sizing: content-box;
  border-radius: 50%;
  flex-grow: 0;
  overflow: hidden;
  border: 1px solid rgba(22, 24, 35, .06) !important;
}
.dwad {
  flex-direction: column;
  margin-left: 8px;
  display: flex;
}
.awdfw {
  text-overflow: ellipsis;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  line-height: 21px;
  display: -webkit-box;
  overflow: hidden;
  color: #161823;
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
  font-size: 14px;
  font-weight: 400;
}
</style>
