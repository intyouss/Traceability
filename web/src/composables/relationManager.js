import {ref} from 'vue';
import {getFansList, getFocusList, relationAction} from '~/api/relation.js';
import {getToken} from '~/composables/auth.js';
import {notify} from '~/composables/util.js';
import {useRouter} from 'vue-router';


/**
 * 关系逻辑
 * @return {{IsFocus: Ref<Boolean>, handleRelation: (function(Number): void)}}
 */
export function useRelation() {
  const IsFocus = ref(false);
  const router = useRouter();
  const handleRelation = (userId) => {
    if (!getToken()) {
      notify('请先登录', 'warning');
      return router.push('/login').then(() => {});
    }
    let actionType;
    if (IsFocus.value) {
      actionType = 2;
    } else {
      actionType = 1;
    }
    relationAction(userId, actionType).then(() => {
      IsFocus.value = !IsFocus.value;
    });
  };
  const setFocus = (isFocus) => {
    console.log(isFocus);
    IsFocus.value = isFocus;
  };
  return {
    IsFocus,
    setFocus,
    handleRelation,
  };
}

export function useRelationList() {
  const OpenRelationList = ref(false);
  const OpenRelationType = ref(1);
  const RelationList = ref([]);
  const openRelationList = (userId, type) => {
    OpenRelationList.value = true;
    OpenRelationType.value = type;
    getRelationList(userId, type);
  };
  const closeRelationList = () => {
    OpenRelationList.value = false;
  };

  const getRelationList = (userId, type) => {
    switch (type) {
      case 1:
        getFocusList(userId).then((res) => {
          RelationList.value = res.data.users;
        });
        break;
      case 2:
        getFansList(userId).then((res) => {
          RelationList.value = res.data.users;
        });
        break;
      default:
        break;
    }
    OpenRelationType.value = type;
  };
  return {
    RelationList,
    getRelationList,
    OpenRelationList,
    OpenRelationType,
    openRelationList,
    closeRelationList,
  };
}
