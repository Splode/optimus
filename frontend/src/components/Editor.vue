<template>
    <div class="p-10 w-full">
        <header class="border-b-2 border-gray-800 flex flex-wrap">
            <div class="w-5/12">
                <div class="bg-gray-800 border-2 border-dashed cursor-pointer drop-zone flex flex-col items-center justify-center py-10 ta-slow rounded-sm"
                     :class="isDragging ? 'border-green' : 'border-gray-400'"
                     ref="dropZone">
                    <svg version="1.1" id="dropZone-plus"
                         xmlns="http://www.w3.org/2000/svg"
                         xmlns:xlink="http://www.w3.org/1999/xlink" x="0px"
                         y="0px"
                         viewBox="0 0 48 48" enable-background="new 0 0 48 48"
                         width="48px" height="48px"
                         xml:space="preserve">
                        <path fill="#808080" d="M47,20.6H28.4c-0.6,0-1-0.4-1-1V1c0-0.6-0.4-1-1-1h-4.9c-0.6,0-1,0.4-1,1v18.6c0,0.6-0.4,1-1,1H1
                c-0.6,0-1,0.4-1,1v4.9c0,0.6,0.4,1,1,1h18.6c0.6,0,1,0.4,1,1V47c0,0.6,0.4,1,1,1h4.9c0.6,0,1-0.4,1-1V28.4c0-0.6,0.4-1,1-1H47
                c0.6,0,1-0.4,1-1v-4.9C48,21,47.6,20.6,47,20.6z"/>
            </svg>
                    <p class="mt-6 text-gray-200">Drag and drop or select
                        images</p>
                </div>
            </div>
            <div class="pl-6 w-7/12">
                <transition name="fade-fast" mode="out-in">
                    <div v-if="!stats.time"
                         key="intro"
                         class="flex h-full items-center justify-center">
                        <h2 class="leading-none text-4xl text-center text-green">
                            Add
                            image files<br>to get started</h2>
                    </div>
                    <div v-else
                         key="stats"
                         class="flex flex-wrap items-center justify-center h-full">
                        <div class="px-3 w-5/12">
                            <h2 class="font-bold leading-none text-5xl text-green tracking-tight">
                                {{
                                getPrettySize(stats.savings) }}</h2>
                            <p class="font-medium text-gray-300 tracking-wider uppercase">
                                Saved</p>
                        </div>
                        <div class="px-3 w-3/12">
                            <p class="font-bold text-2xl">{{ stats.count }}</p>
                            <p class="font-medium text-gray-300 text-sm tracking-wider uppercase">
                                {{
                                stats.count > 1 ? 'Images' : 'Image'}}</p>
                            <p class="font-bold text-2xl">{{
                                getPrettyTime(stats.time)[0] }}</p>
                            <p class="font-medium text-gray-300 text-sm tracking-wider uppercase">
                                {{
                                getPrettyTime(stats.time)[1] }}</p>
                        </div>
                        <div class="px-3 w-4/12">
                            <p class="font-bold text-2xl">2.25 GB</p>
                            <p class="font-medium text-gray-300 text-sm tracking-wider uppercase">
                                All time Savings</p>
                            <p class="font-bold text-2xl">2,204</p>
                            <p class="font-medium text-gray-300 text-sm tracking-wider uppercase">
                                All time Images</p>
                        </div>
                        <div class="px-3 w-full">
                            <p>Optimized {{ lastStat.count }} {{ lastStat.count
                                > 1
                                ? 'images' : 'image '}} in {{
                                getPrettyTime(lastStat.time)[0] }} {{
                                getPrettyTime(lastStat.time)[1].toLowerCase()
                                }}</p>
                        </div>
                    </div>
                </transition>
            </div>
            <footer class="w-full">
                <section class="flex justify-between py-6 w-5/12">
                    <button
                            class="btn focus:outline-none ta-slow"
                            :class="canConvert ? 'border-purple hover:bg-purple hover:text-gray-900 text-gray-200' : 'btn--disabled'"
                            @click="convert"
                            :disabled="!canConvert"
                    >
                        Optimize
                    </button>
                    <button
                            class="btn focus:outline-none ta-slow"
                            :class="files.length > 0 ? 'border-gray-400 hover:bg-gray-400 hover:text-gray-900' : 'btn--disabled'"
                            @click="clear"
                            :disabled="files.length === 0"
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
            <div v-if="files.length > 0" class="table-wrapper">
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
                        <td><p class="cell-l">{{ file.filename }}</p></td>
                        <td><p>{{ getPrettySize(file.size) }}</p>
                        </td>
                        <td><p>{{
                            getPrettySize(file.convertedSize) }}</p></td>
                        <td><p>{{ getSavings(file) }}</p></td>
                        <!--                    <td @click="openFile(file)"><p>{{ file.convertedPath }}</p>-->
                        <!--                    </td>-->
                        <td>
                            <p v-if="file.isConverted"
                               class="cell-r flex items-center justify-center">
                                <svg version="1.1" :id="`${i}-check`"
                                     xmlns="http://www.w3.org/2000/svg"
                                     xmlns:xlink="http://www.w3.org/1999/xlink"
                                     x="0px" y="0px"
                                     viewBox="0 0 20 20"
                                     enable-background="new 0 0 20 20"
                                     width="20px" height="20px"
                                     xml:space="preserve">
                                <path fill="#07FDBC" d="M10,0C4.5,0,0,4.5,0,10s4.5,10,10,10s10-4.5,10-10S15.5,0,10,0z M8,14.4l-3.7-3.7l1.4-1.4L8,11.6l5.3-5.3
                            l1.4,1.4L8,14.4z"/>
                            </svg>
                            </p>
                            <p v-else class="cell-r"></p>
                        </td>
                    </tr>
                    </tbody>
                </table>
            </div>
        </transition>
    </div>
