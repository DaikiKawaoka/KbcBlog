<template>
  <div id="app">
    <Header :isArticle="true" :isQuestion="false"></Header>
    <Article-form :article="article" :user="user" :errors="errors" :create="create" @submit="updateArticle" @cancell="goHome"></Article-form>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import ArticleForm from './../components/ArticleForm.vue'

export default {
  name: 'app',
  data(){
    return {
      user: null,
      article : null,
      errors: [],
      create: false,
    }
  },
  components: {
    Header,
    ArticleForm
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get(`http://localhost/api/restricted/Articles/${this.$route.params.id}/edit`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.user = response.data.user
        this.article = response.data.Article
        if(this.user.id != this.article.userid){
          this.$router.push({ path: "/" });
        }
        console.log(response)
      })
      .catch(error => {
        console.log(error)
        if(error.response.status == 401){
            this.$router.push({ path: "/login" });
          }
        this.errors = error.response.data.ValidationErrors;
      })
  },
  methods: {
    updateArticle: function() {
      this.$axios
        .patch(`http://localhost/api/restricted/Articles/${this.$route.params.id}`, this.article,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          this.$router.push({ path: `/Articles/${this.article.id}` });
          console.log(response)
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },
    goHome: function(){
      this.$router.push({ path: "/" });
    }
  }
}
</script>

<style>

</style>
