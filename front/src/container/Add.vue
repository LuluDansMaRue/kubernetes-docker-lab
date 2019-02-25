<template>
  <div class="add__container">
    <div class="add__container--whitebox">
      <h2>Add a bobba</h2>
      <field
        v-for="(field, idx) in constraint"
        v-bind:key="idx"
        :value="field.value"
        :label="field.name"
        @input="event => updateField(field.name, event)"
      />
    </div>
    <submit-button
      name="submit"
      :callback="postForm"
      color="green"
      :activate="validate"
    />
  </div>
</template>

<style lang="scss">
.add__container {
  padding: 50px;

  .add__container--whitebox {
    background-color: white;
    padding: 15px;
  }
}
</style>


<script>
import { post } from '../core/request'
import Validation from '../mixins/validation'
import Modal from '../mixins/modal'

// Components
import Field from '../component/Field.vue'
import SubmitButton from '../component/Button.vue'


// FormData a list of constraint with it's set of Value
const formData = [
  { name: 'name', type: 'string', value: '' },
  { name: 'flavor', type: 'string', value: '' },
  { name: 'calory', type: 'number', value: 0 },
  { name: 'price', type: 'number', value: 0 },
  { name: 'shop', type: 'string', value: '' }
]

/**
 * @TODO refactor the HTML integration.
 */
export default {
  mixins: [
    Validation,
    Modal
  ],
  components: {
    Field,
    SubmitButton
  },
  data() {
    return {
      validate: true,
      constraint: formData.slice(),
    }
  },
  methods: {
    /**
     * Update Field
     * 
     * @TODO do a validation on the field... for the time being it's going to be done manually
     * @param {Number} name
     * @param {String|Number} value
     */
    updateField(name, value) {
      this.setValue(name, value)
      this.validate = this.checkAllConstraint()
    },
    /**
     * Build Object
     *    Build the obejct that is going to be send to the APi
     * 
     * @return {Object} data
     */
    buildObject() {
      const data = this.constraint
        .map(c => ({ name: c.name, value: c.value }))
        .reduce((prev, curr) => {
          prev[curr.name] = curr.value 
          return prev
        }, {})

      return data
    },
    /**
     * Post Form
     */
    postForm() {
      const data = this.buildObject()
      // @TODO do a validation on the form...
      post('bobba/add', data)
        .then(() => {
          this.$store.dispatch('addBobba', data)
          this.$router.push({
            path: '/'
          })
        })
        .catch(err => this.show('error-modal', {
          content: err.message
        }))
    }
  },
  created: function() {
    this.setModal(this.$modal)
  }
}
</script>

