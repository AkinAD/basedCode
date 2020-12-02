import axios from 'axios'

const state = {
    employeeList = []
}


const actions = {



    async deleteEmployee({commit}, id) {
        await axios.delete("...").then({
            //do something
        }).catch({
          //do something else  
        });
    } 


}


const mutations = {

}