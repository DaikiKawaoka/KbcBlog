<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="false" :user="myUser"></Header>
    <div class="user-show-all">

      <div class="user-show-header">
        <div class="user-show-icon">
          <img class="user-show-icon-img" src="../assets/IMG_6217.jpeg" @click="imgClick()">
          <el-dialog :visible.sync="dialogVisible" width="390px">
              <img width="350px" height="350px" style="border-radius:3px;" src="../assets/IMG_6217.jpeg" alt="">
              <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="centerDialogVisible = false">変更</el-button>
                <el-button type="danger" @click="centerDialogVisible = false">削除</el-button>
              </span>
          </el-dialog>

          <!-- <div class="demo-image__preview">
          <el-image
            style="width: 150px; height: 150px; border-radius: 50%;"
            :src="url"
            :preview-src-list="srcList">
          </el-image>
          </div> -->
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
                  <router-link class="a-tag2" v-bind:to="{ name : 'PassEdit', params : { id: user.id }}"><el-dropdown-item>パスワード変更</el-dropdown-item></router-link>
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
                <el-button type="primary" size="medium" v-if="!follow.isfollow" @click="Follow_UnFollw">
                  <i class="el-icon-user"></i>フォロー
                </el-button>

                <el-button size="medium" v-else @click="Follow_UnFollw">
                  <i class="el-icon-user"></i>フォロー解除
                </el-button>
              </div>
            </div>


            <el-dialog title="フォロワー" :visible.sync="followerdialog" width="400px" :center="true">
              <div class="dialog-all">
                <ul class="dialog-ul">
                  <li class="dialog-li" v-for="(f,index) in this.followers" :key="f.id">
                    <div class="dialog-main">
                      <router-link v-bind:to="{ name : 'UserShow', params : { id: f.id }}" class="a-tag">
                        <div class="dialog-main">
                          <div>
                            <!-- アイコン -->
                            <div class="f-user-icon"></div>
                          </div>
                          <div class="f-info-div">
                            <div>
                              <span class="f-user-name">{{f.name}}</span>
                            </div>
                            <div class="string-out">
                              <span class="f-user-comment">{{f.comment.String}}</span>
                            </div>
                          </div>
                        </div>
                      </router-link>
                      <div class="follow-btn-div" v-if="f.id !== myUser.id">
                        <el-button size="small" v-if="f.isfollowing" @click="followConfirmationOpen(f,index,'followers','unfo')">
                          <i class="el-icon-user"></i>フォロー中
                        </el-button>
                        <el-button type="primary" size="small" v-else @click="ChangeFollow(f,index,'followers','fo')">
                          <i class="el-icon-user"></i>フォロー
                        </el-button>
                      </div>
                    </div>
                  </li>
                </ul>
              </div>
            </el-dialog>

            <el-dialog title="フォロー" :visible.sync="followdialog" width="400px" :center="true">
              <div class="dialog-all">
                <ul class="dialog-ul">
                  <li class="dialog-li" v-for="(f,index) in this.follows" :key="f.id">
                    <div class="dialog-main">
                      <router-link v-bind:to="{ name : 'UserShow', params : { id: f.id }}" class="a-tag">
                        <div class="dialog-main">
                          <div>
                            <!-- アイコン -->
                            <div class="f-user-icon"></div>
                          </div>
                          <div class="f-info-div">
                            <div>
                              <span class="f-user-name">{{f.name}}</span>
                            </div>
                            <div class="string-out">
                              <span class="f-user-comment">{{f.comment.String}}</span>
                            </div>
                          </div>
                        </div>
                      </router-link>
                      <div class="follow-btn-div" v-if="f.id !== myUser.id">
                        <!-- フォローボタン -->
                        <el-button size="small" v-if="f.isfollowing" @click="followConfirmationOpen(f,index,'follows','unfo')">
                          <i class="el-icon-user"></i>フォロー中
                        </el-button>
                        <el-button type="primary" size="small" v-else @click="ChangeFollow(f,index,'follows','fo')">
                          <i class="el-icon-user"></i>フォロー
                        </el-button>
                      </div>
                    </div>
                  </li>
                </ul>
              </div>
            </el-dialog>



          </div>
          <div class="user-show-info-main">
            <div class="user-show-article-count-div">
              <!-- UserArticleCount -->
              <p class="edit-margin-p">{{user.postCount}}</p>
              <p class="edit-margin-p">投稿<i class="el-icon-document"></i></p>
            </div>
            <div class="user-show-follower-count-div" @click="followerdialog = true">
              <!-- Count -->
              <p class="edit-margin-p">{{follow.followedCount}}</p>
              <p class="edit-margin-p">フォロワー</p>
            </div>
            <div class="user-show-follow-count-div" @click="followdialog = true">
              <!-- Count -->
              <p class="edit-margin-p">{{follow.followerCount}}</p>
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
              <span v-if="user.github.Valid">
                <p class="user-show-link-a string-out" @click="open('github')">{{user.github.String}}</p>
              </span>
            <p v-else class="not-conf-link">未設定</p>
          </div>
          <div class="user-show-link-div">
            <p class="link-tag">webサイト</p>
              <span v-if="user.website.Valid">
                <p class="user-show-link-a string-out" @click="open('website')">{{user.website.String}}</p>
              </span>
            <p v-else class="not-conf-link">未設定</p>
          </div>
        </div>
        <div class="user-show-body-info">
          <div class="user-show-body-indiv">
            <p class="user-show-body-div-in-p">好きな言語</p>
            <div v-if="user.languages.Valid" class="user-like-lang-div">
              <div v-for="(lang, key) in langarray" :key="key" class="user-like-lang-for-div">
                <img :src="getImgUrl(lang)" width="50" height="50" class="lang-image">
                <p class="user-like-languages-p">{{lang}}</p>
              </div>
            </div>
            <span v-else>
              <span class="not-conf lang-like-not-conf">未設定</span>
            </span>
          </div>
          <div class="user-post-article-div">
            <div class="div-in-span-article">投稿記事</div>
            <div v-if="parcentArray.length !== 0" style="display: flex;">
              <div v-if="parcentArray.length > 0" class="article-circle-div">
                <el-progress type="circle" :percentage="Math.round(parcentArray[0].count / user.postCount*100) " color="#67C23A" :width="90"></el-progress>
                <p class="article-circle-p">{{ parcentArray[0].value }}</p>
              </div>
              <div v-if="parcentArray.length > 1" class="article-circle-div">
                <el-progress type="circle" :percentage=" Math.round(parcentArray[1].count / user.postCount*100) " color="#ff1493" :width="90"></el-progress>
                <p class="article-circle-p">{{ parcentArray[1].value }}</p>
              </div>
              <div v-if="parcentArray.length > 2" class="article-circle-div">
                <el-progress type="circle" :percentage=" Math.round(parcentArray[2].count / user.postCount*100) " :width="90"></el-progress>
                <p class="article-circle-p">{{ parcentArray[2].value }}</p>
              </div>
              <div v-if="parcentArray.length > 3" class="article-circle-div">
                <el-progress type="circle" :percentage=" Math.round(parcentArray[3].count / user.postCount*100) " color="#ff8c00" :width="90"></el-progress>
                <p class="article-circle-p">{{ parcentArray[3].value }}</p>
              </div>
            </div>
            <span v-else class="not-conf article-parcent-not-conf">No data</span>
          </div>
          <!-- <div class="user-iine-div">
            <span class="user-show-body-div-in-span">Good記事</span>
            <span class="not-conf">No data</span>
          </div> -->
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
                    <div class="article-index-tag-div">
                      <i class="el-icon-collection-tag tag-icon"></i>
                      <span class="article-tag no-magin">{{ article.tag }}</span>
                    </div>
                    <h2 class="post-title-index article-title-index"> {{article.title}} </h2>
                  <div class="post-index-username-updated">
                    <i class="el-icon-star-on star-i"></i>
                    <span class="likecount-span">{{article.likecount}}</span>
                    <p class="post-index-update">{{ article.Updated | moment }}</p>
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
                  <div class="article-index-tag-div">
                    <i class="el-icon-collection-tag tag-icon"></i>
                    <span class="article-tag no-magin">{{ question.tag }}</span>
                  </div>
                    <h2 class="post-title-index"> {{question.title}} </h2>
                  <div class="post-index-username-updated">
                    <div>
                      <p v-if="question.category === 'Q&A'" class="question-index-category p-QandA"> {{question.category}} </p>
                      <p v-else class="question-index-category p-ikenkoukan"> {{question.category}} </p>
                    </div>
                    <i class="el-icon-star-on star-i"></i>
                    <span class="likecount-span">{{ question.likecount}}</span>
                    <p class="post-index-update">{{ question.Updated  | moment }}</p>
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
                    <h3 class="post-index-update">{{ answer.comment_created | moment }}</h3>
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
    <Footer></Footer>
  </div>
