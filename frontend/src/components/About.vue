<template>
    <section>
        <h1>About</h1>
        <p @click="closeView">Close</p>
    </section>
</template>

<script>
  export default {
    name: 'About',
    methods: {
      closeView() {
        this.$emit('close-view')
      }
    }
  }
</script>

<style scoped>

</style>