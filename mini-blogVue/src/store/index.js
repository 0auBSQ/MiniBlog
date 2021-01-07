import Vue from "vue"
//import axios from 'axios'
import Vuex from 'vuex'



Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        status: '',
    },
    getters: {
        isAuth: (state)/*state*/ => {
            
            if(state.status == "success") return true
            else return false

            /*const is_auth = { }
            const url = '127.0.0.1:8888/api/is_auth'
      
            try {
                const res = await this.axios.get(url, {auth}).then(res => res.data)
                if(res.status == 200){
                    return true
                }
                else return false
                
            } catch(err){
                state.status = err.message
            }*/

            //!!state.token,
        },
        authStatus: state => state.status,
        
    },
    mutations: {
        authRequest: (state) => {
            state.status = 'loading'
        },
        authSuccess : (state) => {
            state.status = 'success'
        },
        authError: (state) => {
            state.status = 'error'
        },
        authLogOut: (state) => {
            state.token = ''
        }
    },
    actions: {
        

    },
    modules: {
        
    }
})