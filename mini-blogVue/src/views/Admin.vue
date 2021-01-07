<template>
    <div class="containerAdmin">
        <h2>Admin</h2>
        <div class="banned">

        </div>
        <div class="addArticle">
            <h2>Add Article</h2>
            <form class="article" @submiut.prevent="addarticle">
                <textarea type="text" class="textarea" v-model="text"></textarea>
                <button type="submit">Submit</button>
            </form>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Admin',
    data () {
        return {
            text: null,
        }
    },
    methods: {
        addarticle: function () {

        }
    }
}
</script>

<style>
    @import "../styles/admin.css";
</style>