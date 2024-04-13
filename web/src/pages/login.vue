<script setup>
import { reactive, ref, onMounted, onBeforeUnmount } from 'vue'
import { User, Lock } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { notify } from '~/composables/util.js'

const router = useRouter()
const store = useStore()

// do not use same name with ref
const form = reactive({
  username: '',
  password: ''
})

const registerForm = reactive({
  username: '',
  password: '',
  repeatPassword: ''
})

const registerRules = {
  username: [
    { required: true, message: '用户名不能为空', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '密码不能为空', trigger: 'blur' }
  ],
  repeatPassword: [
    {
      trigger: 'blur',
      validator: (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'))
        } else
          if (value !== registerForm.password) {
            callback(new Error('两次输入密码不一致'))
          } else {
            callback()
          }
      }
    }
  ]
}

const rules = {
  username: [
    { required: true, message: '用户名不能为空', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '密码不能为空', trigger: 'blur' }
  ]
}

const registerFormRef = ref(null)
const formRef = ref(null)
const loading = ref(false)
const onSubmit = () => {
  formRef.value.validate((valid) => {
    if (!valid) {
      return false
    }
    loading.value = true
    store.dispatch('login', form).then(() => {
      notify('登陆成功', 'success')
      router.push('/')
    }).finally(() => {
      loading.value = false
    })
  })
}

const onRegister = () => {
  registerFormRef.value.validate((valid) => {
    if (!valid) {
      return false
    }
    loading.value = true
    store.dispatch('register', registerForm).then(() => {
      notify('注册成功', 'success')
      router.push('/')
    }).finally(() => {
      loading.value = false
    })
  })
}

// 监听回车事件
function onKeyup (e) {
  if (e.key === 'Enter') onSubmit()
}

// 添加键盘监听
onMounted(() => {
  document.addEventListener('keyup', onKeyup)
})
onBeforeUnmount(() => {
  document.removeEventListener('keyup', onKeyup)
})

const Repeat = ref(false)

const clickRepeat = () => {
  Repeat.value = !Repeat.value
}
</script>

<template>
  <div class="container" onclick="onclick">
    <div class="top"></div>
    <div class="bottom"></div>
    <div class="center">
      <div class="rgewgw">
        <font-awesome-icon :icon="['fas', 'repeat']" @click="clickRepeat"/>
      </div>
        <div class="animated fadeIn" v-if="!Repeat">
          <h2 class="item-center justify-center flex font-bold text-3xl text-gray-800">欢迎回来</h2>
          <div
              class="flex items-center justify-center my-5 text-gray-300 space-x-2"
          >
            <span class="h-[1px] w-16 bg-gray-200"></span>
            <span>账号密码登录</span>
            <span class="h-[1px] w-16 bg-gray-200"></span>
          </div>
          <el-form ref="formRef" :rules="rules" :model="form" class="w-[250px]">
            <el-form-item prop="username">
              <el-input v-model="form.username" placeholder="请输入用户名">
                <template #prefix>
                  <el-icon>
                    <user/>
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input
                  type="password"
                  v-model="form.password"
                  placeholder="请输入密码"
                  show-password
              >
                <template #prefix>
                  <el-icon>
                    <lock/>
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item>
              <el-button
                  round
                  color="#626aef"
                  class="w-[250px]"
                  type="primary"
                  @click="onSubmit"
                  :loading="loading">
                登 录
              </el-button>
            </el-form-item>
          </el-form>
        </div>
        <div v-else class="animated fadeIn">
          <h2 class="item-center justify-center flex font-bold text-3xl text-gray-800">加入我们</h2>
          <div
              class="flex items-center justify-center my-5 text-gray-300 space-x-2"
          >
            <span class="h-[1px] w-16 bg-gray-200"></span>
            <span>注册</span>
            <span class="h-[1px] w-16 bg-gray-200"></span>
          </div>
          <el-form ref="registerFormRef" :rules="registerRules" :model="registerForm" class="w-[250px]">
            <el-form-item prop="username">
              <el-input v-model="registerForm.username" placeholder="请输入用户名">
                <template #prefix>
                  <el-icon>
                    <user/>
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input
                  type="password"
                  v-model="registerForm.password"
                  placeholder="请输入密码"
                  show-password
              >
                <template #prefix>
                  <el-icon>
                    <lock/>
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item prop="repeatPassword">
              <el-input
                  type="password"
                  v-model="registerForm.repeatPassword"
                  placeholder="请输入确认密码"
                  show-password
              >
                <template #prefix>
                  <el-icon>
                    <lock/>
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item>
              <el-button
                  round
                  color="#626aef"
                  class="w-[250px]"
                  type="primary"
                  @click="onRegister"
                  :loading="loading">
                注册
              </el-button>
            </el-form-item>
          </el-form>
        </div>

    </div>
  </div>
</template>

<style scoped>
@import url("https://fonts.googleapis.com/css?family=Raleway:400,700");

*,
*:before,
*:after {
  box-sizing: border-box;
}

body {
  min-height: 100vh;
  font-family: 'Raleway', sans-serif;
  margin: 0;
}

.container {
  position: absolute;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.container:hover .top:before,
.container:hover .top:after,
.container:hover .bottom:before,
.container:hover .bottom:after,
.container:active .top:before,
.container:active .top:after,
.container:active .bottom:before,
.container:active .bottom:after {
  margin-left: 200px;
  transform-origin: -200px 50%;
  transition-delay: 0s;
}

.container:hover .center,
.container:active .center {
  opacity: 1;
  transition-delay: 0.2s;
}

.top:before,
.top:after,
.bottom:before,
.bottom:after {
  content: '';
  display: block;
  position: absolute;
  width: 200vmax;
  height: 200vmax;
  top: 50%;
  left: 50%;
  margin-top: -100vmax;
  transform-origin: 0 50%;
  transition: all 0.5s cubic-bezier(0.445, 0.05, 0, 1);
  z-index: 10;
  opacity: 0.65;
  transition-delay: 0.2s;
}

.top:before {
  transform: rotate(45deg);
  background: #e46569;
}

.top:after {
  transform: rotate(135deg);
  background: #ecaf81;
}

.bottom:before {
  transform: rotate(-45deg);
  background: #60b8d4;
}

.bottom:after {
  transform: rotate(-135deg);
  background: #3745b5;
}

.center {
  position: absolute;
  width: 400px;
  height: 400px;
  top: 50%;
  left: 50%;
  margin-left: -200px;
  margin-top: -200px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 30px;
  opacity: 0;
  transition: all 0.5s cubic-bezier(0.445, 0.05, 0, 1);
  transition-delay: 0s;
  color: #333;
}

.center input {
  width: 100%;
  padding: 15px;
  margin: 5px;
  border-radius: 1px;
  border: 1px solid #ccc;
  font-family: inherit;
}

.rgewgw {
  margin-bottom: 10px;
  cursor: pointer;
}
</style>