</template>

<script>
  import { fExt, fName, fSize } from '../lib/file'
  import Wails from '@wailsapp/runtime'
  import { prettyTime } from '../lib/time'

  export default {
    name: 'Editor',

    data() {
      return {
        files: [],
        isDragging: false,
        lastStat: {
          count: 0,
          time: 0
        },
        stats: {
          count: 0,
          savings: 0,
          time: 0
        }
      }
    },

    computed: {
      /**
       * canConvert returns true if the file list satisfies conditions for
       * conversion.
       * @returns {boolean}
       */
      canConvert() {
        if (this.files.length === 0) return false
        return this.files.some(f => {
          return !f.isConverted
        })
      }
    },

    methods: {
      /**
       * calcTotalSavings sums the total difference between original and
       * converted file sizes.
       */
      calcTotalSavings() {
        this.files.forEach(f => {
          if (!f.isConverted || f.convertedSize > f.size) return
          this.stats.savings += f.size - f.convertedSize
        })
      },

      /**
       * clear removes the files from the file list and the FileManager.
       */
      clear() {
        this.files = []
        this.$refs['fileInput'].value = null
        window.backend.FileManager.Clear().then(res => {
          console.log(res)
        }).catch(err => {
          console.error(err)
        })
      },

      /**
       * convert calls the Convert method on the FileManager.
       */
      convert() {
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
       * getPrettyTime returns an array of human-friendly strings representing
       * millisecond time.
       * @param {number} ms
       * @returns {string[]}
       */
      getPrettyTime(ms) {
        return prettyTime(ms)
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
        window.backend.FileManager.OpenFile(file.convertedPath).then(res => {
          console.log(res)
        }).catch(err => {
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
          if (!type || this.hasFile(id)) return
          this.processFile(f, id, type)
          this.files.push({
            convertedPath: '',
            convertedSize: 0,
            filename: f.name,
            id,
            isConverted: false,
            name,
            size
          })
        })
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
              type
            })
          )
        }
        reader.readAsDataURL(file)
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
        const t = e.time
        this.stats.count += c
        this.stats.time += t
        this.lastStat.count = c
        this.lastStat.time = t
        this.calcTotalSavings()
      })

      const dz = this.$refs['dropZone']

      dz.addEventListener('click', () => {
        this.$refs['fileInput'].click()
      }, false)
      dz.addEventListener('dragenter', e => {
        e.stopPropagation()
        e.preventDefault()
      }, false)
      dz.addEventListener('dragover', e => {
        e.stopPropagation()
        e.preventDefault()
        if (this.isDragging) return
        this.isDragging = true
      }, false)
      dz.addEventListener('dragend', e => {
        e.stopPropagation()
        e.preventDefault()
        this.isDragging = false
      }, false)
      dz.addEventListener('dragleave', e => {
        e.stopPropagation()
        e.preventDefault()
        this.isDragging = false
      }, false)
      dz.addEventListener('drop', e => {
        e.stopPropagation()
        e.preventDefault()
        const dt = e.dataTransfer
        const f = dt.files
        this.isDragging = false
        this.processFileInput(f)
      }, false)
    }
  }
</script>

<style scoped>
    .table-wrapper {
        max-width: calc(100vw - 160px);
        min-height: 80px;
        overflow: auto;
        height: calc(100vh - 344px)
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
    }

    .drop-zone > svg path {
        transition: fill 600ms cubic-bezier(.07, .95, 0, 1);
    }

    .drop-zone:hover > svg path {
        fill: #07fdbc;
    }
</style>
