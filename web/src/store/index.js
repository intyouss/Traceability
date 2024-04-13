import { createStore } from 'vuex'
import { getInfo, login, register } from '~/api/user.js'
import { setToken, removeToken } from '~/composables/auth.js'

// 创建一个新的 store 实例
const store = createStore({
  state () {
    return {
      // 用户信息
      user: {},
      // 侧边栏宽度
      asideWidth: '120px',
      tagListWidth: '150px'
    }
  },
  mutations: {
    // 记录用户信息
    setUserInfo (state, user) {
      state.user = user
    },
    // 侧边栏宽度
    handleAsideWidth (state) {
      state.asideWidth = state.asideWidth === '120px' ? '64px' : '120px'
      state.tagListWidth = state.tagListWidth === '150px' ? '95px' : '150px'
    }
  },
  actions: {
    // 用户登录, 保留token，保存用户信息
    login ({ commit }, { username, password }) {
      return new Promise((resolve, reject) => {
        login(username, password).then((res) => {
          setToken(res.data.token)
          commit('setUserInfo', res.data.user)
          resolve(res)
        }).catch((err) => reject(err))
      })
    },
    register ({ commit }, { username, password }) {
      return new Promise((resolve, reject) => {
        register(username, password).then((res) => {
          setToken(res.data.token)
          commit('setUserInfo', res.data.user)
          resolve(res)
        }).catch((err) => reject(err))
      })
    },
    // 获取用户信息并保存
    getUserInfo ({ commit }) {
      return new Promise((resolve, reject) => {
        getInfo(0).then((res) => {
          commit('setUserInfo', res.data.user)
          resolve(res)
        }).catch((err) => reject(err))
      })
    },
    logout ({ commit }) {
      removeToken()
      commit('setUserInfo', {})
    }
  }
})

export default store
