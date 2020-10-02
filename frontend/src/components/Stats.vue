<template>
  <section class="bg-gray-800 overflow-y-auto p-10 w-full">
    <header class="flex items-center justify-between w-full">
      <h1 class="font-medium text-2xl text-blue">Stats</h1>
      <BtnClose @click.native="closeView" color="blue" />
    </header>
    <div class="mt-6 w-full">
      <h2 class="mb-3 text-gray-300 text-xl w-full">Current Session</h2>
      <div class="flex flex-wrap w-full">
        <div class="w-full sm:w-1/3 xl:w-64">
          <p
            class="font-bold leading-none text-4xl md:text-5xl text-green tracking-tight"
          >
            {{ session.savings }}
          </p>
          <p class="font-medium text-gray-200 tracking-wider uppercase">
            Saved
          </p>
        </div>
        <div class="w-full sm:w-1/3 xl:w-64">
          <p
            class="font-bold leading-none text-4xl md:text-5xl text-blue tracking-tight"
          >
            {{ session.count }}
          </p>
          <p class="font-medium text-gray-200 tracking-wider uppercase">
            Images Optimized
          </p>
        </div>
        <div class="w-full sm:w-1/3 xl:w-64">
          <p
            class="font-bold leading-none text-4xl md:text-5xl text-purple tracking-tight"
          >
            {{ session.time[0] }}
          </p>
          <p class="font-medium text-gray-200 tracking-wider uppercase">
            {{ session.time[1] }}
          </p>
        </div>
      </div>
    </div>
    <div class="mt-8 w-full">
      <h2 class="mb-3 text-gray-300 text-xl w-full">All Time</h2>
      <div class="flex flex-wrap w-full">
        <div class="w-full sm:w-1/3 xl:w-64">
          <p
            class="font-bold leading-none text-4xl md:text-5xl text-orange tracking-tight"
          >
            {{ stats.byteCount }}
          </p>
          <p class="font-medium text-gray-200 tracking-wider uppercase">
            Saved
          </p>
        </div>
        <div class="w-full sm:w-1/3 xl:w-64">
          <p
            class="font-bold leading-none text-4xl md:text-5xl text-yellow tracking-tight"
          >
            {{ stats.imageCount }}
          </p>
          <p class="font-medium text-gray-200 tracking-wider uppercase">
            Images Optimized
          </p>
        </div>
        <div class="w-full sm:w-1/3 xl:w-64">
          <p
            class="font-bold leading-none text-4xl md:text-5xl text-pink tracking-tight"
          >
            {{ stats.timeCount[0] }}
          </p>
          <p class="font-medium text-gray-200 tracking-wider uppercase">
            {{ stats.timeCount[1] }}
          </p>
        </div>
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
