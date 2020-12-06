import axios from "axios";
import Vue from 'vue';

import { Auth } from "aws-amplify";

//each store should have a name, location, and list of items
const state = {
  stores: [],
  items: [],
  selectedStore: null,
  dialog: false,
};

const getters = {
  getStores() {
    return state.stores;
  },
  getSelectedStore() {
    return state.selectedStore;
  },
  getDialog() {
    return state.dialog;
  },
  getItems() {
    return state.items;
  },
};

var domain;

if (process.env.NODE_ENV === "development") {
  domain = "http://localhost:8081";
} else {
  domain = "https://thesmartshopper.online:8081";
}

const actions = {
  //todo: implement axios calls to get stores
  async updateStores({ commit }) {
    try {
      const res = await axios.get(domain + "/store"); //stores endpoint
      commit("updateStores", res.data);
    } catch {
      commit("updateStores", []);
    }
  },

  async setSelectedStore({ commit }, store) {
    try {
      const res = await axios.get(domain + `/store/${store.storeID}`); //selected store's items
      commit("setSelectedStore", [store, res.data]);
    } catch {
      commit("setSelectedStore", store, []); //no items found
    }
  },
  //items
  async addItem({ commit }, item) {
    let session = await Auth.currentSession();
    item = {
      name : item.name,
      description : item.description,
      price : parseFloat(item.price),
      categoryID : Number(item.categoryID)
    };

    await axios
      .post("/item/", item, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("addItemToState", item))
      .catch(console.log("error writing item to db"));
  },
  async deleteItem({ commit }, id) {
    let session = await Auth.currentSession();
    await axios
      .delete(domain + `/item/${id}` , {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("deleteItemFromState", id))
      .catch(console.log("error deleting item from db"));
  },
  async updateItem({ commit }, item) {
    let session = await Auth.currentSession();
    let id = item.itemID;
    item = {
      name : item.name,
      description : item.description,
      price : parseFloat(item.price),
      categoryID : Number(item.categoryID)
    };
    await axios
      .put(`/item/${id}`, item, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("updateItemInState", item))
      .catch(console.log("error updating item on db"));
  },
};

const mutations = {
  //stores
  updateStores: (state, newStores) => {
    state.stores = newStores;
  },
  setSelectedStore: (state, data) => {
    (state.selectedStore = data[0]), (state.items = data[1]);
  },
  setDialog: (state, bool) => {
    state.dialog = bool;
  },
  //items
  addItemToState: (state, item) => {
    state.items.push(item);
  },
  deleteItemFromState: (state, id) => {
    state.items = state.items.filter((i) => i.itemID != id);
  },
  updateItemInState: (state, item) => {
    for (var i in state.items) {
      if (state.items[i].itemID == item.itemID) {
        Vue.set(state.items, i, item);
        break;
      }
    }
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
