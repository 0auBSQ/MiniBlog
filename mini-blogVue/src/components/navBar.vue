<template>
  <nav class="topnav" id="myTopnav">
    <router-link exact-active-class="active" class="navBarItem" v-for="routes in links" v-bind:key="routes.id" :to="`${routes.page}`">{{routes.text}}</router-link>
    <div class="log">
      <router-link v-if="statusUser != true" exact-active-class="active" class="navBarItem" to="/login">Login</router-link>
      <router-link v-if="statusUser == true" exact-active-class="active" class="navBarItem" to="/account">Account</router-link>
    </div>
  </nav>
</template>

<script>
import { mapGetters } from 'vuex'

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
  },
  async created () {
    if (this.statusAdmin != true)
      await this.isAdmin;
    if (this.statusUser != true)
      await this.isAuth;
  },
  computed: {

    ...mapGetters({
      statusUser: 'statusUser',
      statusAdmin: 'statusAdmin',
      isAuth: 'isAuth',
      isAdmin: 'isAdmin',
    })


  }
}
</script>

<style scoped lang="css">
  @import '../styles/nav.css';
</style>
