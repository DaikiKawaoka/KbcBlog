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


<div class="tag-all">
    <el-form-item label="タグ">
    <div class="select-div">
      <el-select
        style="width: 400px;"
        v-model="tagArray"
        v-on:change="arrayChangeString"
        multiple
        size="large"
        :multiple-limit=5
        :collapse-tags="false"
        no-match-text="一致するタグがありません"
        filterable
        :allow-create="false"
        :default-first-option="true"
        placeholder="関連するタグを最大5つまで選択してください。">
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value">
        </el-option>
      </el-select>
    </div>
  </el-form-item>
  </div>

    <el-row>
      <el-col :span="2"><div class="grid-content"></div></el-col>
      <el-col :span="4"><div class="grid-content"></div></el-col>
      <el-col :span="4"><div class="grid-content"></div></el-col>
      <el-col :span="4"><div class="grid-content"></div></el-col>
      <el-col :span="6"><div class="grid-content"></div></el-col>
      <el-col :span="4">
        <!-- <el-button v-if="isActive" type="primary" icon="el-icon-open" circle v-on:click="active"></el-button>
        <el-button v-else icon="el-icon-turn-off" circle v-on:click="active"></el-button> -->
        <el-switch
          v-model="isActive"
          active-text="プレビュー"
          inactive-text="編集">
        </el-switch>
      </el-col>
    </el-row>

    <el-form-item label="本文" v-if="!isActive">
      <el-input
        type="textarea"
        :rows="20"
        placeholder="Please input"
        v-model="article.body"
        font>
      </el-input>
    </el-form-item>

    <div class="article-form__preview" v-else>
      <div class="article-form__label--preview">
        <p class="article-form__label--preview-p">プレビュー</p>
      </div>
      <div class="article-form__preview-body">
        <div class="article-form__preview-body-contents article-body" v-html="compiledMarkdown"></div>
      </div>
    </div>

    <el-row>
      <el-col :span="6"><div class="grid-content"></div></el-col>
      <el-col :span="8"><el-button type="danger" @click="$emit('cancell')">キャンセル</el-button></el-col>
      <el-col :span="8"><el-button type="success" @click="$emit('submit')">セーブ</el-button></el-col>
      <el-col :span="4"><div class="grid-content"></div></el-col>
    </el-row>
  </el-form>
</template>

<script>
import marked from 'marked';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';
import '../css/markdown.css';
import '../css/monokai-sublime.css'

