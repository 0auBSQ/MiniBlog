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
      error: null,
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
      const auth = { email: this.email, password: this.password }
      const url = 'http://localhost:8888/api/register'

      console.log(auth);
      await this.axios.post(url + "?email=" + this.email + "&login=" + this.userName + "&password=" + this.password, {})
        .then(res => {
          if (res.status == 200){
            
            this.$router.push({path: '/tmp'})
          }
          console.log(res.status);
        })
        .catch(err => {
          console.log(err)
          if(err && err.response && err.reponse.status){
            if (err.response.status === 400)
              this.error = "Invalid information";
            else if (err.reponse.status === 500)
              this.error = "Internal server error";
            else
              this.error = err.message;
          }
        })
    }
  }
};
</script>

<style scoped>
    @import "../styles/form.css";
</style>