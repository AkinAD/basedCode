import { Auth } from "aws-amplify";

const state = {
  user: {},
  signedIn: false,
  assignedStore: {}, //assigned store for manager/employee (via admin)
};

const getters = {
  signedIn() {
    return state.signedIn;
  },
  getUser() {
    return state.user;
  },
  getUserGroups() {
    return state.user.signInUserSession.accessToken.payload["cognito:groups"];
  },
  getAssignedStore() {
    return state.assignedStore;
  }
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
  setSignedStore: (state, newStore) => {
    (state.assignedStore = newStore)
  }
};

export default {
  state,
  getters,
  actions,
  mutations,
};
