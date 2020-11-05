<template>
  <div id="app">
    <Header :isArticle="true" :isQuestion="false"></Header>
    <div class="article-show-main">
      <div v-if="this.article != null">
        <div class="comment-header">
          <div class="comment-header-div">
            <div class="comment-user-icon"></div>
            <p class="comment-user-name">{{ article.name }}</p>
          </div>
          <div class="comment-header-div">
            <div class="article-create-update-date-div">
              <p class="comment-create-date article-create-date"> 作成日 {{ article.Created }}</p>
              <p class="comment-create-date article-update-date"> 更新日 {{ article.Updated }}</p>
            </div>
            <span v-if="article.userid === user.id" class="dropdown-span">
              <el-dropdown>
                <span class="el-dropdown-link icon-menu-span">
                  <i class="el-icon-more"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <router-link class="a-tag2" v-bind:to="{ name : 'ArticleEdit', params : { id: article.id }}"><el-dropdown-item>編集</el-dropdown-item></router-link>
                  <el-popconfirm
                  title="本当に削除しますか?"
                  confirm-button-text="Yes"
                  cancel-button-text="No"
                  @confirm="deleteArticle"
                  >
                    <el-dropdown-item slot="reference">削除</el-dropdown-item>
                  </el-popconfirm>
                </el-dropdown-menu>
              </el-dropdown>
            </span>
          </div>
        </div>
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
      user: null,
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
      },})
      .then(response => {
        this.article = response.data.Article
        this.user = response.data.user
        console.log(response.data)
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
        }
        console.log(error.response);
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
  methods: {
    deleteArticle: function() {
      this.$axios
        .delete(`http://localhost/api/restricted/Articles/${this.$route.params.id}`,{
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
.article-show-main,.article-comment-all{
  width: 800px;
  margin: 30px auto 0 auto;
  background-color: #fff;
  padding: 20px;
}
.article-title{
  padding: 0px 20px 20px 10px;
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
  margin-top: 10px;
}
.comment-text{
  margin-bottom: 50px;
}
.a-tag2{
  text-decoration: none;
  color: #606266;
}

</style>
