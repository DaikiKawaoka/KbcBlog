<template>
  <header id="header">

    <!-- ログインしていない -->

    <div class="body-main header-body-main-div" v-if="loginpage || user === null">
      <div class="header-left">
        <router-link to="/login" class="a-tag"><h1 class="title">KBC Blog</h1></router-link>
      </div>
      <div class="header-right">
        <router-link to="/login" class="a-tag2"><span class="header-login-span">ログイン</span></router-link>
      </div>
    </div>

    <!-- ゲストユーザー -->

    <div class="body-main header-body-main-div" v-else-if="user.id === 920437694">
      <div class="header-left">
        <router-link to="/" class="a-tag"><h1 class="title" @click="$emit('reset')">KBC Blog</h1></router-link>
        <router-link to="/" class="a-tag2"><h3 class="header-article" v-bind:class="{ active: isArticle }">記事</h3></router-link>
        <router-link to="/Questions" class="a-tag2"><h3 class="header-question" v-bind:class="{ active: isQuestion }">質問</h3></router-link>
      </div>
      <div class="header-right">
        <div class="header-login-btn-div">
          <router-link to="/login" class="a-tag2"><span class="header-login-span">ログイン</span></router-link>
        </div>
      </div>
    </div>


    <!-- ログイン済み -->

    <div class="body-main header-body-main-div" v-else-if="user.id !== undefined && user.id !== null ">
      <div class="header-left">
        <router-link to="/" class="a-tag"><h1 class="title" @click="$emit('reset')">KBC Blog</h1></router-link>
        <router-link to="/" class="a-tag2"><h3 class="header-article" v-bind:class="{ active: isArticle }">記事</h3></router-link>
        <router-link to="/Questions" class="a-tag2"><h3 class="header-question" v-bind:class="{ active: isQuestion }">質問</h3></router-link>
      </div>
      <nav id="header-nav">
        <ul class="header-right-ul">
          <!-- <li><router-link tag="li" id="gift-nav" to="/Articles/new" class="a-tag">作成</router-link></li> -->
          <li>
            <el-dropdown>
              <span class="el-dropdown-link">
                <i class="el-icon-edit toukou-icon"></i>
                投稿<i class="el-icon-caret-bottom el-icon--right"> </i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <router-link to="/Articles/new" class="a-tag2"><el-dropdown-item>記事</el-dropdown-item></router-link>
                <router-link to="/Questions/new" class="a-tag2"><el-dropdown-item>質問</el-dropdown-item></router-link>
              </el-dropdown-menu>
            </el-dropdown>
          </li>
          <li>

            <el-badge :value="notificationCount" :max="99" class="badge-all">
              <el-popover
                placement="bottom"
                width="320"
                trigger="click"
                v-on:hide="hideNotifications">
                <!-- <el-table :data="user">
                  <el-table-column width="150" property="date" label="date"></el-table-column>
                  <el-table-column width="100" property="name" label="name"></el-table-column>
                  <el-table-column width="300" property="address" label="address"></el-table-column>
                </el-table> -->
                <div  v-if="notifications.length > 0">
                  <span style="margin-right: 230px;"><i class="el-icon-message-solid"></i>お知らせ</span>
                  <span style="margin:0;"><i v-on:click="deleteNotificationConfirmationOpen()" class="el-icon-delete" style="font-size: 1.4em; cursor:pointer;"></i></span>
                  <div class="overflow">
                    <div v-for="notification in notifications" :key="notification.id" class="notification-body-div">
                      <span v-if="notification.checked === true"></span>
                      <span v-else class="notification-badge">●</span>
                      <!-- <div class="user-icon">
                        <div class="notification-user-icon"> -->
                          <!-- <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag">
                          </router-link> -->
                        <!-- </div>
                      </div> -->
                      <div v-if="notification.action === 'alike'">
                        <router-link class="a-tag-article" v-bind:to="{ name : 'ArticleShow', params : { id: notification.articleid.Int32 }}" >
                        <div><p class="notification-text">
                          <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                            <span class="visiter-name">{{notification.name}}</span></router-link> があなたの記事に「いいね！」しました。 <span class="notification-created">{{ notification.Created | moment }}</span></p>
                        </div></router-link>
                      </div>
                      <div v-else-if="notification.action === 'qlike'">
                        <router-link class="a-tag-article" v-bind:to="{ name : 'QuestionShow', params : { id: notification.questionid.Int32 }}" >
                        <div><p class="notification-text">
                          <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                            <span class="visiter-name">{{notification.name}}</span></router-link> があなたの質問に「いいね！」しました。<span class="notification-created">{{ notification.Created | moment }}</span></p>
                          </div></router-link>
                      </div>
                      <div v-else-if="notification.action === 'follow'">
                        <div><p class="notification-text">
                          <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                            <span class="visiter-name">{{notification.name}}</span></router-link> があなたをフォローしました。 <span class="notification-created">{{ notification.Created | moment }}</span></p></div>
                      </div>
                      <div v-else-if="notification.action === 'aclike'">
                        <router-link class="a-tag-article" v-bind:to="{ name : 'ArticleShow', params : { id: notification.articleid.Int32 }}" >
                          <div><p class="notification-text">
                            <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                              <span class="visiter-name">{{notification.name}}</span></router-link> があなたのコメントに「いいね！」しました。 <span class="notification-created">{{ notification.Created | moment }}</span></p>
                              <div class="notification-comment-div">
                                <span class="notification-comment-p">{{notification.a_text.String}}</span>
                              </div>
                          </div>
                        </router-link>
                      </div>
                      <div v-else-if="notification.action === 'qclike'">
                        <router-link class="a-tag-article" v-bind:to="{ name : 'QuestionShow', params : { id: notification.questionid.Int32 }}" >
                          <div><p class="notification-text">
                            <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                              <span class="visiter-name">{{notification.name}}</span></router-link> があなたのコメントに「いいね！」しました。 <span class="notification-created">{{ notification.Created | moment }}</span></p>
                              <div class="notification-comment-div">
                                <span class="notification-comment-p">{{notification.q_text.String}}</span>
                              </div>
                          </div>
                        </router-link>
                      </div>
                      <div v-else-if="notification.action === 'acomment'">
                        <router-link class="a-tag-article" v-bind:to="{ name : 'ArticleShow', params : { id: notification.articleid.Int32 }}" >
                          <div><p class="notification-text">
                            <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                              <span class="visiter-name">{{notification.name}}</span></router-link> があなたの記事に「コメント」しました。 <span class="notification-created">{{ notification.Created | moment }}</span></p>
                              <div class="notification-comment-div">
                                <span class="notification-comment-p">{{notification.a_text.String}}</span>
                              </div>
                          </div>
                        </router-link>
                      </div>
                      <div v-else-if="notification.action === 'qcomment'">
                        <router-link class="a-tag-article" v-bind:to="{ name : 'QuestionShow', params : { id: notification.questionid.Int32 }}" >
                          <div><p class="notification-text">
                            <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                              <span class="visiter-name">{{notification.name}}</span></router-link> があなたの質問に「回答」しました。 <span class="notification-created">{{ notification.Created | moment }}</span></p>
                              <div class="notification-comment-div">
                                <span class="notification-comment-p">{{notification.q_text.String}}</span>
                              </div>
                          </div>
                        </router-link>
                      </div>
                    </div>
                    <infinite-loading @infinite="scrollNotifications" :identifier="infiniteId" spinner="spiral">
                      <span slot="no-more"></span>
                      <span slot="no-results"></span>
                    </infinite-loading>
                  </div>
                  <!-- <router-link to="/Notifications" class="a-tag2"><p class="text-center">通知一覧を見る</p></router-link> -->
                </div>
                <div v-else>
                  <p style="margin:0;"><i class="el-icon-message-solid"></i>お知らせ</p>
                  <p class="text-center">現在通知はございません。</p>
                </div>
                <el-button size="mini" slot="reference" @click="click_notification_btn"><i class="el-icon-message-solid message-solid-i"></i></el-button>
              </el-popover>
            </el-badge>

          </li>
          <li v-if="user.imgdata64">
            <img v-if="user.sex !== 3" class="header-user-icon" :src="image_path()" @click="myUserPage">
          </li>
        </ul>
      </nav>
    </div>
  </header>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading'
