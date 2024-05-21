import { loginApi, getInfoApi } from '@/api/user'
const state = () => ({
  token: '', // 登录token
  info: {} // 用户信息
})

// getters
const getters = {
  token (state) {
    return state.token
  },
  info (state) {
    return state.info
  }
}

// mutations
const mutations = {
  tokenChange (state, token) {
    state.token = token
  },
  infoChange (state, info) {
    state.info = info
  }
}

// actions
const actions = {
  // login by login.vue
  login ({ commit, dispatch }, params) {
    return new Promise((resolve, reject) => {
      loginApi(params)
        .then(res => {
          commit('tokenChange', res.data.token)
          commit('infoChange', res.data.user)
          resolve(res.data.token)
        }).catch(err => {
          reject(err)
        })
    })
  },

  getInfo ({ commit }, params) {
    return new Promise((resolve, reject) => {
      getInfoApi(params)
        .then(res => {
          commit('infoChange', res.data.user)
          resolve(res.data.user)
        })
    })
  },

  loginOut ({ commit }) {
        commit('tokenChange', '')
        commit('infoChange', {})
        localStorage.removeItem('tabs')
        localStorage.removeItem('vuex')
        sessionStorage.removeItem('vuex')
        location.reload()
  }
}

export default {
  namespaced: true,
  state,
  actions,
  getters,
  mutations
}
