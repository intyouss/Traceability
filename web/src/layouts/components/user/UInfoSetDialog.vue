<script setup>
import {useSetUserInfo, useUserByOwner} from '~/composables/userManager.js';
import {ref, watch} from 'vue';
import DefaultAvatar from '~/assets/icon/default_avatar.jpg';

const props = defineProps({
  infoForm: Boolean,
  formClose: Function,
  avatarUrl: String,
  userId: Number,
  signature: String,
});
const {
  getUserInfo,
} = useUserByOwner();

const InfoForm = ref(props.infoForm);
watch(() => props.infoForm, (val) => {
  InfoForm.value = val;
});

const {
  form,
  rules,
  formRef,
  onSubmit,
  loading,
  avatarUpload,
  avatarAbolish,
} = useSetUserInfo();

const submit = () => {
  if (form.avatarUrl === '' && form.signature === '') {
    props.formClose();
    return;
  }
  onSubmit(form.signature, form.avatarUrl);
  getUserInfo();
  form.avatarUrl = '';
  form.signature = '';
  props.formClose();
};

const HttpRequestByAvatar = (param) => {
  const coverImageFile = param.file;
  avatarUpload(coverImageFile);
};

const AvatarUrl = ref(props.avatarUrl);

watch(() => form.avatarUrl, (val) => {
  AvatarUrl.value = val;
});

const setAvatar = () => {
  return AvatarUrl.value === '' ? DefaultAvatar : AvatarUrl.value;
};

const CloseSetInfo = () => {
  if (form.avatarUrl !== '') {
    avatarAbolish(props.userId);
    form.avatarUrl = '';
  }
  if (form.signature !== '') {
    form.signature = '';
  }
  props.formClose();
};
</script>
<template>
  <el-dialog
      v-model="InfoForm"
      title="编辑资料"
      :close-on-click-modal="false"
      :show-close="false"
      width="40%"
      class="gwgwef"
  >
    <div class="pb-12">
      <div class="flex justify-center text-center pb-4">

            <el-upload
                drag
                :limit="1"
                action=""
                :http-request="HttpRequestByAvatar"
                :show-file-list="false"
                accept="image/jpeg,image/png"
                class="kkk"
            >
                <div class="fwefwa">
                  <div  :style="{backgroundImage: 'url(' + setAvatar() + ')', backgroundSize:'100% 100%', backgroundRepeat:'no-repeat'}">
                    <div class="gewfqw" style=" background-color: rgba(200, 200, 200, 0.5);">
                      <font-awesome-icon
                          :icon="['fas', 'camera']"
                          class="camera-icon"
                      />
                    </div>
                  </div>
                </div>
            </el-upload>

      </div>
      <div class="flex justify-center text-center">头像</div>
    </div>
    <el-form :model="form" ref="formRef" :rules="rules">
      <el-form-item prop="signature" label="个性签名">
        <el-input
            type="textarea"
            v-model="form.signature"
            :placeholder="props.signature"
            :rows="4"
            maxlength="100"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="flex justify-center">
        <el-button @click="CloseSetInfo">取消</el-button>
        <el-button
            type="primary"
            @click="submit"
            :loading="loading"
        >
          确认
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<style scoped>
:deep(.kkk .el-upload-dragger) {
  background-color: rgb(128, 128, 128, 0.05);
  border: none;
  padding: initial;
  border-radius: 50%;
}

:deep(.kkk .el-upload){
  --el-upload-dragger-padding-horizontal: 40px;
  --el-upload-dragger-padding-vertical: 40px;
  border-radius: 50%;
}

:deep(.kkk .el-icon) {
  width: 150px;
}
.kkk {
  @apply rounded-1/2 border-2 border-gray-300;
}
.fwefwa {
  width: 150px;
  height: 150px;
}
.camera-icon {
  color: white;
  width: 80px;
  height: 80px;
  margin-top: 30px;
}
.gewfqw {
  width: 150px;
  height: 150px;
}
</style>
<style>
.gwgwef {
  border-radius: 12px;
  min-width: 400px;
  @apply bg-gray-100 border-2 shadow-md
}
</style>
