import axios from "axios";

const state = {
  purchasedItems: [],
};

const getters = {
  allPurchasedItems: (state) => state.purchasedItems,
};

const actions = { 
  async fetchPurchasedItems({ commit }) { //needs to come from backend for current user
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
