<template>
    <div class="container">
        <h1>{{ message }}</h1>
        <a @click="getMessage">Press Me!</a>
        <input type="file" @change="processFile"/>
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
            processFile(e) {
                const file = e.target.files[0]
                const reader = new FileReader()
                reader.onload = () => {
                    console.log(reader.result)
                    window.backend.HandleFile(JSON.stringify({
                        data: reader.result.split(',')[1],
                        name: file.name,
                        type: file.type
                    }))
                }
                // reader.readAsBinaryString(file)
                reader.readAsDataURL(file)
                // reader.readAsArrayBuffer(file)
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
