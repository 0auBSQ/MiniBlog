<template>
    <div class="articleContainer">
        <h1>Edit Article</h1>
        <form class="articleForm" @submit.prevent="edit">
            <div class="head">
                <input v-bind="title" type="text" class="title" v-model="title">
            </div>
            <div class="content">
                <textarea v-bind="content" name="" class="body" v-model="content"></textarea>
            </div>
            <button type="submit">Add Article</button>
        </form>
    </div>
</template>

<script>
export default {
    name: 'editArticle',
    data () {
        return {
            id: 0,
            content: "mon article",
            title: "titre de mon article"
        }
    },
    methods: {
        edit () {

        }
    },
    created: function () {

    }
}
</script>

<style>
    @import "../styles/addArticle.css";
</style>