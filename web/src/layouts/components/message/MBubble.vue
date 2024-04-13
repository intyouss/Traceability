<script setup>
import UAvatar from '~/layouts/components/user/UAvatar.vue'

const props = defineProps({
  fromUser: Object,
  toUser: Object,
  content: String
})
</script>

<template>
  <div class="jytj">
    <template v-if="props.fromUser.id !== $store.state.user.id">
      <div class="to_user">
        <u-avatar
            :user-id="0"
            :avatar="props.fromUser.avatar"
            :mine="false"
            class="w-[30px] h-[30px]"
        />
        <div class="bubble">
          <div class="kuiy">
            <div class="nyttr">
              <span class="trhb">{{props.content}}</span>
            </div>
          </div>
        </div>
      </div>
    </template>
    <template v-else>
      <div class="from_user">
        <div class="bubble1">
          <div class="uyjy">
            <div class="nyttr">
              <span class="trhb">{{props.content}}</span>
            </div>
          </div>
        </div>
        <u-avatar
            :user-id="0"
            :avatar="props.fromUser.avatar"
            :mine="false"
            class="gewf"
        />
      </div>

    </template>
  </div>
</template>

<style scoped>
.to_user {
  justify-content: flex-start;
  margin-bottom: 18px;
  display: flex;
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
  position: relative;
  flex-grow: 1;
}
.from_user {
  justify-content: flex-end;
  margin-bottom: 18px;
  display: flex;
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
  position: relative;
  flex-grow: 1;
}
.bubble {
  color: #fff;
  max-width: 224px;
  color: rgba(22, 24, 35, .75);
  margin-left: 8px;
  margin-right: 8px;
  font-size: 13px;
  font-weight: 400;
  position: relative;
}
.bubble1 {
  color: #fff;
  max-width: 224px;
  color: rgba(22, 24, 35, .75);
  margin-left: 8px;
  margin-right: 8px;
  font-size: 13px;
  font-weight: 400;
  position: relative;
  order: -1;
}
.uyjy {
  align-items: flex-end;
  flex-direction: column;
  display: flex;
}
.kuiy {
  align-items: flex-start;
  flex-direction: column;
  display: flex;
}
.nyttr {
  color: #fff;
  background-color: #298cff;
  word-wrap: break-word;
  word-break: break-all;
  max-width: 224px;
  border-radius: 10px;
  padding: 8px 14px;
  font-size: 14px;
  line-height: 21px;
  display: inline-block;
}
.trhb {
  font-size: 13px;
  line-height: 21px;
}
.jytj {
  display: flex;
  justify-content: space-between;
  position: relative;
}
.gewf {
  @apply w-[30px] h-[30px];
  margin-right: 20px;
}
</style>
