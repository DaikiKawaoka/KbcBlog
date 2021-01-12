<template>
  <div class="comment-form-all" v-if=" user.id !== 920437694 ">

    <div v-if="errors.length != 0">
      <ul class="error-ul" v-for="e in errors" :key="e">
          <li class="error-icon-li"><i class="el-icon-warning-outline"></i></li>
          <li><font color="red">{{ e }}</font></li>
      </ul>
    </div>

    <div class="comment-header">
          <!-- <router-link v-bind:to="{ name : 'UserShow', params : { id: user.id }}" class="a-tag"> -->
        <div class="comment-header-div">
            <div class="comment-user-icon" v-if="user.imgdata64">
              <img class="comment-user-icon" :src="image_path()">
            </div>
            <p class="comment-user-name">{{ user.name }}</p>
        </div>
          <!-- </router-link> -->
        <div class="form-preview-icon">
          <el-button v-if="isActive" type="primary" icon="el-icon-open" circle v-on:click="active"></el-button>
          <el-button v-else icon="el-icon-turn-off" circle v-on:click="active"></el-button>
        </div>
    </div>
    <div class="commnet-form-main">
      <el-form>
        <el-form-item label="投稿" v-if="isActive == false">
            <el-input
              type="textarea"
              :rows="4"
              autosize
              placeholder="Please input"
              v-model="comment.text"
              font>
            </el-input>
        </el-form-item>

        <div class="article-form__preview" v-else>
          <div class="article-form__label--preview"><p class="article-form__label--preview-p">プレビュー</p></div>
          <div class="article-form__preview-body">
            <div class="article-form__preview-body-contents comment-form-text" id="article-form-body" v-html="compiledMarkdown"></div>
          </div>
        </div>

        <el-button type="success" @click="$emit('submit')" class="submit-btn">投稿</el-button>
      </el-form>
    </div>
  </div>
  <div v-else>
    <p class="guestUser-comment-p">コメントをするためにはログインが必要です。</p>
    <div class="guestUser-signin-div"><el-button type="primary" @click="signup" class="guestUser-signin-btn">ユーザー登録</el-button></div>
    <div class="guestUser-login-div">既にアカウントを持っている方は<el-button type="text" @click="login">こちら</el-button></div>
  </div>
</template>

<script>
import marked from 'marked';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';
import '../components/commentFormMarkdown.css';

export default {
  props: {
    comment: Object,
    user: Object,
    errors: Object,
  },
  data(){
    return{
      isActive: false,
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
      return marked(this.comment.text)
    }
  },
  methods: {
    active: function () {
      this.isActive = !this.isActive;
    },
    image_path(){
      if(this.user.imgdata64.Valid === true){
        return 'data:image/jpeg;base64,' + this.user.imgdata64.String
        // return require("@/assets/userIcon/" + this.user.imgdata64.String);
      }else{
        if(this.user.sex === 1){
          return require("@/assets/userIcon/man.png");
        }else if(this.user.sex === 2){
          return require("@/assets/userIcon/woman.png");
        }
      }
    },
    signup: function(){
      this.$router.push("/Users/new");
    },
    login: function(){
      this.$router.push("/Login");
    },
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
  height: 200px;
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
.article-form__label--preview-p{
  line-height: 40px;
  color: #888;
  text-align: left;
  margin: 0;
}
.comment-header{
  display: flex;
}
.form-preview-icon{
  margin: 5px 10px 0 0;
}
.submit-btn{
  margin-left: 720px;
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
.guestUser-comment-p{
  text-align: center;
  font-size: 1.5em;
}
.guestUser-signin-div{
  width: 125px;
  margin-right: auto;
  margin-left: auto;
}
.guestUser-login-div{
  width: 300px;
  margin: 15px auto 0 auto;
}


</style>
