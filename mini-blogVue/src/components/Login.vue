<template>
 <div>
   <form class="login" @submit.prevent="login">
     <h1>Sign in</h1>
     <label>User name</label>
     <input required v-model="userName" type="text" placeholder="Snoopy"/>
     <label>Password</label>
     <input required v-model="password" type="password" placeholder="Password"/>
     <router-link to="/register">Register</router-link>
     <button type="submit">Login</button>
   </form>
 </div>
</template>

<script>
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
    /*login: function () {
      const { username, password } = this
      this.$store.dispatch(AUTH_REQUEST, { username, password }).then(() => {
        this.$router.push('/')
      })
    }*/
  }
}
</script>

<style scoped lang="css">
  @import '../styles/form.css';
</style>
