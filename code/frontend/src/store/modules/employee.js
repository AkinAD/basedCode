import axios from 'axios'

const state = {
    employeeList: []
}

const getters = {
    getEmployees() {
        return state.employeeList;
    },

}

const actions = {

    async addEmployee({ commit }, employee) {
        await axios.post("...").then( (response) => {
            
            response;
            //do something

            commit('addEmployeeToState', employee);
        }).catch({
          //do something else  
        });
        
    },
    async deleteEmployee({ commit }, id) {
        await axios.delete("...").then( (response) => {

            //do something
            response;
            commit('deleteEmployeeFromState', id);
        }).catch({
          //do something else  
        });
        
    },
    async updateEmployee({ commit }, employee) {
        await axios.put("...").then( (response) => {
            

            //do something
            response;
            commit('updateEmployeeInState', employee);
        }).catch({
          //do something else  
        });
    }
}


const mutations = {
    addEmployeeToState : (state, employee) => {
        state.employeeList.push(employee);
    },
    deleteEmployeeFromState: (state, id) => {
        state.employeeList = state.employeeList.filter((i) => i.id != id);
    },
    updateEmployeeInState: (state, employee) => {
        for(var e of state.employeeList) {
            if (e.id == employee.id)
            {
                e = employee;
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