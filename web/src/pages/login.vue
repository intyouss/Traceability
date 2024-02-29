<template>
    <el-row class="login-container">
      <el-col :lg="16" :md="12" class="left">
        <div>
          <div class="font-bold text-5xl text-light-50 mb-4">欢迎光临</div>
          <div class="text-gray-200 text-sm">你的出现 给予我最大的信心~~</div>
        </div>
      </el-col>
      <el-col :lg="8" :md="12" class="right">
        <h2 class="font-bold text-3xl text-gray-800">欢迎回来</h2>
        <div class="flex items-center justify-center my-5 text-gray-300 space-x-2">
          <span class="h-[1px] w-16 bg-gray-200"></span>
          <span>账号密码登录</span>
          <span class="h-[1px] w-16 bg-gray-200"></span>
        </div>
        <el-form ref="formRef" :rules="rules" :model="form" class="w-[250px]">
          <el-form-item prop="username">
            <el-input v-model="form.username" placeholder="请输入用户名">
              <template #prefix>
                <el-icon><user /></el-icon>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input type="password" v-model="form.password" placeholder="请输入密码" show-password>
              <template #prefix>
                <el-icon><lock /></el-icon>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button round color="#626aef" class="w-[250px]" type="primary" @click="onSubmit" :loading="loading">登 录</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
</template>
<script setup>
import {reactive, ref, onMounted, onBeforeUnmount} from 'vue';
import {User, Lock} from '@element-plus/icons-vue';
import {useRouter} from 'vue-router';
import {useStore} from 'vuex';
import {notify} from '~/composables/util.js';
const router = useRouter();
const store = useStore();

// do not use same name with ref
const form = reactive({
  username: '',
  password: '',
});

const rules = {
  username: [
    {required: true, message: '用户名不能为空', trigger: 'blur'},
  ],
  password: [
    {required: true, message: '密码不能为空', trigger: 'blur'},
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
    store.dispatch('login', form).then((res)=>{
      notify('登陆成功', 'success');
      router.push('/');
    }).finally(()=>{
      loading.value = false;
    });
  });
};

// 监听回车事件
function onKeyup(e) {
  if (e.key === 'Enter') onSubmit();
}

// 添加键盘监听
onMounted(()=>{
  document.addEventListener('keyup', onKeyup);
});
onBeforeUnmount(()=>{
  document.removeEventListener('keyup', onKeyup);
});
</script>

<style>
.login-container {
  @apply min-h-screen bg-indigo-500;
}
.login-container .left, .login-container .right {
  @apply flex items-center justify-center;
}
.login-container .right {
  @apply bg-light-50 flex-col;
}
</style>
