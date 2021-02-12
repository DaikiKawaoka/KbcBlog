<template>
  <div id="app" v-if="article">
    <Header :isArticle="true" :isQuestion="false" :user="user" :url="url"></Header>
    <div class="body-main">
      <div class="article-show-main">
        <div v-if="this.article != null">
          <div class="comment-header">
            <router-link v-bind:to="{ name : 'UserShow', params : { id: article.userid }}" class="a-tag">
              <div class="comment-header-div">
                <img class="comment-user-icon" :src="article.imgpath">
                <p class="comment-user-name">{{ article.name }} <i v-if="teacherMatch(article.KBC_mail)" class="el-icon-success teacher" title="先生マーク"></i></p>
              </div>
            </router-link>
            <div class="comment-header-div">
              <div class="article-create-update-date-div">
                <p class="article-create-date"> 作成日 {{ article.Created | moment }}</p>
                <p class="article-create-date"> 更新日 {{ article.Updated | moment }}</p>
              </div>
              <span class="dropdown-span">
                <el-dropdown>
                  <span class="el-dropdown-link icon-menu-span">
                    <i class="el-icon-more"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <router-link v-if="article.userid === user.id" class="a-tag2" v-bind:to="{ name : 'ArticleEdit', params : { id: article.id }}"><el-dropdown-item>編集</el-dropdown-item></router-link>
                    <el-popconfirm
                    v-if="article.userid === user.id"
                    title="本当に削除しますか?"
                    confirm-button-text="Yes"
                    cancel-button-text="No"
                    @confirm="deleteArticle"
                    >
                      <el-dropdown-item slot="reference">削除</el-dropdown-item>
                    </el-popconfirm>
                    <router-link class="a-tag2" target="_blank" v-bind:to="{ name : 'ArticleMarkdown', params : { id: article.id }}">
                      <el-dropdown-item>ソースを見る</el-dropdown-item>
                    </router-link>
                  </el-dropdown-menu>
                </el-dropdown>
              </span>
            </div>
          </div>
          <h1 class="article-title">{{article.title}}</h1>
          <div class="article-tags-div">
            <div v-for="(tag,index) in tags" :key="index" class="article-tag-div">
              <span>{{tag}}</span>
            </div>
          </div>
          <div class="article-form__preview-body">
            <div class="article-body article-form__preview-body-contents" v-html="compiledMarkdown"></div>
          </div>

          <div class="article-like-btn-div">
            <el-row v-if="user.id === 920437694">
              <button @click="loginDialog" class="like-btn">いいね <i class="el-icon-star-off"></i></button>
              <span class="like-count-span">{{like.likeCount}}</span>
            </el-row>

            <el-row v-else>
              <el-button v-if="like.isLike" @click="DeleteLikeConfirmationOpen" type="warning">解除 <i class="el-icon-star-on"></i></el-button>
              <button v-else @click="CreateArticleLike" class="like-btn">いいね <i class="el-icon-star-off"></i></button>
              <span class="like-count-span">{{like.likeCount}}</span>
            </el-row>
          </div>

        </div>
      </div>
      <div class="article-comment-all">
        <div class="article-comment-header">
          <i class="el-icon-chat-dot-round comment-icon"></i>
          <p class="comment">コメント</p>
        </div>
        <div v-for="(c,index) in comments" :key="c.id" class="article-comment-main">
          <div class="comment-div">
            <div class="comment-header">
              <router-link v-bind:to="{ name : 'UserShow', params : { id: c.userid }}" class="a-tag">
                <div class="comment-header-div">
                    <img class="comment-user-icon" :src="c.imgpath">
                    <p class="comment-user-name">{{c.name}} <i v-if="teacherMatch(c.KBC_mail)" class="el-icon-success teacher" title="先生マーク"></i></p>
                </div>
              </router-link>
              <div class="comment-header-div">
                <div>
                  <p class="comment-create-date">投稿日 {{c.Created | moment}}</p>
                </div>
                <span v-if="c.userid === user.id" class="dropdown-span">
                  <el-dropdown>
                    <span class="el-dropdown-link icon-menu-span">
                      <i class="el-icon-more comment-edit-icon"></i>
                    </span>
                    <el-dropdown-menu slot="dropdown">
                      <el-popconfirm
                      title="本当に削除しますか?"
                      confirm-button-text="Yes"
                      cancel-button-text="No"
                      @confirm="deleteArticleComment(c.id,index)"
                      >
                        <el-dropdown-item slot="reference">削除</el-dropdown-item>
                      </el-popconfirm>
                    </el-dropdown-menu>
                  </el-dropdown>
                </span>
              </div>
            </div>
            <div class="comment-main">
              <div class="article-body comment-text" v-html="compiledMarkdownComment(c.text)"></div>
            </div>

            <el-row v-if="user.id === 920437694">
              <!-- <i class="el-icon-star-on"></i> -->
              <i class="el-icon-star-off not-comment-ster-i" @click="loginDialog"></i>
              <span class="comment-like-count-span">{{c.Like.likeCount}}</span>
            </el-row>

            <el-row v-else>
              <!-- <i class="el-icon-star-on"></i> -->
              <i v-if="c.Like.isLike" class="el-icon-star-on comment-ster-i" @click="DeleteArticleCommentLike(index)"></i>
              <i v-else class="el-icon-star-off not-comment-ster-i" @click="CreateArticleCommentLike(index)"></i>
              <span class="comment-like-count-span">{{c.Like.likeCount}}</span>
            </el-row>
          </div>
        </div>

        <CommentForm :comment="comment" :errors="errors" :user="user" @submit="createArticleComment"></CommentForm>
      </div>
    </div>
    <Footer :user="user" :url="url"></Footer>
  </div>
