<template>
  <div id="app">
    <Header :isArticle="true" :isQuestion="false" :user="user"></Header>
    <div class="index-main body-main">
      <Tag @scopetag="scopetag" :tag="tag"></Tag>
      <div class="article-all-div">
        <div class="article-search-div">
          <el-input placeholder="キーワード検索" v-model="searchText" suffix-icon="el-icon-search" style="width:200px; margin-left: 30px;"></el-input>
          <div class="order-div">
            <span class="order-span-left">並び替え</span>
            <span class="order-span" v-bind:class="{'order-span-color': orderNew }" @click="articleOrder('new')">新着</span>
            <span class="order-span" v-bind:class="{ 'order-span-color': orderLike }" @click="articleOrder('like')">人気</span>
          </div>
        </div>

        <div v-for="article in articles" :key="article.id" class="article-show-div">
          <div class="article-user-icon">
            <router-link v-bind:to="{ name : 'UserShow', params : { id: article.userid }}" class="a-tag">
              <!-- image -->
            </router-link>
          </div>
          <div class="article-index-body">
            <div class="article-tag-div">
              <i class="el-icon-collection-tag tag-icon"></i>
              <span class="article-tag no-magin">{{ article.tag }}</span>
            </div>
            <router-link v-bind:to="{ name : 'ArticleShow', params : { id: article.id }}" class="a-tag">
              <h2 class="article-title-index"> {{article.title}} </h2>
            </router-link>
            <div class="article-index-username-updated">
              <router-link v-bind:to="{ name : 'UserShow', params : { id: article.userid }}" class="a-tag">
                <h3 class="article-index-username">by {{ article.name }}</h3>
              </router-link>
              <h3 class="article-index-update" v-if="orderNew"> {{ article.Updated | moment }}</h3>
              <h3 class="article-index-update" v-else> {{ article.Updated | moment2 }}</h3>
              <i class="el-icon-star-on article-star-i"></i>
              <span class="article-likecount-span">{{article.likecount}}</span>
            </div>
          </div>
        </div>

      </div>
      <div>
        <Ranking :likeRanking="likeRanking" :postRanking="postRanking"></Ranking>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
// import axios from "axios";
import Header from './../components/Header.vue'
import Footer from './../components/Footer.vue'
import Tag from './../components/Tag.vue'
import Ranking from './../components/Ranking.vue'
import moment from 'moment'

export default {
  name: 'app',
  data(){
    return {
      user: {},
      articles: [],
      likeRanking: [],
      postRanking: [],
      searchText:"",
      tag: "",
      order: "new",
      orderNew: true,
      orderLike: false,
    }
  },
  components: {
    Header,
    Footer,
    Tag,
    Ranking,
  },
  filters: {
    moment: function (date) {
      // locale関数を使って言語を設定すると、日本語で出力される
      // moment.locale( 'ja' );
      return moment(date).fromNow();
      //return moment(date).utc().format('YYYY/MM/DD');
    },
    moment2: function (date) {
      // locale関数を使って言語を設定すると、日本語で出力される
      return moment(date).utc().format('YYYY/MM/DD');
    }
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    const jst = this.$cookies.get("JWT");
    if(jst === null){
      this.$router.push({ path: "/login" });
    }
    this.$axios.get('http://localhost/api/restricted/Articles',{
      headers: {
        Authorization: `Bearer ${jst}`
      },
    })
      .then(response => {
        this.articles = response.data.Articles
        this.likeRanking = response.data.LikeRanking
        this.postRanking = response.data.PostRanking
        this.user = response.data.user
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
  },
  methods:{
    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },

    articleOrder(c) {
      this.$axios.get('http://localhost/api/restricted/Articles/scope', {
      params: {
        // ここにクエリパラメータを指定する
        order: c,
        tag: this.tag
      },
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        if(c === "new"){
          this.order = "new"
          this.orderNew = true;
          this.orderLike = false;
        }else{
          this.order = "like"
          this.orderNew = false;
          this.orderLike = true;
        }
        this.articles = response.data.Articles
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
    },

    scopetag() {
      this.$axios.get('http://localhost/api/restricted/Articles/scope', {
      params: {
        order: this.order,
        tag: this.tag
      },
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.articles = response.data.Articles
        console.log("スコープしました"+this.tag+this.order)
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
    },
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
  background-color: #F6F6F4;
  /* background-color: #15202b; */
}
.body-main{
  width: 1000px;
  margin: 0 auto 0 auto;
}
.index-main{
  display: flex;
  margin-top: 30px;
}
.article-all-div{
  width: 500px;
  background-color: #F6F6F4;
  /* background-color: #15202b; */
}
.article-show-div{
  background-color: #FFF;
  padding-left: 10px;
  border: solid 1px #eee;
  border-radius: 6px;
  margin-bottom: 5px;
  display: flex;
}
.article-show-div:hover {
  transition: 0.15s ;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
  transform: translateY(-0.1875em);
}
.article-user-icon{
  margin-top: 25px;
  background-color: #ccc;
  width: 30px;
  height: 30px;
  border-radius: 50%;
}
.article-index-body{
  margin-left: 10px;
}
.article-index-username-updated,.article-search-div{
  display: flex;
}
.article-title-index{
  font-size: 1.1em;
  width: 440px;
  text-align: left;
  margin: 0;
}
.article-index-username,.article-index-update{
  font-size: 13px;
  color: #999;
}
.article-index-update{
  margin-left: 20px;
}
.a-tag{
  color: #000;
  text-decoration: none;
  cursor: pointer;
}
.article-search-div{
  background-color: #F6F6F4;
  /* background-color: #15202b; */
  padding-bottom: 8px;
}
.order-div{
  padding-top: 20px;
}
.order-span{
  margin-left: 10px;
  cursor: pointer;
}
.order-span-left{
  margin-left: 100px;
  color: #777;
}
.order-span:hover{
  opacity: 0.7;
}
.order-span-color{
  color: #2ee002;
}
.article-star-i{
  color: orange;
  margin-top: 13px;
  margin-left: 20px;
  font-size: 1.3em;
}
.article-likecount-span{
  margin-top: 15px;
  font-size: 0.8em;
  color: #777;
}
.no-magin{
  margin: 0;
}
.article-tag-div{
  display: flex;
}
.tag-icon{
  color: blue;
  margin-top: 2px;
  font-size: 0.6em;
}
.article-tag{
  font-size: 0.6em;
  color: #999;
  width: 440px;
}
</style>
