<template>
    <div class="commentContainer">
        <div class="commentHeader">
            <h4 class="user"> {{user}}</h4>
            
            <p class="dateComment"><i> - {{date}}</i></p>
        </div>
        <div class="thecomment">
            <p>
                {{content}}
            </p>
        </div>
    </div>
</template>

<script>
export default {
    name: "Comment",
    props: {
        user: String,
        date: String,
        content: String,
        id: Number
    }
}
</script>

<style scoped>
    @import "../styles/comment.css";
</style>