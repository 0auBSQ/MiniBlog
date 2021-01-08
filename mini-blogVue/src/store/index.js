import Vue from "vue"
import Vuex from 'vuex'
import axios from 'axios'
import VueAxios from 'vue-axios'


Vue.use(Vuex)
Vue.use(VueAxios,axios)

export default new Vuex.Store({
    state: {
        status: '',
        isAdmin: false
    },
    getters: {
        isAuth: async function (state)  {

            const url = 'http://localhost:8888/api/is_auth'

            await axios.post(url, {})
            .then(res => {
                if (res.status == 200){
                    state.status = "success"
                    return true
                }
            })
            .catch(err => {
                
                if(err && err.response && err.reponse.status){
                    if (err.response.status === 404){
                        state.status = "The requested account doesn't exist";
                    }
                    else if (err.reponse.status === 500)
                        state.status = "Internal server error";
                    else
                        state.status = err.message;
                }
            })
        },
        authStatus: state => state.status,
        
    },
    mutations: {
        
        authSuccess : (state) => {
            state.status = 'success'
        }
        
    },
    actions: {
        

    },
    modules: {
        
    }
})