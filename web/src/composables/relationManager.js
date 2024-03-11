import {ref} from 'vue';
import {relationAction} from '~/api/relation.js';
import {getToken} from '~/composables/auth.js';
import {notify} from '~/composables/util.js';
import {useRouter} from 'vue-router';


/**
 * 关系逻辑
 * @param {Boolean} isFocus
 * @return {{IsFocus: Ref<Boolean>, handleRelation: (function(Number): void)}}
 */
export function useRelation(isFocus) {
  const IsFocus = ref(isFocus);
  const router = useRouter();

  const actionType = IsFocus.value ? 2 : 1;
  const handleRelation = (userId) => {
    if (!getToken()) {
      notify('请先登录', 'warning');
      return router.push('/login').then(() => {});
    }
    relationAction(userId, actionType).then(() => {
      IsFocus.value = !IsFocus.value;
    });
  };
  return {
    IsFocus,
    handleRelation,
  };
}
