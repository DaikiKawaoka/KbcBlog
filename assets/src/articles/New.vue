<template>
  <div id="app">
    <Header :isArticle="true" :isQuestion="false" :user="user"></Header>
    <Article-form :article="article" :errors="errors" :create="create" @submit="createArticle" @cancell="goHome"></Article-form>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import Footer from './../components/Footer.vue';
import ArticleForm from './../components/ArticleForm.vue'

export default {
  name: 'app',
  data(){
    return {
      user: {
        id: 0,
        name: "",
        KBC_mail: "",
      },
      article :{
        userid: 0,
        title: "",
        body: "",
        tag: "",
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
        this.user.id = response.data.user.id
        this.user.KBC_mail = response.data.user.KBC_mail
        this.user.name = response.data.user.name
        this.article.userid = this.user.id
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
  },
  components: {
    Header,
    Footer,
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
        .then(() => {
          this.$router.push({ path: "/" });
          this.createArticleAlert();
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    createArticleAlert() {
      this.$notify({
        title: 'Success',
        message: '記事を投稿しました。',
        type: 'success'
      });
    },

    goHome: function(){
      this.$router.push({ path: "/" });
    },

    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },
  }
}
</script>

<style>

</style>
