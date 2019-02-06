<template>
  <button
    @click="callback"
    :class="'button '+color"
    :disabled="!activate"
  >
  {{ name }}
  </button>
</template>

<style lang="scss">
.button {
  width: 100%;
  height: 40px;
  text-align: center;
  padding: 10px 0px;
  transition: all ease .25s;
  font-weight: 600;
  font-size: 1em;
  font-family: 'Major Mono Display', 'sans-serif';

  &.white {
    background-color: white;
  }

  &.red {
    background-color: rgb(216, 61, 33);

    &:hover {
      background-color: rgb(173, 51, 29);
    }
  }

  &.green {
    background-color: rgb(89, 197, 89);

    &:hover {
      background-color: rgb(61, 145, 61);
    }
  }

  &:hover {
    cursor: pointer;
  }

  &:disabled {
    background-color: #5E5E5E;

    &:hover {
      background-color: rgb(56, 56, 56);
    }
  }
}
</style>


<script>
export default {
  name: 'Button',
  props: {
    name: {
      type: String,
      required: true
    },
    callback: {
      type: Function,
      required: false,
      default: () => {}
    },
    color: {
      type: String,
      required: false,
      default: 'white'
    },
    activate: {
      type: Boolean,
      required: false,
      default: true
    }
  }
}
</script>

