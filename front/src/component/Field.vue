<template>
  <div class="field__wrapper">
    <p v-if="label">{{ label }}</p>
    <input
      type="text"
      :value="value"
      @input="$emit('input', $event.target.value)"
      placeholder="type your favorite name"
    />
  </div>
</template>

<style lang="scss">
.field__wrapper {
  display: flex;
  align-items: center;

  p {
    width: 20%;
  }
}

input {
  display: block;
  width: 500px;
  height: 25px;
  border-radius: 7px;
  border: 1px solid #d4a883;
  margin: 0 auto;
}
</style>

<script>
export default {
  name: 'Field',
  props: {
    value: {
      type: null,
      default: undefined
    },
    label: {
      type: String,
      required: false
    }
  }
}
</script>

