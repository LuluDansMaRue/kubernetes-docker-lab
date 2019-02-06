<template>
  <div class="add__container">
    <div class="add__container--whitebox">
      <h2>Add a bobba</h2>
      <field
        :value="form.name"
        label="name"
        @input="event => updateField('name', event)"
      />
      <field
        :value="form.flavor"
        label="flavor"
        @input="event => updateField('flavor', event)"
      />
      <field
        :value="form.calory"
        label="calory"
        @input="event => updateField('calory', event)"
      />
      <field
        :value="form.price"
        label="price"
        @input="event => updateField('price', event)"
      />
      <field
        :value="form.shop"
        label="shop"
        @input="event => updateField('shop', event)"
      />
    </div>
    <submit-button
      name="submit"
      :callback="postForm"
      color="green"
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
import Field from '../component/Field.vue'
import SubmitButton from '../component/Button.vue'
import { post } from '../core/request'

/**
 * @TODO refactor the HTML integration.
 */
export default {
  components: {
    Field,
    SubmitButton
  },
  data() {
    return {
      form: {
        name: '',
        flavor: '',
        calory: 0,
        price: 0,
        shop: ''
      }
    }
  },
  methods: {
    /**
     * Update Field
     * 
     * @TODO do a validation on the field... for the time being it's going to be done manually
     * @param {String} name
     * @param {String|Number} value
     */
    updateField(name, value) {
      if (name === 'calory' || name === 'price') {
        this.form[name] = parseFloat(value, 10)
        return
      }

      this.form[name] = value
    },
    /**
     * Post Form
     */
    postForm() {
      // @TODO do a validation on the form...
      post('bobba/add', this.form)
        .then(res => {
          this.$router.push({
            path: '/'
          })
        })
        .catch(err => console.warn(err))
    }
  }
}
</script>

