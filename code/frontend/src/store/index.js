import Vue from "vue";
import Vuex from "vuex";
import recommendations from "./modules/recommendations";
import notification from "./modules/notification";
import purchases from "./modules/purchases";
import login from "./modules/login";
import location from "./modules/location";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    recommendations,
    notification,
    login,
    location,
    purchases,
  },
});
