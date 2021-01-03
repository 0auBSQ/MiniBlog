<template>
    <div class="homeContainer">
        <h2>Articles</h2>
        <div class="postContainer">
          <postCard class="card" v-for="article in articleList" :key="article.title" v-bind:title="article.title" v-bind:author="article.author"></postCard>
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
          title: "mon incroyable article 1",
          author: "author 1"
        },
        {
          title: "mon incroyable article 2",
          author: "author 2"
        },
        {
          title: "mon incroyable article 3",
          author: "author 3"
        },
        {
          title: "mon incroyable article 4, le plus beau du monde",
          author: "author 4"
        },
        {
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