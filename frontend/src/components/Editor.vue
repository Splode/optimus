<template>
  <section class="p-10 w-full">
    <header class="border-b-2 border-gray-800 flex flex-wrap" ref="header">
      <div class="w-full md:w-1/2 lg:w-5/12">
        <div
          class="bg-gray-800 border-2 border-dashed cursor-pointer drop-zone flex flex-col items-center justify-center py-6 md:py-10 ta-slow rounded-sm h-full"
          :class="isDragging ? 'border-green' : 'border-gray-400'"
          ref="dropZone"
        >
          <svg
            version="1.1"
            id="dropZone-plus"
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            x="0px"
            y="0px"
            viewBox="0 0 48 48"
            enable-background="new 0 0 48 48"
            width="48px"
            height="48px"
            xml:space="preserve"
          >
            <path
              fill="#808080"
              d="M47,20.6H28.4c-0.6,0-1-0.4-1-1V1c0-0.6-0.4-1-1-1h-4.9c-0.6,0-1,0.4-1,1v18.6c0,0.6-0.4,1-1,1H1
                c-0.6,0-1,0.4-1,1v4.9c0,0.6,0.4,1,1,1h18.6c0.6,0,1,0.4,1,1V47c0,0.6,0.4,1,1,1h4.9c0.6,0,1-0.4,1-1V28.4c0-0.6,0.4-1,1-1H47
                c0.6,0,1-0.4,1-1v-4.9C48,21,47.6,20.6,47,20.6z"
            />
          </svg>
          <p class="mt-6 text-gray-200">Drag and drop or select images</p>
        </div>
      </div>
      <div class="lg:w-7/12 md:pl-6 md:w-1/2 my-3 w-full">
        <transition name="fade" mode="out-in">
          <div
            v-if="!session.count"
            key="intro"
            class="flex h-full items-center justify-center"
          >
            <h2 class="leading-none text-4xl text-center text-yellow">
              Add image files<br />to get started
            </h2>
          </div>
          <div v-else key="stats" class="flex flex-wrap items-end h-full">
            <div class="px-3 w-full lg:w-5/12">
              <div v-if="session.hasSavings">
                <h2
                  class="font-bold leading-none text-4xl md:text-5xl text-green tracking-tight"
                >
                  {{ session.savings }}
                </h2>
                <p class="font-medium text-gray-300 tracking-wider uppercase">
                  Saved
                </p>
              </div>
              <div v-else>
                <h2
                  class="font-bold leading-none text-5xl text-pink tracking-tight uppercase"
                >
                  No Savings
                </h2>
                <p class="font-medium text-gray-300 tracking-wider uppercase">
                  Try adjusting options
                </p>
              </div>
            </div>
            <div class="flex lg:flex-col justify-between px-3 w-full lg:w-3/12">
              <div class="w-1/2 lg:w-full">
                <p class="font-bold text-xl lg:text-2xl text-blue">
                  {{ session.count }}
                </p>
                <p
                  class="font-medium text-gray-300 text-sm tracking-wider uppercase"
                >
                  {{ session.count > 1 ? 'Images' : 'Image' }}
                </p>
              </div>
              <div class="w-1/2 lg:w-full">
                <p class="font-bold text-xl lg:text-2xl text-yellow">
                  {{ session.time[0] }}
                </p>
                <p
                  class="font-medium text-gray-300 text-sm tracking-wider uppercase"
                >
                  {{ session.time[1] }}
                </p>
              </div>
            </div>
            <div class="flex lg:flex-col justify-between px-3 w-full lg:w-4/12">
              <div class="w-1/2 lg:w-full">
                <p class="font-bold text-xl lg:text-2xl text-pink">
                  {{ totalStats.byteCount }}
                </p>
                <p
                  class="font-medium text-gray-300 text-sm tracking-wider uppercase"
                >
                  All time Savings
                </p>
              </div>
              <div class="w-1/2 lg:w-full">
                <p class="font-bold text-xl lg:text-2xl text-purple">
                  {{ totalStats.imageCount }}
                </p>
                <p
                  class="font-medium text-gray-300 text-sm tracking-wider uppercase"
                >
                  All time Images
                </p>
              </div>
            </div>
          </div>
        </transition>
      </div>
      <footer class="w-full">
        <section
          class="flex justify-between lg:w-5/12 md:w-1/2 py-3 md:py-6 w-full"
        >
          <button
            class="btn focus:outline-none ta-slow"
            :class="
              canConvert
                ? 'border-green hover:bg-green hover:text-gray-900 text-gray-200'
                : 'btn--disabled'
            "
            @click="convert"
            :disabled="!canConvert"
          >
            {{ optBtnTxt }}
          </button>
          <button
            class="btn focus:outline-none ta-slow"
            :class="
              canClear
                ? 'border-gray-400 hover:bg-gray-800 hover:text-green'
                : 'btn--disabled'
            "
            @click="clear"
            :disabled="!canClear"
          >
            Clear
          </button>
        </section>
      </footer>
    </header>
    <input
      type="file"
      accept="image/jpeg, image/jpg, image/png, image/webp"
      multiple
      class="hidden"
      @input="handleFileInput"
      ref="fileInput"
    />
    <!-- file table -->
    <transition name="fade" mode="out-in">
      <div
        v-if="files.length > 0"
        class="table-wrapper"
        :style="{ height: `calc(100vh - ${headerHeight + 80}px)` }"
      >
        <table class="table-auto w-full text-left whitespace-no-wrap">
          <thead>
            <tr>
              <th
                class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
              >
                File
              </th>
              <th
                class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
              >
                Size
              </th>
              <th
                class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
              >
                Compressed
              </th>
              <th
                class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
              >
                Ratio
              </th>
              <!--                    <th-->
              <!--                            class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"-->
              <!--                    >-->
              <!--                        Result-->
              <!--                    </th>-->
              <th
                class="font-medium pl-3 pt-6 text-center text-gray-400 text-left text-sm tracking-wider uppercase"
              >
                Status
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(file, i) in files" :key="`${i}-${file.name}`">
              <td>
                <p
                  class="cell-l ta"
                  :class="{
                    'text-gray-400': !file.isProcessed,
                    'text-gray-200': file.isProcessed,
                    'anime-txt-success': file.isConverted
                  }"
                >
                  {{ file.filename }}
                </p>
              </td>
              <td>
                <p
                  class="ta"
                  :class="{
                    'text-gray-400': !file.isProcessed,
                    'text-gray-200': file.isProcessed,
                    'anime-txt-success': file.isConverted
                  }"
                >
                  {{ getPrettySize(file.size) }}
                </p>
              </td>
              <td>
                <p :class="{ 'anime-txt-success': file.isConverted }">
                  {{ getPrettySize(file.convertedSize) }}
                </p>
              </td>
              <td>
                <p :class="{ 'anime-txt-success': file.isConverted }">
                  {{ getSavings(file) }}
                </p>
              </td>
              <!--                    <td @click="openFile(file)"><p>{{ file.convertedPath }}</p>-->
              <!--                    </td>-->
              <td>
                <p
                  v-if="file.isConverted"
                  class="cell-r flex items-center justify-center"
                >
                  <svg
                    version="1.1"
                    :id="`${i}-check`"
                    xmlns="http://www.w3.org/2000/svg"
                    xmlns:xlink="http://www.w3.org/1999/xlink"
                    x="0px"
                    y="0px"
                    viewBox="0 0 20 20"
                    enable-background="new 0 0 20 20"
                    width="20px"
                    height="20px"
                    xml:space="preserve"
                  >
                    <path
                      fill="#27ffa7"
                      d="M10,0C4.5,0,0,4.5,0,10s4.5,10,10,10s10-4.5,10-10S15.5,0,10,0z M8,14.4l-3.7-3.7l1.4-1.4L8,11.6l5.3-5.3
                            l1.4,1.4L8,14.4z"
                    />
                  </svg>
                </p>
                <p v-else class="cell-r"></p>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </transition>
  </section>
