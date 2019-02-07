<template>
  <modal 
    name="error-modal"
    :width="300"
    height="auto"
    @before-open="beforeOpen"
  >
    <p class="error-text">An error happened: {{ content }}</p>
    <close-button
      color="red"
      name="close"
      :callback="close"
    />
  </modal>
</template>

<style lang="scss">
.v--modal {
  padding: 15px;
}

.error-text {
  font-family: 'Raleway', 'sans-serif'
}
</style>


<script>
import CloseButton from './Button.vue'

export default {
  name: 'Alert',
  components: {
    CloseButton,
  },
  data() {
    return {
      content: ''
    }
  },
  methods: {
    /**
     * Close
     */
    close() {
      this.$modal.hide('error-modal')
    },
    /**
     * Before Open
     */
    beforeOpen(event) {
      this.content = event.params.content
    }
  }
}
</script>
