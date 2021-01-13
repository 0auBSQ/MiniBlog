import Vue from "vue"
import Vuex from 'vuex'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.use(Vuex)
Vue.use(VueAxios,axios)

export default new Vuex.Store({
    state: {
        status: '',
        user: false,
        admin: false
    },
    getters: {
        statusUser: function (state) {
          return state.user;
        },
        statusAdmin: function (state) {
          return state.admin;
        },
        isAuth: async function (state)  {

            const url = 'http://localhost:8888/api/is_auth/user'

            await axios.get(url, {withCredentials: true})
            .then(res => {
                if (res.status == 200){
                    state.status = "success"
                    state.user = true
                    return true
                }
            })
            .catch(err => {

                if(err && err.response && err.response.status){
                    if (err.response.status === 401){
                        state.status = "The requested account doesn't exist";
                    }
                    else if (err.response.status === 500)
                        state.status = "Internal server error";
                    else
                        state.status = err.message;
                }
                state.user = false
            })
        },
        isAdmin: async function (state)  {

            const url = 'http://localhost:8888/api/is_auth/admin'

            await axios.get(url, {withCredentials: true})
            .then(res => {
                if (res.status == 200){
                    state.status = "success"
                    state.admin = true
                    state.user = true
                    return true
                }
            })
            .catch(err => {

                if(err && err.response && err.response.status){
                    if (err.response.status === 401){
                        state.status = "The requested account doesn't exist";
                    }
                    else if (err.response.status === 500)
                        state.status = "Internal server error";
                    else
                        state.status = err.message;
                }
                state.admin = false
            })
        },
        authStatus: state => state.status,

    },
    mutations: {

        authSuccess : async (state) => {
            state.status = 'success'
            state.user = true;
        },

        logoutSuccess : (state) => {
          state.admin = false;
          state.user = false;
        }

    },
    actions: {


    },
    modules: {

    }
})
