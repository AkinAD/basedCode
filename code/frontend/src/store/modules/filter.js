const state = {
  filters: ["Any", "Under $25", "$25 to $100", "$100 to $500", "Over $500"],
  selectedFilter: null,
};

const getters = {
  getFilters() {
    return state.filters;
  },
  getSelectedFilter() {
    return state.selectedFilter;
  },
};

const mutations = {
  setSelectedFilter: (state, newFilter) => {
    state.selectedFilter = newFilter;
  },
};

export default {
  state,
  getters,
  mutations,
};
