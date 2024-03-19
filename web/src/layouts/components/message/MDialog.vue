<script setup>

import MContainer from '~/layouts/components/message/MContainer.vue';
import MMenu from '~/layouts/components/message/MMenu.vue';
import {ref, watch} from 'vue';
import {useMessage} from '~/composables/messageManager.js';
const {
  Messages,
  getOpenUser,
  getMsg,
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

const emit = defineEmits(['messageClose']);
const handleClose = () => {
  emit('messageClose', 'close');
};


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
  if (!props.privateEmile && value) {
    getOpenUser();
  }
  message.value = value;
});

const handleSelect = (index) => {
  ContainerOpen.value = true;
  getMsg(Users.value[index].id);
  User.value = Users.value[index];
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
            :message="activeMessages"
            :user="User"
            :open="ContainerOpen"
            :empty="!Users"
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
</style>
