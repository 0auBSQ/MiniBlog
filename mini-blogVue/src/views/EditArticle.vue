<template>
    <div class="articleContainer">
        <h1>Edit Article</h1>
        <form class="articleForm" @submit.prevent="edit">
            <div class="head">
                <input v-bind="title" type="text" class="title" v-model="title">
                <input v-bind="link" type="text" class="title" placeholder="Link to your image">
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

    methods: {
        async edit () {
            const url = 'http://localhost:8888/api/article/update/'
            console.log(url)

            await this.axios.patch(url + + "?aid=" + this.$route.params.id + "&title=" + this.title + "&content=" + this.content + "&img_link=" + this.link, {headers: {Authorization: this.$store.state.token}})
            .then(res => {
                if (res.status == 200){
                    this.$router.push({path: `/article/${this.id}`})
                }
            })
            .catch(err => {

            if(err && err.response && err.response.status){
              if (err.response.status === 401){
                this.error = "Erreur";
              }
              else if (err.response.status === 500)
                this.error = "Internal server error";
              else
                this.error = err.message;
            }
          })
        }
    },
    created: async function() {

          const url = 'http://localhost:8888/api/article/read/'
          console.log(url)

          await this.axios.get(url + "?aid=" + this.$route.params.id, {})
          .then(res => {
            if (res.status == 200){
                this.id = res.data.id
                this.title = res.data.title
                this.content = res.data.content
                this.link = res.data.img_link
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
        id: 0,
        content: null,
        link: null,
        title: null
    }
}
</script>

<style>
    @import "../styles/addArticle.css";
</style>
