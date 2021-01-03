<template>
    <div class="containerPost">
        <div class="postTitle">
            <h1>{{title}}</h1>
        </div>
        <div class="author">
            <p>{{author}}</p>
        </div>
        <a class="see" href="/"><i>See more...</i></a>
    </div>
</template>

<script>
export default {
    name: "PostCard",
    props: {
        title: String,
        author: String
    }

}
</script>

<style>
    @import "../styles/postCard.css";
</style>