import Vue from 'vue'
import Vuex from 'vuex'
import bobba from './modules/bobba'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    bobba
  }
})