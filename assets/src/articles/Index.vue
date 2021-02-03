<template>
  <div id="app" v-if="user">
    <Header :isArticle="true" :isQuestion="false" :user="user" @reset="reset" :url="url"></Header>
    <div class="index-main body-main">
      <Tag @scopetag="scopetag" :tag="tag" :friendsOnly="friendsOnly" @reset="reset" @update:friendsOnly ="friendsOnly=$event" :url="url"></Tag>
      <div class="article-all-div">
        <div class="article-search-div">
          <el-input placeholder="キーワード検索" v-model="searchText" suffix-icon="el-icon-search" style="width:200px; margin-left: 30px;" @input="scopetag"></el-input>
          <div class="order-div">
            <span class="order-span-left">並び替え</span>
            <span class="order-span" v-bind:class="{'order-span-color': order === 'new' }" @click="articleOrder('new')">新着</span>
            <span class="order-span" v-bind:class="{ 'order-span-color': order === 'like' }" @click="articleOrder('like')">人気</span>
          </div>
        </div>

        <div v-if="articles.length !== 0">
          <div v-for="article in articles" :key="article.id" class="article-show-div">
            <div class="article-user-icon-div">
              <router-link v-bind:to="{ name : 'UserShow', params : { id: article.userid }}" class="a-tag">
                <!-- image -->
                <img class="article-user-icon" :src="article.imgpath">
              </router-link>
            </div>
            <div class="article-index-body">
              <div class="article-index-tag-div">
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
                <h3 class="article-index-update" v-if="order === 'new' "> {{ article.Created | moment }}</h3>
                <h3 class="article-index-update" v-else> {{ article.Created | moment2 }}</h3>
                <i class="el-icon-star-on article-star-i"></i>
                <span class="article-likecount-span">{{article.likecount}}</span>
              </div>
            </div>
          </div>


          <infinite-loading @infinite="scrollArticles" :identifier="infiniteId" spinner="spiral">
            <span slot="no-more">----- 検索結果は以上です-----</span>
            <span slot="no-results">----- 検索結果は以上です -----</span>
          </infinite-loading>
        </div>

        <div v-else>
          <div class="article-show-div not-tag-div" v-if="friendsOnly === false ">
            <p class="not-tag" v-if="searchText === '' && tag !== '全て' ">タグ『{{ tag }}』の記事はまだありません。</p>
            <p class="not-tag" v-else-if="tag === '全て' && searchText !== '' ">キーワード『{{ searchText }}』の記事はありません。</p>
            <p class="not-tag"  v-else-if="tag !== '全て' && searchText !== '' ">タグ『{{ tag }}』,キーワード『{{ searchText }}』の記事はありません。</p>
            <p class="not-tag" v-else>まだ記事が投稿されていません。</p>
          </div>
          <div class="article-show-div not-tag-div" v-else>
            <p class="not-tag" v-if="searchText === '' && tag !== '全て' ">タグ『{{ tag }}』の記事はまだありません。</p>
            <p class="not-tag" v-else-if="tag === '全て' && searchText !== '' ">キーワード『{{ searchText }}』の記事はありません。</p>
            <p class="not-tag"  v-else-if="tag !== '全て' && searchText !== ''">タグ『{{ tag }}』,キーワード『{{ searchText }}』の記事はありません。</p>
            <p class="not-tag" v-else >フレンドの記事はまだありません。</p>
          </div>
        </div>

      </div>
      <div>
        <Ranking :Ranking="Ranking" :rankingType="rankingType" :isArticle="true" :user="user" :url="url"></Ranking>
      </div>
    </div>
    <Footer :user="user" :url="url"></Footer>
  </div>
</template>

<script>
// import axios from "axios";
import InfiniteLoading from 'vue-infinite-loading'
import Header from './../components/Header.vue'
import Footer from './../components/Footer.vue'
import Tag from './../components/Tag.vue'
import Ranking from './../components/Ranking.vue'
import moment from 'moment'

