import axios from 'axios'

const state = {
    categories: []
}

const getters = {
    getcategories() {
        return state.categories;
    },

}

const actions = {

    async addCategory({ commit }, category) {
        await axios.post("...").then( (response) => {
            
            response;
            //do something

            commit('addCategoryToState', category);
        }).catch({
          //do something else  
        });
        
    },
    async deleteCategory({ commit }, id) {
        await axios.delete("...").then( (response) => {

            //do something
            response;
            commit('deleteCategoryFromState', id);
        }).catch({
          //do something else  
        });
        
    },
    async updateCategory({ commit }, category) {
        await axios.put("...").then( (response) => {
            

            //do something
            response;
            commit('updateCategoryInState', category);
        }).catch({
          //do something else  
        });
    }
}


const mutations = {
    addCategoryToState : (state, category) => {
        state.categories.push(category);
    },
    deleteCategoryFromState: (state, id) => {
        state.categories = state.categories.filter((i) => i.id != id);
    },
    updateCategoryInState: (state, category) => {
        for(var c of state.categories) {
            if (c.id == category.id)
            {
                c = category;
                break;
            }
        }
    },
}

export default {
    state,
    getters,
    actions,
    mutations,
  };