</template>

<script>
// import countArrayValues from 'count-array-values'
import moment from 'moment'
import Header from './../components/Header.vue';
import Footer from './../components/Footer.vue';

export default {
  name: 'app',
  data(){
    return {
      articles: Array,
      questions: Array,
      answerQuestions: Array,
      myUser: {},
      follow:{},
      follows:[],
      followers:[],
      langarray: [],
      tagArray: [],
      parcentArray: [],
      dialogVisible: false,
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
      followerdialog: false,
      followdialog: false,
      notificationCount: localStorage.notificationCount,
    }
  },
  components: {
    Header,
    Footer,
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
        this.follow = response.data.Follow
        this.follows = response.data.Follows
        this.followers = response.data.Followers
        this.notificationCount = response.data.NotificationCount
        if(this.followers === null){
          this.followers = [];
        }
        // 文字列のlangsを配列に変換
        this.langarray = this.user.languages.String.split(',');
        this.tagArray = response.data.Tags
        this.tagCountPercent()
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
  },
  filters: {
    moment: function (date) {
      // locale関数を使って言語を設定すると、日本語で出力される
      moment.locale( 'ja' );
      return moment(date).format('YYYY/MM/DD HH:mm');
    }
  },
  computed: {
  },
  methods: {

    tagCountPercent: function() {
      let x = [];
      let w;
      this.tagArray.forEach(function( value ) {
        w = value.split(',')
        if(w.length > 1){
          w.forEach(function(value2){
            x.push(value2);
          })
        }else{
          x.push(w[0])
        }
      });
      // console.log(x)
      var count = require('count-array-values')
      w = count(x)
      if(w.length > 3){
        for(let i = 0; i < 4; i++){
          this.parcentArray.push(w[i])
        }
      }else{
        for(let i = 0; i < w.length; i++){
          this.parcentArray.push(w[i])
        }
      }
    },

    // 大きいフォローボタンを押した時
    Follow_UnFollw(){
      this.$axios
        .post(`http://localhost/api/restricted/Users/${this.user.id}/Follow`,{
            visiterid: this.myUser.id,
            visitedid: this.user.id,
            action: "follow"
          },{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          this.follow.isfollow = !this.follow.isfollow;
          if(this.follow.isfollow === false){
            const w = this.followers.findIndex(item => item.id === this.myUser.id)
            if(w !== -1){
              //ユーザのフォロワー配列から自分を削除
              this.followers.splice(w, 1);
            }
            this.follow.followedCount--;
          }else{
            const user={id: this.myUser.id,comment:this.myUser.comment, isfollowing: true, name: this.myUser.name}
            this.followers.push(user)
            this.follow.followedCount++;
          }
          console.log(response)
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    open(value) {
      this.$confirm('外部ページへの移動を許可しますか?', 'Warning', {
        confirmButtonText: '許可',
        cancelButtonText: 'キャンセル',
        type: 'warning',
        center: true
      }).then(() => {
        if( value === "github"){
          window.open(this.user.github.String, '_blank');
        }else{
          window.open(this.user.website.String, '_blank');
        }
      }).catch(() => {
        this.$notify({
          title: 'Cancell',
          type: 'info',
          message: '外部ページへの移動をキャンセルしました。'
        });
      });
    },

    followConfirmationOpen(f,index,c,c2) {
      this.$confirm(`${f.name}さんのフォローをやめますか?`, '確認', {
        confirmButtonText: 'フォローをやめる',
        cancelButtonText: 'キャンセル',
        center: true
      }).then(() => {
        this.ChangeFollow(f,index,c,c2)
      });
    },

    editUser: function() {
      this.$router.push({ path: `/Users/${this.$route.params.id}/edit` });
    },
    handleClick(tab) {
      this.click_tab = tab.paneName;
    },
    logout: function() {
      this.$cookies.remove("JWT");
      this.$router.push({ path: "/login" });
    },

    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },

    // ChangeFollow(f,index,'follows','unfo')
    ChangeFollow(f,index,c,c2){
      this.$axios
        .post(`http://localhost/api/restricted/Users/${f.id}/Follow`,{
            visiterid: this.myUser.id,
            visitedid: this.f.id,
            action: "follow"
          },{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          if(this.user.id === this.myUser.id){
            if(c === "followers"){
              if(c2 === 'fo'){
                this.followers[index].isfollowing = true;
                const w = this.follows.findIndex(item => item.id === this.followers[index].id)
                if(w !== -1){
                  this.follows[w].isfollowing = true
                }else{
                  //フォロワーリストから初めてフォローボタンを押した時に自分のフォローリストに追加処理
                  this.follows.push(f)
                }
                this.follow.followerCount++;
              }else{
                this.followers[index].isfollowing = false;
                const w = this.follows.findIndex(item => item.id === this.followers[index].id)
                if(w !== -1){
                  this.follows[w].isfollowing = false
                }
                this.follow.followerCount--;
              }
            }else{ // フォローリストのボタン
              if(c2 === "fo"){
                this.follows[index].isfollowing = true;
                const w = this.followers.findIndex(item => item.id === this.follows[index].id)
                if(w !== -1){
                  this.followers[w].isfollowing = true;
                }
                this.follow.followerCount++;
              }else{ //アンフォロー
                this.follows[index].isfollowing = false;
                const w = this.followers.findIndex(item => item.id === this.follows[index].id)
                if(w !== -1){
                  this.followers[w].isfollowing = false
                }
                this.follow.followerCount--;
              }
            }
          }else{
            if(c === "followers"){
              this.followers[index].isfollowing = !this.followers[index].isfollowing;
            }else{
              this.follows[index].isfollowing = !this.follows[index].isfollowing;
            }
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
    getImgUrl(lang) {
      // #はパス名で使えないから変換
      if(lang === "C#"){
        lang = "C-sharp"
      }
      return require('@/assets/pgsvg/'+lang+".svg")
    },
    imgClick(){
      this.dialogVisible = true
    },
  },
  watch: {
    notificationCount(newNotificationCount) {
      localStorage.notificationCount = newNotificationCount;
    },
  },
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
  margin-top: 50px;
  border-radius: 50%;
}
.user-show-icon-img {
  width: 180px;
  height: 180px;
  border-radius: 50%;
}
.user-show-info{
  margin-top: 20px;
  margin-left: 22px;
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
.a-tag2{
  text-decoration: none;
  color: #606266;
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
  margin-left: 20px;
  width: 80px;
  text-align: center;
}
.user-show-follower-count-div,.user-show-follow-count-div{
  cursor: pointer;
}

/* ダイアログ */
.el-dialog--center{
  padding: 0px;
}
.el-dialog__body{
  padding: 0px;
  height: 320px;
}
.dialog-all{
  height: 100%;
  overflow-y: scroll;
}
.dialog-main{
  display: flex;
}
.dialog-ul{
  margin: 0;
  padding: 0;
}
.dialog-li{
  margin-bottom: 7px;
}
.f-user-icon{
  background-color: #ccc;
  width: 35px;
  height: 35px;
  border-radius: 50%;
}
.f-user-name{
  color: #000;
  font-weight: bold;
}
.f-user-comment{
  font-size: 0.85em;
}
.f-info-div{
  margin-left: 10px;
  width: 190px;
}
.string-out{
  overflow: hidden;
  white-space: nowrap;
	text-overflow: ellipsis;
	-webkit-text-overflow: ellipsis; /* Safari */
	-o-text-overflow: ellipsis; /* Opera */
}
.follow-btn-div{
  margin-left: 10px;
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
  height: 50px;
  margin-top: 25px;
  border-radius: 4px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
}
.link-tag{
  font-weight:bold;
  margin: 5px;
  font-size: 0.9em;
}
.not-conf-link{
  color: #777;
  font-size: 0.9em;
  margin: 0 0 0 5px;
}

.user-show-link-a{
  color: #333;
  cursor: pointer;
  width: 160px;
  margin: 0 0 0 5px;
  font-size: 0.9em;
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
  display: flex;
  height: 100px;
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
  width: 80px;
  margin: 0 0 0 20px;
  font-size: 0.7em;
  color: #fff;
  text-align: center;
  text-overflow: ellipsis;
  /* white-space: nowrap; */
}

.user-show-body-div-in-p{
  color: #94ff5f;
  font-size: 0.9em;
  margin: 0 10px 0 10px;
  line-height: 100px;
  width: 100px;
}
.div-in-span-article{
  color: #94ff5f;
  font-size: 0.9em;
  margin: 22px 10px 0 10px;
  width: 72px;
  height: 20px;
}
.user-like-languages-p{
  font-size: 0.8em;
  color: #fff;
  font-weight: bold;
  margin: 0;
  text-align: center;
}
.not-conf{
  color: #ccc;
  /* margin-left: 10px; */
}
.lang-like-not-conf{
  line-height: 100px;
}
.article-parcent-not-conf{
  height: 20px;
  margin-top: 22px;
  margin-left: 30px;
}
.el-progress{
  width: 80px;
  height: 80px;
  margin-left: 20px;
}
svg{
  width: 80px;
}
.lang-image{
  padding-top: 10px;
}
.user-like-lang-div{
  display: flex;
  margin-left: 45px;
}
.user-like-lang-for-div{
  margin-right: 50px;
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
  background: #fff;
  border-radius: 6px;
}
.post-show-div:hover{
  transition: 0.2s ;
  background: #fff;
  border: solid 1px #409eff;
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
  margin-left: 300px;
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
