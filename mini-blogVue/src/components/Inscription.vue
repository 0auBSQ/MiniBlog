<template>
 <div>
   <form class="login" @submit.prevent="login">
     <h1>Register</h1>
     <label>User name</label>
     <input required v-model="username" type="text" placeholder="Snoopy"/>
     <label>E-mail</label>
     <input required v-model="password" type="email" placeholder="snoopy.snoopy@gmail.com"/>
     <label>Comfirm E-mail</label>
     <input required v-model="password" type="email"/>
     <label>Password</label>
     <input required v-model="password" type="password" placeholder="Password"/>
     <label>Comfirm Password</label>
     <input required v-model="password" type="password"/>

     <button type="submit">Register</button>
   </form>
 </div>
</template>

<script>
export default {
  name: "inscription",
  props: {
    msg: String
  }
};
</script>

<style scoped lang="css">
    @import '../styles/form.css';
</style>