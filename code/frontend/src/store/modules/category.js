import axios from "axios";
import Vue from 'vue';

import { Auth } from "aws-amplify";

const state = {
  categories: [],
  selectedCategory: 0,
};

const getters = {
  getCategories() {
    return state.categories;
  },
  getSelectedCategory() {
    return state.selectedCategory;
  },
};

var domain;

if (process.env.NODE_ENV === 'development') {
  domain = '';
} else {
  domain = 'https://thesmartshopper.online:8081';
}

const actions = {
  async addCategory({ commit }, category) {
    let session = await Auth.currentSession();
    let newCategory = { name: category.category }
  
    await axios
      .post(domain + "/category", newCategory, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then((response) => {
        commit("addCategoryToState", response.data)
      })
      .catch(console.log("error adding category"));
  },
  async deleteCategory({ commit }, id) {
    let session = await Auth.currentSession();
    await axios
      .delete(domain + `/category/${id}` , {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("deleteCategoryFromState", id))
      .catch(console.log("error deleting category"));
  },
  async updateCategory({ commit }, category) {
    let session = await Auth.currentSession();
    category = {
      category : category.category,
      categoryID : Number(category.categoryID)
    };
    console.log(category);
    await axios
      .put(domain + "/category", category , {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("updateCategoryInState", category))
      .catch(console.log("error updating category"));
  },
  async fetchCategories({ commit }) {
    await axios
      .get(domain + "/category")
      .then((res) => commit("updateCategories", res.data))
      .catch(console.log("error fetching categories"));
  },
};

const mutations = {
  addCategoryToState: (state, category) => {
    state.categories.push(category);
  },
  deleteCategoryFromState: (state, id) => {
    state.categories = state.categories.filter((i) => i.categoryID != id);
  },
  updateCategoryInState: (state, category) => {
    for (var i in state.categories) {
      if (state.categories[i].categoryID == category.categoryID) {
        Vue.set(state.categories, i, category)
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
