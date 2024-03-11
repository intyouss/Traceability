import {getInfo} from '~/api/user.js';
import {useStore} from 'vuex';
import {ref} from 'vue';
import {getFocusList} from '~/api/relation.js';

export function useUserByOwner() {
  const store = useStore();
  const getUserInfo = () => {
    getInfo(store.state.user.id).then((res) => {
      store.state.user = res.data.user;
    });
  };

  const Users = ref([]);
  const getUserFocusList = () => {
    return getFocusList(store.state.user.id).then((res)=>{
      Users.value = res.data.users;
    });
  };

  return {
    Users,
    getUserFocusList,
    getUserInfo,
  };
}

export function useUserByOther(userId) {
  const User = ref({});
  const getUserInfo = () => {
    getInfo(userId).then((res) => {
      User.value = res.data.user;
    });
  };

  return {
    User,
    getUserInfo,
  };
}
