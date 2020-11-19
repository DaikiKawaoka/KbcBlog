<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="false" :user="myUser"></Header>
    <div class="user-show-all">

      <div class="user-show-header">
        <div class="user-show-icon">
          <!-- <img class="user-show-icon-img" src="../assets/IMG_6217.jpeg"> -->
          <div class="demo-image__preview">
          <el-image
            style="width: 150px; height: 150px; border-radius: 50%;"
            :src="url"
            :preview-src-list="srcList">
          </el-image>
          </div>
        </div>
        <div class="user-show-info">
          <div class="user-show-info-header">
            <span class="user-show-name">{{user.name}}</span>

            <div class="user-show-info-div" v-if="myUser.id===user.id">
              <div class="user-show-btn-div"><el-button type="info" size="mini" @click="editUser">プロフィール編集</el-button></div>


              <el-dropdown>
                <span class="el-dropdown-link icon-menu-span">
                  <i class="el-icon-setting"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <router-link class="a-tag2" v-bind:to="{ name : 'UserEdit', params : { id: user.id }}"><el-dropdown-item>パスワード変更</el-dropdown-item></router-link>
                  <el-popconfirm
                  title="本当にログアウトしますか?"
                  confirm-button-text="Yes"
                  cancel-button-text="No"
                  @confirm="logout"
                  >
                    <el-dropdown-item slot="reference">ログアウト</el-dropdown-item>
                  </el-popconfirm>
                </el-dropdown-menu>
              </el-dropdown>


            </div>

            <div class="user-show-info-div" v-else>
              <div class="user-show-btn-div">
                <el-button type="primary" size="medium" v-if="true">
                  <i class="el-icon-user"></i>フォロー
                </el-button>

                <el-button size="medium" v-else>
                  <i class="el-icon-user"></i>フォロー解除
                </el-button>
              </div>
            </div>

          </div>
          <div class="user-show-info-main">
            <div class="user-show-article-count-div">
              <!-- UserArticleCount -->
              <p class="edit-margin-p">{{user.postCount}}</p>
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
            <p class="user-show-comment-p">{{user.comment.String}}</p>
          </div>
        </div>
      </div>

      <div class="user-show-body">
        <div class="user-show-link-info">
          <div class="user-show-link-div">
            <p class="link-tag">github</p>
              <p class="link-p" v-if="user.github.Valid">
                <span class="user-show-link-a" @click="open('github')">{{user.github.String}}</span>
              </p>
            <p class="link-p" v-else>未設定</p>
          </div>
          <div class="user-show-link-div">
            <p class="link-tag">webサイト</p>
              <p class="link-p" v-if="user.website.Valid">
                <span class="user-show-link-a" @click="open('website')">{{user.website.String}}</span>
              </p>
            <p class="link-p" v-else>未設定</p>
          </div>
        </div>
        <div class="user-show-body-info">
          <div class="user-show-body-indiv user-like-lang-div">
            <span class="user-show-body-div-in-span">好きな言語</span>
            <span v-if="user.languages.Valid">
              <span v-for="(lang, key) in langarray" :key="key" class="user-like-languages-span">
                {{lang}}
              </span>
            </span>
            <span v-else>
              <span class="not-conf">未設定</span>
            </span>
          </div>
          <div class="user-post-article-div">
            <div class="div-in-span-article">投稿記事</div>
            <!-- <span class="not-conf">未設定</span> -->
            <div class="article-circle-div">
              <el-progress type="circle" :percentage="33" color="#67C23A" :width="90"></el-progress>
              <p class="article-circle-p">Java</p>
            </div>
            <div class="article-circle-div">
              <el-progress type="circle" :percentage="25" color="#ff1493" :width="90"></el-progress>
              <p class="article-circle-p">JavaScript</p>
            </div>
            <div class="article-circle-div">
              <el-progress type="circle" :percentage="25" :width="90"></el-progress>
              <p class="article-circle-p">JavaScript</p>
            </div>
            <div class="article-circle-div">
              <el-progress type="circle" :percentage="25" color="#ff8c00" :width="90"></el-progress>
              <p class="article-circle-p">JavaScript</p>
            </div>
          </div>
          <div class="user-iine-div">
            <span class="user-show-body-div-in-span">Good記事</span>
            <span class="not-conf">未設定</span>
          </div>
        </div>
      </div>

      <div class="user-show-footer">
        <div class="user-show-post-menu">
          <el-tabs :tab-position="tabPosition" style="height: 150px; width: 170px;" @tab-click="handleClick">
            <el-tab-pane label="記事"><i class="el-icon-document"></i> 記事</el-tab-pane>
            <el-tab-pane label="質問"><i class="el-icon-chat-round"></i> 質問</el-tab-pane>
            <el-tab-pane label="回答"><i class="el-icon-s-promotion"></i> 回答</el-tab-pane>
            <!-- <el-tab-pane label="いいね"><i class="el-icon-star-off"></i> いいね</el-tab-pane> -->
          </el-tabs>
        </div>

        <div class="post-all-div" v-if="click_tab == 0">
          <div v-if="articles.length > 0">
            <div v-for="article in articles" :key="article.id" class="post-show-div">
              <router-link v-bind:to="{ name : 'ArticleShow', params : { id: article.id }}" class="a-tag">
                <div class="post-index-body">
                    <h2 class="post-title-index"> {{article.title}} </h2>
                  <div class="post-index-username-updated">
                    <h3 class="post-index-update">{{ article.Updated }}</h3>
                  </div>
                </div>
              </router-link>
            </div>
          </div>
          <div v-else>
            <p class="no-post-array">記事はありません。</p>
          </div>
        </div>
        <div class="post-all-div" v-else-if="click_tab == 1">
          <div v-if="questions.length > 0">
            <div v-for="question in questions" :key="question.id" class="post-show-div">
              <router-link v-bind:to="{ name : 'QuestionShow', params : { id: question.id }}" class="a-tag">
                <div class="post-index-body">
                    <h2 class="post-title-index"> {{question.title}} </h2>
                  <div class="post-index-username-updated">
                    <h3 class="post-index-update">{{ question.Updated }}</h3>
                  </div>
                </div>
              </router-link>
            </div>
          </div>
          <div v-else>
            <p class="no-post-array">質問はありません。</p>
          </div>
        </div>
        <div class="post-all-div" v-else>
          <div v-if="answerQuestions.length > 0">
            <div v-for="(answer,index) in answerQuestions" :key="index" class="post-show-div">
              <router-link v-bind:to="{ name : 'QuestionShow', params : { id: answer.id }}" class="a-tag">
                <div class="answer-index-body">
                    <span class="answer-question-title">「{{answer.question_title}}」</span> <span class="answer-question-title-last">に回答しました。</span>
                    <p class="answer-comment-text"> {{answer.comment_text}} </p>
                  <div class="post-index-username-updated">
                    <h3 class="post-index-update">{{ answer.comment_created }}</h3>
                  </div>
                </div>
              </router-link>
            </div>
          </div>
          <div v-else>
            <p class="no-post-array">回答はありません。</p>
          </div>
        </div>

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
      articles: Array,
      questions: Array,
      answerQuestions: Array,
      myUser: {},
      langarray: [],
      user: {
        KBCmail: "",
        id : 0,
        name: "",
        postCount: 0,
        comment: {
          String: String,
          Valid: Boolean
        },
        github: {
          String: "",
          Valid: Boolean
        },
        website: {
          String: "",
          Valid: Boolean
        },
        languages: {
          String: "",
          Valid: Boolean
        },
      },
      errors: {},
      url :"https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg",
      srcList :["https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg"],
      tabPosition: 'left',
      click_tab: 0,
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
        this.questions = response.data.Questions
        this.answerQuestions = response.data.AnswerQuestions
        this.user = response.data.User
        this.myUser = response.data.MyUser
        // 文字列のlangsを配列に変換
        this.langarray = this.user.languages.String.split(',');
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
    open(value) {
      this.$confirm('外部ページへの移動を許可しますか?', 'Warning', {
        confirmButtonText: '許可',
        cancelButtonText: 'キャンセル',
        type: 'warning',
        center: true
      }).then(() => {
        // this.$message({
        //   type: 'success',
        //   message: 'Delete completed'
        // });
        if( value === "github"){
          window.open(this.user.github.String, '_blank');
        }else{
          window.open(this.user.website.String, '_blank');
        }
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '外部ページへの移動をキャンセルしました。'
        });
      });
    },
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
    handleClick(tab) {
      this.click_tab = tab.paneName;
      console.log(this.click_tab);
    },
    logout: function() {
      this.$cookies.remove("JWT");
      this.$router.push({ path: "/login" });
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

/* header */

.user-show-header{
  display: flex;
  padding-bottom: 10px;
}
.user-show-icon{
  margin-top: 25px;
  border-radius: 50%;
}
.user-show-icon-img {
  width: 150px;
  height: 150px;
  border-radius: 50%;
}
.user-show-info{
  margin-top: 20px;
  margin-left: 50px;
  padding: 0 25px 10px 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
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
.user-show-comment-p{
  font-size: 1em;
  width: 480px;
  height: 100%;
  margin: 10px 0 0 15px;
  word-wrap: break-word;
}

/* body */

.user-show-body{
  display: flex;
  width: 800px;
  height: 230px;
  border-bottom: #000 1px solid;
  padding: 0 0 30px 0;
  margin: 0 auto 30px auto;
}
.user-show-link-info{
  width: 180px;
  height: 210px;
  margin-right: 20px;
  word-wrap: break-word;
  border-radius: 2px;
}
.user-show-link-div{
  margin-top: 25px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
}
.link-tag{
  font-weight:bold;
  margin: 5px;
  font-size: 0.9em;
}
.link-p{
  margin: 5px;
  font-size: 0.8em;
}
.user-show-link-a{
  color: #333;
  cursor: pointer;
}
.user-show-link-a:hover{
  color: #409eff;
}
.user-show-body-info{
  width: 550px;
  background-color: rgb(43, 43, 43);
  border-radius: 2px;
}
.user-show-body-indiv{
  height: 70px;
}
.user-iine-div{
  height: 60px;
}
.user-post-article-div{
  display: flex;
  height: 100px;
}
.article-circle-div{
  height: 100px;
}
.article-circle-p{
  margin: 0 0 0 20px;
  font-size: 0.7em;
  color: #fff;
  text-align: center;
}
.user-show-body-div-in-span{
  color: #94ff5f;
  line-height: 70px;
  font-size: 0.9em;
  margin: 0 10px 0 10px;
  width: 100px;
}
.div-in-span-article{
  color: #94ff5f;
  font-size: 0.9em;
  margin: 22px 10px 0 10px;
  width: 72px;
  height: 20px;
}
.user-like-languages-span{
  font-size: 1.4em;
  color: #fff;
  font-weight: bold;
  margin-left: 10px;
}
.not-conf{
  color: #ccc;
  margin-left: 10px;
}

/* footer */
.user-show-footer{
  display: flex;
}

.user-show-post-menu{
  margin-left: 40px;
}


.post-all-div{
  width: 550px;
  margin: 30px 50px 0 0;
}
.post-show-div{
  margin-bottom: 5px;
  padding-left: 10px;
  border: solid 1px #ccc;
  display: flex;
  background: #eee;
  border-radius: 6px;
}
.post-show-div:hover{
  transition: 0.2s ;
  background: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
}
.post-user-icon{
  margin-top: 25px;
  background-color: #ccc;
  width: 40px;
  height: 40px;
  border-radius: 50%;
}
.post-index-body{
  margin-left: 20px;
}
.post-index-username-updated{
  display: flex;
}
.post-title-index{
  font-size: 20px;
  width: 500px;
  text-align: left;
  margin-bottom: 0;
}
.post-index-username,.post-index-update{
  font-size: 13px;
  color: #999;
}
.post-index-update{
  margin-left: 350px;
}
.a-tag{
  color: #000;
  text-decoration: none;
  cursor: pointer;
}
.answer-index-body{
  margin-top: 10px;
  width: 500px;
  word-wrap: break-word;
}
.answer-question-title{
  font-weight: bold;
  font-size: 1.1em;
}
.answer-question-title-last{
  font-size: 0.8em;
  color: #555;
}
.answer-comment-text{
  margin: 0;
  width: 500px;
  word-break: normal;
}
.no-post-array{
  text-align: center;
  color: #777;
}
.el-progress{
  width: 80px;
  height: 80px;
  margin-left: 20px;
}
svg{
  width: 80px;
}
</style>
