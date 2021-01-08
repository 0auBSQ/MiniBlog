<template>
    <div class="containerAdmin">
        <h2>Admin</h2>
        <div class="banned">
            <h3>Ban Machine</h3>
            <div class="user" v-for="user in users" :key="user.id">
                <p>{{user.userName}}</p>
                <button class="ban">Ban</button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Admin',
    data () {
        return {
            users: [
                {
                    id: 0,
                    userName: "jean",
                },
                {
                    id: 1,
                    userName: "jacque",
                },
                {
                    id: 2,
                    userName: "maxime",
                },
            ]
        }
    },
    methods: {
        banUser: function () {
            
        }
    }
}
</script>

<style>
    @import "../styles/admin.css";
</style>