</template>

<script>
import moment from 'moment'
import Header from './../components/Header.vue'
import Footer from './../components/Footer.vue';
import CommentForm from './../components/CommentForm.vue'
import marked from 'marked';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';
import '../css/monokai-sublime.css';
import '../css/markdown.css';

export default {
  name: 'app',
  data(){
    return {
      article: null,
      like: {
        isLike: Boolean,
        likeCount: 0,
      },
      user: {},
      comment:{
        userid: 0,
        name: "",
        articleid: 0,
        text: "",
      },
      comments: [],
      errors: {},
      notificationCount: localStorage.notificationCount,
      tags:[],
      url: null,
    }
  },
  components: {
    Header,
    Footer,
    CommentForm,
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.openFullScreen()
    marked.setOptions({
      // code要素にdefaultで付くlangage-を削除
      langPrefix: '',
      // highlightjsを使用したハイライト処理を追加
      highlight: function(code, lang) {
        return hljs.highlightAuto(code, [lang]).value
      }
    });
    this.url = process.env.VUE_APP_URL
    this.$axios.get(this.url+`api/restricted/Articles/${this.$route.params.id}`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },})
      .then(response => {
        this.article = response.data.Article
        this.user = response.data.user
        this.comments = response.data.Comments
        this.like = response.data.Like
        this.notificationCount = response.data.NotificationCount
        this.tags = this.article.tag.split(',');

        //commentに各idを代入
        this.comment.articleid = this.article.id
        this.comment.userid = this.user.id
        this.closeFullScreen()
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
        this.closeFullScreen()
      })
  },
  filters: {
    moment: function (date) {
      moment.locale( 'ja' );
      return moment(date).utc().format('YYYY/MM/DD HH:mm');
    }
  },
  computed: {
    compiledMarkdown: function () {
      return marked(this.article.body)
    },
    compiledMarkdownComment: function(){
      return function(value) {
        return marked(value)
      }
    }
  },
  beforeRouteLeave (to, from, next) {
    if (this.comment.text !== ""){
      this.$confirm('編集中のものは保存されませんが、よろしいですか？', 'Warning', {
          confirmButtonText: 'OK',
          cancelButtonText: 'Cancel',
          type: 'warning',
          center: true
        }).then(() => {
          next()
        }).catch(() => {
          next(false)
        });
    }else{
      next();
    }
  },
  methods: {
    createArticleComment: function() {
      this.comment.name = this.user.name;
      this.$axios
        .post(this.url+`api/restricted/Articles/${this.article.id}/Comments`, this.comment,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          if(this.comments == null){
            this.comments = [];
          }
          this.comments.unshift(response.data.Comment);
          // 日本時間に変換
          this.comments[0].Created = moment.utc(this.comments[0].Created).local().format();
          this.comments[0].imgpath = this.user.imgpath
          this.comments[0].KBC_mail = this.user.KBC_mail
          this.comment.text = "";
          this.createCommentAlert();
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },
    createCommentAlert() {
      this.$notify({
        title: 'Success',
        message: '記事にコメントを送信しました。',
        type: 'success'
      });
    },

    deleteArticle: function() {
      this.$axios
        .delete(this.url+`api/restricted/Articles/${this.article.id}`,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.$router.push({ path: '/' });
          this.deleteArticleAlert();
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    deleteArticleAlert() {
      this.$notify({
        title: 'Success',
        message: '記事を削除しました。',
        type: 'success'
      });
    },

    deleteArticleComment: function(commentId,index) {
      this.$axios
        .delete(this.url+`api/restricted/Articles/${this.article.id}/Comments/${commentId}`,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(()=> {
          this.comments.splice(index, 1);
          this.deleteCommentAlert()
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    deleteCommentAlert() {
      this.$notify({
        title: 'Success',
        message: 'コメントを削除しました。',
        type: 'success'
      });
    },

    CreateArticleLike(){
      this.$axios
        .post(this.url+`api/restricted/Articles/${this.article.id}/Likes`,this.article.userid,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.like.isLike = true
          this.like.likeCount++;
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    DeleteLikeConfirmationOpen() {
      this.$confirm(`いいねを解除しますか?`, '確認', {
        confirmButtonText: '解除する',
        cancelButtonText: 'キャンセル',
        center: true
      }).then(() => {
        this.DeleteArticleLike()
      });
    },


    DeleteArticleLike(){
      this.$axios
        .delete(this.url+`api/restricted/Articles/${this.article.id}/Likes`,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.like.isLike = false
          this.like.likeCount--;
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    CreateArticleCommentLike(index){
      this.$axios
        .post(this.url+`api/restricted/Articles/Comments/${this.comments[index].id}/Likes`,{articleid:this.article.id,visitedid:this.comments[index].userid},{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.comments[index].Like.isLike = true
          this.comments[index].Like.likeCount++;
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    DeleteArticleCommentLike(index){
      this.$axios
        .delete(this.url+`api/restricted/Articles/Comments/${this.comments[index].id}/Likes`,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.comments[index].Like.isLike = false
          this.comments[index].Like.likeCount--;
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
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

    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },

    loginDialog() {
      this.$confirm('この機能を利用するにはログインする必要があります', 'ユーザー登録してみませんか？', {
        confirmButtonText: 'ログイン',
        cancelButtonText: 'キャンセル',
        type: 'success',
        center: true
      }).then(() => {
        this.$router.push("/Login");
      }).catch(() => {
      });
    },
    teacherMatch(kbcmail) {
      const mail = String(kbcmail)
      if (mail.match(/[\w\-._]+@kawahara.ac.jp/) !== null ){
        return true
      }
      return false
    },
  },

  watch: {
    notificationCount(newNotificationCount) {
      localStorage.notificationCount = newNotificationCount;
    },
  },
}
</script>

<style scoped>
.body-main{
  width: 1000px;
  padding: 15px 0 40px 0;
  background-color: #F6F6F4;
}
.article-show-main,.article-comment-all{
  width: 800px;
  margin: 0 auto;
  background-color: #fff;
}
.article-show-main{
  padding: 20px;
}
.article-comment-all{
  margin-top: 15px;
  padding: 0px 20px 20px 20px;
}
.article-title{
  padding: 0px 20px 0px 10px;
  font-size: 2.5em;
  margin-bottom: 20px;
}
.article-comment-header{
  display: flex;
  height: 50px;
  margin-bottom: 50px;
}
.article-create-date,.article-update-date{
  font-size: 10px;
  line-height: 10px;
}
.comment-icon{
  font-size: 30px;
  line-height: 60px;
  padding-top: 27px;
}
.comment{
  font-size: 20px;
  font-weight: 600;
  margin-left: 10px;
  margin-top: 33px;
  line-height: 50px;
}
.comment-div{
  width: 800px;
  border-bottom: #F6F6F4 solid 2px;
  margin-bottom: 30px;
}
.comment-header{
  display: flex;
  justify-content:space-between;
  height: 50px;
}
.comment-header-div{
  display: flex;
}
.comment-user-icon{
  width: 40px;
  height: 40px;
  object-fit: cover; /* 画像トリミング */
  border-radius: 50%;
}
.comment-user-name{
  margin-left: 10px;
  line-height: 10px;
}
.comment-create-date{
  font-size: 8px;
}
.comment-create-date,.comment-edit-icon{
  text-align: right;
  /* margin-left: 15px; */
}
.article-create-update-date-div{
  margin-right: 10px;
}
.article-create-date,.article-update-date{
  margin-top: 0;
  font-size: 14px;
}
.dropdown-span{
  margin-right: 20px;
  margin-left: 20px;
}
.icon-menu-span{
  text-align: center;
  font-size: 22px;
}
.comment-edit-icon{
  margin-top: 15px;
}
.el-icon-more{
  cursor: pointer;
}
.comment-text{
  margin: 10px 15px 50px 15px;
}
.a-tag2{
  text-decoration: none;
  color: #606266;
}
.like-btn{
  display: inline-block;
  line-height: 1;
  white-space: nowrap;
  cursor: pointer;
  background: #FFF;
  border: 1px solid #DCDFE6;
  color: #606266;
  -webkit-appearance: none;
  text-align: center;
  box-sizing: border-box;
  outline: 0;
  margin: 0;
  transition: .1s;
  font-weight: 500;
  padding: 12px 20px;
  font-size: 14px;
  border-radius: 4px;
}
.like-btn:hover{
  color: #E6A23C;
  border-color: #f5dab1;
  background: #fdf6ec;
}
.like-count-span{
  margin-left: 10px;
  font-size: 1.1em;
  color:#E6A23C;
  font-weight: bold;
}
.comment-like-count-span{
  margin-left: 1px;
  font-size: 0.9em;
  user-select: none;
}
.comment-ster-i{
  margin-left: 10px;
  font-size: 2em;
  cursor: pointer;
  color:#E6A23C;
}
.not-comment-ster-i{
  margin-left: 15px;
  font-size: 1.6em;
  cursor: pointer;
}
.article-tags-div{
  display: flex;
  margin-bottom: 50px;
  flex-wrap:wrap
}
.article-tag-div{
  margin-left: 10px;
  margin-top: 10px;
  padding: 4px 8px;
  border-radius: 4px;
  background: #eee;
  color: #666;
}
.article-tag-div > span{
  font-size: 0.9em;
}
.article-like-btn-div{
  padding-top: 50px;
}
.teacher{
  font-size: 1em;
  color: #409eff;
  cursor: pointer;
}
</style>
