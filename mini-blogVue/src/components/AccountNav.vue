<template>
    <div class="containerNav">
        <router-link class="navItem" to="/account/userinfo"><div class="text">Profile Information</div></router-link>
        <router-link class="navItem" v-if="statusAdmin == true" to="/account/admin"><div class="text">Admin</div></router-link>
        <router-link event="" class="navItem disconnect" to="" @click.native.prevent="logout"><div class="text">Disconnect</div></router-link>
    </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'

export default {
    name: "AccountNav",

    data () {
        return {
            admin: 1
        }
    },
    methods: {
        logout: async function () {
          console.log("test logout")

          this.$store.commit('logoutSuccess');
          this.$router.push({path: '/login'});
        }
    },
    computed: {
      ...mapGetters({
         statusAdmin: 'statusAdmin',
      }),
      ...mapMutations([
        'logoutSuccess'
      ])
    }
}
</script>

<style scoped>
    @import "../styles/accountnav.css";
</style>
