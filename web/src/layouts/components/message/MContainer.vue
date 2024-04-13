<script setup>
import { reactive, ref, watch } from 'vue'
import MBubble from '~/layouts/components/message/MBubble.vue'
import { useMessage } from '~/composables/messageManager.js'

const props = defineProps({
  message: ref([]),
  user: Object,
  sendMessage: Function,
  open: Boolean,
  empty: Boolean
})

const { delOpenUser } = useMessage()
const form = reactive({
  text: ''
})

const emit = defineEmits(['deleteOpenUser'])
const deleteOpenUser = () => {
  delOpenUser(props.user.id)
  emit('deleteOpenUser', props.user.id)
}

const Message = ref(props.message)

watch(() => props.message, (value) => {
  Message.value = value
})

const sendMessage = (toUser, content) => {
  props.sendMessage(toUser, content)
  form.text = ''
}
</script>

<template>
  <div>
    <el-container v-if="props.open">
      <el-header>
        <div class="azu4pICV">
          <div class="IrHbKVto">
            <div class="ZJm3Obdh">
              <span>{{props.user.username}}</span>
            </div>
          </div>
          <el-dropdown style="cursor: pointer;margin-right: 20px;">
            <font-awesome-icon
                :icon="['fas', 'ellipsis']"
                size="lg"
                class="begrdsg"
            />
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                    class="gbwvefs"
                    @click="deleteOpenUser"
                >
                  删除聊天
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>
      <el-main style="height: 330px;overflow: hidden; padding: 0 10px 0 26px">
        <el-scrollbar>
          <ul>
            <template v-for="item in Message" :key="item.id">
              <li>
                <m-bubble
                    :from-user="item.from_user"
                    :to-user="item.to_user"
                    :content="item.content"
                />
              </li>
            </template>
          </ul>
        </el-scrollbar>
      </el-main>
      <el-footer style="padding: initial">
        <el-input
            v-model="form.text"
            placeholder="发送消息"
            maxlength="100"
            class=" w-[100%] p-4"
        >
          <template #suffix>
            <el-button text class="w-[10px]" @click="sendMessage(props.user.id, form.text)">
              <font-awesome-icon :icon="['fas', 'paper-plane']"/>
            </el-button>
          </template>
        </el-input>
      </el-footer>
    </el-container>
    <el-empty v-if="props.empty" :image-size="200"  description="暂时没有消息哦"/>
  </div>
</template>

<style scoped>
.el-input__wrapper {
  @apply bg-gray-200 shadow-sm;
}

:deep(.el-input__wrapper.is-focus) {
  @apply border broder-gray-500/50;
}

.el-input {
  --el-input-focus-border-color: rgba(185, 185, 185) !important;
}

:deep(.el-input__inner) {
  @apply h-[45px];
}

.azu4pICV {
  width: 100%;
  height: 52px;
  color: #161823;
  flex-grow: 0;
  flex-shrink: 0;
  justify-content: space-between;
  align-items: center;
  font-size: 18px;
  font-weight: 500;
  display: flex;
}

.IrHbKVto {
  justify-content: flex-start;
  align-items: center;
  display: flex;
}

.ZJm3Obdh {
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #161823;
  flex-grow: 0;
  font-weight: 500;
  overflow: hidden;
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
}
.begrdsg {
  color: #c2c2c2;
}
.begrdsg:hover {
  color: #000000;
}
</style>
<style>
.el-tooltip__trigger:focus-visible {
  outline: none;
}
.el-tooltip__trigger {
  outline: none;
}
.gbwvefs {
  @apply font-bold;
  font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif !important;
  font-size: 14px;
}
.el-dropdown__popper {
  border-radius: 10px;
}
.el-dropdown-menu {
  border-radius: 10px;
}
.el-popper[data-popper-placement^=bottom]>.el-popper__arrow {
  top: initial;
}
.el-popper__arrow::before{
  display: none;
}
.el-dropdown-menu__item {
  border-radius: 10px;
  margin-left: 4px;
  margin-right: 4px;
}
</style>
