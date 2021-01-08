<template>
    <div class="postContainer">
        <div class="header">
            <img :src="article.Img_link" class="topart" />
            <br />
            <h1 class="titre">{{article.Title}}</h1>
            <h3 class="author"><i>{{article.Author}}</i></h3>
            <h4 class="date">{{article.Creation_date.split("T")[0]}}</h4>
        </div>
        <div class="content">
            <p class="text" >
                {{article.Contents}}
            </p>
        </div>
        <div class="commentSec">
            <h3 class="commentTitle">Comments</h3>
            <div v-if="statusUser == true" class="writeCom">
                <form class="com">
                    <label for="">Leave a comment</label>
                    <textarea type="text" class="cominput"></textarea>
                    <button type="submit">Comment</button>
                </form>
            </div>
            <hr class="sep">

            <comment v-for="comment in comments" :key="comment.Id" :user="comment.Author" :content="comment.Contents" :date="comment.Creation_date" :id="0"></comment>

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
            comments: [],
            article : {}
        }
    },
    created: async function() {

          const url = 'http://localhost:8888/api/comment/fetch/'+ this.$route.params.id
          const art = 'http://localhost:8888/api/article/read/'+ this.$route.params.id

          console.log(url)
          console.log(art)

          await this.axios.get(art, {})
          .then(res => {
            if (res.status == 200){
              this.article = res.data
              console.log(res.data)
            }
          })
          .catch(err => {
            if(err && err.response && err.response.status){
              if (err.response.status === 404){
                this.error = "The requested article doesn't exist";
              }
              else if (err.response.status === 500)
                this.error = "Internal server error";
              else
                this.error = err.message;
            }
          })


          await this.axios.get(url, {})
          .then(res => {
            if (res.status == 200){
              this.comments = res.data
            }
          })
          .catch(err => {
            if(err && err.response && err.response.status){
              if (err.response.status === 404){
                this.error = "The requested ressource doesn't exist";
              }
              else if (err.response.status === 500)
                this.error = "Internal server error";
              else
                this.error = err.message;
            }

          })

    },
    computed: {

    ...mapGetters({
        statusUser:'statusUser',
    })


  }
}

</script>

<style>
    @import "../styles/post.css";
</style>
