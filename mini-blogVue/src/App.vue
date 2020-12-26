<template>
  <div id="app">
    <div class="circle"></div>
    <div class="circle1"></div>
    <nav-bar></nav-bar>
    <router-view/>
    
  </div>
  
</template>

<script>
//import Login from "./components/Login.vue";
import navBar from "./components/navBar.vue";
//import Inscription from "./components/Inscription.vue";
//import Home from "./views/Home.vue"

export default {
  name: "App",
  components: {
    navBar,
   
    //Login,
    
    //Inscription
  }
};
</script>

<style lang="css">

  @import './styles/app.css';

</style>
