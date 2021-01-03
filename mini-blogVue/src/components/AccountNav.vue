<template>
    <div class="containerNav">
        <router-link class="navItem" to="/account/userinfo"><div class="text">Profile Information</div></router-link>
        <router-link class="navItem" v-if="admin > 0" to="/account/admin"><div class="text">Admin</div></router-link>
        <router-link class="navItem disconnect" to="" onclick="logout()"><div class="text">Disconnect</div></router-link>
    </div>
</template>

<script>

const AUTH_LOGOUT = 'AUTH_LOGOUT'

export default {
    name: "AccountNav",
   
    data () {
        return {
            admin: 1
        }
    },
    methods: {
        logout: () => {
            
            this.$store.dispatch(AUTH_LOGOUT).then(() => {
                this.$router.push("/login")
            })
        }
    }
}
</script>

<style scoped>
    @import "../styles/accountnav.css";
</style>