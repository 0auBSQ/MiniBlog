<template>
    <div class="container">
        <router-link to="/account/userinfo">Profile Information</router-link>
        <router-link v-if="admin > 0" to="/account/admin">Admin</router-link>
        <router-link to="/login">Disconnect</router-link>
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

<style>
    
</style>