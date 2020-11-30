const state = {
  cart: [],
};

const getters = {
  getCart() {
    return state.cart;
  },
};

const actions = {
  submitCart() {
    //perform axios post of cart
  },
  addToCart: ({ commit }, item) => {
    commit("addToCart", item);
  },
};

const mutations = {
  //This should be private and not used outside the action. Unfortunately there is no way to enforce this
  setCart: (state, items) => {
    state.cart = items;
  },
  addToCart: (state, item) => {
    state.cart.push(item);
  },
  removeFromCart: (state, item) => {
    state.cart = state.cart.filter((i) => i != item);
  },
  emptyCart: (state) => {
    state.cart = [];
  },
};

export default {
  state,
  getters,
  mutations,
  actions,
};
