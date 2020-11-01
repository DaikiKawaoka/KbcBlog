<template>
  <div id="app">
    <Header></Header>
    <div class="index-main">
      <div class="index-menu">
        <div></div>
      </div>
      <div class="article-all-div">
        <div v-for="article in articles" :key="article.id" class="article-show-div">
          <router-link v-bind:to="{ name : 'ArticleShow', params : { id: article.id }}">
            <h2> {{article.title}}</h2>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import axios from "axios";
import Header from './../components/Header.vue'
// import VueCookie from 'vue-cookie';

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
    this.$axios.get('http://localhost/api/restricted/Articles',{
      headers: {
        Authorization: `Bearer ${this.$cookie.get("JWT")}`
      },
    })
      .then(response => {
        this.articles = response.data.Articles
        console.log(response.data)
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
        }
        console.log(error.response);
      })
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-left: auto;
  margin-right: auto;
  padding-bottom: 30px;
  background-color: #F6F6F4;
}
h2{
  text-align: left;
}
.index-main{
  display: flex;
  margin-top: 30px;
}
.index-menu{
  width: 200px;
  background-color: #eee;
}
.article-all-div{
  width: 800px;
  background-color: #fff;
}
.article-show-div{
  margin-left: 20px;
  padding-left: 20px;
  border: solid 1px #eee;
}
</style>
