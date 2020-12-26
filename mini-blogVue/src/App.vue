<template>
  <div id="app">
    <div class="circle"></div>
    <div class="circle1"></div>
    <nav-bar></nav-bar>
    <login></login> 
    <!--<inscription></inscription>-->
  </div>
</template>

<script>
import Login from "./components/Login.vue";
import navBar from "./components/navBar.vue";
//import Inscription from "./components/Inscription.vue";

export default {
  name: "App",
  components: {
    Login,
    navBar
    //Inscription
  }
};
</script>

<style lang="css">

  @import './styles/app.css';

</style>
