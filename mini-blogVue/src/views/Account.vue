<template>
    <div class="container">
        <account-nav/>
        <router-view name="info"/>
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

</style>