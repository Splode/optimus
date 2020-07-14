<template>
    <section>
        <h1>Options</h1>
        <p @click="closeView">Close</p>
    </section>
</template>

<script>
  export default {
    name: 'Settings',
    methods: {
      closeView() {
        this.$emit('close-view')
      }
    }
  }
</script>

<style scoped>

</style>