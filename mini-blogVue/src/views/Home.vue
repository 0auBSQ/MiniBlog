<template>
    <div class="homeContainer">

        <div class="postContainer">
          <postCard class="card" :style="{ backgroundImage: 'url(' + article.Img_link + ')', backgroundSize : 'cover',  boxShadow : 'inset 0 0 0 2000px rgba(150, 150, 150, 0.7)'}" v-for="article in articleList" :key="article.Id" :title="article.Title" :author="article.Author" :id="article.Id"></postCard>
        </div>
        <div class="aside">
          <div class="search">
            <h3 class="menu">Search</h3>
            <form action="">
              <input v-model="search_bar" type="text" v-on:keyup="getData">
            </form>
          </div>

          <div class="dateQ">

          </div>
        </div>

        <button v-if="statusAdmin == true" class="bouton1" @click="goAdd">+</button>

    </div>
</template>
<script>
import PostCard from "../components/PostCard.vue"
import { mapGetters } from 'vuex'


export default {
  name: 'Home',
  data () {
    return {
      articleList: [],
      search_bar: ""
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
          const url = 'http://localhost:8888/api/article/fetch?search=' + this.search_bar;

          console.log(this.search_bar);
          await this.axios.get(url, {})
          .then(res => {
            console.log(res)
            if (res.status == 200){
              this.articleList = res.data
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
        }

  },
  async created () {
    this.getData();
  },
  computed: {
    ...mapGetters({
      statusAdmin: 'statusAdmin',
    })
  }
}
</script>

<style>
  @import "../styles/home.css";
</style>
