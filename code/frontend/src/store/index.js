import Vue from "vue";
import Vuex from "vuex";
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
    user,
    stores,
    filter,
    cart,
    employee,
    category
  },
});
