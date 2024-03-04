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

<template>
  <router-view></router-view>
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