export default {
  props: {
    article: Object,
    errors: Array,
    create: Boolean,
    // tagArray: Array,
  },
  data(){
    return{
      isActive: false,
      tagArray: [],
        options: [
          {
            value: ' プログラミング ',
            label: 'プログラミング'
          }, {
            value: ' 言語 ',
            label: '言語',
          }, {
            value: ' HTML ',
            label: 'HTML'
          }, {
            value: ' CSS ',
            label: 'CSS'
          }, {
            value: ' Java ',
            label: 'Java'
          }, {
            value: ' Python ',
            label: 'Python'
          }, {
            value: ' JavaScript ',
            label: 'JavaScript'
          }, {
            value: ' Node.js ',
            label: 'Node.js'
          }, {
            value: ' C/C++ ',
            label: 'C/C++'
          }, {
            value: ' C# ',
            label: 'C#'
          }, {
            value: ' SQL ',
            label: 'SQL'
          }, {
            value: ' PHP ',
            label: 'PHP'
          }, {
            value: ' Ruby ',
            label: 'Ruby'
          }, {
            value: ' Go ',
            label: 'Go'
          }, {
            value: ' TypeScript ',
            label: 'TypeScript'
          }, {
            value: ' Kotlin ',
            label: 'Kotlin'
          }, {
            value: ' Swift ',
            label: 'Swift'
          },{
            value: ' フレームワーク ',
            label: 'フレームワーク',
          } ,{
            value: ' Rails ',
            label: 'Rails'
          }, {
            value: ' Vue.js ',
            label: 'Vue.js'
          }, {
            value: ' React ',
            label: 'React'
          }, {
            value: ' AngularJS ',
            label: 'AngularJS'
          }, {
            value: ' Spring Framework ',
            label: 'Spring Framework'
          }, {
            value: ' Play Framework ',
            label: 'Play Framework'
          },{
            value: ' Bootstrap ',
            label: 'Bootstrap'
          }, {
            value: ' CakePHP ',
            label: 'CakePHP'
          }, {
            value: ' Laravel ',
            label: 'Laravel'
          }, {
            value: ' Django ',
            label: 'Django'
          }, {
            value: ' OS ',
            label: 'OS',
          }, {
            value: ' Linux ',
            label: 'Linux'
          }, {
            value: ' Windows ',
            label: 'Windows'
          }, {
            value: ' macOS ',
            label: 'macOS'
          }, {
            value: ' iOS ',
            label: 'iOS'
          }, {
            value: ' Android ',
            label: 'Android'
          }, {
            value: ' Webサーバ ',
            label: 'Webサーバ',
          }, {
            value: ' Apache ',
            label: 'Apache'
          }, {
            value: ' Nginx ',
            label: 'Nginx'
          }, {
            value: ' ネットワーク ',
            label: 'ネットワーク',
          }, {
            value: ' セキュリティ ',
            label: 'セキュリティ',
          },{
            value: ' データベース ',
            label: 'データベース',
          }, {
            value: ' MySQL ',
            label: 'MySQL'
          }, {
            value: ' PostgreSQL ',
            label: 'PostgreSQL'
          }, {
            value: ' Oracle Database ',
            label: 'Oracle Database'
          }, {
            value: ' SQLite ',
            label: 'SQLite'
          }, {
            value: ' MongoDB ',
            label: 'MongoDB'
          }, {
            value: ' 技術 ',
            label: '技術',
          }, {
            value: ' AWS ',
            label: 'AWS'
          }, {
            value: ' Docker ',
            label: 'Docker'
          }, {
            value: ' kubernetes ',
            label: 'kubernetes'
          }, {
            value: ' Git/GitHub ',
            label: 'Git/GitHub'
          }, {
            value: ' WordPress ',
            label: 'WordPress'
          }, {
            value: ' Firebase ',
            label: 'Firebase'
          }, {
            value: ' 国家試験 ',
            label: '国家試験',
          }, {
            value: ' ITパスポート試験 ',
            label: 'ITパスポート試験'
          },{
            value: ' 基本情報技術者試験 ',
            label: '基本情報技術者試験'
          }, {
            value: ' 応用情報技術者試験 ',
            label: '応用情報技術者試験'
          }, {
            value: ' 情報処理安全確保支援士試験 ',
            label: '情報処理安全確保支援士試験'
          },{
            value: ' KBC ',
            label: 'KBC',
          }, {
            value: ' ITイノベーション科 ',
            label: 'ITイノベーション科'
          }, {
            value: ' ITエンジニア科 ',
            label: 'ITエンジニア科'
          }, {
            value: ' ゲームクリエイター科 ',
            label: 'ゲームクリエイター科'
          }, {
            value: ' 情報ビジネス科 ',
            label: '情報ビジネス科'
          }, {
            value: ' その他 ',
            label: 'その他',
          }, {
            value: ' 初心者 ',
            label: '初心者'
          }, {
            value: ' 就活 ',
            label: '就活'
          }]
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
    this.tagArray = this.article.tag.split(',');
    if(this.tagArray[0].length == 0){
      this.tagArray = [];
    }
  },
  computed: {
    compiledMarkdown: function () {
      return marked(this.article.body)
    }
  },
  methods: {
    active: function () {
      this.isActive = !this.isActive;
    },
    arrayChangeString(){
      this.article.tag = this.tagArray.join(',');
    },
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
.tag-all{
  width: 210px;
  background-color: #F6F6F4;
  /* background-color: #15202b; */
}
.tag-header{
  margin-top: 20px;
}
.tag-header-text{
  color: #999;
}
ul{
  padding: 0;
}

</style>
