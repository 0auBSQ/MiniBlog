<template>
  <div id="app">
    <div class="circle"></div>
    <div class="circle1"></div>
    <nav-bar></nav-bar>
    <router-view/>
    
  </div>
  
</template>

<script>
import navBar from "./components/navBar.vue";


export default {
  name: "App",
  components: {
    navBar
  }
};
</script>

<style lang="css">

  @import './styles/app.css';

</style>
