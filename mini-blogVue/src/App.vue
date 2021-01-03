<template>
  <div id="app">
    <div class="circle"></div>
    <div class="circle1"></div>
    <nav-bar></nav-bar>
    <transition name="anim" mode="out-in" enter-active-class="animated fadeInLeft" leave-active-class="animated fadeOutRight">
      <router-view/>
    </transition>
    
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
  @import "https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.7.0/animate.min.css";
  @import './styles/app.css';

  .animated {
    -webkit-animation-duration: 0.5s;
    animation-duration: 0.5s;
  }

</style>    
