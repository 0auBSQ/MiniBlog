<template>
 <div>
   <form class="login" @submit.prevent="log">
     <h1>Sign in</h1>
     <label>Email</label>
     <input required v-model="email" type="email" placeholder="Snoopy@gmail.com"/>
     <label>Password</label>
     <input required v-model="password" type="password" placeholder="Password"/>
     <router-link to="/register" class="regis">Register</router-link>
     <button type="submit" >Login</button>
     <div v-if="error" class="error">
          {{error}}
        </div>
   </form>
 </div>
</template>

<script>
import { mapMutations } from 'vuex'
//import file from "../API/api"

export default {
  name: "login",
  props: {
    
  },
  data () {
    return {
      email: null,
      error: null,
      password: null,
      res: null
    }
  },
  computed: {

    ...mapMutations([
      'authSuccess'
    ])

  },
  methods: {
    log: async function () {
      //const auth = { email: this.email, password: this.password }
      //const url = '127.0.0.1:8888/api/login'
      
      try {
        //const res = await this.axios.post(url, {auth}).then(res => res.data)
        //if(res.status == 200){
          this.$store.commit('authSuccess')
          this.$router.push({path: '/account'})
        //}
        
      } catch(err){
        this.error = err.message
      }
    }
  }
}
</script>

<style scoped >
  @import '../styles/form.css';
</style>
