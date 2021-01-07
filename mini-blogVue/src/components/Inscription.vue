<template>
  <div>
    <form class="login register" @submit.prevent="checkForm">
      <h1>Register</h1>
      <label>User name</label>
      <input required v-model="userName" type="text" placeholder="Snoopy" />
      <label>E-mail</label>
      <input required v-model="email" type="email" placeholder="snoopy.snoopy@gmail.com" />
      <label>Comfirm E-mail</label>
      <input required v-model="confemail" type="email" />
      <label>Password</label>
      <input required v-model="password" type="password" placeholder="Password" />
      <label>Comfirm Password</label>
      <input required v-model="confpassword" type="password" />

      <button type="submit">Register</button>
    </form>
  </div>
</template>

<script>
export default {
  name: "inscription",
  props: {
    msg: String
  },
  data () {
    return {
      errors: [],
      userName: null,
      email: null,
      confemail: null,
      password: null,
      confpassword: null
    }
  },
  methods: {
    checkForm: function () {
      this.errors = [];

      if (!this.userName) {
        this.errors.push("User Name required.");
      }
      if (!this.email) {
        this.errors.push("Email required.");
      } else if (!this.validEmail(this.email)) {
        this.errors.push("Valid email required.");
      } else if(this.confemail != this.email){
        this.errors.push("Email must be the same");
      }
      if(this.password != this.confpassword){
        this.errors.push("password must be the same");
      }

      if (!this.errors.length) {
        this.register()
      }
      
      
    },
    validEmail: function (email) {
      var re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    },
    register: async function () {
      //const regist = { email: this.email, login: this.userName, password: this.password }
      //const url = '127.0.0.1:8888/api/register'
      
      try {
        //const res = await this.axios.post(url, {regist}).then(res => res.data)
        //this.res = res.data
        
        this.$router.push({path: '/login'})
        
      } catch(err){
        this.error = err.message
      }
    }
  }
};
</script>

<style scoped>
    @import "../styles/form.css";
</style>