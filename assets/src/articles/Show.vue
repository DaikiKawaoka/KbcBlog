<template>
  <div id="app">
    <Header></Header>
    <div class="article-show-main body-main">
      <div v-if="this.article != null">
        <h1 class="article-title">{{article.title}}</h1>
        <div class="article-form__preview-body">
          <div class="article-form__preview-body-contents" id="article-body" v-html="compiledMarkdown"></div>
        </div>
      </div>
    </div>
    <div class="article-comment-all">
      <div class="article-comment-header">
        <i class="el-icon-chat-dot-round comment-icon"></i>
        <p class="comment">コメント</p>
      </div>
      <div class="article-comment-main">
        <div class="comment-div">
          <div class="comment-header">
            <div class="comment-header-div">
              <div class="comment-user-icon"></div>
              <p class="comment-user-name">name</p>
            </div>
            <div class="comment-header-div">
              <p class="comment-create-date">2020-10-10 10:59</p>
              <i class="el-icon-more comment-edit-icon"></i>
            </div>
          </div>
          <div class="comment-main">
            <p class="comment-text">aaaaaaiiiiiiuuuuuueeeeeeoooooo</p>
          </div>
        </div>
      </div>
      <div class="article-comment-main">
        <div class="comment-div">
          <div class="comment-header">
            <div class="comment-header-div">
              <div class="comment-user-icon"></div>
              <p class="comment-user-name">name</p>
            </div>
            <div class="comment-header-div">
              <p class="comment-create-date">2020-10-10 10:59</p>
              <i class="el-icon-more comment-edit-icon"></i>
            </div>
          </div>
          <div class="comment-main">
            <p class="comment-text">aaaaaaiiiiiiuuuuuueeeeeeoooooo</p>
          </div>
        </div>
      </div>
      <CommentForm></CommentForm>
    </div>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import CommentForm from './../components/CommentForm.vue'
import marked from 'marked';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';
import '../components/markdown.css';

export default {
  name: 'app',
  data(){
    return {
      article: null,
      comments: [],
    }
  },
  components: {
    Header,
    CommentForm,
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get(`http://localhost/api/restricted/Articles/${this.$route.params.id}`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.article = response.data.Article
        console.log(this.article)
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
      return marked(this.article.body)
    }
  },
}
</script>

<style>
.article-show-main,.article-comment-all{
  width: 800px;
  margin: 30px auto 0 auto;
  background-color: #fff;
  padding: 20px;
}
.article-title{
  padding: 20px;
}
.article-comment-header{
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
.comment-create-date,.comment-user-name{
  line-height: 10px;
}
.comment-create-date,.comment-edit-icon{
  text-align: right;
  margin-left: 15px;
}
.comment-edit-icon{
  margin-top: 10px;
}
.comment-text{
  margin-bottom: 50px;
}

</style>
