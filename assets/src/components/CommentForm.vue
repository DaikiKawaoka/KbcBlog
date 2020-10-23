<template>
  <div class="comment-form-all">
    <div class="comment-header">
        <div class="comment-header-div">
          <div class="comment-user-icon"></div>
          <p class="comment-user-name">name</p>
        </div>
    </div>
    <div class="commnet-form-main">
      <el-button v-if="isActive" type="primary" icon="el-icon-open" circle v-on:click="active"></el-button>
        <el-button v-else icon="el-icon-turn-off" circle v-on:click="active"></el-button>
      <el-form :model="comment">
        <el-form-item label="投稿" v-if="isActive == false">
            <el-input
              type="textarea"
              :rows="4"
              placeholder="Please input"
              v-model="comment"
              font>
            </el-input>
        </el-form-item>

        <div class="article-form__preview" v-else>
          <div class="article-form__label--preview"><p class="article-form__label--preview-p">プレビュー</p></div>
          <div class="article-form__preview-body">
            <div class="article-form__preview-body-contents" id="article-body" v-html="compiledMarkdown"></div>
          </div>
        </div>

        <el-button type="success" @click="$emit('submit')">投稿</el-button>
      </el-form>
    </div>
  </div>
</template>

<script>
import marked from 'marked';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';
import './markdown.css';

export default {
  props: {
  },
  data(){
    return{
      isActive: false,
      comment: "",
    }
  },
  created: function () {
    marked.setOptions({
      // code要素にdefaultで付くlangage-を削除
      langPrefix: '',
      // highlightjsを使用したハイライト処理を追加
      highlight: function(code, lang) {
        return hljs.highlightAuto(code, [lang]).value
      }
    });
  },
  computed: {
    compiledMarkdown: function () {
      return marked(this.comment)
    }
  },
  methods: {
    active: function () {
      this.isActive = !this.isActive;
    }
  }
}
</script>

<style scoped>
.el-textarea >>> .el-textarea__inner {
  font-family: inherit;
  resize: none;
}
.article-form__preview-body {
  grid-area: text;
  background: #fff;
  border: 1px solid #ddd;
  border-radius: 4px;
  width: 800px;
  height: 94px;
  word-break: break-all;
  white-space: normal;
  position: relative;
  margin-bottom: 20px;
}
.article-form__preview-body-contents{
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  text-align: left;
  width: 750px;
  padding: 20px 20px 20px 20px;
  font-size: 110%;
  overflow: scroll;
}
.article-form__label--preview{
  height: 40px;
}
.article-form__label--preview-p{
  line-height: 40px;
  color: #888;
  text-align: left;
  margin: 0;
}
</style>
