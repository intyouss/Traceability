<script setup>
import { onBeforeMount, ref } from 'vue'
import VPlayer from '~/layouts/components/video/VPlayer.vue'
import { useUserByOwner } from '~/composables/userManager.js'
import UAvatar from '~/layouts/components/user/UAvatar.vue'
import { useVideoByPage, useVideoByOther } from '~/composables/videoManager.js'

const {
  Users,
  getUserFocusList
} = useUserByOwner()

const {
  Videos: userVideos,
  getPublishVideos
} = useVideoByOther()

const {
  Videos,
  getFocusVideos
} = useVideoByPage()

onBeforeMount(() => {
  getUserFocusList()
  getFocusVideos()
})

const activeIndex = ref(null)

const handleClick = (index, userId) => {
  activeIndex.value = index
  getPublishVideos(userId)
}

</script>

<template>
  <div>
    <el-empty
        v-if="Users.length === 0"
        description="你还没有关注其他用户哦"
    />
    <el-row v-else :gutter="20">
      <el-col :span="4">
          <ul
              class="hwedg"
          >
            <li
                v-for="(item,index) in Users"
                :key="item.id"
            >
              <div
                  class="dawdaff"
                  :style="activeIndex === index?{boxShadow: '0 1px 2px 0 rgba(0, 0, 0, 0.05)',backgroundColor: 'white'}:''"
                  @click="handleClick(index, item.id)"
              >
                <u-avatar
                    :user-id="item.id"
                    :avatar="item.avatar"
                    :mine="false"
                    class="h-[40px] w-[40px]"
                />
                <div class="dwad">
                  <span class="awdfw">{{ item.username }}</span>
                </div>
              </div>
            </li>
          </ul>
      </el-col>
      <el-col :span="20">
        <v-player
            v-if="(activeIndex !== null && userVideos.length > 0) || (activeIndex === null && Videos.length > 0)"
            :videos="activeIndex !== null?userVideos:Videos"
        />
        <el-empty
            v-else
            description="这个用户暂未发布过视频"/>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
.hwedg{
  overflow-y: auto;
  overflow-x: hidden;
  flex-direction: column;
  display: flex;
  @apply bg-gray-300 shadow-lg h-[600px] border rounded-2xl;
}
.dawdaff {
  @apply rounded-xl ml-2 mr-3;
  height: 55px;
  margin: 10px 10px 5px 10px;
  cursor: pointer;
  user-select: none;
  align-items: center;
  padding-left: 10px;
  display: flex;
  position: relative;
  list-style: none;
}
.dawdaff:hover {
  @apply shadow-sm;
  background-color: white;
  transition: 0.4s;
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
