<template>
  <div>
    <navigation />
    <router-view></router-view>
    <alert />
  </div>
</template>

<style lang="scss">
body {
  background-color: #eddfd4;
  font-family: 'Major Mono Display', 'sans-serif';
  padding: 0;
  margin: 0;
}
</style>

<script>
import axios from 'axios'
import { fetch } from './core/request'
import Modal from './mixins/modal'

// components
import Navigation from './component/Navigation.vue'
import Alert from './component/Alert.vue'

export default {
  name: 'App',
  mixins: [
    Modal
  ],
  components: {
    Navigation,
    Alert
  },
  /**
   * Created
   *    Trigger by the VueJS hook
   *    At this stage we're fetching the API in order to populate the store
   */
  created: function() {
    this.setModal(this.$modal)

    fetch('bobba')
      .then(res => this.$store.dispatch('populateBobba', res.data))
      .catch(err => this.show('error-modal', {
        content: err.message
      }))
  }
}
</script>
