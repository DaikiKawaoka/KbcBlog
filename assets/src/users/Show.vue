<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="false"></Header>
    <div class="user-show-all">

      <div class="user-show-header">
        <div class="user-show-icon"></div>
        <div class="user-show-info">
          <div class="user-show-info-header">
            <span class="user-show-name">{{user.name}}</span>
            <div class="user-show-info-div">
              <div class="user-show-btn-div"><el-button type="info" size="mini" @click="editUser">プロフィール編集</el-button></div>
              <i class="el-icon-setting"></i>
            </div>
          </div>
          <div class="user-show-info-main">
            <div class="user-show-article-count-div">
              <!-- UserArticleCount -->
              <p class="edit-margin-p">10000</p>
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
            <p class="user-show-comment-p">aaaaaaaaaa</p>
          </div>
        </div>
      </div>

      <div class="user-show-body">

      </div>
      <div class="user-show-footer">

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
      articles: null,
      myUser: Object,
      user: Object,
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
        this.myUser = response.data.MyUser
        this.user = response.data.User
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
  padding-bottom: 30px;
  border-bottom: 3px solid #fff;
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
.user-show-info-footer{
  
}
.user-show-comment-p{
  font-size: 1em;
  width: 480px;
  height: 100%;
  margin-bottom: 0;
  word-wrap: break-word;
}
</style>
