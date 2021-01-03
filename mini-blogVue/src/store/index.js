import Vue from "vue"
import axios from 'axios'
import Vuex from 'vuex'

const AUTH_ERROR = 'AUTH_ERROR'
const AUTH_SUCCESS = 'AUTH_SUCCESS'
const AUTH_REQUEST = 'AUTH_REQUEST'
const AUTH_LOGOUT = 'AUTH_LOGOUT'
const USER_REQUEST = 'USER_REQUEST'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        token: localStorage.getItem('user-token') || '',
        status: '',
    },
    getters: {
        isAuth: ()/*state*/ => true,//!!state.token,
        authStatus: state => state.status,
        account: (getters) => {
            if(getters.isAuth) return false
            else return true
        },
        login: (getters) => {
            if(getters.isAuth) return true
            else return false
        }
    },
    mutations: {
        [AUTH_REQUEST]: (state) => {
            state.status = 'loading'
        },
        [AUTH_SUCCESS] : (state, token) => {
            state.status = 'success',
            state.token = token
        },
        [AUTH_ERROR]: (state) => {
            state.status = 'error'
        },
        [AUTH_LOGOUT]: (state) => {
            state.token = ''
        }
    },
    actions: {
        [AUTH_REQUEST]: ({commit, dispatch}, user) => {
            return new Promise((resolve, reject) => {
                commit(AUTH_REQUEST)
                axios({url:'auth',data: user, method: 'POST'})
                .then(resp => {
                    const token = resp.data.token
                    localStorage.setItem('user-token',token)
                    commit(AUTH_SUCCESS, token)
                    dispatch(USER_REQUEST)
                    resolve(resp)
                })
                .catch(err => {
                    commit(AUTH_ERROR, err)
                    localStorage.removeItem('user-token')
                    reject(err)
                })
            })
        },
        [AUTH_LOGOUT]: ({commit}) => {
            return new Promise((resolve) => {
                commit(AUTH_LOGOUT)
                localStorage.removeItem('user-token')
                resolve()
            })
           
        }

    },
    modules: {
        
    }
})