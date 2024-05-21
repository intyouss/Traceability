import { inject, reactive, ref } from 'vue'
import { getAuthUserSearch, getPublicUserSearch, updateUser } from '~/api/user.js'
import { confirm, notify } from '~/composables/util.js'
import { useRouter, useRoute } from 'vue-router'
import { useStore } from 'vuex'
import { getVideoSearch } from '~/api/video.js'
import { getToken } from '~/composables/auth.js'

export function useRePassword () {
  const rePasswordForm = ref(false)
  const formLabelWidth = '140px'
  const store = useStore()
  const router = useRouter()
  const form = reactive({
    oldPassword: '',
    newPassword: '',
    enterPassword: ''
  })
  const rules = {
    oldPassword: [
      { required: true, message: '旧密码不能为空', trigger: 'blur' }
    ],
    newPassword: [
      { required: true, message: '新密码不能为空', trigger: 'blur' }
    ],
    enterPassword: [
      {
        trigger: 'blur',
        validator: (rule, value, callback) => {
          if (value === '') {
            callback(new Error('确认密码不能为空'))
          } else if (value !== form.newPassword) {
            callback(new Error('两次输入密码不一致'))
          } else {
            callback()
          }
        }
      }
    ]
  }
  const formRef = ref(null)
  const loading = ref(false)
  const onSubmit = () => {
    formRef.value.validate((valid) => {
      if (!valid) {
        return false
      }
      loading.value = true
      updateUser({ userId: store.state.user.id, password: form.oldPassword, newPassword: form.enterPassword })
        .then(() => {
          notify('修改成功', 'success')
          store.dispatch('logout').then()
          router.push('/login').then()
          rePasswordFormClose()
        }).finally(() => {
          loading.value = false
        })
    })
  }
  const rePasswordFormOpen = () => rePasswordForm.value = true
  const rePasswordFormClose = () => rePasswordForm.value = false
  return {
    rePasswordForm,
    rePasswordFormOpen,
    rePasswordFormClose,
    formLabelWidth,
    form,
    rules,
    formRef,
    onSubmit
  }
}

export function useLogout () {
  const router = useRouter()
  const route = useRoute()
  const store = useStore()
  const reload = inject('reload')
  function handleLogout () {
    confirm('确定退出登录吗?').then(() => {
      // logout()
      //   .finally(() => {
      //     store.dispatch('logout')
      //     if (route.path !== '/') {
      //       router.push('/login')
      //     }
      //     notify('退出成功')
      //   })
      store.dispatch('logout')
      if (route.path !== '/') {
        router.push('/login')
      }
      notify('退出成功')
      reload()
    })
  }
  return {
    handleLogout
  }
}

export function useSearch () {
  const getSearch = (tag, key) => {
    switch (tag) {
      case '综合':
        return getVideoSearch(key, 1).then((res) => {
          return res.data.videos
        })
      case '视频':
        return getVideoSearch(key, 2).then((res) => {
          return res.data.videos
        })
      case '用户':
        if (getToken()) {
          return getAuthUserSearch(key).then((res) => {
            return res.data.users
          })
        }
        return getPublicUserSearch(key).then((res) => {
          return res.data.users
        })
    }
  }
  return {
    getSearch
  }
}
