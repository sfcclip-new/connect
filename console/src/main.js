import Vue from 'vue'
import App from './App.vue'
import Group from './Group.vue'
import VueRouter from 'vue-router'
import Vuetify from 'vuetify'

Vue.use(VueRouter)
Vue.use(Vuetify)

const router = new VueRouter({
  routes: [
    { path: '/group', component: Group },
    { path: '/group/:id', component: Group },
  ]
})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
