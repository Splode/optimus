<template>
    <section class="bg-gray-800 p-10 w-full">
        <header class="flex items-center justify-between w-full">
            <h1 class="font-medium text-2xl text-yellow">Options</h1>
            <p @click="closeView">Close</p>
        </header>

        <div class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full">
            <h2 class="mb-3 text-gray-200 text-xl w-full">General</h2>
            <div class="flex items-center mr-6 my-2">
                <label for="target">Target</label>
                <select name="target" id="target" v-model="config.target"
                        class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green mx-4 px-3 py-2 rounded-md ta-color-slow"
                        @change="selectTarget">
                    <option value="webp">WebP</option>
                    <option value="jpg">JPG</option>
                    <option value="png">PNG</option>
                </select>
            </div>
            <div class="flex flex-wrap items-center mr-8 my-2">
                <p>Destination</p>
                <p class="bg-gray-900 cursor-pointer font-mono hover:text-green mx-4 px-4 py-2 rounded-md ta-color-slow"
                   @click="selectOutDir"
                >{{ config.outDir }}</p>
                <button class="cursor-pointer"
                        title="Open destination directory" @click="openDir">
                    <svg version="1.1" id="Layer_1"
                         xmlns="http://www.w3.org/2000/svg"
                         xmlns:xlink="http://www.w3.org/1999/xlink" x="0px"
                         y="0px"
                         viewBox="0 0 24 24"
                         style="enable-background:new 0 0 24 24;"
                         xml:space="preserve" width="20" height="18">
                    <path fill="#b3b3b3"
                          d="M20,3H4C2.9,3,2,3.9,2,5v14c0,1.1,0.9,2,2,2h5v-2H4V7h16v12h-5v2h5c1.1,0,2-0.9,2-2V5C22,3.9,21.1,3,20,3z"/>
                        <path fill="#b3b3b3" d="M13,21v-5h3l-4-5l-4,5h3v5H13z"/>
                </svg>
                </button>
            </div>
            <div class="flex flex-wrap items-center my-2">
                <label for="prefix">Prefix</label>
                <input type="text" id="prefix"
                       class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green mx-4 px-4 py-2 rounded-md ta-color-slow"
                       maxlength="16">
            </div>
            <div class="flex flex-wrap items-center my-2">
                <label for="suffix">Suffix</label>
                <input type="text" id="suffix"
                       class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green mx-4 px-4 py-2 rounded-md ta-color-slow"
                       maxlength="16">
            </div>
        </div>
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
    button > svg > path {
        transition: fill .3s cubic-bezier(.07, .95, 0, 1);
    }

    button:hover > svg > path {
        fill: #07fdbc;
    }

    input:focus {
        color: #07fdbc;
    }
</style>