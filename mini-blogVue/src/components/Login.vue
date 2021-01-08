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
      const url = 'http://localhost:8888/api/login'

      await this.axios.get(url + "?email=" + this.email + "&password=" + this.password, {withCredentials: true})
        .then(res => {
          if (res.status == 200){
            this.$store.commit('authSuccess')
            this.$router.push({path: '/account'})
          }
        })
        .catch(err => {

          if(err && err.response && err.response.status){
          if (err.response.status === 404){
            this.error = "The requested account doesn't exist";
          }
          else if (err.response.status === 500)
            this.error = "Internal server error";
          else
            this.error = err.message;
          }
        })
    }
  }
}
</script>

<style scoped >
  @import '../styles/form.css';
</style>
