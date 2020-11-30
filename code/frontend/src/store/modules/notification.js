const state = {
  type: "blue-grey",
  message: "",
  timeout: 3000,
  show: false,
};

const mutations = {
  SHOW_NOTIFICATION(state, notification) {
    state.type = notification.type;
    state.message = notification.message;
    if (notification.type && notification.message) {
      state.show = true;
    }
    if (notification.timeout) {
      state.timeout = notification.timeout;
    }
  },
};

const actions = {
  showNotification({ commit }, notification) {
    commit("SHOW_NOTIFICATION", notification);
  },
};

export default {
  state,
  actions,
  mutations,
};
