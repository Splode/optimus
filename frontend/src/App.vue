<template>
  <div id="app" class="bg-gray-900 flex relative h-full">
    <Sidebar :active="currentView" v-on:select-view="handleViewSelect" />
    <keep-alive>
      <component :is="currentView" v-on:close-view="handleViewClose" />
    </keep-alive>
    <Notification />
  </div>
</template>

<script>
import About from './components/About.vue'
import Editor from './components/Editor.vue'
import Notification from './components/Notification.vue'
import Settings from './components/Settings.vue'
import Sidebar from './components/Sidebar.vue'
import Stats from './components/Stats.vue'
import './assets/css/main.css'

export default {
  name: 'app',

  components: {
    About,
    Editor,
    Notification,
    Settings,
    Sidebar,
    Stats
  },

  data() {
    return {
      currentView: 'Editor'
    }
  },

  methods: {
    handleViewClose() {
      this.currentView = 'Editor'
    },

    handleViewSelect(e) {
      if (this.currentView === e) {
        this.currentView = 'Editor'
      } else {
        this.currentView = e
      }
    }
  },

  mounted() {
    this.$store.dispatch('getConfig')
    this.$store.dispatch('getStats')
  }
}
</script>
