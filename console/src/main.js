import Vue from 'vue'
import VueRouter from 'vue-router'
import Vuetify from 'vuetify'
import axios from 'axios';

import App from './components/App.vue'
import Unit from './components/Unit.vue'
import Group from './components/Group.vue'
import Record from './components/Record.vue'

Vue.use(VueRouter)
Vue.use(Vuetify)

const router = new VueRouter({
  routes: [
    { path: '/unit/:id?', component: Unit, props: true },
    { path: '/group/:id?', component: Group, props: true },
    { path: '/record', component: Record },
  ]
})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
