//import axios from "axios";

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
  selectedStore: null,
  dialog: false
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
  }
};

const actions = {
  //todo: implement axios calls to get stores
};

const mutations = {
  updateStores: (state, newStores) => {
    state.stores = newStores;
  },
  setSelectedStore: (state, store) => {
    state.selectedStore = store;
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