import moment from 'moment'
export default {
  props: {
    loginpage: Boolean,
    isArticle: Boolean,
    isQuestion: Boolean,
    user: {},
  },
  components: {
    InfiniteLoading,
  },
  data(){
    return{
      notificationCount: localStorage.notificationCount,
      notifications: [],
      notificationCursor: 0,
      page: 1,
      infiniteId: +new Date(),
    }
  },
  filters: {
    moment: function (date) {
      moment.locale('ja')
      return moment(date).fromNow();
    },
  },
  watch: {
    notificationCount(newNotificationCount) {
      localStorage.notificationCount = newNotificationCount;
    },
  },
  methods: {

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

    myUserPage: function() {
      if(this.$router.currentRoute.path !== `/Users/${this.user.id}`){
        this.$router.push({ name : 'UserShow', params : { id: this.user.id }});
      }else{
        this.$router.go({ name : 'UserShow', params : { id: this.user.id }})
      }
    },

    click_notification_btn: function(){
      this.changeType()
      const jst = this.$cookies.get("JWT");
      this.$axios.get('http://localhost/api/restricted/Notifications',{
        headers: {
          Authorization: `Bearer ${jst}`
        },
      })
      .then(response => {
        this.notifications = response.data.Notifications
        this.notificationCursor += 4
      })
      .catch(error => {
        if(error.response.status == 401){
          if(jst !== null){
            this.errorNotify();
          }
          this.$router.push({ path: "/login" });
        }
      })
    },
    // 通知一覧を閉じたとき、通知数を0にする処理
    hideNotifications() {
      this.notificationCount = 0;
    },

    scrollNotifications($state) {
      this.$axios.get('http://localhost/api/restricted/Notifications/scope', {
        params: {
          notificationCursor: this.notificationCursor,
        },
        headers: {
          Authorization: `Bearer ${this.$cookies.get("JWT")}`
        },
      })
      .then(response => {
        if (response.data.Notifications.length) {
          this.notifications = this.notifications.concat(response.data.Notifications);
          this.notificationCursor = response.data.NotificationCursor
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

    deleteNotificationConfirmationOpen() {
      this.$confirm(`通知を全て削除しますか？`, '確認', {
        confirmButtonText: '削除',
        cancelButtonText: 'キャンセル',
        center: true
      }).then(() => {
        this.deleteNotification()
      });
    },

    deleteNotification() {
      this.$axios
        .delete(`http://localhost/api/restricted/Notifications`,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(()=> {
          this.notifications = []
          this.deleteNotificationAlert()
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
        });
    },

    deleteNotificationAlert() {
      this.$notify({
        title: 'Success',
        message: '通知を削除しました。',
        type: 'success'
      });
    },

    changeType() {
      this.page = 1;
      this.infiniteId += 1;
      this.notifications = [];
      this.notificationCursor = 0;
    },

    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },
  }
}
</script>

<style scoped>
#header{
  height: 60px;
  /* background: #2ee002; */
  background: #24292e;
  border-bottom: #bbb solid 1px;
}
.header-body-main-div{
  display: flex;
}
.header-left{
  margin: 0;
  width: 50%;
  display: flex;
}
.header-right{
  width: 50%;
}
.title{
  line-height: 60px;
  margin: 0;
  color: #fff;
  text-decoration: none;
  cursor: pointer;
}
.header-article,.header-question{
  margin-left: 30px;
  margin-bottom: 10px;
  padding: 0 12px 7px 12px;
  color: #fff;
  cursor: pointer;
}
.header-login-span{
  float: right;
  margin-right: 30px;
  margin-top: 15px;
  padding: 7px 12px 7px 12px;
  background: #1e90ff;
  border-radius: 8%;
  color: #fff;
  font-weight: bold;
  cursor: pointer;
}

.active{
  border-bottom: 2px solid #2ee002;
  border-radius: 2px;
}
#header-nav{
  width: 50%;
}

.a-tag{
  line-height: 60px;
  font-size: 15px;
  text-align: right;
  color: #606566;
  text-decoration: none;
  cursor: pointer;
}
.a-tag2{
  text-decoration: none;
  color: #606266;
}
.el-dropdown-link {
  cursor: pointer;
  color: #fff;
}
.el-icon-arrow-down {
  font-size: 12px;
}
.toukou-icon{
  font-size: 18px;
}
.header-right-ul{
  margin: 10px 0 0 0;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  list-style: none;
  text-align: right;
}
.header-user-icon{
  width: 40px;
  height: 40px;
  border-radius: 50%;
  cursor: pointer;
  margin-left: 15px;
}
.badge-all{
  margin-left: 15px;
  margin-right: 10px;
}
.message-solid-i{
  font-size: 1.5em;
}


.notification-body-div{
  display: flex;
  position: relative;
  border-bottom: #ccc 1px solid;
  height: 80px;
}
.overflow{
  height: 350px;
  /* height: 100%; */
  overflow-y: scroll;
}
.notification-body-div > .user-icon{
  margin: auto 10px auto 0;
}
.notification-body-div > .user-icon > .notification-user-icon{
  background-color: #ccc;
  width: 30px;
  height: 30px;
  border-radius: 50%;
}
.a-tag-article{
  text-decoration: none;
  cursor: pointer;
}
.notification-text{
  margin: 0;
  padding: 10px 10px 5px 10px;
  color: #000;
}
.notification-text > .a-tag-n {
  text-decoration: none;
  cursor: pointer;
}
.notification-comment-div{
  overflow: hidden;
  white-space: nowrap;
	text-overflow: ellipsis;
	-webkit-text-overflow: ellipsis; /* Safari */
	-o-text-overflow: ellipsis; /* Opera */
  width: 250px;
  height: 20px;
}
.notification-comment-p{
  margin: 0;
  padding-left: 10px;
  padding-right: 10px;
  color: #666;
}
.visiter-name{
  font-weight: bold;
  color:#000;
}
.text-center{
  text-align: center;
}
.notification-created{
  color: #606266;
}

.notification-badge {
  position:relative;
  font-family: 'Courier New', Courier, monospace;
  color: #ed4956;
  /* position: relative; */
  position: absolute;
  top: 0px;
  font-size: 1.2em;
  /* text-shadow: 0px 0px 2px #008489; */
}
</style>
