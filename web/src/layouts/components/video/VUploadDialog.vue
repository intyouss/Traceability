<script setup>
import {reactive, ref, watch} from 'vue';
import {useVideoUpload} from '~/composables/videoManager.js';
import {notify} from '~/composables/util.js';

const props = defineProps({
  open: Boolean,
  close: Function,
});

const Open = ref(props.open);
watch(() => props.open, (value) => {
  Open.value = value;
});

const fileListByVideo = ref([]);
const fileListByImage = ref([]);

const form = reactive({
  title: '',
});

const {
  PlayUrl,
  CoverUrl,
  uploadAbolish,
  videoPublish,
  videoUpload,
  imageUpload,
} = useVideoUpload();

const ImageLoading = ref(false);
const VideoLoading = ref(false);

const HttpRequestByImage = (param) => {
  ImageLoading.value = true;
  const coverImageFile = param.file;
  imageUpload(form.title, coverImageFile);
  ImageLoading.value = false;
};

const HttpRequestByVideo = (param) => {
  VideoLoading.value = true;
  const file = param.file;
  videoUpload(form.title, file);
  VideoLoading.value = false;
};

const Publish = (title, playUrl, coverUrl) => {
  if (ImageLoading.value || VideoLoading.value) {
    notify('正在上传中，请稍后', 'warning');
  } else if (title === '') {
    notify('请填写作品描述', 'warning');
  } else if (playUrl === '' && coverUrl === '') {
    notify('请上传视频和封面', 'warning');
  } else if (playUrl === '') {
    notify('请上传视频', 'warning');
  } else if (coverUrl === '') {
    notify('请上传封面', 'warning');
  } else {
    videoPublish(title, playUrl, coverUrl);
    notify('发布成功', 'success');
    form.title = '';
    props.close();
  }
};

const Abolish = (title) => {
  if (ImageLoading.value || VideoLoading.value) {
    notify('正在上传中，请稍后', 'warning');
  } else if (title === '') {
    props.close();
  } else if (PlayUrl === '' && CoverUrl === '') {
    props.close();
    form.title = '';
  } else if (PlayUrl === '') {
    uploadAbolish(title, 2);
    props.close();
    form.title = '';
  } else if (CoverUrl === '') {
    uploadAbolish(title, 3);
    form.title = '';
    props.close();
  } else {
    uploadAbolish(title, 1);
    form.title = '';
    props.close();
  }
};

const imageListChange = (file, fileList) => {
  fileListByImage.value = fileList;
};

const videoListChange = (file, fileList) => {
  fileListByVideo.value = fileList;
};
</script>
<template>
  <div>
    <el-dialog
        v-model="Open"
        title="发布视频"
        :close-on-click-modal="false"
        :modal="false"
        :before-close="props.close"
        style="pointer-events: auto;"
        :show-close="false"
        class="nytr"
    >
      <template #header>
        <div class="text-xl font-bold">发布视频</div>
      </template>
      <div class="bg-gray-200 p-2 rounded-xl">
        <el-form>
          <el-form-item>
            <p class="hergwwf">作品描述</p>
            <el-input
                type="textarea"
                v-model="form.title"
                placeholder="添加作品简介"
                maxlength="100"
            />
          </el-form-item>
        </el-form>
        <div v-if="form.title">
          <p class="hergwwf">作品上传</p>
          <el-row :gutter="20">
            <el-col :span="8">
              <div v-loading="ImageLoading">
                <template v-if="fileListByImage.length === 0">
                  <el-upload
                      drag
                      :limit="1"
                      action=""
                      :on-change="imageListChange"
                      :file-list="fileListByImage"
                      :http-request="HttpRequestByImage"
                      :show-file-list="false"
                      accept="image/jpeg,image/png"
                  >
                    <el-icon class="el-icon--upload">
                      <upload-filled/>
                    </el-icon>
                    <div class="el-upload__text">
                      拖动图片或者 <em>点击上传</em>
                    </div>
                    <template #tip>
                      <div class="el-upload__tip">
                        封面文件限制jpg和png格式
                      </div>
                    </template>
                  </el-upload>
                </template>
                <template v-else>
                  <img
                      class="rounded-2xl border-2 border-gray-500"
                      style="width: 100%; height: 200px;"
                      :src="CoverUrl"
                      alt=""
                  />
                </template>
              </div>
            </el-col>
            <el-col :span="16">
              <div v-loading="VideoLoading">
                <template v-if="fileListByVideo.length === 0">
                  <el-upload
                      drag
                      :limit="1"
                      action=""
                      :show-file-list="false"
                      :on-change="videoListChange"
                      :file-list="fileListByVideo"
                      :http-request="HttpRequestByVideo"
                      accept="video/mp4"
                  >
                    <el-icon class="el-icon--upload">
                      <upload-filled/>
                    </el-icon>
                    <div class="el-upload__text">
                      拖动视频或者 <em>点击上传</em>
                    </div>
                    <template #tip>
                      <div class="el-upload__tip">
                        视频文件限制MP4格式,并不能超过50M
                      </div>
                    </template>
                  </el-upload>
                </template>
                <template v-else>
                  <video
                      class="rounded-2xl border-2 border-gray-500"
                      style="width: 100%; height: 200px;"
                      :src="PlayUrl"
                      controls
                  />
                </template>
              </div>
            </el-col>
          </el-row>
        </div>
      </div>
      <div class="flex item-center justify-center">
        <el-button type="danger" class="mt-4" @click="Publish(form.title, PlayUrl, CoverUrl)">发布</el-button>
        <el-button type="info" class="mt-4" @click="Abolish(form.title)">取消</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<style>
.nytr {
  border-radius: 12px;
  min-width: 600px;
  @apply bg-gray-100 border-2 shadow-md
}

.el-textarea__inner {
  resize: none;
}
</style>

<style scoped>
.hergwwf {
  display: flex;
  align-items: center;
  font-size: 14px;
  line-height: 20px;
  font-weight: bold;
  color: #404346;
  margin-bottom: 10px;
}
</style>
