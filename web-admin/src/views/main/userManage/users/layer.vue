<template>
  <Layer :layer="props.layer" @confirm="submit">
    <template v-if="props.layer.type !== 2">
      <el-form :model="ruleForm" :rules="rules" ref="form" label-width="120px" style="margin-right:30px;">
        <el-form-item label="用户名：" prop="username">
          <el-input v-model="ruleForm.username" maxlength="10"></el-input>
        </el-form-item>
        <el-form-item label="密码：" prop="password">
          <el-input v-model="ruleForm.password" show-password maxlength="10"></el-input>
        </el-form-item>
        <el-form-item label="确认密码：" prop="rePassword">
          <el-input v-model="ruleForm.rePassword" show-password maxlength="10"></el-input>
        </el-form-item>
        <el-form-item label="角色：" prop="role">
          <el-select v-model="ruleForm.role" placeholder="请选择" clearable>
            <el-option
                v-for="item in option"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
    </template>
    <template v-else>
      <el-form :model="editRuleForm" :rules="editRules" ref="form" label-width="120px" style="margin-right:30px;">
        <el-form-item label="用户名：" prop="username">
          <el-input v-model="editRuleForm.username" maxlength="10"></el-input>
        </el-form-item>
        <el-form-item label="角色：" prop="role">
          <el-select v-model="editRuleForm.role" placeholder="请选择" clearable>
            <el-option
                v-for="item in option"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
    </template>
  </Layer>
</template>

<script setup>
import { reactive, ref } from 'vue'
import Layer from '@/components/layer/index.vue'
import { notify } from '@/composables/util'
import {addUser} from "@/api/system/user";
import {updateUserApi} from "@/api/user";

const props = defineProps({
  layer: {
    type: Object,
    default: () => {
      return {
        show: false,
        title: '',
        row: {},
        showButton: true,
        type: 1
      }
    }
  }
})

const ruleForm = reactive({
  username: '',
  password: '',
  rePassword: '',
  role: null
})
const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  rePassword: [{
    required: true,
    trigger: 'blur',
    validator: (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== ruleForm.password) {
        callback(new Error('两次输入密码不一致'))
      } else {
        callback()
      }
    }
  }],
  role: [{ required: true, message: '请选择', trigger: 'blur' }]
}

const editRuleForm = reactive({
  user_id: 0,
  username: '',
  role: null
})

const editRules = {
  username: [{ message: '请输入用户名', trigger: 'blur' }],
  role: [{ required: false, trigger: 'blur' }]
}



const option = [
  { value: 2, label: '管理员' },
  { value: 1, label: '普通用户' }
]

const emits = defineEmits(['getTableData'])

const form = ref(null)
const submit = () => {
  form.value.validate((valid) => {
    if (valid) {
      if (props.layer.type === 2) {
        editRuleForm.user_id = props.layer.row.id
        updateForm(editRuleForm)
      } else {
        addForm(ruleForm)
      }
    } else {
      return false
    }
  })
}
// 新增提交事件
const addForm = (params) => {
  addUser(params)
    .then(() => {
      notify('新增成功', 'success')
      props.layer.show = false
      emits('getTableData', true)
    })
}
// 编辑提交事件
const updateForm = (params) => {
  updateUserApi(params)
    .then(() => {
      notify('编辑成功', 'success')
      emits('getTableData', false)
    })
}
</script>

<style lang="scss" scoped>

</style>
