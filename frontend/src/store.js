import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    config: {
      outDir: '',
      target: '',
      prefix: '',
      suffix: '',
      jpegOpt: { quality: 0 },
      pngOpt: { quality: 0 },
      webpOpt: { lossless: false, quality: 0 }
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
    setConfig(context, c) {
      window.backend.Config.SetConfig(JSON.stringify(c)).then(() => {
        context.dispatch('getConfig')
      }).catch(err => {
        console.error(err)
      })
    },

    setConfigProp(context, payload) {
      context.commit('setConfigProp', payload)
    },

    toggleWebpLossless(context) {
      context.commit('toggleWebpLossless')
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
    setConfig(state, c) {
      state.config = c
    },

    setConfigProp(state, payload) {
      state.config[payload.key] = payload.value
    },

    setStats(state, s) {
      state.stats = s
    },

    toggleWebpLossless(state) {
      state.config.webpOpt.lossless = !state.config.webpOpt.lossless
    }
  }
})

export default store
