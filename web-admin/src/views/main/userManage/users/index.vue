<template>
  <div class="layout-container">
    <div class="layout-container-form flex space-between">
      <div class="layout-container-form-handle">
        <el-button type="primary" :icon="Plus" @click="handleAdd">{{
          $t("message.common.add")
        }}</el-button>
        <el-popconfirm
          :title="$t('message.common.delTip')"
          @confirm="handleDel(chooseData)"
        >
          <template #reference>
            <el-button
              type="danger"
              :icon="Delete"
              :disabled="chooseData.length === 0"
              >{{ $t("message.common.delBat") }}</el-button
            >
          </template>
        </el-popconfirm>
      </div>
      <div class="layout-container-form-search">
        <el-input
          v-model="query.key"
          :placeholder="$t('message.common.searchTip')"
        ></el-input>
        <el-button
          type="primary"
          :icon="Search"
          class="search-btn"
          @click="getTableData(true)"
          >{{ $t("message.common.search") }}</el-button
        >
      </div>
    </div>
    <div class="layout-container-table">
      <Table
        ref="table"
        v-model:page="page"
        v-loading="loading"
        :showSelection="true"
        :data="tableData"
        @getTableData="getTableData"
        @selection-change="handleSelectionChange"
      >
        <el-table-column prop="id" label="数据库Id" align="center" width="80" />
        <el-table-column prop="username" label="用户名" align="center" />
<!--        <el-table-column prop="nickName" label="昵称" align="center" />-->
<!--        <el-table-column prop="email" label="邮箱" align="center" />-->
<!--        <el-table-column prop="mobile" label="手机号" align="center" />-->
        <el-table-column prop="role.name" label="角色" align="center" />
        <el-table-column prop="status" label="状态" align="center">
          <template #default="scope">
            <span class="statusName">{{ scope.row.status === 1 ? "正常" : "封禁" }}</span>
            <el-switch
              v-model="scope.row.status"
              active-color="#13ce66"
              inactive-color="#ff4949"
              :active-value="1"
              :inactive-value="2"
              :loading="scope.row.loading"
              @change="handleUpdateStatus(scope.row)"
            ></el-switch>
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('message.common.handle')"
          align="center"
          fixed="right"
          width="200"
        >
          <template #default="scope">
            <el-button @click="handleEdit(scope.row)">{{
              $t("message.common.update")
            }}</el-button>
            <el-popconfirm
              :title="$t('message.common.delTip')"
              @confirm="handleDel([scope.row])"
            >
              <template #reference>
                <el-button type="danger">{{ $t("message.common.del") }}</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </Table>
      <Layer :layer="layer" @getTableData="getTableData" v-if="layer.show" />
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { del } from '@/api/system/user'
import Table from '@/components/table/index.vue'
import Layer from './layer.vue'
import { Plus, Delete, Search } from '@element-plus/icons'
import { getUserListApi, updateUserApi } from '@/api/user'
import { notify } from '@/composables/util'

const query = reactive({
  key: ''
})
// 弹窗控制器
const layer = reactive({
  show: false,
  title: '新增',
  showButton: true,
  type: 1
})
// 分页参数, 供table使用
const page = reactive({
  index: 1,
  size: 20,
  total: 0
})
const loading = ref(true)
const tableData = ref([])
const chooseData = ref([])
const handleSelectionChange = (val) => {
  chooseData.value = val
}

// 获取表格数据
const getTableData = (init) => {
  loading.value = true
  if (init) {
    page.index = 1
  }
  const params = {
    page: page.index,
    limit: page.size,
    ...query
  }
  getUserListApi(params)
    .then((res) => {
      const data = res.data.users
      data.forEach((d) => {
        d.loading = false
      })
      tableData.value = data
      page.total = Number(res.total)
    })
    .catch(() => {
      tableData.value = []
      page.index = 1
      page.total = 0
    })
    .finally(() => {
      loading.value = false
    })
}
// 删除功能
const handleDel = (data) => {
  const params = {
    ids: data
      .map((e) => {
        return e.id
      })
  }
  del(params).then(() => {
    notify('删除成功', 'success')
    getTableData(tableData.value.length === 1)
  })
}
// 新增弹窗功能
const handleAdd = () => {
  layer.title = '新增数据'
  layer.show = true
  layer.type = 1
  delete layer.row
}
// 编辑弹窗功能
const handleEdit = (row) => {
  layer.title = '编辑数据'
  layer.type = 2
  layer.row = row
  layer.show = true
}
// 状态编辑功能
const handleUpdateStatus = (row) => {
  if (!row.id) {
    return
  }
  row.loading = true
  const params = {
    user_id: row.id,
    status: row.status
  }
  // TODO: 需要admin jwtToken 配合
  updateUserApi(params)
    .then(() => {
      notify('状态变更成功', 'success')
    })
    .catch(() => {
      notify('状态变更失败', 'error')
    })
    .finally(() => {
      row.loading = false
    })
}

onMounted(() => {
  getTableData(true)
})
</script>

<style lang="scss" scoped>
.statusName {
  margin-right: 10px;
}
</style>
