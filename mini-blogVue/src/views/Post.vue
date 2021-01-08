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
            <div v-if="isAuth == true" class="writeCom">
                <form class="com">
                    <label for="">Leave a comment</label>
                    <textarea type="text" class="cominput"></textarea>
                    <button type="submit">Comment</button>
                </form>
            </div>
            <hr class="sep">
            <comment v-for="comment in comments" :key="comment.id" :user="comment.user" :content="comment.content" :date="comment.date" :id="comment.id"></comment>
        </div>
    </div>
</template>

<script>
import Comment from "../components/Comment.vue"
import { mapGetters } from 'vuex'



export default {
    name: "Post",
    components: {
        Comment
    },
    data () {
        return {
            comments: []
        }
    },
    created: async function() {
       
          const url = 'http://localhost:8888/api/article/read/  '+ this.$route.params.id
          console.log(url)

          await this.axios.get(url, {})
          .then(res => {
            if (res.status == 200){
              this.comments.push(res.data)
            }
          })
          .catch(err => {
            
            if(err && err.response && err.response.status){
              if (err.response.status === 404){
                this.error = "The requested account doesn't exist";
              }
              else if (err.response.status === 500)
                this.error = "Internal server error";
              else
                this.error = err.message;
            }
          })    
        
    },
    computed: {

    ...mapGetters([
      'isAuth',
    ])

    
  }
}

</script>

<style>
    @import "../styles/post.css";
</style>