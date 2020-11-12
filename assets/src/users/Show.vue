<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="false" :user="myUser"></Header>
    <div class="user-show-all">

      <div class="user-show-header">
        <div class="user-show-icon">
          <!-- <img class="user-show-icon-img" src="../assets/IMG_6217.jpeg"> -->
          <div class="demo-image__preview">
          <el-image 
            style="width: 200px; height: 200px; border-radius: 50%;"
            :src="url" 
            :preview-src-list="srcList">
          </el-image>
          </div>
        </div>
        <div class="user-show-info">
          <div class="user-show-info-header">
            <span class="user-show-name">{{user.name}}</span>

            <div class="user-show-info-div" v-if="myUser.id===user.id">
              <div class="user-show-btn-div"><el-button type="info" size="mini" @click="editUser">プロフィール編集</el-button></div>


              <el-dropdown>
                <span class="el-dropdown-link icon-menu-span">
                  <i class="el-icon-setting"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <router-link class="a-tag2" v-bind:to="{ name : 'UserEdit', params : { id: user.id }}"><el-dropdown-item>パスワード変更</el-dropdown-item></router-link>
                  <el-popconfirm
                  title="本当にログアウトしますか?"
                  confirm-button-text="Yes"
                  cancel-button-text="No"
                  @confirm="logout"
                  >
                    <el-dropdown-item slot="reference">ログアウト</el-dropdown-item>
                  </el-popconfirm>
                </el-dropdown-menu>
              </el-dropdown>


            </div>

            <div class="user-show-info-div" v-else>
              <div class="user-show-btn-div">
                <el-button type="primary" size="medium" v-if="true">
                  <i class="el-icon-user"></i>フォロー
                </el-button>

                <el-button size="medium" v-else>
                  <i class="el-icon-user"></i>フォロー解除
                </el-button>
              </div>
            </div>

          </div>
          <div class="user-show-info-main">
            <div class="user-show-article-count-div">
              <!-- UserArticleCount -->
              <p class="edit-margin-p">{{user.postCount}}</p>
              <p class="edit-margin-p">投稿</p>
            </div>
            <div class="user-show-follower-count-div">
              <!-- Count -->
              <p class="edit-margin-p">0</p>
              <p class="edit-margin-p">フォロワー</p>
            </div>
            <div class="user-show-follow-count-div">
              <!-- Count -->
              <p class="edit-margin-p">0</p>
              <p class="edit-margin-p">フォロー</p>
            </div>
          </div>
          <div class="user-show-info-footer">
            <p class="user-show-comment-p">{{user.comment.String}}</p>
          </div>
        </div>
      </div>

      <div class="user-show-body">
        <div class="user-show-link-info">
          <div class="user-show-link-div">
            <p class="link-tag">github</p>
            <p class="link-p"><a href="https://github.com/DaikiKawaoka" class="user-show-link-a">https://github.com/DaikiKawaokagit</a></p>
          </div>
          <div class="user-show-link-div">
            <p class="link-tag">webサイト</p>
            <p class="link-p"><a href="https://github.com/DaikiKawaoka" class="user-show-link-a">https://github.com/DaikiKaw</a></p>
          </div>
        </div>
        <div class="user-show-body-info">
          <div class="user-show-body-indiv user-like-lang-div">
            <p class="user-show-body-div-in-p">好きな言語</p>
          </div>
          <div class="user-show-body-indiv user-post-article-div">
            <p class="user-show-body-div-in-p">投稿記事</p>
          </div>
          <div class="user-show-body-indiv user-iine-div">
            <p class="user-show-body-div-in-p">Good記事</p>
          </div>
        </div>
      </div>

      <div class="user-show-footer">
        <div class="user-show-post-menu">
          <el-tabs :tab-position="tabPosition" style="height: 200px;">
            <el-tab-pane label="記事">記事</el-tab-pane>
            <el-tab-pane label="質問">質問</el-tab-pane>
            <el-tab-pane label="いいね">いいね</el-tab-pane>
            <el-tab-pane label="質問">質問</el-tab-pane>
          </el-tabs>
        </div>
        <div class="post-all-div">
          <div v-for="article in articles" :key="article.id" class="post-show-div">
            <router-link v-bind:to="{ name : 'ArticleShow', params : { id: article.id }}" class="a-tag">
              <div class="post-index-body">
                  <h2 class="post-title-index"> {{article.title}} </h2>
                <div class="post-index-username-updated">
                  <h3 class="post-index-update">投稿日 {{ article.Updated }}</h3>
                </div>
              </div>
            </router-link>
          </div>
        </div>
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
      articles: Array,
      myUser: {},
      user: {
        KBCmail: "",
        id : 0,
        name: "",
        comment: {
          String: String,
          Valid: Boolean
        }
      },
      errors: {},
      url :"https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg",
      srcList :["https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg"],
      tabPosition: 'left'
    }
  },
  components: {
    Header,
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get(`http://localhost/api/restricted/Users/${this.$route.params.id}`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },})
      .then(response => {
        this.articles = response.data.Articles
        this.user = response.data.User
        this.myUser = response.data.MyUser
        console.log(response.data)
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
        }
        console.log(error.response);
      })
  },
  computed: {
  },
  methods: {
    editUser: function() {
      this.$router.push({ path: `/Users/${this.$route.params.id}/edit` });
    },
    deleteArticle: function() {
      this.$axios
        .delete(`http://localhost/api/restricted/Articles/${this.article.id}`,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          this.$router.push({ path: '/' });
          console.log(response)
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },
    logout: function() {
      this.$cookies.remove("JWT");
      this.$router.push({ path: "/login" });
    },
  }
}
</script>

