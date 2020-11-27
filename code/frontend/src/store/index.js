import Vue from "vue";
import Vuex from "vuex";
import recommendations from "./modules/recommendations";
import purchases from "./modules/purchases";
import user from "./modules/user";
import stores from "./modules/stores";
import filter from "./modules/filter";
import cart from "./modules/cart";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    recommendations,
    user,
    purchases,
    stores,
    filter,
    cart
  },
});
