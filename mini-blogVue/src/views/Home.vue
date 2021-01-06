<template>
    <div class="homeContainer">
        
        <div class="postContainer">
          <postCard class="card" v-for="article in articleList" :key="article.title" v-bind:title="article.title" v-bind:author="article.author" v-bind:id="article.id"></postCard>
        </div>
        <div class="aside">
          <div class="search">
            <h3 class="menu">Menu</h3>
            <form action="">
              <input type="text">
              <button type="submit">Search</button>
            </form>
          </div>
          <hr class="sep">
          <div class="dateQ">
            
          </div>
        </div>
    </div>
</template>
<script>
import PostCard from "../components/PostCard.vue"

export default {
  name: 'Home',
  data () {
    return {
      articleList: [
        {
          id: 0,
          title: "mon incroyable article 1",
          author: "author 1"
        },
        {
          id: 1,
          title: "mon incroyable article 2",
          author: "author 2"
        },
        {
          id: 2,
          title: "mon incroyable article 3",
          author: "author 3"
        },
        {
          id: 3,
          title: "mon incroyable article 4, le plus beau du monde",
          author: "author 4"
        },
        {
          id: 4,
          title: "mon incroyable dernier article",
          author: "author 3"
        },
      ]
    }
  },
  components: {
    PostCard,
  },
  created: {

  }

}
</script>

<style>
  @import "../styles/home.css";
</style>