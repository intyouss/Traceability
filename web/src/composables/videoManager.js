import {useStore} from 'vuex';
import {ref} from 'vue';
import {
  abolishVideoUpload,
  getAuthVideo,
  getIndexVideo,
  getUserVideoList,
  publishVideo,
  uploadImage,
  uploadVideo,
} from '~/api/video.js';
import {getLikeList} from '~/api/like.js';
import {getCollectList} from '~/api/collect.js';

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
 * @return {{getVideos: (function(): void), Videos: Ref<[]>}}
 */
export function useVideoByOther() {
  const Videos = ref([]);
  const getPublishVideos = (userId) => {
    getUserVideoList(userId).then((res)=>{
      Videos.value = res.data.videos;
      console.log(Videos.value);
    });
  };
  const getLikeVideos = (userId) => {
    getLikeList(userId).then((res)=>{
      Videos.value = res.data.videos;
    });
  };
  const getCollectVideos = (userId) => {
    getCollectList(userId).then((res)=>{
      Videos.value = res.data.videos;
    });
  };

  const getVideos = (type, userId) => {
    // 根据路由获取视频列表
    switch (type) {
      case '作品':
        return getPublishVideos(userId);
      case '喜爱':
        return getLikeVideos(userId);
      case '收藏':
        return getCollectVideos(userId);
    }
  };
  return {
    Videos,
    getVideos,
    getPublishVideos,
    getLikeVideos,
    getCollectVideos,
  };
}

export function useVideoByPage() {
  const Videos = ref([]);
  const getIndexVideos = () => {
    getIndexVideo(1).then((res)=>{
      Videos.value = res.data.videos;
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

  // const getVideos = (userId=0) => {
  //   // 根据路由获取视频列表
  //   switch (routerName) {
  //     case 'index':
  //       return getIndexVideos();
  //     case 'focus':
  //       return getFocusVideos(userId);
  //     case 'recommend':
  //       return getRecommendVideos();
  //     case 'friend':
  //       return getFriendVideos();
  //   }
  // };
  return {
    Videos,
    getIndexVideos,
    getFocusVideos,
    getFriendVideos,
    getRecommendVideos,
  };
}

export function useVideoUploadDialog() {
  const Upload = ref(false);
  const UploadOpen = () => Upload.value = true;
  const UploadClose = () => Upload.value = false;
  return {
    Upload,
    UploadClose,
    UploadOpen,
  };
}

export function useVideoUpload() {
  const PlayUrl = ref('');
  const CoverUrl = ref('');
  const videoUpload = (title, data) => {
    uploadVideo(title, data).then((res)=>{
      PlayUrl.value = res.data.play_url;
    });
  };

  const imageUpload = (title, data) => {
    uploadImage(title, data).then((res)=>{
      CoverUrl.value = res.data.cover_image_url;
    });
  };

  const videoPublish = (title, playUrl, coverUrl) => {
    publishVideo(title, playUrl, coverUrl).then(()=>{});
  };

  const uploadAbolish = (title, type) => {
    abolishVideoUpload(title, type).then(()=>{});
  };
  return {
    PlayUrl,
    CoverUrl,
    uploadAbolish,
    videoPublish,
    videoUpload,
    imageUpload,
  };
}
