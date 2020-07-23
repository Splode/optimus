<template>
    <section class="bg-gray-800 p-10 w-full">
        <header class="flex items-center justify-between w-full">
            <h1 class="font-medium text-2xl text-gray-100">Options</h1>
            <p @click="closeView">Close</p>
        </header>

        <div class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full">
            <h2 class="mb-3 text-green text-xl w-full">General</h2>
            <div class="flex items-center mr-6 my-2 px-4 text-gray-100">
                <p class="mr-4">Target</p>
                <Dropdown :options="targets"
                          :selected="target"
                          v-on:updateOption="selectTarget"
                          class="m-0 text-gray-200"></Dropdown>
            </div>
            <div class="flex flex-wrap items-center mr-8 my-2 px-4 text-gray-100">
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
            <div class="flex flex-wrap items-center my-2 px-4 text-gray-100">
                <label for="prefix">Prefix</label>
                <input type="text" id="prefix" v-model="config.prefix"
                       @change="setConfig"
                       class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green mx-4 px-4 py-2 rounded-md ta-color-slow"
                       maxlength="16">
            </div>
            <div class="flex flex-wrap items-center my-2 px-4 text-gray-100">
                <label for="suffix">Suffix</label>
                <input type="text" id="suffix" v-model="config.suffix"
                       @change="setConfig"
                       class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green mx-4 px-4 py-2 rounded-md ta-color-slow"
                       maxlength="16">
            </div>
        </div>

        <div class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full">
            <h2 class="mb-3 text-yellow text-xl w-full">WebP</h2>
            <div class="px-4 text-gray-100 w-1/2">
                <div class="flex items-center w-full">
                    <p class="mr-6">Quality</p>
                    <div class="w-full">
                        <vue-slider v-model="config.webpOpt.quality"
                                    @change="setConfig" class="slider-yellow"/>
                    </div>
                </div>
            </div>
            <div class="px-4 text-gray-100 w-1/2">
                <div class="flex items-center w-full">
                    <p class="mr-4">Lossless</p>
                    <div @click="toggleWebpLossless"
                         class="bg-gray-900 check-wrapper flex items-center justify-center rounded-md">
                        <transition name="fade" mode="out-in">
                            <svg v-if="config.webpOpt.lossless" version="1.1"
                                 id="check-icon"
                                 xmlns="http://www.w3.org/2000/svg" x="0px"
                                 y="0px"
                                 viewBox="0 0 24 24"
                                 style="enable-background:new 0 0 24 24;"
                                 width="24" height="24"
                                 xml:space="preserve">
                                <path fill="#ffe027"
                                      d="M10,15.6l-3.3-3.3l-1.4,1.4l4.7,4.7l9.7-9.7l-1.4-1.4L10,15.6z"/>
                            </svg>
                        </transition>
                    </div>
                </div>
            </div>
        </div>

        <div class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full">
            <h2 class="mb-3 text-purple-400 text-xl w-full">JPEG</h2>
            <div class="px-4 text-gray-100 w-1/2">
                <div class="flex items-center w-full">
                    <p class="mr-6">Quality</p>
                    <div class="w-full">
                        <vue-slider v-model="config.jpegOpt.quality"
                                    @change="setConfig" class="slider-purple"/>
                    </div>
                </div>
            </div>
        </div>

        <div class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full">
            <h2 class="mb-3 text-blue text-xl w-full">PNG</h2>
            <div class="px-4 text-gray-100 w-1/2">
                <div class="flex items-center w-full">
                    <p class="mr-6">Quality</p>
                    <div class="w-full">
                        <vue-slider v-model="config.pngOpt.quality"
                                    @change="setConfig" class="slider-blue"/>
                    </div>
                </div>
            </div>
        </div>

        <div class="w-full">
            <button class="btn border-gray-400 focus:outline-none hover:bg-gray-400 hover:text-gray-900 ml-auto ta-slow"
                    @click="restoreDefaults">Restore Defaults
            </button>
        </div>
    </section>
</template>

<script>
  import Dropdown from './Dropdown'
  import VueSlider from 'vue-slider-component'
  import 'vue-slider-component/theme/antd.css'

  export default {
    name: 'Settings',

    components: { Dropdown, VueSlider },

    data() {
      return {
        targets: [{ name: 'WebP', value: 'webp' }, {
          name: 'JPG',
          value: 'jpg'
        }, { name: 'PNG', value: 'png' }],
        value: 0
      }
    },

    computed: {
      /**
       * config returns the app configuration from state.
       * @returns {object}
       */
      config() {
        return this.$store.getters.config
      },

      target() {
        return this.targets.find(t => this.config.target === t.value)
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
       * restoreDefaults resets the app configuration to defaults.
       */
      restoreDefaults() {
        window.backend.Config.RestoreDefaults().then(() => {
          this.$store.dispatch('getConfig')
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
       * selectTarget updates the config target.
       * @param {object} e - The selected target.
       */
      selectTarget(e) {
        this.$store.dispatch('setConfigProp', { key: 'target', value: e.value })
          .then(() => {
            this.setConfig()
          })
          .catch(err => {
            console.error(err)
          })
      },

      /**
       * setConfig updates the configuration with the current state.
       */
      setConfig() {
        this.$store.dispatch('setConfig', this.config)
      },

      /**
       * toggleWebpLossless toggles the lossless property of the WebP config.
       */
      toggleWebpLossless() {
        this.$store.dispatch('toggleWebpLossless')
          .then(() => {
            this.setConfig()
          })
          .catch(err => {
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
        fill: #27ffa7;
    }

    input:focus {
        color: #27ffa7;
    }

    .check-wrapper {
        border: 2px solid transparent;
        cursor: pointer;
        width: 24px;
        height: 24px;
    }
</style>