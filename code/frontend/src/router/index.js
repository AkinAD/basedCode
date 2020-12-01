import Vue from "vue";
import VueRouter from "vue-router";
import { Auth } from "aws-amplify";

Vue.use(VueRouter);

const routerOptions = [
  {
    path: "/browse",
    component: "browse",
  },
  {
    path: "/cart",
    component: "cart",
    meta: {
      requiresAuth: true,
      group: "user",
    },
  },
  {
    path: "/employee",
    component: "employee",
    meta: {
      requiresAuth: true,
      group: "employee",
    },
  },
  {
    path: "/Home",
    component: "Home",
  },
  {
    path: "/login",
    component: "login",
    meta: {
      requiresNonAuth: true,
    },
  },
  {
    path: "/cart",
    component: "cart",
    meta: {
      requiresAuth: true,
      group: "user",
    },
  },
  {
    path: "*",
    component: "Home",
  },
];

const routes = routerOptions.map((route) => {
  return {
    ...route,
    component: () => import(`../views/${route.component}.vue`),
  };
});

const router = new VueRouter({
  mode: "history",
  routes,
});

router.beforeResolve((to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    Auth.currentSession()
      .then((result) => {
        if (
          result.accessToken.payload["cognito:groups"].includes(
            to.matched[0].meta.group
          )
        ) {
          next();
        } else {
          //display popup
          console.log("user does not have permission");
          next({
            path: "/home",
          });
        }
      })
      .catch(() => {
        next({
          path: "/login",
        });
      });
  }
  next();
});

export default router;
