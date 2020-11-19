import axios from 'axios';

const state = {
    accounts: []
};

const getters = {
    allAccounts: state => state.accounts
};

const actions = {
    async fetchAccounts({commit}) {
        const response = await axios.get('https://jsonplaceholder.typicode.com/users');
    
        commit('setAccounts', response.data);
    }
};

const mutations = {
    setAccounts: (state, accounts) => (state.accounts = accounts)
};

export default {
    state,
    getters,
    actions,
    mutations
};