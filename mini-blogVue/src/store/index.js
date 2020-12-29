import axios from 'axios'
import { createStore } from 'vuex'

export default createStore({
    state: {
        token: localStorage.getItem('user-token') || '',
        status: '',
    },
    getters: {
        isAuth: state => !!state.token,
        authStatus: state => state.status,

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
        }
    },
    actions: {
        [AUTH_REQUEST]: ({commit, dispatch}, user) => {
            return new Promise((resolve, reject) => {
                commit(AUTH_REQUEST)
                axios({url:'auth',data: user, method: 'POST'})
                .then(resp => {
                    const token = rep.data.token
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
        }
    },
    modules: {
        
    }
})