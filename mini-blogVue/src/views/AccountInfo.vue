<template>
    <div class="containerInfo">
        <h2>Profile Information</h2>
        <form class="formaccount">
            <div class="field">
                <label>First Name</label>
                <input v-model="firstName" type="text" placeholder=""/>
            </div>
            <div class="field">
                <label>Last Name</label>
                <input v-model="lastName" type="text" placeholder=""/>
            </div>
            <div class="field">
                <label>User name</label>
                <input v-model="userName" type="text" placeholder=""/>
            </div>
            <hr/>
            <div class="field">
                <label>Register Date: </label>
                <p>{{registerDate}}</p>
            </div>
            <hr/>
            <div class="field">
                <label>Last login Date: </label>
                <p>{{lastloginDate}}</p>
            </div>
            <!--<hr/>
            <div class="field">
                <label>Gender</label>
                <select v-model="gender" placeholder="">
                    <option value="male">Male</option>
                    <option value="female">Female</option>
                    <option value="other">Other</option>
                </select>
            </div>-->
            <hr/>
            <div class="field">
                <label>Date of Birth</label>
                <input required v-model="dateBirth" type="date" placeholder=""/>
            </div>
            <hr/>
            <button type="submit">Save</button>
        </form>
    </div>
</template>

<script>
export default {
    name: 'AccountInfo',
    data () {
        return {
            firstName: null,
            lastName: null,
            userName: null,
            registerDate: null,
            gender: null,
            dateBirth: null,
            lastloginDate: null,
        }
    },
    methods: {
        
    }
}
</script>

<style>
    @import "../styles/accountInfo.css"; 
</style>