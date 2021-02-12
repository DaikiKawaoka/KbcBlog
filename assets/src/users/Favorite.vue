<template>
  <div id="app" v-if="user.id !== 920437694 && user.id !== 0 ">
    <Header :isArticle="false" :isQuestion="false" :user="myUser" :url="url"></Header>
    <div class="favorite-show-all">
      <div class="favorite-show-header">
        <div class="user-show-icon">
          <router-link v-bind:to="{ name : 'UserShow', params : { id: user.id }}" class="a-tag">
            <img class="user-show-icon-img" :src="user.imgpath">
          </router-link>
        </div>
        <div class="user-name-div">
          <div class="user-span-div">
          <span class="user-name-span">{{user.name}} </span><span><i v-if="teacherMatch(user.KBC_mail)" class="el-icon-success teacher" title="先生マーク"></i></span>
          </div>
        </div>
        <div class="user-show-info-div" v-if="user.id !== myUser.id">
          <div v-if="myUser.id === 920437694" class="user-show-btn-div">
            <el-button type="primary" size="medium" @click="loginDialog">
              <i class="el-icon-user"></i>フォロー
            </el-button>
          </div>

          <div v-else class="user-show-btn-div">
            <el-button type="primary" size="medium" v-if="!follow.isfollow" @click="Follow">
              <i class="el-icon-user"></i>フォロー
            </el-button>

            <el-button size="medium" v-else @click="BigBtnfollowConfirmationOpen()">
              <i class="el-icon-user"></i>フォロー解除
            </el-button>
          </div>
        </div>
      </div>

      <div class="favorite-show-footer">
        <p v-if="user.id !== myUser.id"><span class="user-name-span">{{user.name}}</span> さんの <span class="favorite-span">お気に入り投稿</span></p>
        <p v-else><span class="user-name-span">あなた</span> の <span class="favorite-span">お気に入り投稿</span></p>
        <div>
          <el-tabs type="card" @tab-click="handleClick">
            <el-tab-pane label="記事"></el-tab-pane>
            <el-tab-pane label="質問"></el-tab-pane>
            <el-tab-pane label="回答"></el-tab-pane>
          </el-tabs>
        </div>

        <div class="post-all-div" v-if="click_tab == 0">
          <div v-if="articles.length > 0">
            <div v-for="article in articles" :key="article.id" class="post-show-div">
              <div class="article-user-icon-div">
                <router-link v-bind:to="{ name : 'UserShow', params : { id: article.userid }}" class="a-tag">
                  <!-- image -->
                  <img class="article-user-icon" :src="article.imgpath">
                </router-link>
              </div>
              <router-link v-bind:to="{ name : 'ArticleShow', params : { id: article.id }}" class="a-tag">
                <div class="post-index-body">
                    <div class="article-index-tag-div">
                      <i class="el-icon-collection-tag tag-icon"></i>
                      <span class="article-tag no-magin">{{ article.tag }}</span>
                    </div>
                    <h2 class="post-title-index article-title-index"> {{article.title}} </h2>
                  <div class="post-index-username-updated">
                    <div class="post-index-name-like-div">
                      <router-link v-bind:to="{ name : 'UserShow', params : { id: article.userid }}" class="a-tag">
                        <h3 class="article-index-username">by {{ article.name }}</h3>
                      </router-link>
                      <i class="el-icon-star-on star-i"></i>
                      <span class="likecount-span">{{article.likecount}}</span>
                    </div>
                    <p class="post-index-update">{{ article.Updated | moment }}</p>
                  </div>
                </div>
              </router-link>
            </div>
            <infinite-loading @infinite="scrollArticles" :identifier="infiniteId" spinner="spiral" class="infinite-loading">
              <span slot="no-more" class="thats-all"><i class="el-icon-caret-top"></i> That’s all  <i class="el-icon-caret-top"></i></span>
              <span slot="no-results" class="thats-all"><i class="el-icon-caret-top"></i> That’s all  <i class="el-icon-caret-top"></i></span>
            </infinite-loading>
          </div>
          <div v-else class="no-post-array-div">
            <p class="no-post-array">記事はありません。</p>
          </div>
        </div>
        <div class="post-all-div" v-else-if="click_tab == 1">
          <div v-if="questions.length > 0">
            <div v-for="question in questions" :key="question.id" class="post-show-div">
              <div class="article-user-icon-div">
                <router-link v-bind:to="{ name : 'UserShow', params : { id: question.userid }}" class="a-tag">
                  <!-- image -->
                  <img class="article-user-icon" :src="question.imgpath">
                </router-link>
              </div>
              <router-link v-bind:to="{ name : 'QuestionShow', params : { id: question.id }}" class="a-tag">
                <div class="post-index-body">
                  <div class="article-index-tag-div">
                    <i class="el-icon-collection-tag tag-icon"></i>
                    <span class="article-tag no-magin">{{ question.tag }}</span>
                  </div>
                    <h2 class="post-title-index"> {{question.title}} </h2>
                  <div class="post-index-username-updated">
                    <div class="question-index-name-like-div">
                      <router-link v-bind:to="{ name : 'UserShow', params : { id: question.userid }}" class="a-tag">
                        <h3 class="article-index-username">by {{ question.name }} </h3>
                      </router-link>
                      <div>
                        <p v-if="question.category === 'Q&A'" class="question-index-category p-QandA"> {{question.category}} </p>
                        <p v-else class="question-index-category p-ikenkoukan"> {{question.category}} </p>
                      </div>
                      <i class="el-icon-star-on star-i"></i>
                      <span class="likecount-span">{{ question.likecount}}</span>
                    </div>
                    <p class="question-index-update">{{ question.Updated  | moment }}</p>
                  </div>
                </div>
              </router-link>
            </div>
          </div>
          <div v-else class="no-post-array-div">
            <p class="no-post-array">質問はありません。</p>
          </div>
        </div>
        <div class="post-all-div" v-else>
          <div v-if="answerQuestions.length > 0">
            <div v-for="(answer,index) in answerQuestions" :key="index" class="post-show-div">
              <div class="article-user-icon-div">
                <router-link v-bind:to="{ name : 'UserShow', params : { id: answer.userid }}" class="a-tag">
                  <!-- image -->
                  <img class="article-user-icon" :src="answer.imgpath">
                </router-link>
              </div>
              <router-link v-bind:to="{ name : 'QuestionShow', params : { id: answer.id }}" class="a-tag">
                <div class="answer-index-body">
                    <span class="answer-question-title">「{{answer.question_title}}」</span> <span class="answer-question-title-last"> への回答</span>
                    <p class="answer-comment-text"> {{answer.comment_text}} </p>
                  <div class="post-index-username-updated">
                    <div class="post-index-name-like-div">
                      <router-link v-bind:to="{ name : 'UserShow', params : { id: answer.userid }}" class="a-tag">
                        <h3 class="article-index-username">by {{ answer.name }}</h3>
                      </router-link>
                    </div>
                    <h3 class="post-index-update">{{ answer.comment_created | moment }}</h3>
                  </div>
                </div>
              </router-link>
            </div>
          </div>
          <div v-else class="no-post-array-div">
            <p class="no-post-array">回答はありません。</p>
          </div>
        </div>

      </div>
    </div>
    <Footer :user="myUser" :url="url"></Footer>
  </div>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading'
