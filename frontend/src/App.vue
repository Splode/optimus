<template>
    <div id="app" class="bg-gray-900 flex h-full">
        <Sidebar/>
        <!--    <img alt="Wails logo" src="./assets/images/logo.png" class="logo" />-->
        <Editor/>
    </div>
</template>

<script>
  import Editor from './components/Editor.vue'
  import Sidebar from './components/Sidebar.vue'
  import './assets/css/main.css'

  export default {
    name: 'app',
    components: {
      Editor,
      Sidebar
    },

    mounted() {
      this.$store.dispatch('getConfig')
    }
  }
</script>
