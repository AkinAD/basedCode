import axios from "axios";

const state = {
  purchasedItems: [],
};

const getters = {
  allPurchasedItems: (state) => state.purchasedItems,
};

const actions = {
  async fetchPurchasedItems({ commit }) {
    const response = await axios.get("https://fakestoreapi.com/products/");

    commit("setPurchasedItems", response.data);
  },
};

const mutations = {
  setPurchasedItems: (state, purchasedItems) =>
    (state.purchasedItems = purchasedItems),
};

export default {
  state,
  getters,
  actions,
  mutations,
};
