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
  methods: {
    createArticle: function() {
      this.$axios
        .post('http://localhost/api/restricted/Articles', this.article,{
          headers: {
            Authorization: `Bearer ${this.$cookie.get("JWT")}`
          },
        })
        .then(response => {
          this.$router.push({ path: "/" });
          console.log(response)
        })
        .catch(error => {
          console.log(error.response.data.ValidationErrors);
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
