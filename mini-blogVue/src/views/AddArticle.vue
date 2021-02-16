<template>
    <div class="articleContainer">
        <h1>Add Article</h1>
        <form class="articleForm" @submit.prevent="add">
            <div class="head">
                <input v-model="title" type="text" class="title" placeholder="Your title">
                <input v-model="link" type="text" class="title" placeholder="Article top image link (optional)">
            </div>
            <div class="content">
                <textarea v-model="content" name="" class="body" placeholder="Your content"></textarea>
            </div>
            <button type="submit">Add Article</button>
            <div v-if="error" class="error">
                {{error}}
            </div>
        </form>
    </div>
</template>

<script>
export default {
    name: 'addArticle',
    data () {
        return {
            id: 0,
            content: "Article content",
            link: "",
            title: "Article title",
            createDate: null,
            author: null,
            owned: null,
            uid: null,
            error: null,
        }
    },
    methods: {

        async add () {
            var date = new Date()
            this.createDate = date.getDate()+"/"+date.getMonth()+"/"+date.getFullYear()

            const url = 'http://localhost:8888/api/article/create?title=' + this.title + '&content=' + this.content + '&img_link=' + this.link;

            console.log(this.title);
            console.log(this.link);
            console.log(this.content);

            console.log(url);
            console.log(this.axios);

            await this.axios.post(url, {}, {headers: {Authorization: this.$store.state.token}})
            .then(res => {
                if (res.status == 200){
                    console.log("test")
                }
            })
            .catch(err => {
                console.log(this.$store.state);
                if(err && err.response && err.response.status){
                if (err.response.status === 404){
                    this.error = "The requested account doesn't exist";
                }
                else if (err.response.status === 500)
                    this.error = "Internal server error";
                else if (err.response.status === 401)
                    this.error = "Only admins can post articles";
                else
                    this.error = err.message;
                }
            })
        }
    }
}
</script>

<style>
    @import "../styles/addArticle.css";
</style>
