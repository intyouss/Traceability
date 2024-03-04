<script setup>

import {
  useRePassword, useLogout,
  useMessage, useVideoUpload,
} from '~/composables/useManager.js';
import {ref} from 'vue';
import DefaultAvatar from '~/assets/icon/default_avatar.jpg';
import {Search} from '@element-plus/icons-vue';
import MContainer from '~/layouts/components/message/MContainer.vue';
import MMenu from '~/layouts/components/message/MMenu.vue';
import VUpload from '~/layouts/components/video/VUpload.vue';
import {getToken} from '~/composables/auth.js';
import {useRoute, useRouter} from 'vue-router';
const router = useRouter();
const route = useRoute();
const {handleLogout} = useLogout();
const {
  rePasswordForm,
  rePasswordFormOpen,
  rePasswordFormClose,
  formLabelWidth,
  form,
  rules,
  formRef,
  onSubmit,
} = useRePassword();
const {
  Message,
  MessageOpen,
} = useMessage();
const {
  Upload,
  UploadOpen,
} = useVideoUpload();
const handleRefresh = () => {
  location.reload();
};

const key = ref('');
const enKey = ref('');

const showSearch = ref(false);

const handleSearch = () => {
  if (showSearch.value) {
    key.value = '';
    return showSearch.value = false;
  }
  showSearch.value = true;
};

const search = () => {
  if (key.value === '') {
    return;
  }
  if (key.value === enKey.value && route.path === '/search') {
    return;
  }
  router.push({path: '/search', query: {key: key.value}});
  enKey.value = key.value;
};
</script>

<template>
  <div class="f-header">
    <span class="logo">
      <img src="../../assets/icon/logo.svg" class="mr-1" alt=""
           style="height:50px; width:50px" />
      溯源
    </span>
    <el-icon class="icon-btn" @click="$store.commit('handleAsideWidth')">
      <Fold v-if="$store.state.asideWidth === '120px'"/>
      <Expand v-else />
    </el-icon>
    <el-tooltip effect="dark" content="刷新" placement="bottom">
      <el-icon class="icon-btn" @click="handleRefresh"><Refresh /></el-icon>
    </el-tooltip>
    <el-tooltip effect="dark" content="搜索" placement="bottom"
                v-if="!showSearch">
      <el-icon class="icon-btn" @click="handleSearch">
        <Search />
      </el-icon>
    </el-tooltip>
    <el-input
        v-model="key"
        placeholder="Every day is great."
        class="ml-auto animated fadeInLeft"
        clearable v-else
    >
      <template #append>
          <el-button :icon="Search" size="large" @click="search"></el-button>
      </template>
    </el-input>
    <div class="ml-auto flex items-center">
      <el-dropdown>
          <el-icon class="icon-btn">
            <font-awesome-icon :icon="['far', 'envelope']" size="lg"/>
          </el-icon>
        <template #dropdown>
          <el-dropdown-menu v-if="!getToken()">
            <el-dropdown-item class="d-item" @click="$router.push('/login')">
              登录账户
            </el-dropdown-item>
          </el-dropdown-menu>
          <el-dropdown-menu v-else>
            <el-dropdown-item class="d-item" @click="MessageOpen">
              我的消息
            </el-dropdown-item>
            <el-dropdown-item class="d-item">收到的赞</el-dropdown-item>
            <el-dropdown-item class="d-item">系统消息</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <el-dropdown>
        <el-icon class="icon-btn">
          <font-awesome-icon :icon="['far', 'square-plus']" size="lg"/>
        </el-icon>
        <template #dropdown>
          <el-dropdown-menu v-if="!getToken()">
            <el-dropdown-item class="d-item" @click="$router.push('/login')">
              登录账户
            </el-dropdown-item>
          </el-dropdown-menu>
          <el-dropdown-menu v-else>
            <el-dropdown-item class="d-item" @click="UploadOpen">
              发布视频
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <el-dropdown class="dropdown">
        <span class="mr-2 flex items-center" style="outline: none;">
          <el-avatar :size="40" :src="$store.state.user.avatar === ''?DefaultAvatar:$store.state.user.avatar" />
        </span>
        <template #dropdown>
        <el-dropdown-menu v-if="!getToken()">
          <el-dropdown-item class="d-item" @click="$router.push('/login')">
            登录账户
          </el-dropdown-item>
        </el-dropdown-menu>
          <el-dropdown-menu v-else>
            <el-dropdown-item class="d-item" @click="rePasswordFormOpen">
              修改密码
            </el-dropdown-item>
            <el-dropdown-item class="d-item" @click="handleLogout">
              退出登录
            </el-dropdown-item>
          </el-dropdown-menu>
      </template>
      </el-dropdown>
    </div>
  </div>

  <el-dialog
      v-model="rePasswordForm"
      title="修改密码"
      :close-on-click-modal="false"
      :show-close="false"
      width="40%"
  >
    <el-form :model="form" ref="formRef" :rules="rules">
      <el-form-item prop="oldPassword" label="旧密码"
                    :label-width="formLabelWidth">
        <el-input v-model="form.oldPassword" />
      </el-form-item>
      <el-form-item
          prop="newPassword"
          label="新密码"
          :label-width="formLabelWidth"
      >
        <el-input v-model="form.newPassword" />
      </el-form-item>
      <el-form-item
          prop="enterPassword"
          label="确认密码"
          :label-width="formLabelWidth"
      >
        <el-input v-model="form.enterPassword" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="rePasswordFormClose">取消</el-button>
        <el-button type="primary" @click="onSubmit">确认</el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog
      v-model="Message"
      title="我的消息"
      :close-on-click-modal="false"
      :modal="false" id="modelDialog"
      style="pointer-events: auto;"
      class="dialog"
  >
    <template #header>
        <div class="text-xl font-bold">我的消息</div>
    </template>
      <el-row class="rounded-lg border h-[400px]">
        <el-col :span="8" class="bg-light-50 rounded-l-lg border">
          <m-menu />
        </el-col>
        <el-col :span="16" class="bg-light-50 rounded-r-lg border">
          <m-container />
        </el-col>
      </el-row>
  </el-dialog>

  <el-dialog
      v-model="Upload"
      title="发布视频"
      :close-on-click-modal="false"
      :modal="false" id="modelDialog"
      style="pointer-events: auto;"
      class="dialog"
  >
    <template #header>
      <div class="text-xl font-bold">发布视频</div>
    </template>
    <v-upload></v-upload>
  </el-dialog>
</template>

<style>
  .f-header {
    @apply flex items-center shadow-md fixed top-0 left-0 right-0 bg-light-50;
    height: 60px;
    z-index: 100;
    min-width: 1010px;
  }

  .logo{
    width: 130px;
    @apply flex justify-center items-center text-xl font-thin;
  }
  .icon-btn{
    @apply flex justify-center items-center;
    width: 55px;
    outline: none;
    height: 60px;
    cursor: pointer;
  }
  .icon-btn:hover{
    @apply bg-gray-300;
  }
  .f-header .dropdown{
    height: 60px;
    cursor: pointer;
    @apply flex justify-center items-center mx-5;
  }
  .el-input {
    width: 360px;
  }
  .d-item {
    @apply font-bold;
  }
  .dialog {
    border-radius: 12px;
    min-width: 650px;
    @apply bg-gray-100 border-2 shadow-md
  }
</style>