</template>

<script>
import { fExt, fName, fSize } from '@/lib/file'
import { EventBus } from '@/lib/event-bus'
import Wails from '@wailsapp/runtime'
import { prettyTime } from '@/lib/time'

export default {
  name: 'Editor',

  data() {
    return {
      files: [],
      headerHeight: 0,
      isConverting: false,
      isDragging: false
    }
  },

  computed: {
    /**
     * canClear returns true if the file list can be cleared.
     * @returns {boolean}
     */
    canClear() {
      if (this.files.length === 0) return false
      return !this.filesPending
    },

    /**
     * canConvert returns true if the file list satisfies conditions for
     * conversion.
     * @returns {boolean}
     */
    canConvert() {
      if (this.files.length === 0 || this.isConverting) return false
      return this.filesConverted && !this.filesPending
    },

    /**
     * filesConverted returns true if all files have been converted.
     * @returns {boolean}
     */
    filesConverted() {
      if (this.files.length === 0) return false
      return this.files.some(f => !f.isConverted)
    },

    /**
     * filesPending returns true if there are files not processed.
     * @returns {boolean}
     */
    filesPending() {
      if (this.files.length === 0) return false
      return this.files.some(f => !f.isProcessed)
    },

    /**
     * optBtnTxt returns the text string for the Optimize button per the
     * current app state.
     * @returns {string}
     */
    optBtnTxt() {
      const d = 'Optimize'
      if (this.files.length === 0) return d
      if (this.filesPending) return 'Processing...'
      return d
    },

    /**
     * session returns the current session stats.
     * @returns {object}
     */
    session() {
      return this.$store.getters.session
    },

    /**
     * totalStats returns the all-time stats.
     */
    totalStats() {
      return this.$store.getters.stats
    }
  },

  methods: {
    /**
     * clear removes the files from the file list and the FileManager.
     */
    clear() {
      this.files = []
      this.$refs['fileInput'].value = null
      window.backend.FileManager.Clear()
        .then(res => {
          console.log(res)
        })
        .catch(err => {
          console.error(err)
        })
    },

    /**
     * convert calls the Convert method on the FileManager.
     */
    convert() {
      this.isConverting = true
      window.backend.FileManager.Convert()
        .then(res => {
          console.log(res)
        })
        .catch(err => {
          console.error(err)
        })
    },

    /**
     * createFileId creates a file ID based on its name and size.
     * @param {string} name - The filename without extension.
     * @param {number} size - The file size in bytes.
     * @returns {string} - The file ID.
     */
    createFileId(name, size) {
      return name + size.toString()
    },

    /**
     * getFileById returns the file from the file list with the given
     * ID.
     * @param {string} id - The file ID.
     * @returns {object} - The matched file.
     */
    getFileById(id) {
      if (this.files.length === 0) return
      return this.files.find(f => {
        return f.id === id
      })
    },

    /**
     * getPrettySize returns a pretty string of file size given bytes.
     * @param {number} size - Size in bytes.
     * @returns {string} - The pretty size.
     */
    getPrettySize(size) {
      if (size === 0) return ''
      return fSize(size)
    },

    /**
     * getSavings returns the percentage difference between a file's
     * original and converted sizes as a pretty string.
     * @param {object} file - A file in the file list.
     * @returns {string} - The file savings as string.
     */
    getSavings(file) {
      if (file.convertedSize === 0) return ''
      const p = Math.floor(100 - (file.convertedSize / file.size) * 100)
      return `${p.toString()}%`
    },

    /**
     * getFileType attempts to determine the file image type based on the
     * given file type and extension. This method exists due to IE's lack of
     * support for WebP.
     * @param {string} type - The file type.
     * @param {string} ext - The file extension.
     * @returns {string}
     */
    getFileType(type, ext) {
      if (this.isValidType(type)) return type
      if (this.isValidExt(ext)) return `image/${ext}`
      return ''
    },

    /**
     * handleFileInput handles the file submission via form input.
     * @param {Event} e
     */
    handleFileInput(e) {
      const f = e.target.files
      this.processFileInput(f)
    },

    /**
     * hasFile returns true if the file is in the file list.
     * @param {string} id - The file ID.
     * @returns {boolean}
     */
    hasFile(id) {
      if (this.files.length === 0) return false
      let e = false
      this.files.forEach(f => {
        if (f.id === id) {
          e = true
        }
      })
      return e
    },

    /**
     * handleWinResize calculates the height of the table based on the height
     * of the header.
     */
    handleWinResize() {
      const hh = this.$refs['header'].clientHeight
      if (hh === this.headerHeight) return
      this.headerHeight = hh
    },

    /**
     * isValidExt returns true if the given file extension is of an accepted
     * set of extensions.
     * @param {string} ext - A file extension
     * @returns {boolean}
     */
    isValidExt(ext) {
      const v = ['jpg', 'jpeg', 'png', 'webp']
      return v.indexOf(ext) >= 0
    },

    /**
     * isValidType returns true if the given type is one of an accepted
     * set of mime types.
     * @param {string} type - A file's mime type.
     * @return {boolean}
     */
    isValidType(type) {
      const v = [
        'image/jpg',
        'image/jpeg',
        'image/.jpg',
        'image/png',
        'image/webp'
      ]
      return v.indexOf(type) >= 0
    },

    /**
     * openFile opens the file at the given file path.
     */
    openFile(file) {
      // TODO: can this be called directly from Wails?
      window.backend.FileManager.OpenFile(file.convertedPath)
        .then(res => {
          console.log(res)
        })
        .catch(err => {
          console.error(err)
        })
    },

    /**
     * processFileInput processes a list of submitted files.
     * @param {FileList} f - A FileList array of files.
     */
    processFileInput(f) {
      f.forEach(f => {
        const name = fName(f.name)
        const ext = fExt(f.name)
        const type = this.getFileType(f.type, ext)
        const size = f.size
        const id = this.createFileId(name, size)
        // reject if wrong mime or already exists
        if (!type) {
          EventBus.$emit('notify', {
            msg: `File type not supported. Valid file types include JPG, PNG, and WebP.`,
            type: 'warn'
          })
          return
        }
        if (this.hasFile(id)) {
          EventBus.$emit('notify', {
            msg: `Image file has already been added to the file list.`,
            type: 'warn'
          })
          return
        }
        this.processFile(f, id, type)
        this.files.push({
          convertedPath: '',
          convertedSize: 0,
          filename: f.name,
          id,
          isConverted: false,
          isProcessed: false,
          name,
          size
        })
      })
      this.$refs['fileInput'].value = null
    },

    /**
     * processFile encodes a file and sends it to the backend for further
     * processing.
     * @param {File} file
     * @param {string} id
     * @param {string} type
     */
    processFile(file, id, type) {
      const reader = new FileReader()
      reader.onload = () => {
        const name = file.name
        window.backend.FileManager.HandleFile(
          JSON.stringify({
            data: reader.result.split(',')[1],
            ext: fExt(name),
            id,
            name: fName(name),
            size: file.size,
            type
          })
        )
          .then(() => {
            const f = this.getFileById(id)
            f.isProcessed = true
          })
          .catch(err => {
            this.removeFileById(id)
            EventBus.$emit('notify', {
              msg: err,
              type: 'warn'
            })
            console.error(err)
          })
      }
      reader.readAsDataURL(file)
    },

    /**
     * removeFileById removes a file from the file list by the given ID.
     * @param {string} id
     */
    removeFileById(id) {
      const i = this.files.findIndex(f => f.id === id)
      if (i < 0) return
      this.files.splice(i, 1)
    }
  },

  mounted() {
    Wails.Events.On('conversion:complete', e => {
      const f = this.getFileById(e.id)
      if (!f) return
      f.convertedPath = e.path
      f.isConverted = true
      f.convertedSize = e.size
    })

    Wails.Events.On('conversion:stat', e => {
      const c = e.count
      const r = e.resizes
      const s = e.savings
      const t = e.time
      this.$store.dispatch('setSessionProp', { key: 'count', value: c })
      this.$store.dispatch('setSessionProp', { key: 'time', value: t })
      this.$store.dispatch('getStats')
      if (s > 0) {
        this.$store.dispatch('setSessionProp', { key: 'savings', value: s })
      }
      if (r > 0) {
        EventBus.$emit('notify', {
          msg: `Optimized ${c} ${c > 1 ? 'images' : 'image'} and made ${r} ${
            r > 1 ? 'resizes' : 'resize'
          } in ${prettyTime(t)[0]} ${prettyTime(t)[1].toLowerCase()}.`
        })
      } else {
        EventBus.$emit('notify', {
          msg: `Optimized ${c} ${c > 1 ? 'images' : 'image'} in ${
            prettyTime(t)[0]
          } ${prettyTime(t)[1].toLowerCase()}.`
        })
      }
      this.isConverting = false
    })

    const dz = this.$refs['dropZone']

    dz.addEventListener(
      'click',
      () => {
        this.$refs['fileInput'].click()
      },
      false
    )
    dz.addEventListener(
      'dragenter',
      e => {
        e.stopPropagation()
        e.preventDefault()
      },
      false
    )
    dz.addEventListener(
      'dragover',
      e => {
        e.stopPropagation()
        e.preventDefault()
        if (this.isDragging) return
        this.isDragging = true
      },
      false
    )
    dz.addEventListener(
      'dragend',
      e => {
        e.stopPropagation()
        e.preventDefault()
        this.isDragging = false
      },
      false
    )
    dz.addEventListener(
      'dragleave',
      e => {
        e.stopPropagation()
        e.preventDefault()
        this.isDragging = false
      },
      false
    )
    dz.addEventListener(
      'drop',
      e => {
        e.stopPropagation()
        e.preventDefault()
        const dt = e.dataTransfer
        const f = dt.files
        this.isDragging = false
        this.processFileInput(f)
      },
      false
    )

    window.addEventListener('resize', this.handleWinResize)
    this.handleWinResize()
  }
}
</script>

<style scoped>
.table-wrapper {
  max-width: calc(100vw - 160px);
  min-height: 80px;
  overflow: auto;
}

td {
  margin: 0;
  padding: 0;
}

table tr:nth-child(odd) p {
  @apply bg-gray-800;
}

td p {
  @apply my-1 pl-3 py-2;
  min-height: 40px;
}

td p.cell-l {
  border-top-left-radius: 6px;
  border-bottom-left-radius: 6px;
}

td p.cell-r {
  border-top-right-radius: 6px;
  border-bottom-right-radius: 6px;
  margin-right: 1rem;
}

.drop-zone > svg path {
  transition: fill 600ms cubic-bezier(0.07, 0.95, 0, 1);
}

.drop-zone:hover > svg path {
  fill: #27ffa7;
}
</style>
