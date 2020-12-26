<template>
  <nav class="topnav" id="myTopnav">
    <router-link exact-active-class="active" class="navBarItem" v-for="routes in links" v-bind:key="routes.id" :to="`${routes.page}`">{{routes.text}}</router-link>
    <div class="log">
      <router-link exact-active-class="active" class="navBarItem" to="/login">Login</router-link>
    </div>
  </nav>
</template>

<script>
export default {
  name: "navBar",
  props: {
    
  },
  data () {
    return {
      links: [
        {
          id: 0,
          text: 'Home',
          page: '/'
        },
        {
          id: 1,
          text: 'News',
          page: '/news'
        }

      ]
    }
  }
}
</script>

<style scoped lang="css">
  @import '../styles/nav.css';
</style>