<template>
  <div>
    <div class="head">
      <h1>Bubble tea</h1>
      <p>Check for your favorite bubble tea</p>
    </div>
    <div class="content">
      <div class="field__container">
        <field :value="name" @input="refreshData" />
      </div>
      <div class="card__container">
        <div
          v-for="bobba in bobbas"
          v-bind:key="bobba.id"
          class="wrapper"
          @click="toDetailRoute(bobba.id)"
        >
          <card
            :bobba="bobba"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
.head {
  h1, p {
    text-align: center;
  }
}

.content {    
  .card__container {
    display: flex;
    flex-wrap: wrap;

    .wrapper {
      padding: 10px;
    }
  }

  .field__container {
    padding: 10px;
  }
}
</style>


<script>
import Card from '../component/Card.vue'
import Field from '../component/Field.vue'

export default {
  name: 'Home',
  components: {
    Field,
    Card
  },
  data: () => {
    return {
      name: ''
    }
  },
  computed: {
    bobbas: function() {
      return this.$store.getters.filteredBobbas(this.name)
    }
  },
  methods: {
    /**
     * Refresh Data
     * 
     * @param {String} value
     */
    refreshData(value) {
      this.name = value
    },
    /**
     * To Detail Route
     * 
     * @param {String} id
     */
    toDetailRoute(id) {
      this.$router.push({
        path: `bobba/${id}`,
      })
    }
  }
}
</script>
