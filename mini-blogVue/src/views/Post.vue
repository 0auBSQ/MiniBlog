<template>
    <div class="postContainer">
        <div class="header">
            <h1 class="titre">Mon article</h1>
            <h3 class="author"><i>Author 1</i></h3>
            <h4 class="date">16/09/2020</h4>
        </div>
        <div class="content">
            <p class="text" >
                Lorem ipsum dolor sit, amet consectetur adipisicing elit. 
                Ullam autem voluptatem quam iusto perspiciatis dolor aspernatur, 
                corrupti nesciunt natus voluptatum beatae sint placeat incidunt 
                cumque porro voluptas adipisci? Odio, omnis.
            </p>
        </div>
        <div class="commentSec">
            <h3 class="commentTitle">Comments</h3>
            <hr class="sep">
            <comment v-for="comment in comments" :key="comment.id" v-bind:user="comment.user" v-bind:content="comment.content" v-bind:date="comment.date"></comment>
        </div>
    </div>
</template>

<script>
import Comment from "../components/Comment.vue"


export default {
    name: "Post",
    components: {
        Comment
    },
    data () {
        return {
            comments: [
                {
                    id: 0,
                    user: 'user 1',
                    date: '13/03/09',
                    content: "nice!"
                },
                {
                    id: 1,
                    user: 'user 1',
                    date: '13/03/09',
                    content: "so bad!"
                },
                {
                    id: 2,
                    date: '13/03/09',
                    user: 'user 1',
                    content: `Lorem ipsum dolor sit, amet consectetur adipisicing elit. 
                                Ullam autem voluptatem quam iusto perspiciatis dolor aspernatur, 
                                corrupti nesciunt natus voluptatum beatae sint placeat incidunt 
                                cumque porro voluptas adipisci? Odio, omnis.`
                },
                {
                    id: 3,
                    date: '13/03/09',
                    user: 'user 1',
                    content: `Lorem ipsum dolor sit, amet consectetur adipisicing elit. 
                                Ullam autem voluptatem quam iusto perspiciatis dolor aspernatur`
                },
            ]
        }
    },
    created: {
        
    }
}

</script>

<style>
    @import "../styles/post.css";
</style>