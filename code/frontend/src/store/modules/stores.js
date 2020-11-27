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
  }
};

const mutations = {
  updateStores: (state, newStores) => {
    state.stores = newStores;
  },
  setSelectedStore: (state, data) => {
    (state.selectedStore = data[0]),
    (state.items = data[1])
  },
  setDialog: (state, bool) => {
    state.dialog = bool;
  }
};

export default {
  state,
  getters,
  actions,
  mutations,
};
