<template>
  <nav class="bg-gray-900 flex flex-col items-center pt-10 w-24 z-40">
    <div
      @click="selectView('Settings')"
      class="cursor-pointer mb-2 p-2 rounded-full w-10 h-10"
    >
      <svg
        version="1.1"
        id="cog"
        xmlns="http://www.w3.org/2000/svg"
        :class="{ active: active === 'Settings' }"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        x="0px"
        y="0px"
        viewBox="0 0 19.6 20"
        enable-background="new 0 0 19.6 20"
        xml:space="preserve"
      >
        <path
          fill="#b3b3b3"
          d="M0.1,13.3l2,3.5c0.3,0.5,0.9,0.6,1.4,0.4l1.4-0.8c0.6,0.5,1.2,0.8,1.9,1.1V19c0,0.6,0.4,1,1,1h4
                c0.6,0,1-0.4,1-1v-1.6c0.7-0.3,1.3-0.7,1.9-1.1l1.4,0.8c0.5,0.3,1.1,0.1,1.4-0.4l2-3.5c0.3-0.5,0.1-1.1-0.4-1.4l-1.4-0.8
                c0.1-0.4,0.1-0.7,0.1-1.1s0-0.7-0.1-1.1l1.4-0.8c0.5-0.3,0.6-0.9,0.4-1.4l-2-3.5c-0.3-0.5-0.9-0.6-1.4-0.4l-1.4,0.8
                c-0.6-0.5-1.2-0.8-1.9-1.1V1c0-0.6-0.4-1-1-1h-4c-0.6,0-1,0.4-1,1v1.6C6.1,2.9,5.5,3.3,4.9,3.7L3.5,2.9C3,2.6,2.4,2.8,2.1,3.3
                l-2,3.5C-0.1,7.2,0,7.8,0.5,8.1l1.4,0.8C1.8,9.3,1.8,9.6,1.8,10s0,0.7,0.1,1.1l-1.4,0.8C0,12.2-0.1,12.8,0.1,13.3z M9.8,6
                c2.2,0,4,1.8,4,4s-1.8,4-4,4s-4-1.8-4-4S7.6,6,9.8,6z"
        />
      </svg>
    </div>
    <div
      @click="selectView('Stats')"
      class="cursor-pointer mb-4 p-2 rounded-full w-10 h-10"
    >
      <svg
        version="1.1"
        id="graph"
        xmlns="http://www.w3.org/2000/svg"
        :class="{ active: active === 'Stats' }"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        x="0px"
        y="0px"
        viewBox="0 0 17 20"
        style="enable-background:new 0 0 17 20;"
        xml:space="preserve"
      >
        <rect y="10" fill="#b3b3b3" width="2" height="10" />
        <rect x="5" y="2" fill="#b3b3b3" width="2" height="18" />
        <rect x="10" y="5" fill="#b3b3b3" width="2" height="15" />
        <rect x="15" fill="#b3b3b3" width="2" height="20" />
      </svg>
    </div>
    <div
      @click="selectView('About')"
      class="cursor-pointer mb-2 p-2 rounded-full w-10 h-10"
    >
      <svg
        version="1.1"
        id="info"
        :class="{ active: active === 'About' }"
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        x="0px"
        y="0px"
        viewBox="0 0 20 20"
        style="enable-background:new 0 0 20 20;"
        xml:space="preserve"
      >
        <g>
          <path
            fill="#b3b3b3"
            d="M10,0C4.5,0,0,4.5,0,10s4.5,10,10,10s10-4.5,10-10S15.5,0,10,0z M10,18c-4.4,0-8-3.6-8-8s3.6-8,8-8s8,3.6,8,8
                        S14.4,18,10,18z"
          />
          <path fill="#b3b3b3" d="M9,9h2v6H9V9z M9,5h2v2H9V5z" />
        </g>
      </svg>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'Sidebar',

  props: {
    active: {
      type: String,
      required: true
    }
  },

  methods: {
    selectView(v) {
      this.$emit('select-view', v)
    }
  }
}
</script>

<style scoped>
#cog,
#info {
  transition: transform 0.3s cubic-bezier(0.07, 0.95, 0, 1);
}

#cog path,
#graph rect,
#info path {
  transition: fill 1s cubic-bezier(0.07, 0.95, 0, 1);
}

#cog.active {
  transform: rotate(270deg);
}

#cog.active path,
#info.active path {
  fill: #27ffa7;
}

#info.active path {
  fill: #d690ff;
}

#graph:hover rect:first-of-type,
#graph.active rect:first-of-type {
  fill: #27d1ff;
}

#graph:hover rect:nth-of-type(2),
#graph.active rect:nth-of-type(2) {
  fill: #27ffa7;
}

#graph:hover rect:nth-of-type(3),
#graph.active rect:nth-of-type(3) {
  fill: #ffe027;
}

#graph:hover rect:nth-of-type(4),
#graph.active rect:nth-of-type(4) {
  fill: #ff9b45;
}

div:hover > #cog path {
  fill: #27ffa7;
}

div:hover > #info path {
  fill: #d690ff;
}
</style>
