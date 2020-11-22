import { Auth } from "aws-amplify";

const state = {
  user: {},
  signedIn: false,
};

const getters = {
  signedIn() {
    return state.signedIn;
  },
  getUser() {
    return state.user;
  },
};

const actions = {
  async signIn({ commit }) {
    Auth.currentAuthenticatedUser()
      .then((user) => {
        commit("signIn", user);
      })
      .catch(() => console.log("error signing in"));
  },
  signOut({ commit }) {
    commit("signOut");
  },
};

const mutations = {
  signIn: (state, user) => {
    (state.signedIn = true), (state.user = user);
  },
  signOut: () => {
    (state.signedIn = false), (state.user = {});
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
