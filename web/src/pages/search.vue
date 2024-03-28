<script setup>
import VSearchPlayer from '~/layouts/components/video/VSearchPlayer.vue';
import {onMounted, reactive, ref, watch} from 'vue';
import {useSearch} from '~/composables/useManager.js';
import {useRoute} from 'vue-router';
import USearchCard from '~/layouts/components/user/USearchCard.vue';
import NoResult from '~/assets/icon/no-result.svg';

const {getSearch} = useSearch();
const route = useRoute();
const tag = reactive({
  '综合': {
    result: [],
    isClick: false,
    isEmpty: false,
  },
  '视频': {
    result: [],
    isClick: false,
    isEmpty: false,
  },
  '用户': {
    result: [],
    isClick: false,
    isEmpty: false,
  }});
const enTag = ref('');
const enResult = ref([]);
const loading = ref(true);
const isEmpty = ref(false);
const handleTag = (t, value) => {
  if (value.isClick) {
    return;
  }

  const key = route.query.key;
  if (tag[t].result.length === 0 && !tag[t].isEmpty) {
    loading.value = true;
    getSearch(t, key).then((res) => {
      tag[t].isEmpty = res.length === 0;
      isEmpty.value = res.length === 0;
      tag[t].result = res;
      enResult.value = res;
    });
    setTimeout(() => {
      loading.value = false;
    }, 1000);
  } else {
    isEmpty.value = tag[t].isEmpty;
    enResult.value = tag[t].result;
  }
  tag[t].isClick = true;
  tag[enTag.value].isClick = false;
  enTag.value = t;
};

onMounted(() => {
  enTag.value = '综合';
  const key = route.query.key;
  getSearch(enTag.value, key).then((res) => {
    tag[enTag.value].isEmpty = res.length === 0;
    isEmpty.value = res.length === 0;
    tag[enTag.value].result = res;
    enResult.value = res;
  });
  tag[enTag.value].isClick = true;
  setTimeout(() => {
    loading.value = false;
  }, 1000);
});
watch(() => route.query.key, (key) => {
  if (route.path !== '/search') {
    return;
  }
  if (key === '') {
    return;
  }
  console.log(key);
  loading.value = true;
  enTag.value = '综合';
  for (const i in tag) {
    tag[i].isEmpty = false;
    tag[i].isClick = false;
    tag[i].result = [];
  }
  tag[enTag.value].isClick = true;
  getSearch(enTag.value, key).then((res) => {
    tag[enTag.value].isEmpty = res.length === 0;
    isEmpty.value = res.length === 0;
    tag[enTag.value].result = res;
    enResult.value = res;
  });
  setTimeout(() => {
    loading.value = false;
  }, 1000);
});
</script>

<template>
  <el-scrollbar>
    <el-container>
      <el-header class="header">
        <div class="div1">
          <span
              class="tag"
              @click="handleTag(key, value)"
              v-for="(value,key) in tag"
              :key="key"
              :style="{color: enTag === key?'#3b82f6':''}">
            {{key}}
          </span>
        </div>
      </el-header>
      <el-main style="margin-top: 45px" >
        <div v-if = "enTag !== '用户'">
          <template v-if="!isEmpty">
            <v-search-player
                v-for="item in enResult" :key="item.id"
                :title="item.title"
                :created-at="item.created_at"
                :avatar="item.author.avatar"
                :username="item.author.username"
                :play-url="item.play_url"
                :loading="loading"
                :user-id="item.author.id"
                :cover-url="item.cover_url"
            />
          </template>
          <template v-else>
            <el-empty :image="NoResult" style="height: 400px"/>
          </template>
        </div>
        <div v-else>
          <div v-if="!isEmpty">
            <el-row :gutter="20">
              <el-col :span="8" v-for="item in enResult" :key="item.id">
                <u-search-card
                    :signature="item.signature"
                    :avatar="item.avatar"
                    :username="item.username"
                    :fans-count="item.fans_count"
                    :liked-count="item.liked_count"
                    :loading="loading"
                    :is-focus="item.is_focus"
                    :userId="item.id"
                />
              </el-col>
            </el-row>
          </div>
          <div style="margin-left: 35%" v-else>
            <img style="height: 400px" :src="NoResult" alt="" />
          </div>
        </div>
      </el-main>
    </el-container>
  </el-scrollbar>
</template>

<style scoped>
.div1 {
  align-items: center;
  display: flex;
  flex-direction: row;
  height: 26px;
}
.header {
  @apply bg-light-50 w-[100%];
  box-sizing: content-box;
  margin-left: -14px;
  padding: 0 34px;
  position: fixed;
  transition: transform .4s;
  z-index: 10;
}
.tag {
  color: rgba(22,24,35,.75);
  cursor: pointer;
  font-family: PingFang SC,DFPKingGothicGB-Regular,sans-serif;
  font-size: 18px;
  font-weight: 700;
  line-height: 26px;
  margin-right: 40px;
  vertical-align: middle;
  width: 67px;
}
.tag:hover {
  @apply text-blue-500;
}
:deep(.el-scrollbar__bar) {
  z-index: 100;
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
  display: none;
}
</style>
