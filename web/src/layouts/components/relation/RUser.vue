<script setup>

import UAvatar from '~/layouts/components/user/UAvatar.vue';
import {useRelation} from '~/composables/relationManager.js';
import {onBeforeMount} from 'vue';

const props = defineProps({
  user: Object,
});
const {
  IsFocus,
  setFocus,
  handleRelation,
} = useRelation();

onBeforeMount(() => {
  setFocus(props.user.is_focus);
});
</script>

<template>
  <div class="hrbegw">
    <u-avatar
        class="hr4eg3rh"
        :user-id="props.user.id"
        :avatar="props.user.avatar"
        :mine="false"
    />
    <div class="SGccCZRQ">
      <div class="htrebghe">
        <span class="hbtre">
          {{ props.user.username }}
        </span>
      </div>
    </div>
    <div class="regerg">
      <span class="htregy">
        {{ props.user.signature }}
      </span>
    </div>
    <el-button
        class="button"
        type="primary"
        size="default"
        round
        @click="handleRelation(props.user.id)"
        v-if="!IsFocus"
    >
      关注
    </el-button>
    <el-button
        class="button"
        type="primary"
        size="default"
        round
        @click="handleRelation(props.user.id)"
        v-else
    >
      已关注
    </el-button>
  </div>
</template>

<style scoped>
.hrbegw {
  align-items: center;
  margin: 17px 0;
  display: flex;
}

.hr4eg3rh {
  height: 70px;
  width: 70px;
}

.SGccCZRQ {
  width: 0;
  flex: 1;
  margin: 0 12px;
}

.htrebghe {
  margin: 4px 0;
  display: flex;
}

.hbtre {
  text-overflow: ellipsis;
  white-space: nowrap;
  cursor: pointer;
  font-family: PingFang SC, DFPKingGothicGB-Medium, sans-serif;
  font-size: 16px;
  font-weight: 500;
  line-height: 24px;
  overflow: hidden;
  color: #161823;
}

.htregy {
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #161823;
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
  font-size: 12px;
  font-weight: 400;
  line-height: 20px;
  overflow: hidden;
}

.regerg {
  color: #161823;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
  margin: 4px 0;
}
</style>
