<template>
  <section
    class="bg-gray-800 overflow-y-auto p-10 w-full"
    id="settings"
    ref="section"
  >
    <header class="flex items-center justify-between w-full">
      <h1 class="font-medium text-2xl text-green">Options</h1>
      <BtnClose @click.native="closeView" />
    </header>

    <div
      class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full"
    >
      <h2 class="mb-3 text-gray-200 text-xl w-full">General</h2>
      <div class="flex items-center mr-6 my-2 px-4 text-gray-100">
        <p class="mr-4" v-tooltip.right-end="'Image conversion file target.'">
          Target
        </p>
        <Dropdown
          :options="targets"
          :selected="target"
          class="m-0 text-gray-200"
          v-on:update-option="selectTarget"
        ></Dropdown>
      </div>
      <div class="flex flex-wrap items-center mr-8 my-2 px-4 text-gray-100">
        <p v-tooltip="'The output directory for converted images.'">
          Destination
        </p>
        <p
          @click="selectOutDir"
          class="bg-gray-900 cursor-pointer font-mono hover:text-green mx-4 px-4 py-2 rounded-md ta-color-slow"
        >
          {{ config.outDir }}
        </p>
        <button
          @click="openDir"
          class="cursor-pointer"
          v-tooltip.right-end="'Open destination directory'"
        >
          <svg
            height="18"
            id="Layer_1"
            style="enable-background:new 0 0 24 24;"
            version="1.1"
            viewBox="0 0 24 24"
            width="20"
            x="0px"
            xml:space="preserve"
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            y="0px"
          >
            <path
              d="M20,3H4C2.9,3,2,3.9,2,5v14c0,1.1,0.9,2,2,2h5v-2H4V7h16v12h-5v2h5c1.1,0,2-0.9,2-2V5C22,3.9,21.1,3,20,3z"
              fill="#b3b3b3"
            />
            <path d="M13,21v-5h3l-4-5l-4,5h3v5H13z" fill="#b3b3b3" />
          </svg>
        </button>
      </div>
      <div class="flex flex-wrap items-center my-2 px-4 text-gray-100">
        <label
          for="prefix"
          v-tooltip.right-end="{
            content: 'A string to prepend to the converted image filename.'
          }"
          >Prefix</label
        >
        <input
          @change="setConfig"
          class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green mx-4 px-4 py-2 rounded-md ta-color-slow"
          id="prefix"
          maxlength="16"
          type="text"
          v-model="config.prefix"
        />
      </div>
      <div class="flex flex-wrap items-center my-2 px-4 text-gray-100">
        <label
          for="suffix"
          v-tooltip.right-end="{
            content: 'A string to append to the converted image filename.'
          }"
          >Suffix</label
        >
        <input
          @change="setConfig"
          class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green mx-4 px-4 py-2 rounded-md ta-color-slow"
          id="suffix"
          maxlength="16"
          type="text"
          v-model="config.suffix"
        />
      </div>
    </div>

    <div
      class="border-2 border-gray-700 flex flex-wrap my-4 px-2 py-4 rounded-md w-full"
    >
      <header class="flex px-2 w-full">
        <h2 class="text-gray-200 text-xl">Resizing</h2>
        <button
          @click="addSize"
          class="border-0 focus:outline-none hover:green hover:text-green ml-8 ta-slow"
          v-tooltip.right-end="{
            content: 'Add an image size to convert.',
            delay: 600
          }"
        >
          Add Size +
        </button>
      </header>
      <div
        :key="i"
        class="flex flex-wrap items-center my-2 px-4 text-gray-100 w-full"
        v-for="(s, i) in config.sizes"
      >
        <div class="mb-2 lg:mb-0 px-2 w-full sm:w-1/2 lg:w-auto">
          <label :for="`width-${i}`" class="block xl:inline-block mb-2 mr-4"
            ><span
              v-tooltip.right-end="
                'Width in pixels. Required and must be a positive integer.'
              "
              >Width</span
            ></label
          >
          <input
            :id="`width-${i}`"
            @blur="setConfig"
            @input="handleNumber"
            class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green px-4 py-2 rounded-md ta-color-slow w-full lg:w-auto"
            type="text"
            v-model.number="s.width"
          />
        </div>
        <div class="mb-2 lg:mb-0 px-2 w-full sm:w-1/2 lg:w-auto">
          <label :for="`height-${i}`" class="block xl:inline-block mb-2 mr-4"
            ><span
              v-tooltip.right-end="
                'Height in pixels. Required and must be a positive integer.'
              "
              >Height</span
            ></label
          >
          <input
            :id="`height-${i}`"
            @blur="setConfig"
            @input="handleNumber"
            class="bg-gray-900 cursor-pointer focus:outline-none hover:text-green px-4 py-2 rounded-md ta-color-slow w-full lg:w-auto"
            type="text"
            v-model.number="s.height"
          />
        </div>
        <div class="mb-2 lg:mb-0 px-2 w-full sm:w-1/2 lg:w-auto">
          <label class="block xl:inline-block mb-2 mr-4">Strategy</label>
          <div class="flex lg:inline-flex w-full lg:w-auto">
            <Dropdown
              :options="strategies"
              :selected="strategy(s)"
              @click.native="handleSelectStrategy(i)"
              class="my-0 mr-4 text-gray-200 w-full"
              v-on:update-option="selectStrategy"
              v-tooltip="strategyTooltip(s)"
            ></Dropdown>
            <button @click="removeSize(i)">
              <svg
                height="10"
                id="x"
                style="enable-background:new 0 0 11.9 11.9;"
                version="1.1"
                viewBox="0 0 11.9 11.9"
                width="10"
                x="0px"
                xml:space="preserve"
                xmlns="http://www.w3.org/2000/svg"
                xmlns:xlink="http://www.w3.org/1999/xlink"
                y="0px"
              >
                <path
                  d="M10.4,0L6,4.5L1.5,0L0,1.5L4.5,6L0,10.4l1.5,1.5L6,7.5l4.5,4.5l1.5-1.5L7.5,6l4.5-4.5L10.4,0z"
                  fill="#b3b3b3"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <div
      class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full"
    >
      <h2 class="mb-3 text-gray-200 text-xl w-full">WebP</h2>
      <div class="px-4 text-gray-100 w-1/2">
        <div class="flex items-center w-full">
          <p class="mr-6" v-tooltip.right-end="{ content: tooltips.quality }">
            Quality
          </p>
          <div class="w-full">
            <vue-slider @change="setConfig" v-model="config.webpOpt.quality" />
          </div>
        </div>
      </div>
      <div class="px-4 text-gray-100 w-1/2">
        <div class="flex items-center w-full">
          <p class="mr-4">Lossless</p>
          <div
            @click="toggleWebpLossless"
            class="bg-gray-900 check-wrapper flex items-center justify-center rounded-md"
          >
            <transition mode="out-in" name="fade">
              <svg
                height="24"
                id="check-icon"
                style="enable-background:new 0 0 24 24;"
                v-if="config.webpOpt.lossless"
                version="1.1"
                viewBox="0 0 24 24"
                width="24"
                x="0px"
                xml:space="preserve"
                xmlns="http://www.w3.org/2000/svg"
                y="0px"
              >
                <path
                  d="M10,15.6l-3.3-3.3l-1.4,1.4l4.7,4.7l9.7-9.7l-1.4-1.4L10,15.6z"
                  fill="#27ffa7"
                />
              </svg>
            </transition>
          </div>
        </div>
      </div>
    </div>

    <div
      class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full"
    >
      <h2 class="mb-3 text-gray-200 text-xl w-full">JPEG</h2>
      <div class="px-4 text-gray-100 w-1/2">
        <div class="flex items-center w-full">
          <p class="mr-6" v-tooltip.right-end="{ content: tooltips.quality }">
            Quality
          </p>
          <div class="w-full">
            <vue-slider @change="setConfig" v-model="config.jpegOpt.quality" />
          </div>
        </div>
      </div>
    </div>

    <div
      class="border-2 border-gray-700 flex flex-wrap my-4 p-4 rounded-md w-full"
    >
      <h2 class="mb-3 text-gray-200 text-xl w-full">PNG</h2>
      <div class="px-4 text-gray-100 w-1/2">
        <div class="flex items-center w-full">
          <p
            class="mr-6"
            title="Specify the image conversion quality."
            v-tooltip.right-end="{ content: tooltips.quality }"
          >
            Quality
          </p>
          <div class="w-full">
            <vue-slider @change="setConfig" v-model="config.pngOpt.quality" />
          </div>
        </div>
      </div>
    </div>

    <div class="mb-4 w-full">
      <button
        @click="restoreDefaults"
        class="border-gray-400 btn focus:outline-none hover:bg-green hover:border-green hover:text-gray-900 ml-auto ta-slow"
      >
        Restore Defaults
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
      strategies: [
        { name: 'Fill', value: 0 },
        { name: 'Fit', value: 1 },
        { name: 'Smart Crop', value: 2 }
      ],
      targets: [
        { name: 'WebP', value: 'webp' },
        {
          name: 'JPEG',
          value: 'jpg'
        },
        { name: 'PNG', value: 'png' }
      ],
      tooltips: {
        quality:
          'Image quality of the converted image where 0 is the lowest and 100 is the highest.'
      }
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
      return (
        this.targets.find(t => this.config.target === t.value) || {
          name: '',
          value: ''
        }
      )
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
      this.$store
        .dispatch('setSizeStrategy', {
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
     * @param {object} s - The current size.
     * @returns {object} The matching strategy.
     */
    strategy(s) {
      return this.strategies.find(o => s.strategy === o.value)
    },

    /**
     * strategyTooltip returns a tooltip content string based on the selected
     * strategy.
     * @param {object} s - Current size.
     * @returns string The tooltip content.
     */
    strategyTooltip(s) {
      switch (s.strategy) {
        case 0: {
          return 'Crop resizes the image to the specified dimensions, cropping excess from the center.'
        }
        case 1: {
          return "Fit resizes the image to fit within the specified dimensions, maintaining the image's aspect ratio."
        }
        default: {
          return ''
        }
      }
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
      window.backend.Config.OpenOutputDir()
        .then(res => {
          console.log(res)
        })
        .catch(err => {
          console.error(err)
        })
    },

    /**
     * restoreDefaults resets the app configuration to defaults.
     */
    restoreDefaults() {
      window.backend.Config.RestoreDefaults()
        .then(() => {
          this.$store.dispatch('getConfig')
        })
        .catch(err => {
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
      this.$store
        .dispatch('setConfigProp', { key: 'target', value: e.value })
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
      this.$store
        .dispatch('toggleWebpLossless')
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
  transition: fill 0.6s cubic-bezier(0.07, 0.95, 0, 1);
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
