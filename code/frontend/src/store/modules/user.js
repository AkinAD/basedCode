import { Auth } from "aws-amplify";

const state = {
  user: {},
  signedIn: false,
  assignedStore: {}, //assigned store for manager/employee (via admin)
  userGroups: [],
};

const getters = {
  signedIn() {
    return state.signedIn;
  },
  getUser() {
    return state.user;
  },
  getUserGroups() {
    return state.userGroups;
  },
  getAssignedStore() {
    return state.assignedStore;
  },
};

const actions = {
  async signIn({ commit }) {
    Auth.currentAuthenticatedUser()
      .then((user) => {
        commit("signIn", user);
        Auth.currentSession()
          .then((result) => {
            commit(
              "setUserGroups",
              result.accessToken.payload["cognito:groups"]
            );
          })
          .catch(() => {
            commit("setUserGroups", []);
          });
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
    (state.signedIn = false), (state.user = {}), (state.userGroups = []);
  },
  setSignedStore: (state, newStore) => {
    state.assignedStore = newStore;
  },
  setUserGroups: (state, newGroups) => {
    state.userGroups = newGroups;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
