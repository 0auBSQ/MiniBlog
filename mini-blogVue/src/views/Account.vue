<template>
    <div class="container">
        
        <account-nav class="nav"/>
        <transition name="anim" mode="out-in" enter-active-class="animated fadeInLeft" leave-active-class="animated fadeOutRight">
            <router-view class="view" name="info"/>
        </transition>
    </div>
</template>

<script>

import AccountNav from "../components/AccountNav.vue"

export default {
    name: 'Account',
    components: {
       AccountNav
    }
}
</script>

<style>
    @import "https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.7.0/animate.min.css";
    @import "../styles/account.css";
</style>