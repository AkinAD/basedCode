import axios from "axios";
import Vue from 'vue';

import { Auth } from "aws-amplify";

//each store should have a name, location, and list of items
const state = {
  stores: [],
  allItems: [],
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
  getAllItems() {
    return state.allItems;
  },
};

var domain;

if (process.env.NODE_ENV === "development") {
  domain = "http://localhost:8081";
} else {
  domain = "https://thesmartshopper.online:8081";
}

const actions = {
  //************************************************************************************stores
  
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

  //***********************************************************************************stock
  async addStock({ commit }, stockInfo) {
    let session = await Auth.currentSession();
    stockInfo = {
      itemID : Number(stockInfo.itemID),
      storeID : Number(stockInfo.storeID),
      row : Number(stockInfo.row),
      col : Number(stockInfo.col)
    };
    await axios
      .post("/stock/", stockInfo, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("addStockToState", stockInfo))
      .catch(console.log("error writing item to db"));
  },
  async deleteStock({ commit }, stockInfo) {
    let session = await Auth.currentSession();
    stockInfo = {
      itemID : Number(stockInfo.itemID),
      storeID : Number(stockInfo.storeID)
    };
    await axios
      .delete(domain + `/stock/${stockInfo.itemID}/${stockInfo.storeID}` , {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("deleteStockFromState", stockInfo))
      .catch(console.log("error deleting item from db"));
  },
  async updateStock({ commit }, stockInfo) {
    let session = await Auth.currentSession();
    stockInfo = {
      itemID : Number(stockInfo.itemID),
      storeID : Number(stockInfo.storeID),
      row : Number(stockInfo.row),
      col : Number(stockInfo.col)
    };
    await axios
      .put("/stock/", stockInfo, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("updateStockInState", stockInfo))
      .catch(console.log("error updating item on db"));
  },
  
  //****************************************************************************************items
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
      .then((response) => {
        commit("addItemToState", response.data)
      })
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
      .then((response) => {
        commit("updateItemInState", response.data)
      })
      .catch(console.log("error updating item on db"));
  },
  async fetchAllItems({ commit }) {
   
    await axios
      .get("/item/")
      .then((response) => {
        commit("populateAllItems", response.data);
      })
      .catch(console.log("error fetching all items"));
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
  populateAllItems: (state, items) =>
  {
    state.allItems = items;
  },
  addItemToState: (state, item) => {
    state.allItems.push(item);
  },
  deleteItemFromState: (state, id) => {
    state.allItems = state.allItems.filter((i) => i.itemID != id);
  },
  updateItemInState: (state, item) => {
    for (var i in state.allItems) {
      if (state.allItems[i].itemID == item.itemID) {
        Vue.set(state.allItems, i, item);
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
