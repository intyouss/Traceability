<script>
export default {
  created() {
    if (sessionStorage.getItem('store')) {
      this.$store.replaceState(
          Object.assign(
              {},
              this.$store.state,
              JSON.parse(sessionStorage.getItem('store')),
          ),
      );
    }
    window.addEventListener('beforeunload', ()=>{
      sessionStorage.setItem('store', JSON.stringify(this.$store.state));
    });
  },
};
</script>
<script setup>
import {provide, nextTick, ref} from 'vue';

const isRouterActive = ref(true);

provide('reload', () => {
  isRouterActive.value = false;
  nextTick(() => {
    isRouterActive.value = true;
  });
});

</script>

<template>
  <router-view v-if="isRouterActive"></router-view>
</template>

<style>
body {
  @apply bg-light-50;
}

#nprogress .bar{
  background-color: #a09ccf !important;
  height: 3px !important;
}
</style>
