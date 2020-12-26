<template>
  <div>
    <form class="login register" @submit.prevent="login">
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
    checkForm: (e) => {
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
        return true;
      }

      e.preventDefault();
    },
    validEmail: function (email) {
      var re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    }
  }
};
</script>

<style scoped lang="css">
    @import "../styles/form.css";
</style>