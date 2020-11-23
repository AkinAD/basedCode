import Vue from 'vue'
import VueRouter from 'vue-router'
import { Auth } from "aws-amplify";

Vue.use(VueRouter)

const routerOptions = [
  {path: "/account", component: "account", meta: {requiresAuth: true}},
  {path: "/browse", component: "browse" },
  {path: "/cart", component: "cart" },
  {path: "/employees", component: "employees", meta: {requiresAuth: true} },
  {path: "/Home", component: "Home" },
  {path: "/login", component: "login", meta: {requiresNonAuth: true}},
  {path: "/managers", component: "manageItems"},// meta: {requiresAuth: true} }, FOR TESTING, auth removed
  {path: "/recommended", component: "recommended" },
  {path: "/stores", component: "stores" },
  {path: "*", component: "Home"}
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

router.beforeResolve((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    Auth.currentAuthenticatedUser().then(() => {
      next()
    }).catch(() => {
      next({
        path: '/login'
      });
    });
  }
  if (to.matched.some(record => record.meta.requiresNonAuth)) {
    Auth.currentAuthenticatedUser().then(() => {
      next({
        path: '/'
      }).catch(() => {
        next()
      })
    });
  }
  next()
})


export default router
