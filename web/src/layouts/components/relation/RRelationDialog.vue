<script setup>
import { ref, watch } from 'vue'
import RUser from '~/layouts/components/relation/RUser.vue'

const props = defineProps({
  open: Boolean,
  type: Number,
  close: Function,
  user: Object,
  list: Array
})

const Open = ref(props.open)

watch(() => props.open, (newVal) => {
  Open.value = newVal
})

const emit = defineEmits(['click'])
const handleClick = (type) => {
  emit('click', type)
}

const RelationList = ref(props.list)

watch(() => props.list, (newVal) => {
  RelationList.value = newVal
})
</script>

<template>
  <div>
    <el-dialog
        v-model="Open"
        width="40%"
        :close-on-click-modal="false"
        :before-close="props.close"
        class="htreg"
    >
      <template #header>
        <div class="htregw">
          <div class="geriufw">
            <template v-if="props.type === 1">
              <div class="btrfheg">关注({{props.user.focus_count}})</div>
              <div
                  class="breffw"
                  @click="handleClick(2)">粉丝({{props.user.fans_count}})
              </div>
            </template>
            <template v-else>
              <div
                  class="breffw"
                  @click="handleClick(1)">关注({{props.user.focus_count}})
              </div>
              <div class="btrfheg">粉丝({{props.user.fans_count}})</div>
            </template>
          </div>
          <div class="trge"></div>
        </div>
      </template>
      <ul>
        <template v-for="item in RelationList" :key="item.id">
          <li>
            <r-user :user="item"/>
            <div class="trge"></div>
          </li>
        </template>
      </ul>
    </el-dialog>
  </div>
</template>

<style scoped>
.geriufw {
  min-height: 48px;
  flex-direction: row;
  display: flex;
}

.breffw {
  color: rgba(22, 24, 35, .6);
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
  font-weight: 600;
  cursor: pointer;
  margin-bottom: 22px;
  margin-right: 40px;
  font-size: 18px;
  line-height: 26px;
}
.breffw:hover {
  color: #161823;
}
.btrfheg {
  color: #161823;
  font-family: PingFang SC, DFPKingGothicGB-Medium, sans-serif;
  font-weight: 800;
  cursor: pointer;
  margin-bottom: 22px;
  margin-right: 40px;
  font-size: 18px;
  line-height: 26px;
}
.trge {
  width: 100%;
  height: 1px;
  min-height: 1px;
  background-color: rgba(22, 24, 35, .11);
  display: block;
  position: relative;
}
.htregw {
  padding: 20px;
}
</style>

<style>
.htreg .el-dialog__body {
  padding: 0 55px 10px 40px;
}
.htreg {
  border-radius: 12px;
  min-width: 500px;
  @apply bg-gray-100 border-2 shadow-md
}
</style>
