import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    config: {
      outDir: '',
      target: ''
    }
  },
  getters: {
    config(state) {
      return state.config
    }
  },
  actions: {
    getConfig(context) {
      window.backend.Config.GetAppConfig().then(cfg => {
        context.commit('setConfig', cfg)
      }).catch(err => {
        console.error(err)
      })
    },
    setConfig(context, cfg) {
      context.commit('setConfig', cfg)
    }
  },
  mutations: {
    setConfig(state, cfg) {
      state.config = cfg
    }
  }
})

export default store
