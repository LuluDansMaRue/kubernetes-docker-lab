import { isEmpty } from 'lodash'

/**
 * Validation
 *    A mixin that is used to check the inputs
 */
const Validation = {
  data: () => {
    return {
      constraint: []
    }
  },
  methods: {
    /**
     * Set Value
     * 
     * @param {String} name
     * @param {String} value
     */
    setValue(name, value) {
      const prop = this.constraint
        .filter(c => c.name === name)
        .reduce((previous, current) => {
          return Object.assign(previous, current)
        })

      if (isEmpty(prop)) {
        return
      }

      switch (prop.type) {
        case 'string':
          prop.value = value
          return
        case 'number': {
          const float = parseFloat(value, 10)
          prop.value = isNaN(float) ? '' : float
          return
        }
        default:
          prop.value = ''
          return    
      }
    },
    /**
     * Check All Constraint
     *    Checking all the constraint of the inputs
     */
    checkAllConstraint() {
      let idx = 0;
      let validate = true

      while (idx < this.constraint.length) {
        const { type, value } = this.constraint[idx]
        if (typeof value !== type) {
          validate = false
          idx = this.constraint.length + 1
        } 

        idx++;
      }

      return validate
    }
  }
}

export default Validation