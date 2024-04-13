<script setup>
const props = defineProps({
  row: Object
})
</script>

<template>
  <div class="card">
    <a :href="props.row.url" target="_blank" style="color: initial">
      <h1 class="name">{{ props.row.name }}</h1>
      <div class="detail">{{ props.row.desc }}</div>
    </a>
  </div>
</template>

<style lang="scss" scoped>
  .card {
    border-radius: 4px;
    overflow: hidden;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: var(--system-page-background);
    margin: 20px;
    box-shadow: 3px 3px 10px 0 rgba(58, 59, 69, 0.15);
    cursor: pointer;
    outline: rgba(158, 158, 158, 1) solid 2px;
    height: 168px;
  }
  .name {
    font-size: 25px;
    font-weight: bold;
  }
  .detail {
    line-height: 24px;
    font-size: 18px;
  }
  .card:hover {
    background-color: #e8e1e1;
    transition: all 0.3s;
  }
</style>
