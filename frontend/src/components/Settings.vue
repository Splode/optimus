<template>
  <section class="bg-gray-800 overflow-y-auto p-10 w-full">
    <header class="flex items-center justify-between w-full">
      <h1 class="font-medium text-2xl text-green">Options</h1>
      <BtnClose @click.native="closeView"/>
    </header>

    <div
        class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full">
      <h2 class="mb-3 text-gray-200 text-xl w-full">General</h2>
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

    <div
        class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full">
      <h2 class="text-gray-200 text-xl">Resizing</h2>
      <button
          class="border-0 focus:outline-none hover:green hover:text-green ml-8 ta-slow"
          @click="addSize">Add Size +
      </button>
      <div v-for="(s, i) in config.sizes" :key="i"
           class="flex flex-wrap items-center my-2 px-4 text-gray-100 w-full">
        <label :for="`width-${i}`">Width</label>
        <input type="text" :id="`width-${i}`" v-model.number="s.width"
               @input="handleNumber" @blur="setConfig"
               class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green ml-4 mr-8 px-4 py-2 rounded-md ta-color-slow">
        <label :for="`height-${i}`">Height</label>
        <input type="text" :id="`height-${i}`" v-model.number="s.height"
               @input="handleNumber" @blur="setConfig"
               class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green ml-4 mr-8 px-4 py-2 rounded-md ta-color-slow">
        <p class="mr-4">Strategy</p>
        <Dropdown :options="strategies"
                  :selected="strategy(s)"
                  v-on:updateOption="selectStrategy"
                  @click.native="handleSelectStrategy(i)"
                  class="m-0 mr-6 text-gray-200"></Dropdown>
        <button @click="removeSize(i)">
          <svg version="1.1" id="x" xmlns="http://www.w3.org/2000/svg"
               xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
               viewBox="0 0 11.9 11.9"
               style="enable-background:new 0 0 11.9 11.9;"
               width="10" height="10"
               xml:space="preserve">
              <path fill="#b3b3b3"
                    d="M10.4,0L6,4.5L1.5,0L0,1.5L4.5,6L0,10.4l1.5,1.5L6,7.5l4.5,4.5l1.5-1.5L7.5,6l4.5-4.5L10.4,0z"/>
          </svg>
        </button>
      </div>
    </div>

    <div
        class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full">
      <h2 class="mb-3 text-gray-200 text-xl w-full">WebP</h2>
      <div class="px-4 text-gray-100 w-1/2">
        <div class="flex items-center w-full">
          <p class="mr-6">Quality</p>
          <div class="w-full">
            <vue-slider v-model="config.webpOpt.quality"
                        @change="setConfig"/>
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
                  <path fill="#27ffa7"
                        d="M10,15.6l-3.3-3.3l-1.4,1.4l4.7,4.7l9.7-9.7l-1.4-1.4L10,15.6z"/>
              </svg>
            </transition>
          </div>
        </div>
      </div>
    </div>

    <div
        class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full">
      <h2 class="mb-3 text-gray-200 text-xl w-full">JPEG</h2>
      <div class="px-4 text-gray-100 w-1/2">
        <div class="flex items-center w-full">
          <p class="mr-6">Quality</p>
          <div class="w-full">
            <vue-slider v-model="config.jpegOpt.quality"
                        @change="setConfig"/>
          </div>
        </div>
      </div>
    </div>

    <div
        class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full">
      <h2 class="mb-3 text-gray-200 text-xl w-full">PNG</h2>
      <div class="px-4 text-gray-100 w-1/2">
        <div class="flex items-center w-full">
          <p class="mr-6">Quality</p>
          <div class="w-full">
            <vue-slider v-model="config.pngOpt.quality"
                        @change="setConfig"/>
          </div>
        </div>
      </div>
    </div>

    <div class="mb-4 w-full">
      <button
          class="border-gray-400 btn focus:outline-none hover:bg-green hover:border-green hover:text-gray-900 ml-auto ta-slow"
          @click="restoreDefaults">Restore Defaults
      </button>
    </div>
  </section>
</template>

<script>
import BtnClose from './BtnClose'
import Dropdown from './Dropdown'
import VueSlider from 'vue-slider-component'
import 'vue-slider-component/theme/antd.css'
import { EventBus } from '@/lib/event-bus'

export default {
  name: 'Settings',

  components: { BtnClose, Dropdown, VueSlider },

  data() {
    return {
      activeStrategy: 0,
      targets: [{ name: 'WebP', value: 'webp' }, {
        name: 'JPEG',
        value: 'jpg'
      }, { name: 'PNG', value: 'png' }],
      strategies: [{ name: 'Crop', value: 0 }, { name: 'Fit', value: 1 }]
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
      return this.targets.find(t => this.config.target === t.value) || {
        name: '',
        value: ''
      }
    }
  },

  methods: {
    /**
     * addSize adds a blank size to the size config.
     */
    addSize() {
      this.$store.dispatch('addSize')
    },

    /**
     * removeSize removes the selected size rect from the sizes config.
     * @param {number} i - The size index.
     */
    removeSize(i) {
      this.$store.dispatch('removeSize', i)
      this.setConfig()
    },

    /**
     * handleSelectStrategy records the index of the currently active size input.
     * @param {number} i - The size index.
     */
    handleSelectStrategy(i) {
      this.activeStrategy = i
    },

    /**
     * selectStrategy updates the configured size with the selected strategy.
     * @param {object} e - The selected strategy.
     */
    selectStrategy(e) {
      this.$store.dispatch('setSizeStrategy', {
        index: this.activeStrategy,
        value: e.value
      })
          .then(() => {
            this.setConfig()
          })
          .catch(err => {
            console.log(err)
          })
    },

    /**
     * strategy returns the strategy from the list of strategies matching the
     * current strategy.
     * @param {object} st - The current strategy.
     * @returns {object} The matching strategy.
     */
    strategy(st) {
      return this.strategies.find(o => st.strategy === o.value || {
        name: '',
        value: ''
      })
    },

    /**
     * closeView closes the current view.
     */
    closeView() {
      this.$emit('close-view')
    },

    /**
     * handleNumber validates a number input and displays an error message if
     * the input cannot be parsed.
     * @param {InputEvent} e
     */
    handleNumber(e) {
      if (!e.data) return
      const n = parseInt(e.data, 10)
      if (isNaN(n)) {
        EventBus.$emit('notify', {
          msg: `Image size must be a number.`,
          type: 'warn'
        })
      }
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
input:focus {
  color: #27ffa7;
}

button > svg > path {
  transition: fill .6s cubic-bezier(.07, .95, 0, 1);
}

button:hover > svg > path {
  fill: #27ffa7;
}

.check-wrapper {
  border: 2px solid transparent;
  cursor: pointer;
  width: 24px;
  height: 24px;
}
</style>