<template>
 <div>
   <form class="login" action="log()">
     <h1>Sign in</h1>
     <label>Email</label>
     <input required v-model="userName" type="email" placeholder="Snoopy@gmail.com"/>
     <label>Password</label>
     <input required v-model="password" type="password" placeholder="Password"/>
     <router-link to="/register" class="regis">Register</router-link>
     <button type="submit" >Login</button>
   </form>
 </div>
</template>

<script>
import file from "../API/api"

export default {
  name: "login",
  props: {
    
  },
  data () {
    return {
      userName: null,
      password: null
    }
  },
  methods: {
    log: function () {
      file.login()
    }
  }
}
</script>

<style scoped >
  @import '../styles/form.css';
</style>
