<template>
  <div id="app">
    <Header></Header>
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
      article :{
        userid: 1,
        title: "",
        body: "",
      },
      errors: [],
      create: true,
    }
  },
  components: {
    Header,
    ArticleForm
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/Articles/new')
      .then(response => {
        console.log(response)
      })
  },
  methods: {
    createArticle: function() {
      this.$axios
        .post('http://localhost/api/Articles',this.article)
        .then(response => {
          // let e = response.data;
          // this.$router.push({ name: 'itemShow', params: { id: e.id } });
          this.$router.push({ path: "/" });
          console.log(response)
        })
        .catch(error => {
          console.log(error.response.data.ValidationErrors);
          this.errors = error.response.data.ValidationErrors;

          // if (error.response.data && error.response.data.errors) {
          //   this.errors = error.response.data.errors;
          // }
        });
    },
    goHome: function(){
      this.$router.push({ path: "/" });
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
