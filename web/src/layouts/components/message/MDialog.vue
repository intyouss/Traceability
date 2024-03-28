<script setup>

import MContainer from '~/layouts/components/message/MContainer.vue';
import MMenu from '~/layouts/components/message/MMenu.vue';
import {ref, watch} from 'vue';
import {useMessage} from '~/composables/messageManager.js';
const {
  Messages,
  getOpenUser,
  getMsg,
  PreTime,
  Users,
  sendMsg,
  increaseOpenUser,
} = useMessage();

const props = defineProps({
  message: Boolean,
  user: Object,
  privateEmile: Boolean,
  openContainer: Boolean,
});


const activeMessages = ref(Messages);
watch(() => Messages, (value) => {
  activeMessages.value = value;
});

const ContainerOpen = ref(false);
const User = ref({});
const message = ref(props.message);
watch(() => props.message, (value) => {
  if (props.privateEmile && value) {
    increaseOpenUser(props.user.id);
    getMsg(props.user.id);
    ContainerOpen.value = true;
    User.value = props.user;
  }
  getOpenUser();
  message.value = value;
});

const OldUserId = ref(0);
const timer = ref([]);
const emit = defineEmits(['messageClose']);
const handleClose = () => {
  emit('messageClose', 'close');
  ContainerOpen.value = false;
  clearInterval(timer.value[OldUserId.value]);
  timer.value = [];
};
const handleSelect = (index) => {
  if (OldUserId.value !== 0) {
    clearInterval(timer.value[OldUserId.value]);
    timer.value[OldUserId.value] = null;
  }
  ContainerOpen.value = true;
  console.log(index);
  getMsg(Users.value[index].id, PreTime.value[Users.value[index].id]);
  User.value = Users.value[index];
  OldUserId.value = Users.value[index].id;
  timer.value[Users.value[index].id] = setInterval(() => {
    getMsg(Users.value[index].id, PreTime.value[Users.value[index].id]);
  }, 3000);
};

const deleteOpUser = (id) => {
  Users.value.forEach((item, index) => {
    if (item.id === id) {
      Users.value.splice(index, 1);
    }
  });
  ContainerOpen.value = false;
};
</script>

<template>
  <el-dialog
      v-model="message"
      title="我的消息"
      :close-on-click-modal="false"
      :before-close="handleClose"
      :modal="false" id="modelDialog"
      style="pointer-events: auto;"
      class="trerg"
  >
    <template #header>
      <div class="text-xl font-bold">我的消息</div>
    </template>
    <el-row class="rounded-lg border" style="height: 470px;">
      <el-col :span="6" class="bg-light-50 rounded-l-lg border">
        <m-menu @select="handleSelect" :users="Users"/>
      </el-col>
      <el-col :span="18" class="bg-light-50 rounded-r-lg border">
        <m-container
            :message="activeMessages[User.id]"
            :user="User"
            :open="ContainerOpen"
            :empty="Users.length === 0"
            :send-message="sendMsg"
            @deleteOpenUser="deleteOpUser"
        />
      </el-col>
    </el-row>
  </el-dialog>
</template>

<style scoped>
</style>
<style>
.dialog .el-dialog__body {
  padding: 0 10px 20px 10px;
}
.trerg {
  border-radius: 12px;
  min-width: 700px;
  @apply bg-gray-100 border-2 shadow-md
}
.el-overlay-dialog {
  overflow: hidden;
}
</style>
