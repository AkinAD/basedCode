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
  //stores
  
  async addStore({ commit }, store) {
    let session = await Auth.currentSession();
    store = {
      address : store.address,
      storeID : Number(store.storeID)
    };

    await axios
      .post( "/store/", store, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then((response) => {
        
        commit("addStoreToState", response.data);
      })
      .catch(console.log("error writing store to db"));
  },
  async deleteStore({ commit }, storeID) {
    let session = await Auth.currentSession();
    storeID = Number(storeID);
    await axios
      .delete(domain + `/store/${storeID}`, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("removeStoreFromState", storeID))
      .catch(console.log("error writing store to db"));
  },
  async updateStore({ commit }, store) {
    let session = await Auth.currentSession();
   
    store = {
      address : store.address,
      storeID : Number(store.storeID)
    };
    console.log(store);
    await axios
      .put( "/store/", store, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("updateStoreInState", store))
      .catch(console.log("error writing store to db"));
  },

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
  addStoreToState: (state, store) => {
    state.stores.push(store);
  },
  removeStoreFromState: (state, id) => {
    state.stores = state.stores.filter((i) => i.storeID != id);
  },
  updateStoreInState: (state, store) => {
    for (var i in state.stores) {
      if (state.stores[i].storeID == store.storeID) {
        Vue.set(state.stores, i, store);
        break;
      }
    }
  },
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
