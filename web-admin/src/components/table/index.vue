<template>
  <div class="system-table-box">
    <el-table
      v-bind="$attrs"
      ref="table"
      class="system-table"
      border
      height="100%"
      :data="props.data"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" align="center" width="50" v-if="props.showSelection" />
      <el-table-column label="序号" width="60" align="center" v-if="props.showIndex">
        <template #default="scope">
          {{ (props.page.index - 1) * props.page.size + scope.$index + 1 }}
        </template>
      </el-table-column>
      <slot></slot>
    </el-table>
    <el-pagination
      v-if="props.showPage"
      v-model:current-page="props.page.index"
      class="system-page"
      background
      :layout="props.pageLayout"
      :total="props.page.total"
      :page-size="props.page.size"
      :page-sizes="props.pageSizes"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    >
    </el-pagination>
  </div>
</template>

<script setup>
import { ref, onActivated } from 'vue'
const props = defineProps({
  data: { type: Array, default: () => [] }, // 数据源
  select: { type: Array, default: () => [] }, // 已选择的数据，与selection结合使用
  showIndex: { type: Boolean, default: false }, // 是否展示index选择，默认否
  showSelection: { type: Boolean, default: false }, // 是否展示选择框，默认否
  showPage: { type: Boolean, default: true }, // 是否展示页级组件，默认是
  page: { // 分页参数
    type: Object,
    default: () => {
      return { index: 1, size: 20, total: 0 }
    }
  },
  pageLayout: { type: String, default: 'total, sizes, prev, pager, next, jumper' }, // 分页需要显示的东西，默认全部
  pageSizes: { type: Array, default: [10, 20, 50, 100] }
})
const emits = defineEmits(['getTableData', 'selection-change'])
const table = ref(null)
let timer = null
// 分页相关：监听页码切换事件
const handleCurrentChange = (val) => {
  if (timer) {
    props.page.index = 1
  } else {
    props.page.index = val
    emits('getTableData')
  }
}
// 分页相关：监听单页显示数量切换事件
const handleSizeChange = (val) => {
  timer = 'work'
  setTimeout(() => {
    timer = null
  }, 100)
  props.page.size = val
  emits('getTableData', true)
}
// 选择监听器
const handleSelectionChange = (val) => {
  emits('selection-change', val)
}
// 解决BUG：keep-alive组件使用时，表格浮层高度不对的问题
onActivated(() => {
  table.value.doLayout()
})
</script>

<style lang="scss" scoped>
  .system-table-box {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    height: 100%;
    .system-table {
      flex: 1;
      height: 100%;
    }

    .system-page {
      margin-top: 20px;
    }
  }
</style>
