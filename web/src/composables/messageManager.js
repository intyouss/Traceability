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
  const Messages = ref([]);
  const Users = ref([]);
  const store = useStore();

  const sendMsg = (toUserId, content) => {
    sendMessage(toUserId, content).then((res) => {
      Messages.value.push(res.data.message);
    });
  };

  const getMsg = (toUserId) => {
    getMessageList(toUserId).then((res) => {
      Messages.value = res.data.messages;
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
