import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routerOptions = [
  {path: "/account", component: "account" },
  {path: "/browse", component: "browse" },
  {path: "/cart", component: "cart" },
  {path: "/employees", component: "employees" },
  {path: "/Home", component: "Home" },
  {path: "/login", component: "login" },
  {path: "/managers", component: "managers" },
  {path: "/recommended", component: "recommended" },
  {path: "/stores", component: "stores" },
];

const routes = routerOptions.map(route => {
  return {
    ...route,
    component: () => import(`../views/${route.component}.vue`)
  }
});

const router = new VueRouter({
  mode: "history",
  routes
})

export default router
