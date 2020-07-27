<template>
    <section class="bg-gray-800 overflow-y-auto p-10 w-full">
        <header class="flex items-center justify-between w-full">
            <h1 class="font-medium text-2xl text-green">Stats</h1>
            <BtnClose @click.native="closeView"/>
        </header>
        <div class="w-full">
            <h2>Current Session</h2>
            <div>
                <p>{{ session.savings }}</p>
                <p>Saved</p>
            </div>
            <div>
                <p>{{ session.count }}</p>
                <p>Images Optimized</p>
            </div>
            <div>
                <p>{{ session.time[0] }}</p>
                <p>{{ session.time[1] }}</p>
            </div>
        </div>
        <div class="w-full">
            <h2>All Time</h2>
            <div>
                <p>{{ stats.byteCount }}</p>
                <p>Saved</p>
            </div>
            <div>
                <p>{{ stats.imageCount }}</p>
                <p>Images Optimized</p>
            </div>
            <div>
                <p>{{ stats.timeCount[0] }}</p>
                <p>{{ stats.timeCount[1] }}</p>
            </div>
        </div>
    </section>
</template>

<script>
  import BtnClose from './BtnClose'

  export default {
    name: 'Stats',

    components: { BtnClose },

    computed: {
      session() {
        return this.$store.getters.session
      },
      stats() {
        return this.$store.getters.stats
      }
    },

    methods: {
      closeView() {
        this.$emit('close-view')
      }
    }
  }
</script>
