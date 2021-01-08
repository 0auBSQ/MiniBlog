<template>
    <div class="articleContainer">
        <h1>Add Article</h1>
        <form class="articleForm" @submit.prevent="add">
            <div class="head">
                <input v-bind="title" type="text" class="title" placeholder="Your title">
                <input v-bind="link" type="text" class="title" placeholder="Link to your image">
            </div>
            <div class="content">
                <textarea v-bind="content" name="" class="body" placeholder="Your content"></textarea>
                
            </div>
            <button type="submit">Add Article</button>
        </form>
    </div>
</template>

<script>
export default {
    name: 'addArticle',
    data () {
        return {
            id: 0,
            content: "mon article",
            link: null,
            title: "titre de mon article"
        }
    },
    methods: {
        add () {

        }
    }
}
</script>

<style>
    @import "../styles/addArticle.css";
</style>