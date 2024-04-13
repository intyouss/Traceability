<template>
  <div class="box">
    <Chart :option="options"/>
  </div>
</template>

<script setup>
import { onBeforeMount, onBeforeUnmount, reactive } from 'vue'
import Chart from '@/components/charts/index.vue'
import option from './modules/bar'
import { getUserIncreaseList } from '@/api/user'
import { getVideoIncreaseList } from '@/api/video'

const options = reactive(option)
const year = new Date().getFullYear()
const month = new Date().getMonth() + 1

onBeforeMount(async () => {
  await getUserIncreaseList({
    year,
    month
  }).then(res => {
    options.series[0].data = change(res.data.user_increase_list)
  })
  await getVideoIncreaseList({
    year,
    month
  }).then(res => {
    options.series[1].data = change(res.data.video_increase_list)
  })
})

let timer = null
timer = setInterval(async () => {
  await getUserIncreaseList({
    year,
    month
  }).then(res => {
    options.series[0].data = change(res.data.user_increase_list)
  })
  await getVideoIncreaseList({
    year,
    month
  }).then(res => {
    options.series[1].data = change(res.data.video_increase_list)
  })
}, 60000)

onBeforeUnmount(() => {
  clearInterval(timer)
  timer = null
})

const change = (data) => {
  const dayNumber = new Date().getDate()
  const d = []
  if (data.length === dayNumber) {
    for (let i = 0; i < dayNumber; i++) {
      d.push(data[i].count)
    }
    return d
  }
  let j = 1
  let k = 0
  while (j <= dayNumber && k < data.length) {
    if (data[k].day !== j) {
      d.push(0)
      j++
    } else {
      d.push(data[k].count)
      j++
      k++
    }
  }
  for (let i = k; i < dayNumber; i++) {
    d.push(0)
  }
  return d
}
</script>

<style lang="scss" scoped>
.box {
  margin: 10px auto 0;
  width: calc(100% - 40px);
  height: 400px;
  background: var(--system-page-background);
  padding: 20px;
  overflow: hidden;
}
</style>
