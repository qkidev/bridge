import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/pages/hqki_change';

Vue.use(Router);

export default new Router({
  routes: [{
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'home',
    component: Home
  },
  ]
})