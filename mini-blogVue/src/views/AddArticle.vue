<template>
    <div class="articleContainer">
        <h1>Add Article</h1>
        <form class="articleForm" @submit.prevent="add">
            <div class="head">
                <input v-bind="title" type="text" class="title" placeholder="Your title">
                <input v-bind="link" type="text" class="title" placeholder="Link to your image">
            </div>
            <div class="content">
                <textarea v-bind="content" name="" class="body" placeholder="Your content"></textarea>
                
            </div>
            <button type="submit">Add Article</button>
        </form>
    </div>
</template>

<script>
export default {
    name: 'addArticle',
    data () {
        return {
            id: 0,
            content: "mon article",
            link: null,
            title: "titre de mon article",
            createDate: null,
            author: null,
            owned: null,
            uid: null,
        }
    },
    methods: {
        
        async add () {
            var date = new Date()
            this.createDate = date.getDate()+"/"+date.getMonth()+"/"+date.getFullYear()
            
            const url = 'http://localhost:8888/api/article/fetch'

            await this.axios.post(url, {})
            .then(res => {
                if (res.status == 200){
                    console.log("test")
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

<style>
    @import "../styles/addArticle.css";
</style>