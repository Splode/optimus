import 'core-js/stable'
import 'regenerator-runtime/runtime'
import Vue from 'vue'
import App from './App.vue'
// import router from './router'
import store from './store'
import * as Wails from '@wailsapp/runtime'

Vue.config.productionTip = false
Vue.config.devtools = true

Wails.Init(() => {
  new Vue({
    render: h => h(App),
    // router,
    store
  }).$mount('#app')
})
