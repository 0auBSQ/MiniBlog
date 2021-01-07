<template>
    <div class="containerNav">
        <router-link class="navItem" to="/account/userinfo"><div class="text">Profile Information</div></router-link>
        <router-link class="navItem" v-if="admin > 0" to="/account/admin"><div class="text">Admin</div></router-link>
        <router-link class="navItem disconnect" to="" onclick="logout()"><div class="text">Disconnect</div></router-link>
    </div>
</template>

<script>


export default {
    name: "AccountNav",
   
    data () {
        return {
            admin: 1
        }
    },
    methods: {
        logout: async function()  {
             
            //const auth = { email: this.email, password: this.password }
            //const url = '127.0.0.1:8888/api/login'
            
            try {
                //const res = await this.axios.post(url, {auth}).then(res => res.data)
                //if(res.status == 200){
                this.$router.push({path: '/login'})
                //}
                
            } catch(err){
                this.error = err.message
            }
                
        }
    }
}
</script>

<style scoped>
    @import "../styles/accountnav.css";
</style>