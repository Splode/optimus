<template>
    <div class="container">
        <button @click="convert">Convert</button>
        <button @click="selectOutDir">Output Directory</button>
        <input type="file" multiple @change="processFileInput"/>
    </div>
</template>

<script>
    import {fExt, fName} from "../lib/file";

    export default {
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
    h1 {
        margin-top: 2em;
        position: relative;
        min-height: 5rem;
        width: 100%;
    }

    a:hover {
        font-size: 1.7em;
        border-color: blue;
        background-color: blue;
        color: white;
        border: 3px solid white;
        border-radius: 10px;
        padding: 9px;
        cursor: pointer;
        transition: 500ms;
    }

    a {
        font-size: 1.7em;
        border-color: white;
        background-color: #121212;
        color: white;
        border: 3px solid white;
        border-radius: 10px;
        padding: 9px;
        cursor: pointer;
    }
</style>
