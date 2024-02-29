<script setup>
const props = defineProps({
  title: String,
  author: String,
  createdAt: String,
});
</script>

<template>
  <div>
    <div class="text-xl my-2">
      @{{props.author}} · {{props.createdAt}}<br>
      {{props.title}}
    </div>
  </div>
</template>

<style scoped>

</style>
