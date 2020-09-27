<template>
  <div id="app">
    <Header></Header>
    <div>
      <div v-for="article in articles" :key="article.id">
        <h2> {{article.title}} : {{article.body}} </h2>
      </div>
    </div>
  </div>
</template>

<script>
import Header from './../components/Header.vue'

export default {
  name: 'app',
  data(){
    return {
      articles: [],
      user: {},
    }
  },
  components: {
    Header
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/Articles')
      .then(response => {
        this.articles = response.data.Articles
        console.log(response.data.Articles)
      })
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
  width: 1000px;
  margin-left: auto;
  margin-right: auto;
}
h2{
  text-align: left;
}
</style>
