import {abolishAvatar, getInfo, updateUser, uploadAvatar} from '~/api/user.js';
import {useStore} from 'vuex';
import {reactive, ref} from 'vue';
import {getFocusList} from '~/api/relation.js';
import {notify} from '~/composables/util.js';

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
      console.log(Users.value, '11111');
    });
  };

  return {
    Users,
    getUserFocusList,
    getUserInfo,
  };
}

export function useUserByOther() {
  const User = ref({});
  const getUserInfo = (userId) => {
    getInfo(userId).then((res) => {
      User.value = res.data.user;
    });
  };

  return {
    User,
    getUserInfo,
  };
}

export function useSetUserInfo() {
  const form = reactive({
    signature: '',
    avatarUrl: '',
  });
  const rules = {
    signature: [
      {max: 100, message: '个性签名不能超过100个字符', trigger: 'blur'},
    ],
  };
  const formRef = ref(null);
  const loading = ref(false);

  const avatarUpload = (data) => {
    return uploadAvatar(data).then((res)=>{
      form.avatarUrl = res.data.avatar_url;
    });
  };

  const avatarAbolish = (userId) => {
    abolishAvatar(userId).then(()=>{});
  };
  const onSubmit = (signature='', avatarUrl='') => {
    formRef.value.validate((valid)=>{
      if (!valid) {
        return false;
      }
      loading.value = true;
      updateUser({signature: signature, avatarUrl: avatarUrl}).then(()=>{
        notify('修改成功', 'success');
      }).finally(()=>{
        loading.value = false;
      });
    });
  };
  return {
    loading,
    form,
    rules,
    formRef,
    onSubmit,
    avatarUpload,
    avatarAbolish,
  };
}

export function useInfoForm() {
  const InfoForm = ref(false);
  const setInfoFormOpen = () => InfoForm.value = true;
  const setInfoFormClose = () => InfoForm.value = false;
  return {
    InfoForm,
    setInfoFormOpen,
    setInfoFormClose,
  };
}
