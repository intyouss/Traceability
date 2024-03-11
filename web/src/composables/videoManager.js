import {useStore} from 'vuex';
import {ref} from 'vue';
import {getAuthVideo, getIndexVideo, getUserVideoList} from '~/api/video.js';
import {getLikeList} from '~/api/like.js';
import {getCollectList} from '~/api/collect.js';
import {useRouter} from 'vue-router';


/**
 * 根据当前用户获取视频列表
 * @return {{getVideos: (function(): void), Videos: Ref<[]>}}
 */
export function useVideoByOwner() {
  const store = useStore();
  const Videos = ref([]);
  const getPublishVideos = () => {
    getUserVideoList(store.state.user.id).then((res)=>{
      Videos.value = res.data.videos;
    });
  };
  const getLikeVideos = () => {
    getLikeList(store.state.user.id).then((res)=>{
      Videos.value = res.data.videos;
    });
  };
  const getCollectVideos = () => {
    getCollectList(store.state.user.id).then((res)=>{
      Videos.value = res.data.videos;
    });
  };

  const getVideos = (type) => {
    // 根据路由获取视频列表
    switch (type) {
      case '作品':
        return getPublishVideos();
      case '喜爱':
        return getLikeVideos();
      case '收藏':
        return getCollectVideos();
    }
  };
  return {
    Videos,
    getVideos,
  };
}

/**
 * 根据其他用户获取视频列表
 * @param {Number} userId 用户id
 * @return {{getVideos: (function(): void), Videos: Ref<[]>}}
 */
export function useVideoByOther(userId) {
  const Videos = ref([]);
  const getPublishVideos = () => {
    getUserVideoList(userId).then((res)=>{
      Videos.value = res.data.videos;
    });
  };
  const getLikeVideos = () => {
    getLikeList(userId).then((res)=>{
      Videos.value = res.data.videos;
    });
  };
  const getCollectVideos = () => {
    getCollectList(userId).then((res)=>{
      Videos.value = res.data.videos;
    });
  };

  const getVideos = (type) => {
    // 根据路由获取视频列表
    switch (type) {
      case '作品':
        return getPublishVideos();
      case '喜爱':
        return getLikeVideos();
      case '收藏':
        return getCollectVideos();
    }
  };
  return {
    Videos,
    getVideos,
  };
}

/**
 * 根据路由获取视频列表
 * @return {{getVideos: (function(): void), Videos: Ref<[]>}}
 */
export function useVideoByPage() {
  const router = useRouter();
  const routerName = router.currentRoute.value.name;
  const Videos = ref([]);
  const getIndexVideos = () => {
    getIndexVideo(1).then((res)=>{
      Videos.value = res.data.videos;
      console.log(Videos.value);
    });
  };
  const getFocusVideos = () => {
    getAuthVideo(2).then((res)=>{
      Videos.value = res.data.videos;
      console.log(Videos.value);
    });
  };
  const getFriendVideos = () => {
    getAuthVideo(3).then((res)=>{
      Videos.value = res.data.videos;
      console.log(Videos.value);
    });
  };
  const getRecommendVideos = () => {
    getAuthVideo(4).then((res)=>{
      Videos.value = res.data.videos;
      console.log(Videos.value);
    });
  };

  const getVideos = () => {
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
    getVideos,
  };
}
