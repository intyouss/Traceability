<script setup>
import UAvatar from '~/layouts/components/user/UAvatar.vue'
import { ref, watch } from 'vue'

const props = defineProps({
  users: ref([])
})

const emit = defineEmits(['select'])

const handleSelect = (index) => {
  emit('select', index)
}

const Users = ref(props.users)

watch(() => props.users, (value) => {
  console.log(value)
  Users.value = value
})

</script>

<template>
  <div class="f-message-menu">
    <el-menu
        class="fwewf border-0"
        @select="handleSelect"
        active-text-color="#ffffff"
    >
      <template v-for="(item, index) in Users" :key="item.id">
        <el-menu-item class="bre m-1 rounded-lg" :index="index">
          <u-avatar
              :user-id="0"
              :avatar="item.avatar"
              :mine="false"
              class="w-[40px] h-[40px] mr-2"
          />
          <span
              style="text-overflow:ellipsis; width: 8em;white-space: nowrap;overflow: hidden;"
          >
            {{ item.username }}
          </span>
        </el-menu-item>
      </template>
    </el-menu>
  </div>
</template>

<style>
.f-message-menu{
  transition: 0.4s;
  overflow-y: auto;
  overflow-x: hidden;
}
.f-message-menu::-webkit-scrollbar{
  width: 0;
}
.fwewf .el-menu-item.is-active {
  background: linear-gradient(to bottom, #33ccff 0%, #99ccff 100%);
}
</style>
