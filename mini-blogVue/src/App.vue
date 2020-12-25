<template>
  <div id="app">
    <login></login> 
    <!--<inscription></inscription>-->
  </div>
</template>

<script>
import Login from "./components/Login.vue";
//import Inscription from "./components/Inscription.vue";

export default {
  name: "App",
  components: {
    Login
    //Inscription
  }
};
</script>

<style lang="css">

  @import './styles/app.css';

</style>
