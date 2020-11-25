import Vue from "vue";
import Vuex from "vuex";
import recommendations from "./modules/recommendations";
import purchases from "./modules/purchases";
import login from "./modules/login";
import stores from "./modules/stores";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    recommendations,
    login,
    purchases,
    stores
  },
});
