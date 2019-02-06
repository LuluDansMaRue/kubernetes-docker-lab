import VueRouter from 'vue-router'
import Home from '../container/Home.vue'
import Add from '../container/Add.vue'
import Detail from '../container/Detail.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/add', component: Add },
  { path: '/bobba/:id', component: Detail }
]

const router = new VueRouter({
  routes
})

export default router