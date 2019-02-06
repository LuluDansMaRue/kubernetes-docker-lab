<template>
  <div class="detail__container">
    <div class="detail__container--whitebox">
      <h2>Detail of a bobba</h2>
      <p>Name: {{ bobba.name }}</p>
      <p>Flavor: {{ bobba.flavor }}</p>
      <p>Calory: {{ bobba.calory }}</p>
      <p>Price: {{ bobba.price }}</p>
      <p>Shop: {{ bobba.shop }}</p>
    </div>
    <delete-button
      name="delete"
      color="red"
      :callback="deleteAction"
    /> 
  </div>
</template>

<style lang="scss">
.detail__container {
  padding: 15px;

  &--whitebox {
    background-color: white;
    padding: 10px;
  }
}
</style>

<script>
import { isEmpty } from 'lodash'
import { fetch } from '../core/request'
import DeleteButton from '../component/Button.vue'


export default {
  name: 'Detail',
  components: {
    DeleteButton
  },
  computed: {
    /**
     * Bobba
     *    Return a bobba by an ID
     * 
     * @return {Object}
     */
    bobba: function() {
      const tea = this.$store.getters.bobbaById(this.id)
      
      if (isEmpty(tea)) {
        return {}
      }

      return tea[0]
    }
  },
  data: () => {
    return {
      id: 0
    }
  },
  methods: {
    // @TODO add warning notification...
    deleteAction() {
      fetch('bobba', 'delete', this.id)
        .then(res => {
          this.$store.dispatch('deleteBobba', this.id)
          console.warn(this.$router)
          this.$router.push({
            path: '/'
          })
        })
        .catch(err => console.warn(err))
    }
  },
  created() {
    this.id = parseInt(this.$route.params.id, 10)
  }
}
</script>

