<template>
  <div class="box">
    <Card />
    <Charts />
  </div>
</template>

<script setup>
import Card from './components/card/index.vue'
import Charts from './components/charts/index.vue'
</script>

<style lang="scss" scoped>
  .box {
    padding: 15px;
  }
</style>