export default {
  name: 'app',
  data(){
    return {
      user: null,
      articles: [],
      Ranking: [],
      searchText:"",
      tag: String,
      order: String,
      friendsOnly: Boolean,
      cursor: 0,
      page: 1,
      infiniteId: +new Date(),
      notificationCount: localStorage.notificationCount,
      rankingType: Number,
      loading: null,
      url: null,
    }
  },
  components: {
    Header,
    Footer,
    Tag,
    Ranking,
    InfiniteLoading,
  },
  filters: {
    moment: function (date) {
      return moment(date).fromNow();
    },
    moment2: function (date) {
      // locale関数を使って言語を設定すると、日本語で出力される
      return moment(date).utc().format('YYYY/MM/DD');
    }
  },
  watch: {
    immediate: true,
    searchText(newSearchText) {
      localStorage.searchText = newSearchText;
    },
    order(newOrder) {
      localStorage.order = newOrder;
    },
    tag(newTag) {
      localStorage.tag = newTag;
    },
    friendsOnly(newFriendsOnly) {
      localStorage.friendsOnly = newFriendsOnly;
    },
    notificationCount(newNotificationCount) {
      localStorage.notificationCount = newNotificationCount;
    },
  },

  created () {
    document.title = `KBC Blog`;
    this.url = process.env.VUE_APP_URL
    this.openFullScreen()
    const jst = this.$cookies.get("JWT");
    this.setUp()
    this.$axios.get(this.url+'api/restricted/Articles',{
      params: {
        // ここにクエリパラメータを指定する
        friendsOnly: this.friendsOnly,
        searchText: this.searchText,
        order: this.order,
        tag: this.tag,
        cursor: this.cursor,
        rankingType: this.rankingType,
      },
      headers: {
        Authorization: `Bearer ${jst}`
      },
    })
      .then(response => {
        this.articles = response.data.Articles
        this.Ranking = response.data.Ranking
        this.user = response.data.user
        this.cursor = response.data.Cursor
        this.notificationCount = response.data.NotificationCount
        this.closeFullScreen()
      })
      .catch(error => {
        if(error.response.status == 401){
          if(jst !== null){
            this.errorNotify();
          }
          this.$router.push({ path: "/login" });
        }
        this.closeFullScreen()
      })
  },
  methods:{
    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },

    openFullScreen() {
      this.loading = this.$loading({
          lock: true,
          text: 'Loading',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.98)'
        });
    },
    closeFullScreen() {
      this.loading.close();
    },

    articleOrder(c) {
      this.changeType()
      this.$axios.get(this.url+'api/restricted/Articles/scope', {
      params: {
        friendsOnly: this.friendsOnly,
        searchText: this.searchText,
        order: c,
        tag: this.tag,
        cursor: this.cursor,
      },
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        if(c === "new"){
          this.order = "new"
        }else{
          this.order = "like"
        }
        this.articles = response.data.Articles
        this.cursor = response.data.Cursor
        console.log("並べ替え")
        console.log(this.cursor)
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
    },

    scopetag() {
      this.changeType()
      this.$axios.get(this.url+'api/restricted/Articles/scope', {
      params: {
        friendsOnly: this.friendsOnly,
        searchText: this.searchText,
        order: this.order,
        tag: this.tag,
        cursor: this.cursor,
      },
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.articles = response.data.Articles
        this.cursor = response.data.Cursor
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
    },

    scrollArticles($state) {
      this.$axios.get(this.url+'api/restricted/Articles/scope', {
      params: {
        friendsOnly: this.friendsOnly,
        searchText: this.searchText,
        order: this.order,
        tag: this.tag,
        cursor: this.cursor,
      },
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        if (response.data.Articles.length) {
          this.articles = this.articles.concat(response.data.Articles);
          this.cursor = response.data.Cursor
          this.page += 1;
          $state.loaded();
        } else {
          $state.complete();
        }
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
    },
    changeType() {
      this.page = 1;
      this.infiniteId += 1;
      this.articles = [];
      this.cursor = 0;
    },

    setUp() {
      if (localStorage.searchText) {
        this.searchText = localStorage.searchText;
      }
      if (localStorage.tag) {
        this.tag = localStorage.tag;
      }else{
        this.tag = '全て'
      }
      if (localStorage.order) {
        this.order = localStorage.order;
      }else{
        this.order = 'new'
      }
      if (localStorage.friendsOnly) {
        if(localStorage.friendsOnly === "true"){
          this.friendsOnly = true
        }else{
          this.friendsOnly = false
        }
      }else{
        this.friendsOnly = false
      }
      this.rankingType = Math.floor(Math.random() * Math.floor(2));
    },
    reset(){
      this.searchText = ''
      this.tag = '全て'
      this.order = 'new'
      this.friendsOnly = false
      this.scopetag()
    }
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
li{
    list-style: none;
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
  width: 460px;
  /* width: 500px; */
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
  transition: 0.10s ;
  border: solid 1.3px #409eff;
  /* box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04); */
  /* transform: translateY(-0.1875em); */
}
.article-user-icon-div{
  margin-top: 25px;
}
.article-user-icon{
  width: 30px;
  height: 30px;
  object-fit: cover; /* 画像トリミング */
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
  width: 400px;
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
  /* margin-left: 20px; */
  padding-top: 10px;
  position: sticky;
  top: 0;
  /* z-index: 1000px; */
}
.order-div{
  padding-top: 20px;
}
.order-span{
  margin-left: 10px;
  cursor: pointer;
}
.order-span-left{
  margin-left: 80px;
  color: #777;
}
.order-span:hover{
  opacity: 0.7;
}
.order-span-color{
  color: #2ee002;
  font-weight: bold;
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
.article-index-tag-div{
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
  width: 400px;
}
.not-tag{
  width: 100%;
  text-align: center;
  color: #555;
  font-size: 0.9em;
}
.not-tag-div{
  margin-top: 30px;
  padding-top: 15px;
  padding-bottom: 15px;
}
</style>
