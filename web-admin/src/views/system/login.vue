<template>
  <div class="container">
    <div class="box">
      <div class="login-content-left">
        <img :src="loginLeftPng" alt=""/>
        <div class="login-content-left-mask">
          <div>{{ $t(systemTitle) }}</div>
          <div>{{ $t(systemSubTitle) }}</div>
        </div>
      </div>

      <div class="box-inner">
        <h1>{{ $t('message.system.welcome') }}</h1>
        <el-form class="form">
          <el-input
              size="large"
              v-model="form.name"
              :placeholder="$t('message.system.userName')"
              type="text"
              maxlength="50"
          >
            <template #prepend>
              <i class="sfont system-xingmingyonghumingnicheng"></i>
            </template>
          </el-input>
          <el-input
              size="large"
              ref="password"
              v-model="form.password"
              :type="passwordType"
              :placeholder="$t('message.system.password')"
              name="password"
              maxlength="50"
          >
            <template #prepend>
              <i class="sfont system-mima"></i>
            </template>
            <template #append>
              <i class="sfont password-icon" :class="passwordType ? 'system-yanjing-guan': 'system-yanjing'"
                 @click="passwordTypeChange"></i>
            </template>
          </el-input>

          <el-button type="primary" :loading="form.loading" @click="submit" style="width: 100%;" size="medium">
            {{ $t('message.system.login') }}
          </el-button>
        </el-form>
        <div class="fixed-top-right">
          <select-lang/>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="js" setup>
import { systemTitle, systemSubTitle } from '@/config'
import { ref, reactive } from 'vue'
import { useStore } from 'vuex'
import selectLang from '@/layout/Header/functionList/word.vue'
import loginLeftPng from '@/assets/login/left.jpg'
import { notify } from '@/composables/util'

const store = useStore()
const form = reactive({
  name: 'admin',
  password: '123456',
  admin: true,
  loading: false
})
const passwordType = ref('password')
const passwordTypeChange = () => {
  passwordType.value === '' ? passwordType.value = 'password' : passwordType.value = ''
}
const checkForm = () => {
  return new Promise((resolve, reject) => {
    if (form.name === '') {
      notify('用户名不能为空', 'warning')
      return
    }
    if (form.password === '') {
      notify('密码不能为空', 'warning')
      return
    }
    resolve(true)
  })
}
const submit = () => {
  checkForm()
    .then(() => {
      form.loading = true
      const params = {
        username: form.name,
        password: form.password,
        admin: form.admin
      }
      store.dispatch('user/login', params)
        .then(async () => {
          notify('登录成功', 'success')
          location.reload()
          // await getAuthRoutes()
          // await router.push(route.query.redirect || '/')
        }).finally(() => {
          form.loading = false
        })
    })
}
</script>

<style lang="scss" scoped>
.container {
  position: relative;
  width: 100vw;
  height: 100vh;
  background: #fff url('@/assets/login/bg.png') no-repeat center;
  overflow: hidden;
  background-size: cover;
  cursor: pointer;
  user-select: none;

  .box {
    width: 1160px;
    display: flex;
    position: absolute;
    left: 50%;
    top: 50%;
    background: white;
    border-radius: 8px;
    transform: translate(-50%, -50%);
    height: 440px;
    overflow: hidden;
    box-shadow: 0 6px 20px 5px rgba(152, 152, 152, 0.1),
    0 16px 24px 2px rgba(117, 117, 117, 0.14);

    .login-content-left {
      position: relative;

      img {
        height: 440px;
      }

      .login-content-left-mask {
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        background-image: linear-gradient(rgba(0,204,222,0.8), rgba(51,132,224,0.8));
        text-align: center;
        color: #fff;
        font-size: 1.8rem;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        letter-spacing: 2px;

        div:nth-child(1) {
          font-size: 3.5rem;
          margin-bottom: 1em;
        }
      }
    }

    .box-inner {
      width: 500px;

      h1 {
        margin-top: 80px;
        text-align: center;
      }

      .form {
        width: 80%;
        margin: 50px auto 15px;

        .el-input {
          margin-bottom: 20px;
        }

        .password-icon {
          cursor: pointer;
          color: #409eff;
        }
      }

      .fixed-top-right {
        position: absolute;
        top: 10px;
        right: 10px;
      }
    }
  }
}

@media screen and (max-width: 1260px) {
  .login-content-left {
    display: none;
  }
  .box {
    width: 500px !important;
  }
}

@media screen and (max-width: 750px) {
  .container .box, .container .box-inner {
    width: 100vw !important;
    height: 100vh;
    box-shadow: none;
    left: 0;
    top: 0;
    transform: none;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    h1 {
      margin-top: 0;
    }

    .form {
    }
  }
}
</style>
