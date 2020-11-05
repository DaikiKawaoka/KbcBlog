<template>
  <div id="app">
    <Header :isArticle="true" :isQuestion="false"></Header>
    <div class="index-main body-main">
      <div class="index-menu">
        <div></div>
      </div>
      <div class="article-all-div">
        <div v-for="article in articles" :key="article.id" class="article-show-div">
          <div class="article-user-icon"></div>
          <div class="article-index-body">
            <router-link v-bind:to="{ name : 'ArticleShow', params : { id: article.id }}" class="a-tag">
              <h2 class="article-title-index"> {{article.title}} </h2>
            </router-link>
            <div class="article-index-username-updated">
              <h3 class="article-index-username">{{ article.name }}</h3>
              <h3 class="article-index-update">投稿日 {{ article.Updated }}</h3>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import axios from "axios";
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
    this.$axios.get('http://localhost/api/restricted/Articles',{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.articles = response.data.Articles
        this.user = response.data.user
        console.log(response.data)
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
        }
        console.log(error.response);
      })
  },
}
</script>

<style>
body {
  margin: 0px;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin: 0 auto 0 auto;
  padding-bottom: 30px;
  background-color: #F6F6F4;
}
.body-main{
  width: 1000px;
  margin: 0 auto 0 auto;
}
.index-main{
  display: flex;
  margin-top: 30px;
}
.index-menu{
  width: 300px;
  background-color: #F6F6F4;
}
.article-all-div{
  width: 600px;
  background-color: #fff;
}
.article-show-div{
  padding-left: 20px;
  border: solid 1px #eee;
  display: flex;
}
.article-user-icon{
  margin-top: 25px;
  background-color: #ccc;
  width: 40px;
  height: 40px;
  border-radius: 50%;
}
.article-index-body{
  margin-left: 20px;
}
.article-index-username-updated{
  display: flex;
}
.article-title-index{
  font-size: 20px;
  width: 500px;
  text-align: left;
  margin-bottom: 0;
}
.article-index-username,.article-index-update{
  font-size: 13px;
}
.article-index-update{
  margin-left: 20px;
  color: #999;
}
.a-tag{
  color: #000;
  text-decoration: none;
  cursor: pointer;
}
</style>
