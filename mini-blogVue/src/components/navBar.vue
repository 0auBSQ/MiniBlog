<template>
  <nav class="topnav" id="myTopnav">
    <router-link class="navBarItem" v-for="routes in links" v-bind:key="routes.id" :to="`${routes.page}`">{{routes.text}}</router-link>
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
        },
        {
          id: 2,
          text: 'Login',
          page: '/login'
        }

      ]
    }
  }
}
</script>

<style scoped lang="css">
  @import '../styles/nav.css';
</style>