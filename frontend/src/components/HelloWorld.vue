<template>
    <div class="container">
        <h1>{{ message }}</h1>
        <a @click="convert">Press Me!</a>
        <input type="file" multiple @change="processFileInput"/>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                message: ' '
            }
        },
        methods: {
            getMessage: function () {
                var self = this
                const time = Date.now()
                console.log(time)
                window.backend.Toast().then(result => {
                    console.log(Date.now() - time)
                    self.message = result
                })
            },
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
                    window.backend.FileManager.HandleFile(JSON.stringify({
                        data: reader.result.split(',')[1],
                        name: file.name,
                        type: file.type
                    }))
                }
                reader.readAsDataURL(file)
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
