<script setup>
import {ref} from 'vue';
import {uploadVideo} from '~/api/video.js';

const fileList = ref([]);
const uploadRequest = (options) => {
  const file = options.file;
  const submitData = new FormData();
  submitData.append('file', file);
  uploadVideo(submitData).then((res) => {
    const url = res.data.data;
    fileList.value.push({
      name: file.name,
      url,
    });
  });
};
</script>
<template>
<div>
  <el-upload
      class="upload-demo"
      drag
      action="#"
      multiple
  >
    <video v-if="videoForm.showVideoPath !='' && !videoFlag"
           v-bind:src="videoForm.showVideoPath"
           class="avatar video-avatar"
           controls="controls">
      您的浏览器不支持视频播放
    </video>
    <i v-else-if="videoForm.showVideoPath =='' && !videoFlag"
       class="el-icon-plus avatar-uploader-icon"></i>
    <el-progress v-if="videoFlag == true"
                 type="circle"
                 v-bind:percentage="videoUploadPercent"
                 style="margin-top:7px;"></el-progress>
<!--    <el-icon class="el-icon&#45;&#45;upload"><upload-filled /></el-icon>-->
<!--    <div class="el-upload__text">-->
<!--      拖动视频或者 <em>点击上传</em>-->
<!--    </div>-->
<!--    <template #tip>-->
<!--      <div class="el-upload__tip">-->
<!--        视频文件限制MP4格式,并不能超过50M-->
<!--      </div>-->
<!--    </template>-->
  </el-upload>
</div>
</template>

<style scoped>

</style>
