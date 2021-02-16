import Vue from "vue"
import Vuex from 'vuex'
import axios from 'axios'
import VueAxios from 'vue-axios'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)
Vue.use(VueAxios,axios)

export default new Vuex.Store({
    plugins: [createPersistedState({
        storage: window.sessionStorage,
    })],
    state: {
        status: '',
        user: false,
        admin: false,
        token: "None"
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

            console.log(state.token);

            await axios.get(url, {headers: {Authorization: state.token}})
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

            await axios.get(url, {headers: {Authorization: state.token}})
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

        authSuccess : async (state, tok) => {
            console.log("Token : ", tok);
            state.token = tok;
            state.status = 'success';
            state.user = true;
        },

        logoutSuccess : (state) => {
          state.token = "";
          state.admin = false;
          state.user = false;
        }

    },
    actions: {


    },
    modules: {

    }
})
