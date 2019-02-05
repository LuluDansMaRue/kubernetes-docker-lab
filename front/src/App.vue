<template>
  <div>
    <div class="head">
      <h1>Bubble tea</h1>
      <p>Check for your favorite bubble tea</p>
    </div>
    <button @click="testApi">Test API</button>
  </div>
</template>

<style lang="scss">
body {
  background-color: #eddfd4;
  font-family: 'Major Mono Display', 'sans-serif';
}

.head {
  h1, p {
    text-align: center;
  }
}
</style>


<script>
import axios from 'axios'

export default {
  name: 'App',
  methods: {
    testApi() {
      axios.get('http://localhost:9000/bobba')
        .then(res => console.warn(res))
    }
  }
}
</script>
