<template>
  <Layer :layer="props.layer" @confirm="submit">
    <el-form :model="ruleForm" :rules="rules" ref="form" label-width="120px" style="margin-right:30px;">
      <el-form-item label="角色名：" prop="name">
        <el-input v-model="ruleForm.name" maxlength="10"></el-input>
      </el-form-item>
      <el-form-item label="描述：" prop="desc">
        <el-input v-model="ruleForm.desc" maxlength="100" type="textarea" autosize></el-input>
      </el-form-item>
    </el-form>
  </Layer>
</template>

<script setup>
import { reactive, ref } from 'vue'
import Layer from '@/components/layer/index.vue'
import { addRole, updateRole } from '@/api/system/user'
import { notify } from '@/composables/util'

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
  name: '',
  desc: ''
})
const rules = {
  name: [{ required: true, message: '请输入角色名', trigger: 'blur' }],
  desc: [{ required: true, message: '请输入角色描述', trigger: 'blur' }]
}

const emits = defineEmits(['getTableData'])

const form = ref(null)
const submit = () => {
  form.value.validate((valid) => {
    if (valid) {
      const params = ruleForm
      if (props.layer.type === 2) {
        params.id = props.layer.row.id
        updateForm(params)
      } else {
        addForm(params)
      }
    } else {
      return false
    }
  })
}
// 新增提交事件
const addForm = (params) => {
  addRole(params)
    .then(() => {
      notify('新增成功', 'success')
      props.layer.show = false
      emits('getTableData', true)
    })
}
// 编辑提交事件
const updateForm = (params) => {
  updateRole(params)
    .then(() => {
      notify('编辑成功', 'success')
      emits('getTableData', false)
    })
}
</script>

<style lang="scss" scoped>

</style>
