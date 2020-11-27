const state = {
    cart : [{ 
        id : 2,
        name : "default item",
        price : 10.99,
    }],
  };

const getters = {
    cartItems : (state) => state.cart,
    cartCount : (state) => {
        return state.cart.length;
    }
    };

const mutations = {
    //This should be private and not used outside the action. Unfortunately there is no way to enforce this
    _populateCart (state, items) {
        state.cart = items;
    },
    addToCart (state, item) {
        state.cart.push({ 
            id : item.id,
            name : item.name,
            price : item.price,
            image : item.image
        });
    },
    //This needs to be tested more but seems to work
    removeFromCart (state, item) {
        state.cart = state.cart.filter(function(currentItem){
            return currentItem.id != item.id;
        });
    },
    
    emptyCart (state) {
        state.cart = [];
    }
    };


const actions = {
        async fetchSavedCart ({commit}) {
            //logic for getting saved cart

            //can mock data for now
            let fetchedItems = [{ 
                id : 0,
                name : "saved item",
                price : 99.99,
            }];


            commit('_populateCart', fetchedItems);
        },
    };

export default {
state,
getters,
mutations,
actions
};