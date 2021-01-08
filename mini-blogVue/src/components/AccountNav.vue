<template>
    <div class="containerNav">
        <router-link class="navItem" to="/account/userinfo"><div class="text">Profile Information</div></router-link>
        <router-link class="navItem" v-if="isAdmin == true" to="/account/admin"><div class="text">Admin</div></router-link>
        <router-link class="navItem disconnect" to="" @click="logout"><div class="text">Disconnect</div></router-link>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
    name: "AccountNav",
   
    data () {
        return {
            admin: 1
        }
    },
    methods: {
        logout: async function () {
        const url = 'http://localhost:8888/api/logout'

        await this.axios.post(url, {withCredentials: true})
        .then(res => {
          if (res.status == 200){
            this.$router.push({path: '/login'})
          }
        })
        .catch(err => {
          
          if(err && err.response && err.response.status){
          if (err.response.status === 404){
            this.error = "The requested account doesn't exist";
          }
          else if (err.response.status === 500)
            this.error = "Internal server error";
          else
            this.error = err.message;
          }
        })    
    }
    },
    computed: {
      ...mapGetters(
        'isAdmin'
      )
    }
}
</script>

<style scoped>
    @import "../styles/accountnav.css";
</style>