import {ref} from 'vue';
import {useStore} from 'vuex';
import {notify} from '~/composables/util.js';
import {useRouter} from 'vue-router';
import {collectAction} from '~/api/collect.js';

export function useCollect(isCollect) {
  const IsCollected = ref(isCollect);
  const store = useStore();
  const router = useRouter();
  const handleCollectAction = (videoId) => {
    if (store.state.user.id) {
      // 点赞/取消点赞
      const actionType = IsCollected.value ? 2 : 1;
      collectAction(videoId, actionType).then(()=>{
        IsCollected.value = !IsCollected.value;
      });
    } else {
      notify('请先登录', 'warning');
      router.push('/login').then(() => {});
    }
  };
  return {
    IsCollected,
    handleCollectAction,
  };
}
