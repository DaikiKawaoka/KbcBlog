<template>
  <div id="app">
    <Header></Header>
    <Article-form :article="article" :errors="errors" :create="create" @submit="createArticle"></Article-form>
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
      // axios
      //   .post('/api/v1/items',this.item)
      //   .then(response => {
      //     let e = response.data;
      //     this.$router.push({ name: 'itemShow', params: { id: e.id } });
      //   })
      //   .catch(error => {
      //     console.error(error);
      //     if (error.response.data && error.response.data.errors) {
      //       this.errors = error.response.data.errors;
      //     }
      //   });
      console.log("create article")
    },
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
