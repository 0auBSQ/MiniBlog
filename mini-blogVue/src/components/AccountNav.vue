<template>
    <div class="containerNav">
        <router-link class="navItem" to="/account/userinfo"><div class="text">Profile Information</div></router-link>
        <router-link class="navItem" v-if="admin > 0" to="/account/admin"><div class="text">Admin</div></router-link>
        <router-link class="navItem disconnect" to="/login"><div class="text">Disconnect</div></router-link>
    </div>
</template>

<script>

export default {
    name: 'AccountNav',
   
    data () {
        return {
            admin: 1
        }
    }
}
</script>

<style scoped>
    @import "../styles/accountnav.css";
</style>