<template>
    <div id="app" class="bg-gray-900 flex relative h-full">
        <Sidebar v-on:select-view="handleViewSelect"/>
        <transition name="fade-fast" mode="out-in">
            <keep-alive>
                <component :is="currentView" v-on:close-view="handleViewClose"/>
            </keep-alive>
        </transition>
        <Notification/>
    </div>
</template>

<script>
  import About from './components/About.vue'
  import Editor from './components/Editor.vue'
  import Notification from './components/Notification.vue'
  import Settings from './components/Settings.vue'
  import Sidebar from './components/Sidebar.vue'
  import './assets/css/main.css'

  export default {
    name: 'app',

    components: {
      About,
      Editor,
      Notification,
      Settings,
      Sidebar
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
        this.currentView = e
      }
    },

    mounted() {
      this.$store.dispatch('getConfig')
    }
  }
</script>
