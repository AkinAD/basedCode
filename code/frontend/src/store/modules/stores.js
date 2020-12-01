import axios from "axios";

//each store should have a name, location, and list of items
const state = {
  stores: [
    {
      i: 1,
      text: "Shopperz",
      icon: "mdi-basket",
      location: "123 Bay Street",
    },
    {
      i: 2,
      text: "NotLongos",
      icon: "mdi-basket",
      location: "420 Victor Philp Road",
    },
    {
      i: 3,
      text: "JoeUnfresh",
      icon: "mdi-basket",
      location: "666 Allen Avenu",
    },
    {
      i: 4,
      text: "Matt's awesome store",
      icon: "mdi-basket",
      location: "69 Troyer Avenue",
    },
    {
      i: 5,
      text: "Kinny's Konvenience",
      icon: "mdi-basket",
      location: "404 Unknown Street",
    },
  ],
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
  }
};

const actions = {
  //todo: implement axios calls to get stores
  async updateStores({ commit }) {
    try {
      const res = await axios.get(""); //stores endpoint
      commit("updateStores", res.data);
    } catch {
      commit("updateStores", []);
    }
  },
  
  async setSelectedStore({ commit }, store) {
    try {
      const res = await axios.get("https://fakestoreapi.com/products/"); //selected store's items
      commit("setSelectedStore", [store, res.data]);
    } catch {
      commit("setSelectedStore", store, []); //no items found
    }
  },
  //items
  async addItem({ commit }, item) {
    
    await axios.post("...").then( (response) => {

      response;
      //do something

      commit('addItemToState', item);
    }).catch({
      //do something else  
    });
    
},
async deleteItem({ commit }, id) {
    await axios.delete("...").then( (response) => {
      response;
      //do something

      commit('deleteItemFromState', id);
    }).catch({
      //do something else  
    });
    
},
async updateItem({ commit }, item) {
   
    await axios.put("...").then( (response) => {
        
      response;
      //do something

       commit('updateItemInState', item);
    }).catch({
      //do something else  
    });
}
};

const mutations = {
  //stores
  updateStores: (state, newStores) => {
    state.stores = newStores;
  },
  setSelectedStore: (state, data) => {
    (state.selectedStore = data[0]),
    (state.items = data[1])
  },
  setDialog: (state, bool) => {
    state.dialog = bool;
  },
  //items
  addItemToState: (state, item) => {
    state.items.push(item);
  },
  deleteItemFromState: (state, id) => {
    state.items = state.items.filter((i) => i.id != id);
  },
  updateItemInState: (state, item) => {
    for(var i of state.items) {
      if (i.id == item.id)
      {
          i = item;
          break;
      }
  }
  }

};

export default {
  state,
  getters,
  actions,
  mutations,
};
