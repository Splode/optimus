import Vue from 'vue'
import Vuex from 'vuex'
import { fSize } from './lib/file'
import { prettyTime } from './lib/time'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    config: {
      outDir: '',
      target: '',
      prefix: '',
      suffix: '',
      sizes: [],
      jpegOpt: { quality: 0 },
      pngOpt: { quality: 0 },
      webpOpt: { lossless: false, quality: 0 }
    },
    stats: {
      byteCount: 0,
      imageCount: 0,
      timeCount: 0
    },
    session: {
      count: 0,
      savings: 0,
      time: 0
    }
  },
  getters: {
    config(state) {
      return state.config
    },

    session(state) {
      return {
        count: state.session.count,
        hasSavings: state.session.savings > 0,
        savings: fSize(state.session.savings),
        time: prettyTime(state.session.time)
      }
    },

    stats(state) {
      return {
        byteCount: fSize(state.stats.byteCount),
        imageCount: state.stats.imageCount,
        timeCount: prettyTime(state.stats.timeCount)
      }
    }
  },
  actions: {
    addSize(context) {
      context.commit('addSize')
    },
    removeSize(context, index) {
      context.commit('removeSize', index)
    },
    setSizeStrategy(context, payload) {
      context.commit('setSizeStrategy', payload)
    },

    getConfig(context) {
      window.backend.Config.GetAppConfig()
        .then(cfg => {
          context.commit('setConfig', cfg)
        })
        .catch(err => {
          console.error(err)
        })
    },

    setConfig(context, c) {
      window.backend.Config.SetConfig(JSON.stringify(c))
        .then(() => {
          context.dispatch('getConfig')
        })
        .catch(err => {
          console.error(err)
        })
    },

    setConfigProp(context, payload) {
      context.commit('setConfigProp', payload)
    },

    setSessionProp(context, payload) {
      context.commit('setSessionProp', payload)
    },

    toggleWebpLossless(context) {
      context.commit('toggleWebpLossless')
    },

    getStats(context) {
      window.backend.Stat.GetStats()
        .then(s => {
          context.commit('setStats', s)
        })
        .catch(err => {
          console.error(err)
        })
    },
    setStats(context, s) {
      context.commit('setStats', s)
    }
  },
  mutations: {
    addSize(state) {
      const s = { height: null, width: null, strategy: 0 }
      if (!state.config.sizes) {
        state.config.sizes = [s]
      } else {
        state.config.sizes.push(s)
      }
    },
    removeSize(state, index) {
      state.config.sizes.splice(index, 1)
    },
    setSizeStrategy(state, payload) {
      state.config.sizes[payload.index].strategy = payload.value
    },

    setConfig(state, c) {
      state.config = c
    },

    setConfigProp(state, payload) {
      state.config[payload.key] = payload.value
    },

    setSessionProp(state, payload) {
      state.session[payload.key] += payload.value
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
