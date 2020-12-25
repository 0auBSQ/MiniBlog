<template>
 <div>
   <form class="login" @submit.prevent="login">
     <h1>Inscription</h1>
     <label>User name</label>
     <input required v-model="username" type="text" placeholder="Snoopy"/>
     <label>E-mail</label>
     <input required v-model="password" type="email" placeholder="x.x@gmail.com"/>
     <label>Comfirm E-mail</label>
     <input required v-model="password" type="email"/>
     <label>Password</label>
     <input required v-model="password" type="password" placeholder="Password"/>
     <label>Comfirm Password</label>
     <input required v-model="password" type="password"/>
     <hr/>
     <button type="submit">Login</button>
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

<style scoped lang="less">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>