import axios from 'axios'

const state = {
    Categories: []
}

const getters = {
    getCategories() {
        return state.Categories;
    },

}

const actions = {

    async addCategory({ commit }, id) {
        await axios.post("...").then( (response) => {
            
            let Category  = {};
            response;
            //do something

            commit('addCategoryToState', Category);
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
    async updateCategory({ commit }, Category) {
        await axios.put("...").then( (response) => {
            

            //do something
            response;
            commit('updateCategoryInState', Category);
        }).catch({
          //do something else  
        });
    }
}


const mutations = {
    addCategoryToState : (state, Category) => {
        state.Categories.push(Category);
    },
    deleteCategoryFromState: (state, id) => {
        state.Categories = state.Categories.filter((i) => i.id != id);
    },
    updateCategoryInState: (state, Category) => {
        for(var c of state.Categories) {
            if (c.id == Category.id)
            {
                c = Category;
                break;
            }
        }
    },
}