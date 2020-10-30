import Vue from 'vue'
import VueRouter from 'vue-router'
import SampleHome from '../views/sample_list/home'
import Sample1 from '../views/sample_list/sample1'
import Sample2 from '../views/sample_list/sample2'
import Sample3 from '../views/sample_list/sample3'

Vue.use(VueRouter)

const routes = [

  // {
  //   path: '/about',
  //   name: 'About',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
  {
    path: '/sample',
    name: 'SampleHome',
    component: SampleHome
  },
  {
    path: '/sample/1',
    name: 'Sample1',
    component: Sample1
  },
  {
    path: '/sample/2',
    name: 'Sample2',
    component: Sample2
  },
  {
    path: '/sample/3',
    name: 'Sample3',
    component: Sample3
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
