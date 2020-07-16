import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    config: {
      outDir: '',
      target: ''
    },
    stats: {
      byteCount: 0,
      imageCount: 0
    }
  },
  getters: {
    config(state) {
      return state.config
    },

    stats(state) {
      return state.stats
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
    },

    getStats(context) {
      window.backend.Stat.GetStats().then(s => {
        context.commit('setStats', s)
      }).catch(err => {
        console.error(err)
      })
    },
    setStats(context, s) {
      context.commit('setStats', s)
    }
  },
  mutations: {
    setConfig(state, cfg) {
      state.config = cfg
    },

    setStats(state, s) {
      state.stats = s
    }
  }
})

export default store
