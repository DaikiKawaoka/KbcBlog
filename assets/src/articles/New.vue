<template>
  <div id="app">
    <Header :isArticle="true" :isQuestion="false"></Header>
    <Article-form :article="article" :errors="errors" :create="create" @submit="createArticle" @cancell="goHome"></Article-form>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import ArticleForm from './../components/ArticleForm.vue'

export default {
  name: 'app',
  data(){
    return {
      user: {
        id: 0,
        name: "",
        KBC_mail: "",
        mailname: "",
      },
      article :{
        userid: 0,
        title: "",
        body: "",
      },
      errors: [],
      create: true,
    }
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/restricted/Articles/new',{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        console.log(response.data)
        this.user.id = response.data.user.id
        this.user.KBC_mail = response.data.user.KBC_mail
        this.user.name = response.data.user.name
        this.user.mailname = response.data.user.mailname
        this.article.userid = this.user.id
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
        }
        console.log(error.response);
      })
  },
  components: {
    Header,
    ArticleForm
  },
  methods: {
    createArticle: function() {
      this.$axios
        .post('http://localhost/api/restricted/Articles', this.article,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          this.$router.push({ path: "/" });
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
