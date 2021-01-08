<template>
    <div class="homeContainer">
        
        <div class="postContainer">
          <postCard class="card" :style="{ backgroundImage: 'url(' + article.link + ')'}" v-for="article in articleList" :key="article.Id" :title="article.Title" :author="article.Author" :id="article.Id"></postCard>
        </div>
        <div class="aside">
          <div class="search">
            <h3 class="menu">Menu</h3>
            <form action="">
              <input type="text">
              <button type="submit">Search</button>
            </form>
          </div>
          
          <div class="dateQ">
            
          </div>
        </div>
        <button v-if="isAdmin == true" class="bouton1" @click="goAdd">+</button>
    </div>
</template>
<script>
import PostCard from "../components/PostCard.vue"
import { mapGetters } from 'vuex'


export default {
  name: 'Home',
  data () {
    return {
      articleList: []
    }
  },
  components: {
    PostCard,
  },
  methods: {
        goAdd () {
            this.$router.push('/addarticle')
        },
        async getData () {
          const url = 'http://localhost:8888/api/article/fetch'

          await this.axios.get(url, {})
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
        
  },
  created () {
    this.getData()
  },
  computed: {
    ...mapGetters(
      'isAdmin'
    )
  }
}
</script>

<style>
  @import "../styles/home.css";
</style>