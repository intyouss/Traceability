<template>
  <div>
    <barChart />
    <el-row :gutter="20">
      <el-col :lg="12" :md="24">
        <usageCPUChart :cpu-usage="cpuUsage"/>
      </el-col>
      <el-col :lg="12" :md="24">
        <cacheChart :memory-usage="memoryUsage"/>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import barChart from './barChart.vue'
import usageCPUChart from './usageCPUChart.vue'
import cacheChart from './cacheChart.vue'
import { getCpuUsage, getMemoryUsage } from '@/api/system/system'
import { onBeforeMount, onBeforeUnmount, ref } from 'vue'

const cpuUsage = ref(0)
const memoryUsage = ref(0)

onBeforeMount(async () => {
  await getCpuUsage().then(res => {
    cpuUsage.value = +res.data.cpu_usage
  })
  await getMemoryUsage().then(res => {
    memoryUsage.value = +res.data.memory_usage
  })
})

let timer = null
timer = setInterval(async function () {
  await getCpuUsage().then(res => {
    cpuUsage.value = +res.data.cpu_usage
  })
  await getMemoryUsage().then(res => {
    memoryUsage.value = +res.data.memory_usage
  })
}, 6000)

onBeforeUnmount(() => {
  clearInterval(timer)
  timer = null
})
</script>
