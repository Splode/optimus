<template>
  <button
    class="border-2 border-gray-800 cursor-pointer flex focus:outline-none hover:border-gray-400 hover:bg-gray-900 items-center justify-center p-2 ta-slow rounded-full w-10 h-10"
  >
    <svg
      version="1.1"
      id="x"
      xmlns="http://www.w3.org/2000/svg"
      xmlns:xlink="http://www.w3.org/1999/xlink"
      x="0px"
      y="0px"
      viewBox="0 0 11.9 11.9"
      style="enable-background:new 0 0 11.9 11.9;"
      width="14"
      height="14"
      :class="color"
      xml:space="preserve"
    >
      <path
        fill="#b3b3b3"
        d="M10.4,0L6,4.5L1.5,0L0,1.5L4.5,6L0,10.4l1.5,1.5L6,7.5l4.5,4.5l1.5-1.5L7.5,6l4.5-4.5L10.4,0z"
      />
    </svg>
  </button>
</template>

<script>
export default {
  name: 'BtnClose',

  props: {
    color: {
      type: String,
      required: false,
      default: 'green'
    }
  }
}
</script>

<style scoped>
button > svg > path {
  transition: fill 0.6s cubic-bezier(0.07, 0.95, 0, 1);
}

button:hover > svg.blue > path {
  fill: #27d1ff;
}

button:hover > svg.green > path {
  fill: #27ffa7;
}

button:hover > svg.purple > path {
  fill: #d690ff;
}
</style>
