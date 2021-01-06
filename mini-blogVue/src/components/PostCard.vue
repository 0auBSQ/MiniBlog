<template>
    <div class="containerPost">
        <div class="postTitle">
            <h1>{{title}}</h1>
        </div>
        <div class="author">
            <p>{{author}}</p>
        </div>
        <router-link :to="{ path: `/article/${id}` }"><i>See more...</i></router-link>
    </div>
</template>

<script>
export default {
    name: "PostCard",
    props: {
        title: String,
        author: String,
        id: Number
    }

}
</script>

<style>
    @import "../styles/postCard.css";
</style>