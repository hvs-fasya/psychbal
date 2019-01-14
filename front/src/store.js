import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    user: null,
    loggedIn: false,
    socket: null
  },
  mutations: {
    setUser (state, user) {
      state.user = user
    },
    setLoggedIn (state) {
      state.loggedIn = true
    },
    setLoggedOut (state) {
      state.loggedIn = false;
      state.user = null
    }
  },
  actions: {
    setUser({ commit }, user) {
      commit('setUser', user)
    },
    setLoggedIn ({ commit }) {
      commit('setLoggedIn')
    },
    setLoggedOut ({ commit }) {
      commit('setLoggedIn')
    }
  }
})
