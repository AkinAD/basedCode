const state = {
  selectedStore: null,
  location: null,
};

const getters = {
  getSelectedStore() {
    return state.selectedStore;
  },
  getLocation() {
    return state.location;
  },
};

const mutations = {
  setSelectedStore: (state, store) => {
    state.selectedStore = store;
  },
  setLocation: (state, location) => {
    state.location = location;
  },
};

export default {
  state,
  getters,
  mutations,
};
