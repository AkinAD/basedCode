import axios from "axios";

const state = {
  recommendedItems: [],
};

const getters = {
  allRecommendations: (state) => state.recommendedItems,
};

const actions = {
  async fetchRecommendations({ commit }) {
    const response = await axios.get("https://fakestoreapi.com/products/");

    commit("setRecommendations", response.data);
  },
};

const mutations = {
  setRecommendations: (state, recommendedItems) =>
    (state.recommendedItems = recommendedItems),
};

export default {
  state,
  getters,
  actions,
  mutations,
};
