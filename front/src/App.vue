<template>
  <div>
    <navigation />
    <router-view></router-view>
  </div>
</template>

<style lang="scss">
body {
  background-color: #eddfd4;
  font-family: 'Major Mono Display', 'sans-serif';
}
</style>

<script>
import axios from 'axios'
import { fetch } from './core/request'

// components
import Navigation from './component/Navigation.vue'

export default {
  name: 'App',
  components: {
    Navigation
  },
  created: () => {
    fetch('bobba')
      .then(res => console.warn(res))
      .catch(err => console.warn(err))
  }
}
</script>
