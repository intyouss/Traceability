import {reactive, ref} from 'vue';
import {getAuthUserSearch, getPublicUserSearch, logout} from '~/api/user.js';
import {confirm, notify} from '~/composables/util.js';
import {useRouter, useRoute} from 'vue-router';
import {useStore} from 'vuex';
import {getVideoSearch} from '~/api/video.js';
import {getToken} from '~/composables/auth.js';

export function useRePassword() {
  const rePasswordForm = ref(false);
  const formLabelWidth = '140px';
  const form = reactive({
    oldPassword: '',
    newPassword: '',
    enterPassword: '',
  });
  const rules = {
    oldPassword: [
      {required: true, message: '旧密码不能为空', trigger: 'blur'},
    ],
    newPassword: [
      {required: true, message: '新密码不能为空', trigger: 'blur'},
    ],
    enterPassword: [
      {required: true, message: '确认密码不能为空', trigger: 'blur'},
    ],
  };
  const formRef = ref(null);
  const loading = ref(false);
  const onSubmit = () => {
    formRef.value.validate((valid)=>{
      if (!valid) {
        return false;
      }
      loading.value = true;
      // rePassword(form).then((res)=>{
      //   notify('修改成功', 'success');
      //   store.dispatch('logout');
      //   router.push('/login');
      //   rePasswordFormClose();
      // }).finally(()=>{
      //   loading.value = false;
      // });
    });
  };
  const rePasswordFormOpen = () => rePasswordForm.value = true;
  const rePasswordFormClose = () => rePasswordForm.value = false;
  return {
    rePasswordForm,
    rePasswordFormOpen,
    rePasswordFormClose,
    formLabelWidth,
    form,
    rules,
    formRef,
    onSubmit,
  };
}

export function useLogout() {
  const router = useRouter();
  const route = useRoute();
  const store = useStore();
  function handleLogout() {
    confirm('确定退出登录吗?').then(()=>{
      logout()
          .finally(()=>{
            store.dispatch('logout');
            if (route.path !== '/') {
              router.push('/login');
            }
            notify('退出成功');
          });
    });
  }
  return {
    handleLogout,
  };
}

export function useSearch() {
  const getSearch = (tag, key) => {
    switch (tag) {
      case '综合':
        return getVideoSearch(key, 1).then((res)=>{
          return res.data.videos;
        });
      case '视频':
        return getVideoSearch(key, 2).then((res)=>{
          return res.data.videos;
        });
      case '用户':
        if (getToken()) {
          return getAuthUserSearch(key).then((res)=>{
            return res.data.users;
          });
        }
        return getPublicUserSearch(key).then((res)=>{
          return res.data.users;
        });
    }
  };
  return {
    getSearch,
  };
}