import moment from 'moment'
import Header from './../components/Header.vue';
import Footer from './../components/Footer.vue';

export default {
  name: 'app',
  data(){
    return {
      articles: [],
      questions: [],
      answerQuestions: [],
      follow:{},
      myUser: {},
      dialogVisible: false,
      uploadFile:null,
      fileData: null,
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
        imgpath: String,
      },
      errors: {},
      click_tab: 0,
      notificationCount: localStorage.notificationCount,
      loading: null,
      url: null,
      articleCursor: 0,
      page: 1,
      infiniteId: +new Date(),
    }
  },
  components: {
    InfiniteLoading,
    Header,
    Footer,
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.url = process.env.VUE_APP_URL
    this.openFullScreen()
    this.$axios.get(this.url+`api/restricted/Posts/${this.$route.params.id}/Favorite`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },})
      .then(response => {
        if(this.user.id === 920437694){
          this.$router.push("/");
        }
        this.myUser = response.data.MyUser
        this.user = response.data.User
        this.follow = response.data.Follow
        this.articles = response.data.Articles
        this.questions = response.data.Questions
        this.answerQuestions = response.data.AnswerQuestions
        this.articleCursor = response.data.ArticleCursor
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
      // locale関数を使って言語を設定すると、日本語で出力される
      moment.locale( 'ja' );
      return moment(date).utc().format('YYYY/MM/DD HH:mm');
    }
  },
  computed: {
  },
  methods: {

    // 大きいフォローボタンを押した時
    Follow(){
      this.$axios
        .post(this.url+`api/restricted/Users/${this.user.id}/Follow`,{
            visiterid: this.myUser.id,
            visitedid: this.user.id,
            action: "follow"
          },{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          this.follow.isfollow = true;
          this.follow.followedCount++;
          console.log(response);
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    // 大きいフォローボタンを押した時
    UnFollow(){
      this.$axios
        .delete(this.url+`api/restricted/Users/${this.user.id}/Follow`,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          this.follow.isfollow = false;
          this.follow.followedCount--;
          console.log(response);
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    BigBtnfollowConfirmationOpen() {
      this.$confirm(`${this.user.name}さんのフォローをやめますか?`, '確認', {
        confirmButtonText: 'フォローをやめる',
        cancelButtonText: 'キャンセル',
        center: true
      }).then(() => {
        this.UnFollow()
      });
    },

    handleClick(tab) {
      this.click_tab = tab.paneName;
    },

    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },

    imgClick(){
      this.$router.go({ name : 'UserShow', params : { id: this.user.id }})
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
      scrollArticles($state) {
        this.$axios.get(this.url+'api/restricted/Article/Favorites/scroll', {
          params: {
            userid: this.user.id,
            articleCursor: this.articleCursor,
          },
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          if (response.data.Articles.length) {
            this.articles = this.articles.concat(response.data.Articles);
            this.articleCursor = response.data.Cursor
            this.page += 1;
            $state.loaded();
          } else {
            $state.complete();
          }
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
        })
    },

  },
  watch: {
    notificationCount(newNotificationCount) {
      localStorage.notificationCount = newNotificationCount;
    },
    $route (){
      this.userUpdate()
    }
  },

}
</script>

<style scoped>

.favorite-show-all{
  width: 1000px;
  background-color: #F6F6F4;
  margin: 0 auto 0 auto;
  padding-bottom: 40px;
  display: flex;
  /* height: 550px; */
}

/* header */

.favorite-show-header{
  margin-right: 100px;
  padding-bottom: 10px;
  height: 550px;
}

.user-show-icon{
  margin-top: 50px;
  border-radius: 50%;
}
.user-show-icon-img{
  width: 180px;
  height: 180px;
  object-fit: cover; /* 画像トリミング */
  cursor: pointer;
  background: #000;
  border-radius: 50%;
}
.user-name-div{
  display: flex;
  width: 180px;
}
.user-span-div{
  margin: 10px auto;
}

.teacher{
  position: relative;
  top: 2px;
  font-size: 1.5em;
  color: #409eff;
  cursor: pointer;
}
.user-name-span{
  font-weight: bold;
  text-align: center;
  font-size: 1.4em;
  margin-top: 10px;
}

.user-show-btn-div{
  margin: 0 auto;
}

/* body */

.favorite-span{
  color: #ffa500;
  font-size: 1.2em;
  font-weight: bold;
}

.post-all-div{
  width: 600px;
  margin: 0px 0 0 0;
}
.post-show-div{
  margin-bottom: 5px;
  padding-left: 10px;
  border: solid 1px #ccc;
  display: flex;
  background: #fff;
  border-radius: 6px;
}
.post-show-div:hover{
  transition: 0.2s ;
  background: #fff;
  border: solid 1px #409eff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
}
.no-post-array-div{
  height: 450px;
}
.post-index-body{
  margin-left: 20px;
}
.post-index-username-updated,.post-index-name-like-div,.question-index-name-like-div{
  display: flex;
}
.post-index-name-like-div{
  width: 200px;
}
.question-index-name-like-div{
  width: 300px;
}
.article-index-username{
  margin-right: 10px;
}
.post-title-index{
  font-size: 20px;
  width: 500px;
  text-align: left;
  margin-bottom: 0;
  margin-top: 5px;
}
.article-title-index{
  margin-top: 5px;
}
.post-index-username,.post-index-update{
  font-size: 13px;
  color: #999;
}
.post-index-update{
  margin: 0px;
  position: relative;
  top: 15px;
  left: 220px;
}
.question-index-update{
  margin: 0px;
  position: relative;
  top: 15px;
  left: 120px;
}
.star-i{
  color: orange;
  margin-top: 15px;
  font-size: 1.3em;
}
.likecount-span{
  margin-top: 17px;
  font-size: 0.8em;
  color: #444;
}
.a-tag{
  color: #000;
  text-decoration: none;
  cursor: pointer;
}
.answer-index-body{
  margin-top: 10px;
  margin-left: 20px;
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

</style>
