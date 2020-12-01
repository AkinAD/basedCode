import Vue from "vue";
import Vuex from "vuex";
import recommendations from "./modules/recommendations";
import user from "./modules/user";
import stores from "./modules/stores";
import filter from "./modules/filter";
import cart from "./modules/cart";
import employee from "./modules/employee";
import category from "./modules/category";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    recommendations,
    user,
    stores,
    filter,
    cart,
    employee,
    category
  },
});
