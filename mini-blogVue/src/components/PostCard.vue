<template>
    <div class="containerPost" @click="goArticle(id)">
        <div class="postTitle">
            <h1>{{title}}</h1>
        </div>
        <div class="author">
            <p>{{author}}</p>
        </div>
    </div>
</template>

<script>
export default {
    name: "PostCard",
    props: {
        title: String,
        author: String,
        id: Number,
        link: String,
    },
    methods: {
        goArticle(id) {
            console.log(id)
            this.$router.push(`/article/${id}`)
        }
    }

}
</script>

<style>
    @import "../styles/postCard.css";
</style>
