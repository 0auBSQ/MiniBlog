<template>
 <div class="containerTmp"> 
   <p class="success">Your account has been created !</p>
   <p class="time">You will be redirected in 5 seconds</p>
   <p class="time">If not, click on "Go to Login"</p>
   <button @click="redirect">Go to Login</button>
 </div>
</template>

<script>

export default {
    name: "tmp",
    methods: {
        redirect () {
            this.$router.push('/login')
        }
    },
    created: function timer () {
            console.log("test")
            setTimeout(this.redirect, 5000);
    }
    
}
</script>

<style scoped >
  @import '../styles/tmp.css';
</style>
