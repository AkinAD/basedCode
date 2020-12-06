import axios from "axios";
import Vue from 'vue';
import { Auth } from "aws-amplify";

const state = {
  employeeList: [],
};

const getters = {
  getEmployees() {
    return state.employeeList;
  },
};

var domain;

if (process.env.NODE_ENV === 'development') {
  domain = '';
} else {
  domain = 'https://thesmartshopper.online:8081';
}

const actions = {
  async addEmployee({ commit }, employee) {
    let session = await Auth.currentSession();
    employee = {
      username : employee.username,
      email : employee.email,
      storeid : Number(employee.storeid),
      firstname : employee.firstname,
      lastname : employee.lastname
    };
    await axios
      .post(domain + "/employee/", employee, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("addEmployeeToState", employee))
      .catch(console.log("error adding employee"));
  },
  async deleteEmployee({ commit }, username) {
    let session = await Auth.currentSession();
    await axios
      .delete(domain + "/employee/", username, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then((response) => {
        //do something
        response;
        commit("deleteEmployeeFromState", username);
      })
      .catch({
        //do something else
      });
  },
  async updateEmployee({ commit }, employee) {
    let session = await Auth.currentSession();
    await axios
      .put(domain + "/employee/", employee, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then((response) => {
        //do something
        response;
        commit("updateEmployeeInState", employee);
      })
      .catch({
        //do something else
      });
  },
  async promoteToManager(employee) {
    let session = await Auth.currentSession();
    await axios
      .post(domain + "/manager/", employee.username, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then((response) => {
        //do something
        response;
        
      })
      .catch({
        //do something else
      });
  },
  async promoteToAdmin(employee) {
    let session = await Auth.currentSession();
    await axios
      .post(domain + "/admin/", employee.username ,{
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then((response) => {
        //do something
        response;
        
      })
      .catch({
        //do something else
      });
  },
  async fetchEmployees({ commit }) {
    let session = await Auth.currentSession();
    await axios
      .get(domain + "/employee", {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then((res) => commit("updateEmployees", res.data))
      .catch(console.log("error fetching employee list"));
  },
};

const mutations = {
  addEmployeeToState: (state, employee) => {
    state.employeeList.push(employee);
  },
  deleteEmployeeFromState: (state, username) => {
    state.employeeList = state.employeeList.filter((e) => e.username != username);
  },
  updateEmployeeInState: (state, employee) => {
  
    for (var i in state.employeeList) {
      if (state.employeeList[i].username == employee.username) {
        Vue.set(state.employeeList, i, employee);
        break;
      }
    }
  },
  updateEmployees: (state, employees) => {
    state.employeeList = employees;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
