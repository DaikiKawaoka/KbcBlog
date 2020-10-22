<template>
  <el-form :model="article" class="article-form-all">

    <h2 v-if="create">技術記事作成</h2>
    <h2 v-else>技術記事編集</h2>

    <div v-if="errors.length != 0">
      <ul class="error-ul" v-for="e in errors" :key="e">
          <li class="error-icon-li"><i class="el-icon-warning-outline"></i></li>
          <li><font color="red">{{ e }}</font></li>
      </ul>
    </div>

    <el-form-item label="タイトル">
      <el-input
        type="text"
        placeholder="Please input"
        v-model="article.title"
        maxlength="50"
        show-word-limit
      >
      </el-input>
    </el-form-item>

    <el-row>
      <el-col :span="4"><div class="grid-content"></div></el-col>
      <el-col :span="4"><div class="grid-content"></div></el-col>
      <el-col :span="4"><div class="grid-content"></div></el-col>
      <el-col :span="4"><div class="grid-content"></div></el-col>
      <el-col :span="4"><div class="grid-content"></div></el-col>
      <el-col :span="4">
        <el-button v-if="isActive" type="primary" icon="el-icon-open" circle v-on:click="active"></el-button>
        <el-button v-else icon="el-icon-turn-off" circle v-on:click="active"></el-button>
      </el-col>
    </el-row>


    <el-form-item label="本文" v-if="isActive == false">
      <el-input
        type="textarea"
        :rows="20"
        placeholder="Please input"
        v-model="article.body"
        font>
      </el-input>
    </el-form-item>

    <div class="article-form__preview" v-else>
      <div class="article-form__label--preview"><p class="article-form__label--preview-p">プレビュー</p></div>
      <div class="article-form__preview-body">
        <div class="article-form__preview-body-contents" id="article-body" v-html="compiledMarkdown"></div>
      </div>
    </div>

    <el-row>
      <el-col :span="12"><el-button type="danger" @click="$emit('cancell')">キャンセル</el-button></el-col>
      <el-col :span="12"><el-button type="success" @click="$emit('submit')">セーブ</el-button></el-col>
    </el-row>
  </el-form>
</template>

<script>

import marked from 'marked';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';
import './markdown.css';

 export default {
   props: {
    article: Object,
    errors: Array,
    create: Boolean,
  },
  data(){
    return{
      isActive: false
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
      return marked(this.article.body)
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
.article-form-all{
  width: 1000px;
  margin-left: auto;
  margin-right: auto;
}
.el-textarea >>> .el-textarea__inner {
  font-family: inherit;
  font-size: 130%;
  resize: none;
}
.grid-content {
    border-radius: 4px;
    min-height: 36px;
}
.article-form__preview{
  margin-bottom: 22px;
}
.article-form__preview-body {
  grid-area: text;
  background: #fff;
  border: 1px solid #ddd;
  border-radius: 4px;
  width: 998px;
  height: 550px;
  word-break: break-all;
  white-space: normal;
  position: relative;
}
.article-form__preview-body-contents{
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  text-align: left;
  width: 960px;
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
i{
  font-size: 25px;
  color:#F56C6C;
}
li{
  list-style: none;
}
.error-icon-li{
  padding-right: 10px;
}
.error-ul{
  display: flex;
}

</style>
