<template>
    <div class="mx-auto p-10">
        <input type="file" accept="image/png, image/jpeg, image/jpg" multiple
               @change="processFileInput"/>
        <button class="flex ml-auto text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded"
                @click="selectOutDir">Output Directory
        </button>
        <button class="flex ml-auto text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded"
                @click="convert">Convert
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
                    <td class="px-4 py-3">{{ file.size }}</td>
                    <td class="px-4 py-3">{{ file.isConverted }}</td>
                </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script>
    import {fExt, fName, fSize} from "../lib/file";

    export default {
        data() {
            return {
                files: []
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
            processFileInput(e) {
                e.target.files.forEach(f => {
                    this.processFile(f)
                    this.files.push({
                        filename: f.name,
                        isConverted: false,
                        name: fName(f.name),
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
