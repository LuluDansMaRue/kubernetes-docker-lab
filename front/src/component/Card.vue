<template>
  <div class="card">
    <h3>{{ bobba.name }}</h3>
    <h3><span>Flavor:</span> {{ bobba.flavor }}</h3>
    <p>Shop: {{ bobba.shop }}</p>
  </div>
</template>

<style lang="scss">
.card {
  width: 200px;
  height: 200px;
  background-color: #f8eee6;
  box-shadow: 0px 2px 3px 2px rgba(0, 0, 0, 0.1);
  border-radius: 10px;
  padding: 5px;

  h3, p {
    font-family: 'Raleway', 'sans-serif';

    span {
      font-weight: 300;
    }
  }
}
</style>

<script>
export default {
  name: 'Card',
  props: {
    bobba: {
      type: Object,
      required: true
    }
  }
}
</script>
