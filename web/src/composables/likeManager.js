import { ref } from 'vue'
import { useStore } from 'vuex'
import { notify } from '~/composables/util.js'
import { likeAction } from '~/api/like.js'
import { useRouter } from 'vue-router'

export function useLike (isLike) {
  const IsLiked = ref(isLike)
  const store = useStore()
  const router = useRouter()
  const handleLikeAction = (videoId) => {
    if (store.state.user.id) {
      // 点赞/取消点赞
      const actionType = IsLiked.value ? 2 : 1
      likeAction(videoId, actionType).then(() => {
        IsLiked.value = !IsLiked.value
      })
    } else {
      notify('请先登录', 'warning')
      router.push('/login').then(() => {})
    }
  }
  return {
    IsLiked,
    handleLikeAction
  }
}
