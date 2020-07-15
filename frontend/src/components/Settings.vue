<template>
    <section class="p-10 w-full">
        <h1 class="font-medium text-5xl text-green">Options</h1>
        <p @click="closeView">Close</p>

        <p @click="openDir">{{ config.outDir }}</p>
        <label for="target">Target</label>
        <select name="target" id="target" @change="selectTarget">
            <option value="webp">WebP</option>
            <option value="jpg">JPG</option>
            <option value="png">PNG</option>
        </select>
        <button
                class="flex ml-auto bg-purple border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded-full text-gray-900"
                @click="selectOutDir"
        >
            Output Directory
        </button>
    </section>
</template>

<script>
  export default {
    name: 'Settings',

    computed: {
      /**
       * config returns the app configuration from state.
       * @returns {object}
       */
      config() {
        return this.$store.getters.config
      }
    },

    methods: {
      closeView() {
        this.$emit('close-view')
      },

      /**
       * openDir opens the configured output directory.
       */
      openDir() {
        window.backend.Config.OpenOutputDir().then(res => {
          console.log(res)
        }).catch(err => {
          console.error(err)
        })
      },


      /**
       * selectOutDir selects an output directory in the config.
       */
      selectOutDir() {
        window.backend.Config.SetOutDir()
          .then(res => {
            console.log(res)
            this.$store.dispatch('getConfig')
          })
          .catch(err => {
            console.error(err)
          })
      },

      /**
       * selectTarget selects the encoding target in the config.
       * @param {HTMLInputElement} e
       */
      selectTarget(e) {
        const t = e.target.value
        window.backend.Config.SetTarget(t).then(res => {
          console.log(res)
          this.$store.dispatch('getConfig')
        }).catch(err => {
          console.error(err)
        })
      }
    }
  }
</script>

<style scoped>

</style>