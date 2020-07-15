<template>
    <div id="app" class="bg-gray-900 flex h-full">
        <Sidebar v-on:select-view="handleViewSelect"/>
        <keep-alive>
            <transition name="fade-fast" mode="out-in">
                <component :is="currentView" v-on:close-view="handleViewClose"/>
            </transition>
        </keep-alive>
    </div>
</template>

<script>
  import About from './components/About.vue'
  import Editor from './components/Editor.vue'
  import Settings from './components/Settings.vue'
  import Sidebar from './components/Sidebar.vue'
  import './assets/css/main.css'

  export default {
    name: 'app',

    components: {
      About,
      Editor,
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
