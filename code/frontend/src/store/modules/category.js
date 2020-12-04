import axios from "axios";

const state = {
  categories: [],
  selectedCatorgory: 0,
};

const getters = {
  getCategories() {
    return state.categories;
  },
  getSelectedCategory() {
    return state.selectedCatorgory;
  },
};

const actions = {
  async addCategory({ commit }, category) {
    await axios
      .post(`http://localhost:8081/category/${category}`)
      .then(commit("addCategoryToState", category))
      .catch(console.log("error adding category"));
  },
  async deleteCategory({ commit }, id) {
    await axios
      .delete(`http://localhost:8081/category/${id}`)
      .then(commit("deleteCategoryFromState", id))
      .catch(console.log("error deleting category"));
  },
  async updateCategory({ commit }, category) {
    await axios
      .put(`http://localhost:8081/category/${category}`)
      .then(commit("updateCategoryInState", category))
      .catch(console.log("error updating category"));
  },
  async fetchCategories({ commit }) {
    await axios
      .get("http://localhost:8081/category")
      .then((res) => commit("updateCategories", res.data))
      .catch(console.log("error fetching categories"));
  },
};

const mutations = {
  addCategoryToState: (state, category) => {
    state.categories.push(category);
  },
  deleteCategoryFromState: (state, id) => {
    state.categories = state.categories.filter((i) => i.id != id);
  },
  updateCategoryInState: (state, category) => {
    for (var c of state.categories) {
      if (c.id == category.id) {
        c = category;
        break;
      }
    }
  },
  setSelectedCategory: (state, category) => {
    state.selectedCatorgory = category;
  },
  updateCategories: (state, categories) => {
    state.categories = categories;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
