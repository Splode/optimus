import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    config: {
      outDir: ''
    }
  },
  getters: {
    config(state) {
      return state.config
    }
  }
})

export default store
