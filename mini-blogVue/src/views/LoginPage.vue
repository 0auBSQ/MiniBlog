<template>
    <div class="containerlog">
        <login></login>
    </div>
</template>
<script>

import Login from "../components/Login.vue"


export default {
    name: 'LoginPage',
    components: {
        Login
    }

}
</script>
<style>
    
    .containerlog {
        width: 50%;
        padding: 30px;
        margin: 0 auto 0 auto;
        height: calc(60% - 50px);
    }

</style>