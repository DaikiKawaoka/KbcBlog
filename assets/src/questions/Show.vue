<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="true" :user="user"></Header>
    <div class="question-show-main">
      <div v-if="this.question != null">
        <div class="comment-header">
          <div class="comment-header-div">
            <div class="comment-user-icon"></div>
            <p class="comment-user-name">{{ question.name }}</p>
          </div>
          <div class="comment-header-div">
            <div class="article-create-update-date-div">
              <p class="comment-create-date article-create-date"> 作成日 {{ question.Created }}</p>
              <p class="comment-create-date article-update-date"> 更新日 {{ question.Updated }}</p>
            </div>
            <span v-if="question.userid === user.id" class="dropdown-span">
              <el-dropdown>
                <span class="el-dropdown-link icon-menu-span">
                  <i class="el-icon-more"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <router-link class="a-tag2" v-bind:to="{ name : 'QuestionEdit', params : { id: question.id }}"><el-dropdown-item>編集</el-dropdown-item></router-link>
                  <el-popconfirm
                  title="本当に削除しますか?"
                  confirm-button-text="Yes"
                  cancel-button-text="No"
                  @confirm="deleteQuestion"
                  >
                    <el-dropdown-item slot="reference">削除</el-dropdown-item>
                  </el-popconfirm>
                </el-dropdown-menu>
              </el-dropdown>
            </span>
          </div>
        </div>
        <h1 class="article-title">{{question.title}}</h1>
        <div class="article-form__preview-body">
          <div class="article-form__preview-body-contents" id="article-body" v-html="compiledMarkdown"></div>
        </div>

        <el-row>
          <el-button v-if="like.isLike" type="warning" @click="click_like">解除 <i class="el-icon-star-on"></i></el-button>
          <button v-else @click="click_like" class="like-btn">いいね <i class="el-icon-star-off"></i></button>
          <span class="like-count-span">{{like.likeCount}}</span>
        </el-row>

      </div>
    </div>
    <div class="article-comment-all">
      <div class="article-comment-header">
        <i class="el-icon-chat-dot-round comment-icon"></i>
        <p class="comment">回答</p>
      </div>
      <div v-for="(c,index) in comments" :key="c.id" class="article-comment-main">
        <div class="comment-div">
          <div class="comment-header">
            <div class="comment-header-div">
              <div class="comment-user-icon"></div>
              <p class="comment-user-name">{{c.name}}</p>
            </div>
            <div class="comment-header-div">
              <div>
                <p class="comment-create-date">投稿日 {{c.Created}}</p>
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
                    @confirm="deleteQuestionComment(c.id,index)"
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

          <el-row>
            <!-- <i class="el-icon-star-on"></i> -->
            <i v-if="c.Like.isLike" class="el-icon-star-on comment-ster-i" @click="ChangeQuestionCommentLike(index)"></i>
            <i v-else class="el-icon-star-off not-comment-ster-i" @click="ChangeQuestionCommentLike(index)"></i>
            <span class="comment-like-count-span">{{c.Like.likeCount}}</span>
          </el-row>

        </div>
      </div>
      <CommentForm :comment="comment" :errors="errors" :user="user" @submit="createQuestionComment"></CommentForm>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import CommentForm from './../components/CommentForm.vue'
import Footer from './../components/Footer.vue';
import marked from 'marked';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';
import '../components/markdown.css';

export default {
  name: 'app',
  data(){
    return {
      question: null,
      like: {
        isLike: Boolean,
        likeCount: 0,
      },
      user: {},
      comment:{
        userid: 0,
        name: "",
        questionid: 0,
        text: "",
      },
      comments: [],
      errors: {},
    }
  },
  components: {
    Header,
    CommentForm,
    Footer,
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get(`http://localhost/api/restricted/Questions/${this.$route.params.id}`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },})
      .then(response => {
        this.question = response.data.Question
        this.user = response.data.user
        this.comments = response.data.Comments
        this.like = response.data.Like

        //commentに各idを代入
        this.comment.questionid = this.question.id
        this.comment.userid = this.user.id
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      }),
    function () {
      marked.setOptions({
        // code要素にdefaultで付くlangage-を削除
        langPrefix: '',
        // highlightjsを使用したハイライト処理を追加
        highlight: function(code, lang) {
          return hljs.highlightAuto(code, [lang]).value
        }
      });
    }
  },
  computed: {
    compiledMarkdown: function () {
      return marked(this.question.body)
    },
    compiledMarkdownComment: function(){
      return function(value) {
        return marked(value)
      }
    }
  },
  methods: {
    createQuestionComment: function() {
      this.comment.name = this.user.name;
      this.$axios
        .post(`http://localhost/api/restricted/Questions/${this.question.id}/Comments`, this.comment,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          if(this.comments == null){
            this.comments = [];
          }
          this.comments.unshift(response.data.Comment);
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
        message: '質問に回答しました。',
        type: 'success'
      });
    },

    deleteQuestionComment: function(commentId,index) {
      this.$axios
        .delete(`http://localhost/api/restricted/Questions/${this.question.id}/Comments/${commentId}`,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.comments.splice(index, 1);
          this.deleteCommentAlert();
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
        message: '自分の回答を削除しました。',
        type: 'success'
      });
    },

    deleteQuestion: function() {
      this.$axios
        .delete(`http://localhost/api/restricted/Questions/${this.$route.params.id}`,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.$router.push({ path: '/Questions' });
          this.deleteQuestionAlert();
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    deleteQuestionAlert() {
      this.$notify({
        title: 'Success',
        message: '質問を削除しました。',
        type: 'success'
      });
    },

    ChangeQuestionLike(){
      this.$axios
        .post(`http://localhost/api/restricted/Questions/${this.question.id}/Likes`,this.question.id,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.like.isLike = !this.like.isLike;
          if(this.like.isLike === false){
            this.like.likeCount--;
          }else{
            this.like.likeCount++;
          }
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    click_like(){
      this.ChangeQuestionLike();
    },

    ChangeQuestionCommentLike(index){
      this.$axios
        .post(`http://localhost/api/restricted/Questions/Comments/${this.comments[index].id}/Likes`,this.comments[index].id,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.comments[index].Like.isLike = !this.comments[index].Like.isLike;
          if(this.comments[index].Like.isLike === false){
            this.comments[index].Like.likeCount--;
          }else{
            this.comments[index].Like.likeCount++;
          }
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },
  }
}
</script>

<style>
.question-show-main,.question-comment-all{
  width: 800px;
  margin: 30px auto 0 auto;
  background-color: #fff;
  padding: 20px;
}
.question-title{
  padding: 20px;
}
.question-comment-header{
  display: flex;
  height: 50px;
  margin-bottom: 50px;
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
  background-color: #ccc;
  width: 40px;
  height: 40px;
  border-radius: 50%;
}
.comment-user-name{
  margin-left: 10px;
}
.comment-create-date,.comment-edit-icon{
  text-align: right;
  margin-left: 15px;
}
.comment-edit-icon{
  margin-top: 15px;
}
.comment-text{
  margin-bottom: 50px;
}

</style>