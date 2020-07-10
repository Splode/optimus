<template>
    <div class="mx-auto p-10">
        <input type="file" accept="image/png, image/jpeg, image/jpg" multiple
               @change="processFileInput"/>
        <button class="flex ml-auto text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded"
                @click="selectOutDir">Output Directory
        </button>
        <button class="flex ml-auto text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded"
                @click="convert" :disabled="!canConvert">Convert
        </button>
        <div v-if="files.length > 0" class="table-wrapper">
            <table
                    class="table-auto w-full text-left whitespace-no-wrap">
                <thead>
                <tr>
                    <th class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl">
                        File
                    </th>
                    <th class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl">
                        Size
                    </th>
                    <th class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200">
                        Converted
                    </th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(file, i) in files" :key="`${i}-${file.name}`">
                    <td class="px-4 py-3">{{ file.filename }}</td>
                    <td class="font-mono px-4 py-3">{{ file.size }}</td>
                    <td class="px-4 py-3">{{ file.isConverted }}</td>
                </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script>
    import {fExt, fName, fSize} from "../lib/file";
    import Wails from '@wailsapp/runtime'

    export default {
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
            }
        },
        methods: {
            convert() {
                window.backend.FileManager.Convert().then(result => {
                    console.log(result)
                }).catch(err => {
                    console.error(err)
                })
            },
            /**
             * getFileByName returns the file from the file list with the given
             * name.
             * @param {string} name
             * @returns {object}
             */
            getFileByName(name) {
                if (this.files.length === 0) return
                return this.files.find(f => {
                    return f.name === name
                })
            },
            /**
             * hasFile returns true if the file is already in the file list.
             * @param {string} name
             * @returns {boolean}
             */
            hasFile(name) {
                if (this.files.length === 0) return false
                let e = false
                this.files.forEach(f => {
                    if (f.name === name) {
                        e = true
                    }
                })
                return e
            },
            isValidType(type) {
                const v = ["image/png", "image/jpg", "image/jpeg"]
                return v.indexOf(type) >= 0
            },
            // TODO: make an ID by hashing name and size
            processFileInput(e) {
                e.target.files.forEach(f => {
                    const name = fName(f.name)
                    // reject if wrong mime or already exists
                    if (!this.isValidType(f.type) || this.hasFile(name)) return
                    this.processFile(f)
                    this.files.push({
                        filename: f.name,
                        isConverted: false,
                        name,
                        size: fSize(f.size)
                    })
                })
            },
            processFile(file) {
                const reader = new FileReader()
                reader.onload = () => {
                    const name = file.name
                    window.backend.FileManager.HandleFile(JSON.stringify({
                        data: reader.result.split(',')[1],
                        ext: fExt(name),
                        name: fName(name),
                        type: file.type
                    }))
                }
                reader.readAsDataURL(file)
            },
            selectOutDir() {
                window.backend.FileManager.SetOutDir().then(result => {
                    console.log(result)
                }).catch(err => {
                    console.error(err)
                })
            }
        },
        mounted() {
            Wails.Events.On("conversion:complete", name => {
                const f = this.getFileByName(name)
                if (!f) return
                f.isConverted = true
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
