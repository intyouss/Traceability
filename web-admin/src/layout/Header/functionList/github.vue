<template>
  <a href="https://github.com/cmdparkour/vue-admin-box" target="_blank" :title="$t('message.system.github')">
    <i class="sfont system-github" @click="toggle" />
  </a>
</template>

<script setup>
</script>

<style lang="scss" scoped>
  a {
    &:focus {
      outline: none;
    }
  }
  i {
    cursor: pointer;
    font-size: 18px;
    &:focus {
      outline: none;
    }
  }
</style>
