import Vue from "vue";
import VueRouter from "vue-router";
import { Auth } from "aws-amplify";

Vue.use(VueRouter);

const routerOptions = [
  {
    path: "/account",
    component: "account",
    meta: {
      requiresAuth: true,
      groups: ["user"]
    },
  },
  {
    path: "/browse",
    component: "browse",
  },
  {
    path: "/cart",
    component: "cart",
    meta: {
      requiresAuth: true,
      groups: ["user"]
    },
  },
  {
    path: "/employee",
    component: "employee",
    meta: {
      requiresAuth: true,
      groups: ["employee"]
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
    path: "/recommended",
    component: "recommended",
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
  // const user = Auth.currentAuthenticatedUser();
  // if (to.matched.some((record) => record.meta.requiresAuth)) {
  //   if (user !== null) {
  //     if (
  //       user.signInUserSession.accessToken.payload["cognito:groups"].includes(
  //         record.meta.group
  //       )
  //     ) {
  //       next();
  //     }
  //   }
  //   next({
  //     path: "/login",
  //   });
  // }
  // next();

  if (to.matched.some((record) => record.meta.requiresAuth)) {
    Auth.currentAuthenticatedUser()
      .then(() => {
        next();
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
