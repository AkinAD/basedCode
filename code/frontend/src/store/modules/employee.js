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
      firstName : employee.firstName,
      lastName : employee.lastName
    };
    console.log(employee);
    await axios
      .post(domain + "/employee", employee, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("addEmployeeToState", employee))
      .catch(console.log("error adding employee"));
  },
  async deleteEmployee({ commit }, username) {
    let session = await Auth.currentSession();
    let employee = {
      username: username
    };
    console.log(employee);
    await axios
      .delete(domain + "/employee", {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        },
        data: {
          username: employee.username
        }
      })
      .then(commit("deleteEmployeeFromState", username))
      .catch();
  },
  async updateEmployee({ commit }, employee) {
    let session = await Auth.currentSession();
    console.log(employee);
    await axios
      .put(domain + "/employee", employee, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then(commit("updateEmployeeInState", employee))
      .catch();
  },
  async promoteToManager({state}, employee) {
    let session = await Auth.currentSession();
    state;
    let username = {
      username: employee.username
    };
    console.log(username);
    console.log(employee);
    await axios
      .post(domain + "/manager", username, {
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then()
      .catch();
  },
  async promoteToAdmin({state},employee) {
    let session = await Auth.currentSession();
    state;
    let username = {
      username: employee.username
    };
    await axios
      .post(domain + "/admin", username ,{
        headers: {
        'Authorization': `Bearer ${session.getAccessToken().getJwtToken()}`
        }
      })
      .then()
      .catch();
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
