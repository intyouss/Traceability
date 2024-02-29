<script setup>
import {Sunny} from '@element-plus/icons-vue';
import {computed} from 'vue';
import {useStore} from 'vuex';
const store = useStore();

const isCollapse = computed(() => {
  return store.state.asideWidth === '64px';
});
</script>
<script>
import VPlayer from '~/layouts/components/video/VPlayer.vue';


export default {
  components: {
    VPlayer,
  },
  data() {
    return {
      data: [{
        title: 'hahahah',
      }],
    };
  },
};
</script>

<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="3">
        <div class="menu" :style="{ width:$store.state.focusAsideWidth }">
          <el-menu
              class="el-menu-vertical-demo border-0"
              :collapse="isCollapse"
          >
            <template v-for="(item, index) in focusList" :key="index">
              <el-menu-item :index="item.path">
                <el-icon><component :is="item.avatar"></component></el-icon>
                <span>{{ item.name }}</span>
              </el-menu-item>
            </template>
            <el-menu-item>
              <el-icon><Sunny /></el-icon>
              <span>切换</span>
            </el-menu-item>
            <el-menu-item index="/setting">
              <el-icon><Setting /></el-icon>
              <span>设置</span>
            </el-menu-item>
          </el-menu>
        </div>
      </el-col>
      <el-col :span="21">
        <v-player :data="data"/>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
.menu{
  transition: 0.4s;
  overflow-y: auto;
  overflow-x: hidden;
  height: 600px;
  border-bottom-left-radius: 22px;
  border-top-left-radius: 22px;
  @apply shadow-2xl bg-light-100;
}
.menu::-webkit-scrollbar{
  width: 0;
}
</style>
