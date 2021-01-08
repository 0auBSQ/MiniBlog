<template>
    <div class="commentContainer">
        <button class="del" @click="this.delete">X</button>
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
    },
    methods: {
        async delete () {
            const url = 'http://localhost:8888/api/article/fetch'

            await this.axios.post(url, {})
            .then(res => {
            if (res.status == 200){
              this.articleList.push(res.data)
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
        }
    }
}
</script>

<style scoped>
    @import "../styles/comment.css";
</style>