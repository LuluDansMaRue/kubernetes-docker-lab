import Vue from 'vue'
import VueRouter from 'vue-router'
import store from './store'
import router from './core/router'

// Entry point App.vue
import App from './App.vue'

Vue.config.productionTip = false

// Use the router
Vue.use(VueRouter)

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
