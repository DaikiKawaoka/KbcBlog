<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="false" :user="myUser"></Header>
    <div class="user-show-all">

      <div class="user-show-header">
        <div class="user-show-icon"></div>
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
        <div class="article-all-div">
          <div v-for="article in articles" :key="article.id" class="article-show-div">
            <div class="article-index-body">
              <router-link v-bind:to="{ name : 'ArticleShow', params : { id: article.id }}" class="a-tag">
                <h2 class="article-title-index"> {{article.title}} </h2>
              </router-link>
              <div class="article-index-username-updated">
                <h3 class="article-index-update">投稿日 {{ article.Updated }}</h3>
              </div>
            </div>
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
.user-show-header{
  display: flex;
  padding-bottom: 10px;
  /* border-bottom: 3px solid #fff; */
}
.user-show-icon{
  margin-top: 25px;
  background-color: #ccc;
  width: 150px;
  height: 150px;
  border-radius: 50%;
}
.user-show-info{
  margin-left: 100px;
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
/* .user-show-info-footer{
  
} */
.user-show-comment-p{
  font-size: 1em;
  width: 480px;
  height: 100%;
  margin-bottom: 0;
  word-wrap: break-word;
}

.user-show-body{
  display: flex;
  width: 800px;
  height: 210px;
  margin-right: auto;
  margin-left: auto;
}
.user-show-link-info{
  width: 180px;
  height: 210px;
  margin-right: 20px;
  word-wrap: break-word;
  /* background-color: #fff; */
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
</style>
