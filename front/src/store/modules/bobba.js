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
        return state.bobbas
      }

      return state.bobbas.filter(bobba => {
        return bobba.name.toLowerCase().includes(query.toLowerCase())
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

/**
 * Actions
 *    A set of method that can interact with the mutations
 */
const actions = {
  /**
   * Add Bobba
   * 
   * @param {Object}
   * @param {Object} bobba
   */
  addBobba({ commit }, bobba) {
    commit('add', bobba)
  },
  /**
   * Delete Bobba
   * 
   * @param {Object}
   * @param {Number} id
   */
  deleteBobba({ commit }, id) {
    commit('delete', id)
  },
  /**
   * Populate Bobba
   *  
   * @param {Object} 
   * @param {Array} bobbas 
   */
  populateBobba({ commit }, bobbas) {
    if (isEmpty(bobbas)) {
      return
    }

    commit('populate', bobbas)
  }
}

/**
 * Mutations
 *    A set of mthod that interact with the store
 */
const mutations = {
  /**
   * Add
   * 
   * @param {Object} state
   * @param {Object} bobba
   */
  add(state, bobba) {
    state.bobbas.push(bobba)
  },
  /**
   * Delete
   * 
   * @param {Object} state
   * @param {Number} id
   */
  delete(state, id) {
    const bobbas = [...state.bobbas]
    const idx = state.bobbas.findIndex(bobba => bobba.id === id )
    if (idx === -1) {
      return
    }

    bobbas.splice(idx, 1)
    state.bobbas = bobbas
  },
  /**
   * Populate
   *    Populate the bobbas state
   *    Only use during the initialization
   * 
   * @param {Object} state
   * @param {Array} bobbas
   */
  populate(state, bobbas) {
    state.bobbas = bobbas
  }
}

export default {
  state,
  getters,
  actions,
  mutations,
}