import { isEmpty } from 'lodash'

// State
const state = {
  bobbas: [],
}

/**
 * Getters
 *    A list of methods to retrieve the bubble teas
 */
const getters = {
  /**
   * Filtered Bobbas
   * 
   * @param {Object} state
   * @return {Function}
   */
  filteredBobbas(state) {
    return query => {
      if (isEmpty(query)) {
        return state.bobba
      }

      return state.bobbas.filter(bobba => {
        return bobba.name.includes(query.toLowerCase())
      })
    }
  },
  /**
   * Bobba By Id
   *  Return a bubble tea object
   * 
   * @param {Object} state 
   * @return {Function}
   */
  bobbaById(state) {
    return id => state.bobbas.filter(bobba => {
      return bobba.id === id
    })
  }
}

const actions = {
  fetch({ commit }, bobba) {
    // commit
    commit()
  }
}

const mutations = {
  setAction() {

  }
}

export {
  getters,
  actions,
  mutations,
}