<template>
    <div class="commentContainer">
        <button v-if="isAdmin" class="del" @click="this.delete">X</button>
        <button v-if="isAdmin" class="ban" @click="this.ban">Ban</button>
        <div class="commentHeader">
            <h4 class="user">{{user}}</h4>

        </div>
        <div class="thecomment">
            <p>
                {{content}}
            </p>
        </div>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'


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
            const url = 'http://localhost:8888/api/comment/delete'

            await this.axios.delete(url, {})
            .then(res => {
            if (res.status == 200){
                console.log("test")
            }
            })
            .catch(err => {
              if(err && err.response && err.response.status){
                if (err.response.status === 401){
                  this.error = "You don't have the rights to delete the comment";
                }
                else if (err.response.status === 500)
                  this.error = "Internal server error";
                else
                  this.error = err.message;
              }
            })
        },
        async ban () {

        }
    },
    computed: {
        ...mapGetters({
            isAdmin: 'isAdmin'
        })
    }
}
</script>

<style scoped>
    @import "../styles/comment.css";
</style>
