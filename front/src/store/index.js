import Vue from 'vue'
import Vuex from 'vuex'
import bobba from './modules/bobba'
import room from './modules/room'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    bobba,
    room
  }
})