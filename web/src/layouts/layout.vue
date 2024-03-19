<script setup>
import FHeader from './components/FHeader.vue';
import FMenu from './components/FMenu.vue';
</script>
<script>
</script>

<template>
  <el-container>
    <el-header style="z-index: 1">
      <f-header/>
    </el-header>
    <el-container class="gwefdq">
      <el-aside :width="$store.state.asideWidth">
        <f-menu/>
      </el-aside>
      <el-main class="greer">
          <router-view v-slot="{ Component }">
            <transition name="fade">
              <keep-alive :max="4">
                <component :is="Component" />
              </keep-alive>
            </transition>
          </router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<style>
.el-aside{
  transition: 0.4s;
}
.fade-enter-from, .fade-leave-to{
  opacity: 0;
}
.fade-enter-to, .fade-leave-from {
  opacity: 1;
}
.fade-enter-active, .fade-leave-active {
  transition: all 0.4s;
}
.fade-enter-active {
  transition-delay: 0.4s;
}
.el-container {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
body {
  overflow: hidden;
}
.gwefdq {
  height: 647px;
}
.gwefdq .greer {
  overflow: hidden;
  min-width: 860px
}
.el-popup-parent--hidden {
  width: calc(100%) !important;
}
</style>
