<script setup>
const props = defineProps({
  title: String,
  author: String,
  createdAt: String,
});
</script>

<template>
  <div class="all1">
    <div class="all">
      <div class="account">
        <div class="name">
          <span class="name1">
            @{{props.author}}
          </span>
        </div>
        <div class="time">
          <span class="time1">
            &ensp;·&ensp;{{props.createdAt}}
          </span>
        </div>
      </div>
      <div class="title">
        <span class="title1">
          {{props.title}}
        </span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.account {
  @apply flex flex-row items-center justify-start;
  pointer-events: auto;
}
.name {
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
  max-width: 248px !important;
  flex-shrink: 0;
  pointer-events: auto;
}
.name1 {
  font-size: 18px;
  line-height: 26px;
  color: #fff;
  font-family: PingFang SC, DFPKingGothicGB-Medium, sans-serif;
  font-weight: 500;
  text-shadow: 0.5px 0.5px 1px rgba(0,0,0,.7);
}
.time {
  @apply items-center;
  flex-shrink: 0;
  pointer-events: auto;
}
.time1 {
  font-size: 12px;
  line-height: 22px;
  color: rgba(255, 255, 255, .9);
  text-shadow: 0.5px 0.5px 1px rgba(0,0,0,.7);
  margin-left: 5px;
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
  font-weight: 400;
}
.all {
  width: 100%;
  z-index: 2;
  padding: 16px 95px 16px 16px;
  pointer-events: auto;
}
.all1 {
  z-index: 2;
  height: auto;
  width: 100%;
  pointer-events: none;
  flex-direction: column;
  display: flex;
  position: absolute;
  bottom: 80px;
  padding-left: 38px;
}
.all .title {
  cursor: auto;
  text-shadow: 0.5px 0.5px 1px rgba(0,0,0,.7);
  margin-bottom: 0;
  pointer-events: auto;
  display: flex;
  width: 100%;
  text-overflow: ellipsis;
  text-align: justify;
  -webkit-box-orient: vertical;
  display: -webkit-box;
  position: relative;
  overflow: hidden;
}
.title1 {
  color: #fff;
  vertical-align: middle;
  font-size: 14px;
  line-height: 22px;
}
</style>
