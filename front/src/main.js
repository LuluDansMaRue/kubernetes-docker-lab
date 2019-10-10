import Vue from 'vue'
import VueRouter from 'vue-router'
import VModal from 'vue-js-modal'

import store from './store'
import router from './core/router'
import { open } from './core/socket'

// Entry point App.vue
import App from './App.vue'

Vue.config.productionTip = false

open()

// Use the router
Vue.use(VueRouter)
Vue.use(VModal)

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