<style>

.user-show-all{
  width: 1000px;
  margin: 0 auto 0 auto;
}
.user-show-header,.user-show-body,.user-show-footer{
  margin-left: 100px;
  margin-right: 100px;
}

/* header */

.user-show-header{
  display: flex;
  padding-bottom: 10px;
}
.user-show-icon{
  margin-top: 25px;
  border-radius: 50%;
}
.user-show-icon-img {
  width: 200px;
  height: 200px;
  border-radius: 50%;
}
.user-show-info{
  margin-left: 50px;
}
.user-show-info-header{
  display: flex;
  height: 100px;
  margin-bottom: 20px;
}
.user-show-name{
  margin: 30px 30px 30px 0;
  font-size: 1.6em;
  width: 300px;
  height: 60px;
}
.el-button--info{
  height: 30px;
}
.el-icon-setting{
  margin-left: 15px;
  margin-top: 10px;
  font-size: 30px;
  line-height: 100px;
  cursor: pointer;
  height: 30px;
}
.user-show-info-div{
  display: flex;
  height: 50px;
  margin-top: auto;
  margin-bottom: auto;
}
.user-show-btn-div{
  height: 30px;
  margin-top: 10px;
}
.user-show-info-main{
  display: flex;
}
.user-show-article-count-div,.user-show-follower-count-div,.user-show-follow-count-div{
  width: 100px;
  text-align: center;
}
.edit-margin-p{
  margin: 0;
  font-size: 0.9em;
}
.user-show-comment-p{
  font-size: 1em;
  width: 480px;
  height: 100%;
  margin-bottom: 0;
  word-wrap: break-word;
}

/* body */

.user-show-body{
  display: flex;
  width: 800px;
  height: 210px;
  border-bottom: #000 1px solid;
  padding: 0 0 30px 0;
  margin: 0 auto 30px auto;
}
.user-show-link-info{
  width: 180px;
  height: 210px;
  margin-right: 20px;
  word-wrap: break-word;
  border-radius: 2px;
}
.user-show-link-div{
  margin-top: 25px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
}
.link-tag{
  font-weight:bold;
  margin: 5px;
  font-size: 0.9em;
}
.link-p{
  margin: 5px;
  font-size: 0.8em;
}
.user-show-link-a{
  color: #333;
}
.user-show-link-a:hover{
  color: #409eff;
}
.user-show-body-info{
  width: 550px;
  background-color: rgb(43, 43, 43);
  border-radius: 2px;
}
.user-show-body-indiv{
  height: 70px;
}
.user-show-body-div-in-p{
  color: rgb(197, 231, 0);
  line-height: 70px;
  margin: 0 0 0 10px;
  width: 100px;
}

/* footer */
.user-show-footer{
  display: flex;
}

.user-show-post-menu{
  width: 270px;
  margin-left: 50px;
}


.post-all-div{
  width: 550px;
  margin: 30px 50px 0 0;
}
.post-show-div{
  margin-bottom: 5px;
  padding-left: 20px;
  border: solid 1px #ccc;
  display: flex;
  background: #eee;
}
.post-show-div:hover{
  transition: 0.3s ;
  background: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
}
.post-user-icon{
  margin-top: 25px;
  background-color: #ccc;
  width: 40px;
  height: 40px;
  border-radius: 50%;
}
.post-index-body{
  margin-left: 20px;
}
.post-index-username-updated{
  display: flex;
}
.post-title-index{
  font-size: 20px;
  width: 500px;
  text-align: left;
  margin-bottom: 0;
}
.post-index-username,.post-index-update{
  font-size: 13px;
  color: #999;
}
.post-index-update{
  margin-left: 20px;
}
.a-tag{
  color: #000;
  text-decoration: none;
  cursor: pointer;
}
</style>
