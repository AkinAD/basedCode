import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/staticViews/Home.vue'
import views from '../views'

Vue.use(VueRouter)

const routes = [
  //STATIC ROUTES 
  {
    path: '/',
    name: 'Home',
    component: Home //views.Home
  },
  {
    path: '/account',
    name: 'Account',
    component: views.account
  },
  {
    path: '/browse',
    name: 'Browse',
    component: views.browse
  },
  {
    path: '/cart',
    name: 'Cart',
    component: views.cart
  },
  {
    path: '/checkout',
    name: 'Checkout',
    component: views.checkout
  },
  {
    path: '/employees',
    name: 'Employees',
    component: views.employees
  },
  {
    path: '/managers',
    name: 'Managers',
    component: views.managers
  },
  {
    path: '/recommended',
    name: 'Recommended',
    component: views.recommended
  },
  {
    path: '/signin',
    name: 'Signin',
    component: views.signin
  },
  {
    path: '/signup',
    name: 'Signup',
    component: views.signup
  },
  {
    path: '/stores',
    name: 'Stores',
    component: views.stores
  },
  //DYNAMIC ROUTES https://router.vuejs.org/guide/essentials/dynamic-matching.html#reacting-to-params-changes
  //:todo
  {
    path: '/item/:id',
    component: views.item
  },
  {
    path: '/search?query=',
    component: views.search
  }
]

const router = new VueRouter({
  routes
})

export default router
