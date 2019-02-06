<template>
  <button @click="callback">{{ name }}</button>
</template>

<style lang="scss">
button {
  width: 100%;
  height: 40px;
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
    }
  }
}
</script>

