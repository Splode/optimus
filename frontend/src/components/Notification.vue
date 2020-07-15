<template>
    <section class="absolute bottom-0 right-0">
        <div v-if="notifications.length > 0" class="p-10">
            <transition-group name="fade" mode="in-out">
                <div v-for="(n, i) in notifications" :key="`${i}-notification`"
                     class="bg-gray-700 border-l-2 border-purple mt-2 p-4 ta rounded-br-md rounded-tr-md">
                    <p class="text-gray-100">{{ n.msg }}</p>
                </div>
            </transition-group>
        </div>
    </section>
</template>

<script>
  import { EventBus } from '../lib/event-bus'

  export default {
    name: 'Notification',

    data() {
      return {
        isClearing: false,
        notifications: []
      }
    },

    methods: {
      clear() {
        if (this.notifications.length <= 0) {
          this.isClearing = false
          return
        }
        this.isClearing = true
        setTimeout(() => {
          this.notifications.shift()
          this.clear()
        }, 6000)
      }
    },

    mounted() {
      EventBus.$on('notify', n => {
        this.notifications.push({ msg: n.msg })
        if (!this.isClearing) {
          this.clear()
        }
      })
    }
  }
</script>

<style scoped>

</style>