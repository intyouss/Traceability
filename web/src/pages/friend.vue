<script>
import VPlayer from '~/layouts/components/video/VPlayer.vue';

export default {
  components: {
    VPlayer,
  },
  data() {
    return {
      data: [{
        title: 'hahahah',
      }],
    };
  },
};

</script>

<template>
  <div>
    <v-player :data="data"/>
  </div>
</template>

<style scoped>

</style>
