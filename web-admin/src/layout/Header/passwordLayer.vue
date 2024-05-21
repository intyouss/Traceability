<template>
  <Layer :layer="props.layer" @confirm="submit" ref="layerDom">
    <el-form :model="form" :rules="rules" ref="ruleForm" label-width="120px" style="margin-right:30px;">
      <el-form-item label="用户名：" prop="name">
        管理员
      </el-form-item>
      <el-form-item label="原密码：" prop="old">
        <el-input v-model="form.old" placeholder="请输入原密码" show-password />
      </el-form-item>
      <el-form-item label="新密码：" prop="new">
        <el-input v-model="form.new" placeholder="请输入新密码" show-password />
      </el-form-item>
    </el-form>
  </Layer>
</template>

<script lang="js" setup>
import { ref } from 'vue'
import { useStore } from 'vuex'
import { updateUserApi} from '@/api/user'
import Layer from '@/components/layer/index.vue'
import { notify } from '@/composables/util'
import router from "@/router";

const props = defineProps({
  layer: {
    type: Object,
    default: () => {
      return {
        show: false,
        title: '',
        showButton: true
      }
    }
  }
})

const ruleForm = ref(null)
const layerDom = ref(null)
const store = useStore()
const form = ref({
  userId: '',
  name: '',
  old: '',
  new: ''
})
const rules = {
  old: [{ required: true, message: '请输入原密码', trigger: 'blur' }],
  new: [{ required: true, message: '请输入新密码', trigger: 'blur' }]
}
function submit () {
  if (ruleForm.value) {
    ruleForm.value.validate((valid) => {
      if (valid) {
        const params = {
          user_id: store.getters['user/info'].id,
          password: form.value.old,
          new_password: form.value.new
        }
        updateUserApi(params)
          .then(() => {
            notify('密码修改成功，即将跳转到登录页面', 'success')
            layerDom.value
            setTimeout(() => {
              store.dispatch('user/loginOut').then(() => {
                router.push('/login')
              })
            }, 2000)
          })
      } else {
        return false
      }
    })
  }
}
</script>

<style lang="scss" scoped>

</style>
