import {ref} from 'vue';
import {
  addOpenUser,
  deleteOpenUser,
  getMessageList,
  getOpenUsers,
  sendMessage,
} from '~/api/message.js';
import {useStore} from 'vuex';

export function useMessage() {
  const Messages = ref({});
  const PreTime = ref({});
  const Users = ref([]);
  const store = useStore();

  const sendMsg = (toUserId, content) => {
    sendMessage(toUserId, content).then((res) => {
      if (!Messages.value[toUserId]) {
        Messages.value[toUserId] = [];
      }
      Messages.value[toUserId].push(res.data.message);
      PreTime.value[toUserId] = res.data.message.created_at;
    });
  };

  const getMsg = (toUserId, preMsgTime='0') => {
    getMessageList(toUserId, preMsgTime).then((res) => {
      PreTime.value[toUserId] = res.data.pre_msg_time;
      if (res.data.messages.length === 0) {
        return;
      }
      if (!Messages.value[toUserId]) {
        Messages.value[toUserId] = [];
      }
      Messages.value[toUserId] = Messages.value[toUserId].concat(res.data.messages);
    });
  };

  const getOpenUser = () => {
    getOpenUsers(store.state.user.id).then((res) => {
      Users.value = res.data.users;
    });
  };

  const delOpenUser = (openUserId) => {
    deleteOpenUser(openUserId).then(() => {});
  };

  const increaseOpenUser = (openUserId) =>{
    addOpenUser(openUserId).then((res) => {
      if (!res.data) {
        return;
      }
      Users.value.push(res.data.user);
    });
  };
  return {
    Messages,
    sendMsg,
    getMsg,
    Users,
    PreTime,
    delOpenUser,
    increaseOpenUser,
    getOpenUser,
  };
}

export function useMessageDialog() {
  const Message = ref(false);
  const MessageOpen = () =>{
    Message.value = true;
  };
  const MessageClose = () =>{
    Message.value = false;
  };
  return {
    Message,
    MessageOpen,
    MessageClose,
  };
}
