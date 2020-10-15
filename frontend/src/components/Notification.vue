<template>
  <section class="absolute bottom-0 right-0">
    <transition name="fade">
      <div v-if="notifications.length > 0" class="p-10">
        <transition-group name="fade-list">
          <div
            v-for="(n, i) in notifications"
            :key="`${i}-notification`"
            class="bg-gray-700 border-l-2 flex mt-2 p-4 pr-6 ta rounded-br-md rounded-tr-md"
            :class="n.type === 'warn' ? 'border-red' : 'border-green'"
          >
            <div class="mr-3">
              <!-- warn -->
              <svg
                v-if="n.type === 'warn'"
                version="1.1"
                id="warn-icon"
                xmlns="http://www.w3.org/2000/svg"
                x="0px"
                y="0px"
                viewBox="0 0 20 19"
                style="enable-background:new 0 0 20 19;"
                width="24"
                height="24"
                xml:space="preserve"
              >
                <path
                  fill="#F84D53"
                  d="M10.9,0.5c-0.3-0.7-1.4-0.7-1.8,0l-9,17c-0.2,0.3-0.2,0.7,0,1C0.3,18.8,0.6,19,1,19h18c0.4,0,0.7-0.2,0.9-0.5
                            c0.2-0.3,0.2-0.7,0-1L10.9,0.5z M11,16H9v-2h2V16z M9,12V7h2l0,5H9z"
                />
              </svg>
              <!-- success -->
              <svg
                v-else
                version="1.1"
                id="check-icon"
                xmlns="http://www.w3.org/2000/svg"
                x="0px"
                y="0px"
                viewBox="0 0 24 24"
                style="enable-background:new 0 0 24 24;"
                width="24"
                height="24"
                xml:space="preserve"
              >
                <path
                  fill="#27ffa7"
                  d="M10,15.6l-3.3-3.3l-1.4,1.4l4.7,4.7l9.7-9.7l-1.4-1.4L10,15.6z"
                />
              </svg>
            </div>
            <p class="text-gray-100">{{ n.msg }}</p>
          </div>
        </transition-group>
      </div>
    </transition>
  </section>
</template>

<script>
import { EventBus } from '@/lib/event-bus'
import Wails from '@wailsapp/runtime'

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
        this.notifications.pop()
        this.clear()
      }, 6000)
    },
    notify(n) {
      this.notifications.unshift({ msg: n.msg, type: n.type })
      if (!this.isClearing) {
        this.clear()
      }
    }
  },

  mounted() {
    EventBus.$on('notify', this.notify)
    Wails.Events.On('notify', this.notify)
  }
}
</script>

<style scoped></style>
