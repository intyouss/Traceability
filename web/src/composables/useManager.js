import {reactive, ref} from 'vue';
import {getUserSearch, logout} from '~/api/user.js';
import {confirm, notify} from '~/composables/util.js';
import {useRouter} from 'vue-router';
import {useStore} from 'vuex';
import {
  getIndexVideo,
  getAuthVideo, getVideoSearch, getUserVideoList,
} from '~/api/video.js';
import {getLikeList} from '~/api/like.js';

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
  const store = useStore();
  function handleLogout() {
    confirm('确定退出登录吗?').then(()=>{
      logout()
          .finally(()=>{
            store.dispatch('logout');
            router.push('/login');
            notify('退出成功');
          });
    });
  }
  return {
    handleLogout,
  };
}


export function useMessage() {
  const Message = ref(false);
  const MessageOpen = () =>Message.value = true;
  return {
    Message,
    MessageOpen,
  };
}

export function useVideoUpload() {
  const Upload = ref(false);
  const UploadOpen = () => Upload.value = true;
  return {
    Upload,
    UploadOpen,
  };
}

export function useVideos() {
  const router = useRouter();
  const routerName = router.currentRoute.value.name;
  const Videos = ref([]);
  const getIndexVideos = () => {
    getIndexVideo({
      'type': 1,
      'latest_time': 0,
    }).then((res)=>{
      Videos.value = res.data.videos;
      console.log(Videos.value);
    });
  };
  const getFocusVideos = () => {
    getAuthVideo({
      'type': 2,
      'latest_time': 0,
    }).then((res)=>{
      Videos.value = res.data.videos;
      console.log(Videos.value);
    });
  };
  const getFriendVideos = () => {
    getAuthVideo({
      'type': 3,
      'latest_time': 0,
    }).then((res)=>{
      Videos.value = res.data.videos;
      console.log(Videos.value);
    });
  };
  const getRecommendVideos = () => {
    getAuthVideo({
      'type': 4,
      'latest_time': 0,
    }).then((res)=>{
      Videos.value = res.data.videos;
      console.log(Videos.value);
    });
  };

  const getVideoList = () => {
    // 根据路由获取视频列表
    switch (routerName) {
      case 'index':
        return getIndexVideos();
      case 'focus':
        return getFocusVideos();
      case 'recommend':
        return getRecommendVideos();
      case 'friend':
        return getFriendVideos();
    }
  };
  return {
    Videos,
    getVideoList,
  };
}

export function useOwnerVideos() {
  const store = useStore();
  const Videos = ref([]);
  const getSelfVideos = () => {
    getUserVideoList({
      'user_id': store.state.user.id,
    }).then((res)=>{
      Videos.value = res.data.videos;
    });
  };
  const getLikeVideos = () => {
    getLikeList({
      'user_id': store.state.user.id,
    }).then((res)=>{
      Videos.value = res.data.videos;
    });
  };

  const getVideoList = (type) => {
    // 根据路由获取视频列表
    switch (type) {
      case '作品':
        return getSelfVideos();
      case '喜爱':
        return getLikeVideos();
      case '收藏':
        return Videos.value = [];
    }
  };
  return {
    Videos,
    getVideoList,
  };
}

export function useVideoAndUserSearch() {
  const getSearch = (tag, key, page=1, limit=10) => {
    switch (tag) {
      case '综合':
        return getVideoSearch({
          'key': key,
          'type': 1,
        }).then((res)=>{
          return res.data.videos;
        });
      case '视频':
        return getVideoSearch({
          'key': key,
          'type': 2,
        }).then((res)=>{
          return res.data.videos;
        });
      case '用户':
        return getUserSearch({
          'key': key,
          'page': page,
          'limit': limit,
        }).then((res)=>{
          return res.data.users;
        });
    }
  };
  return {
    getSearch,
  };
}
