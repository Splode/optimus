import Vue from 'vue'
import VueRouter from 'vue-router'
import HelloWorld from './components/Editor.vue'

Vue.use(VueRouter)

const routes = [{ component: HelloWorld, name: 'Home', path: '/' }]

const router = new VueRouter({
  mode: 'abstract',
  routes
})

export default router
