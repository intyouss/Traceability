<template>
  <component :is="type" v-bind="linkProps(props.to)" @click="hideMenu" >
    <slot></slot>
  </component>
</template>

<script lang="js" setup>
import { computed } from 'vue'
import { useStore } from 'vuex'

const props = defineProps({
  to: {
    type: String,
    required: true
  }
})

const type = 'router-link'

const store = useStore()
const isCollapse = computed(() => store.state.app.isCollapse)
const linkProps = (to) => {
  return {
    to
  }
}
const hideMenu = () => {
  if (document.body.clientWidth <= 1000 && !isCollapse.value) {
    store.commit('app/isCollapseChange', true)
  }
}
</script>
<style lang="">

</style>
