<template>
  <div id="app">
    <Header></Header>
    <div class="article-show-main">
      <div v-if="this.article != null">
        <h1 class="article-title">{{article.title}}</h1>
        <div class="article-form__preview-body">
          <div class="article-form__preview-body-contents" id="article-body" v-html="compiledMarkdown"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import marked from 'marked';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';
import '../components/markdown.css';

export default {
  name: 'app',
  data(){
    return {
      article: null,
    }
  },
  components: {
    Header
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get(`http://localhost/api/Articles/${this.$route.params.id}`)
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
.article-show-main{
  width: 800px;
  margin-right: auto;
  margin-left: auto;
  margin-top: 30px;
  background-color: #fff;
  padding: 20px;
}
.article-title{
  padding: 20px;
}
</style>
