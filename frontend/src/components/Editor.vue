<template>
    <div class="mx-auto p-10">
        <div class="bg-indigo-500 w-50 h-10" ref="dropZone"></div>
        <input
                type="file"
                accept="image/jpeg, image/jpg, image/png, image/webp"
                multiple
                class="hidden"
                @input="handleFileInput"
                ref="fileInput"
        />
        <p @click="openDir">{{ config.outDir }}</p>
        <label for="target">Target</label>
        <select name="target" id="target" @change="selectTarget">
            <option value="webp">WebP</option>
            <option value="jpg">JPG</option>
            <option value="png">PNG</option>
        </select>
        <button
                class="flex ml-auto text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded"
                @click="selectOutDir"
        >
            Output Directory
        </button>
        <button
                class="flex ml-auto text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded"
                @click="convert"
                :disabled="!canConvert"
        >
            Convert
        </button>
        <button
                class="flex ml-auto text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded"
                @click="clear"
        >
            Clear
        </button>
        <!-- file table -->
        <div v-if="files.length > 0" class="table-wrapper">
            <table class="table-auto w-full text-left whitespace-no-wrap">
                <thead>
                <tr>
                    <th
                            class="tracking-wider font-medium pl-3 text-gray-900 text-left text-sm uppercase"
                    >
                        File
                    </th>
                    <th
                            class="tracking-wider font-medium pl-3 text-gray-900 text-left text-sm uppercase"
                    >
                        Size
                    </th>
                    <th
                            class="tracking-wider font-medium pl-3 text-gray-900 text-left text-sm uppercase"
                    >
                        New Size
                    </th>
                    <th
                            class="tracking-wider font-medium pl-3 text-gray-900 text-left text-sm uppercase"
                    >
                        Ratio
                    </th>
                    <th
                            class="tracking-wider font-medium pl-3 text-gray-900 text-left text-sm uppercase"
                    >
                        Result
                    </th>
                    <th
                            class="tracking-wider font-medium pl-3 text-gray-900 text-left text-sm uppercase"
                    >
                        Status
                    </th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(file, i) in files" :key="`${i}-${file.name}`">
                    <td><p class="cell-l">{{ file.filename }}</p></td>
                    <td class="font-mono "><p>{{ getPrettySize(file.size) }}</p>
                    </td>
                    <td class="font-mono "><p>{{
                        getPrettySize(file.convertedSize) }}</p></td>
                    <td><p>{{ getSavings(file) }}</p></td>
                    <td @click="openFile(file)"><p>{{ file.convertedPath }}</p>
                    </td>
                    <td><p class="cell-r">{{ file.isConverted }}</p></td>
                </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script>
  import { fExt, fName, fSize } from '../lib/file'
  import Wails from '@wailsapp/runtime'

  export default {
    name: 'Editor',

    data() {
      return {
        files: []
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
      },

      /**
       * config returns the app configuration from state.
       * @returns {object}
       */
      config() {
        return this.$store.getters.config
      }
    },

    methods: {
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
    },

    mounted() {
      Wails.Events.On('conversion:complete', e => {
        const f = this.getFileById(e.id)
        if (!f) return
        f.convertedPath = e.path
        f.isConverted = true
        f.convertedSize = e.size
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
      }, false)
      dz.addEventListener('drop', e => {
        e.stopPropagation()
        e.preventDefault()
        const dt = e.dataTransfer
        const f = dt.files
        this.processFileInput(f)
      }, false)
    }
  }
</script>

<style scoped>
    .table-wrapper {
        min-height: calc(100vh / 5);
        max-height: calc(100vh / 2);
        overflow: auto;
    }

    td {
        margin: 0;
        padding: 0;
    }

    table tr:nth-child(odd) p {
        @apply bg-gray-700;
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
</style>
