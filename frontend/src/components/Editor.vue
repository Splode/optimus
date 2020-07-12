<template>
    <div class="mx-auto p-10">
        <p @click="openDir">{{ config.outDir }}</p>
        <input
                type="file"
                accept="image/jpeg, image/jpg, image/png, image/webp"
                multiple
                @change="processFileInput"
        />
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
        <div v-if="files.length > 0" class="table-wrapper">
            <table class="table-auto w-full text-left whitespace-no-wrap">
                <thead>
                <tr>
                    <th
                            class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl"
                    >
                        File
                    </th>
                    <th
                            class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl"
                    >
                        Size
                    </th>
                    <th
                            class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl"
                    >
                        New Size
                    </th>
                    <th
                            class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl"
                    >
                        Savings
                    </th>
                    <th
                            class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl"
                    >
                        Result
                    </th>
                    <th
                            class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200"
                    >
                        Converted
                    </th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(file, i) in files" :key="`${i}-${file.name}`">
                    <td class="px-4 py-3">{{ file.filename }}</td>
                    <td class="font-mono px-4 py-3">{{ getPrettySize(file.size)
                        }}
                    </td>
                    <td class="font-mono px-4 py-3">
                        {{ getPrettySize(file.convertedSize) }}
                    </td>
                    <td class="px-4 py-3">{{ getSavings(file) }}</td>
                    <td class="px-4 py-3" @click="openFile(file)">{{
                        file.convertedPath }}
                    </td>
                    <td class="px-4 py-3">{{ file.isConverted }}</td>
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
          .then(result => {
            console.log(result)
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

      processFileInput(e) {
        e.target.files.forEach(f => {
          const name = fName(f.name)
          const size = f.size
          const id = this.createFileId(name, size)
          // reject if wrong mime or already exists
          if (!this.isValidType(f.type) || this.hasFile(id)) return
          this.processFile(f, id)
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

      processFile(file, id) {
        const reader = new FileReader()
        reader.onload = () => {
          const name = file.name
          window.backend.FileManager.HandleFile(
            JSON.stringify({
              data: reader.result.split(',')[1],
              ext: fExt(name),
              id,
              name: fName(name),
              type: file.type
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
          .then(result => {
            console.log(result)
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
        console.log(t)
        window.backend.Config.SetTarget(t).then(res => {
          console.log(res)
          this.$store.dispatch('getConfig')
        }).catch(err => {
          console.error(err)
        })
      }
    },

    mounted() {
      console.log(this.$store)
      Wails.Events.On('conversion:complete', e => {
        console.log(e)
        const f = this.getFileById(e.id)
        if (!f) return
        f.convertedPath = e.path
        f.isConverted = true
        f.convertedSize = e.size
      })
    }
  }
</script>

<style scoped>
    .table-wrapper {
        min-height: calc(100vh / 5);
        max-height: calc(100vh / 2);
        overflow: auto;
    }
</style>
