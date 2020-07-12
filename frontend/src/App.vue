<template>
    <div id="app" class="bg-gray-800">
        <!--    <img alt="Wails logo" src="./assets/images/logo.png" class="logo" />-->
        <Editor/>
        <!-- <router-link to="/">Home</router-link>
        <router-view></router-view> -->
    </div>
</template>

<script>
  import Editor from './components/Editor.vue'
  import './assets/css/main.css'

  export default {
    name: 'app',
    components: {
      Editor
    },

    mounted() {
      this.$store.dispatch('getConfig')
    }
  }
</script>
