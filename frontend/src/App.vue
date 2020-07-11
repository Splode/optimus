<template>
    <div id="app" class="bg-gray-800">
        <!--    <img alt="Wails logo" src="./assets/images/logo.png" class="logo" />-->
        <HelloWorld/>
        <!-- <router-link to="/">Home</router-link>
        <router-view></router-view> -->
    </div>
</template>

<script>
  import HelloWorld from './components/HelloWorld.vue'
  import './assets/css/main.css'

  export default {
    name: 'app',
    components: {
      HelloWorld
    }
  }
</script>
