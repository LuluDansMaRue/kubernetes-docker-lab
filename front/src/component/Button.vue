<template>
  <div
    @click="callback"
    :class="'button '+color"
  >
  {{ name }}
  </div>
</template>

<style lang="scss">
.button {
  width: 100%;
  height: 20px;
  text-align: center;
  padding: 10px 0px;
  transition: all ease .25s;
  font-weight: 600;

  &.white {
    background-color: white;
  }

  &.red {
    background-color: #E82C0C;
  }

  &.green {
    background-color: rgb(89, 197, 89);
  }

  &:hover {
    background-color: #5E5E5E;
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
    }
  }
}
</script>

