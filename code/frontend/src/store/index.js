import Vue from "vue";
import Vuex from "vuex";
import accounts from "./modules/accounts";
import recommendations from "./modules/recommendations";
import notification from "./modules/notification";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    accounts,
    recommendations,
    notification,
  },
});
