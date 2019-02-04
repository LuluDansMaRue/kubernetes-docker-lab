<template>
  <div>
    <h2>Bubble tea lister</h2>
    <button @click="testApi">Test API</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'App',
  methods: {
    testApi() {
      axios.get('http://localhost:8000')
        .then(res => console.warn(res))
    }
  }
}
</script>
