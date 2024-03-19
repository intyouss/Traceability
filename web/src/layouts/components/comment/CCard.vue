<script setup>
import UAvatar from '~/layouts/components/user/UAvatar.vue';

const props = defineProps({
  avatar: String,
  userId: Number,
  username: String,
  content: String,
  time: String,
});
</script>

<template>
  <el-card>
    <div class="vega">
      <div class="bwged">
        <u-avatar class="veu" :user-id="props.userId" :avatar="props.avatar" />
      </div>
      <div class="gewfas">
        <div class="gewfed">
          <div class="gewfa">
            <div class="hrgeuif">
              {{props.username}}
            </div>
          </div>
          <div class="brwefe">
          <span class="grwefs">
            {{props.content}}
          </span>
          </div>
          <div class="bregw">
            <span>
              {{ props.time }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </el-card>
</template>

<style scoped>
.vega {
  margin-top: 9px;
  display: flex;
  position: relative;
}
.bwged {
  margin-right: 8px;
  cursor: pointer;
  z-index: 1;
}
.gewfas {
  width: 0;
  flex-grow: 1;
}
.gewfed {
  padding: 4px 0;
  position: relative;
}
.gewfa {
  align-items: center;
  margin-bottom: 4px;
  display: flex;
  flex: 1;
  margin-right: 12px;
  position: relative;
  overflow: hidden;
}
.hrgeuif {
  max-width: 100%;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: rgba(22, 24, 35, .6);
  margin-right: 8px;
  line-height: 20px;
  overflow: hidden;
  font-size: 12px;
}
.brwefe {
  text-align: left;
  word-break: break-all;
  position: relative;
}
.grwefs {
  margin-top: 8px;
  margin-right: 8px;
  font-size: 13px;
}
.bregw {
  color: rgba(22, 24, 35, .6);
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
  font-size: 12px;
  font-weight: 400;
  line-height: 20px;
  position: relative;
}
.veu {
  width: 40px;
  height: 40px;
}
:deep(.el-card__body) {
  padding: 5px 10px 5px 10px;
}
</style>
