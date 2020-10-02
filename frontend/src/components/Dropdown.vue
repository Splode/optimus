<template>
  <div class="btn-group">
    <li
      @click="toggleMenu()"
      class="dropdown-toggle py-2"
      v-if="selectedOption.name !== undefined"
      :class="showMenu ? 'rounded-bl-none rounded-br-none' : ''"
    >
      {{ selectedOption.name }}
      <span class="caret"></span>
    </li>

    <li
      @click="toggleMenu()"
      class="dropdown-toggle"
      v-if="selectedOption.name === undefined"
    >
      {{ placeholderText }}
      <span class="caret"></span>
    </li>

    <ul class="dropdown-menu" v-if="showMenu">
      <li v-for="(option, idx) in options" :key="idx">
        <a href="javascript:void(0)" @click="updateOption(option)">
          {{ option.name }}
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Dropdown',

  data() {
    return {
      selectedOption: {
        name: ''
      },
      showMenu: false,
      placeholderText: 'Please select an item'
    }
  },
  props: {
    options: {
      type: [Array, Object]
    },
    selected: {},
    placeholder: [String],
    closeOnOutsideClick: {
      type: [Boolean],
      default: true
    }
  },

  watch: {
    selected(o) {
      this.updateOption(o)
    }
  },

  mounted() {
    this.selectedOption = this.selected
    if (this.placeholder) {
      this.placeholderText = this.placeholder
    }
    if (this.closeOnOutsideClick) {
      document.addEventListener('click', this.clickHandler)
    }
  },
  beforeDestroy() {
    document.removeEventListener('click', this.clickHandler)
  },
  methods: {
    updateOption(option) {
      this.selectedOption = option
      this.showMenu = false
      this.$emit('update-option', this.selectedOption)
    },
    toggleMenu() {
      this.showMenu = !this.showMenu
    },
    clickHandler(event) {
      const { target } = event
      const { $el } = this
      if (!$el.contains(target)) {
        this.showMenu = false
      }
    }
  }
}
</script>

<style>
.btn-group {
  min-width: 160px;
  height: 40px;
  position: relative;
  margin: 10px 1px;
  display: inline-block;
  vertical-align: middle;
}

.btn-group a:hover {
  text-decoration: none;
}

.dropdown-toggle {
  color: #636b6f;
  min-width: 160px;
  padding: 10px 20px 10px 10px;
  text-transform: none;
  font-weight: 300;
  margin-bottom: 7px;
  border: 0;
  background-image: linear-gradient(#009688, #009688),
    linear-gradient(#d2d2d2, #d2d2d2);
  background-size: 0 2px, 100% 1px;
  background-repeat: no-repeat;
  background-position: center bottom, center calc(100% - 1px);
  background-color: transparent;
  transition: background 0s ease-out;
  float: none;
  box-shadow: none;
  border-radius: 0;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
}

.dropdown-toggle:hover {
  background: #e1e1e1;
  cursor: pointer;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  z-index: 1000;
  float: left;
  min-width: 160px;
  /*padding: 5px 0;*/
  /*margin: 2px 0 0;*/
  list-style: none;
  font-size: 14px;
  text-align: left;
  background-color: #fff;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.175);
  background-clip: padding-box;
}

.dropdown-menu > li > a {
  padding: 10px 30px;
  display: block;
  clear: both;
  font-weight: normal;
  line-height: 1.6;
  color: #333333;
  white-space: nowrap;
  text-decoration: none;
}

.dropdown-menu > li > a:hover {
  background: #efefef;
  color: #409fcb;
}

.dropdown-menu > li {
  overflow: hidden;
  width: 100%;
  position: relative;
  margin: 0;
}

.caret {
  width: 0;
  position: absolute;
  top: 19px;
  height: 0;
  margin-left: -24px;
  vertical-align: middle;
  border-top: 4px dashed;
  border-right: 4px solid transparent;
  border-left: 4px solid transparent;
  right: 10px;
}

li {
  list-style: none;
}
</style>
