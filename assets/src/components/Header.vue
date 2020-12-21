<template>
  <header id="header">
    <div class="body-main header-body-main-div">
      <div class="header-left">
        <router-link to="/" class="a-tag"><h1 class="title" @click="$emit('reset')">KBC Blog</h1></router-link>
        <router-link to="/" class="a-tag2"><h3 v-if="!loginpage" class="header-article" v-bind:class="{ active: isArticle }">記事</h3></router-link>
        <router-link to="/Questions" class="a-tag2"><h3 v-if="!loginpage" class="header-question" v-bind:class="{ active: isQuestion }">質問</h3></router-link>
      </div>
      <nav v-if="!loginpage" id="header-nav">
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
                trigger="click">
                <!-- <el-table :data="user">
                  <el-table-column width="150" property="date" label="date"></el-table-column>
                  <el-table-column width="100" property="name" label="name"></el-table-column>
                  <el-table-column width="300" property="address" label="address"></el-table-column>
                </el-table> -->
                <div  v-if="notifications.length > 0">
                  <p style="margin:0;"><i class="el-icon-message-solid"></i>お知らせ</p>
                  <div>
                    <div v-for="notification in notifications" :key="notification.id" class="notification-body-div">
                      <div class="user-icon">
                        <div class="notification-user-icon">
                          <!-- <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag">
                          </router-link> -->
                        </div>
                      </div>
                      <div v-if="notification.action === 'alike'">
                        <router-link class="a-tag-article" v-bind:to="{ name : 'ArticleShow', params : { id: notification.articleid.Int32 }}" >
                        <div><p class="notification-text">
                          <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                            <span class="visiter-name">{{notification.name}}</span></router-link> があなたの記事に「いいね！」しました。 <span class="notification-created">{{ notification.Created | moment }}</span></p>
                        </div></router-link>
                      </div>
                      <div v-else-if="notification.action === 'qlike'">
                        <div><p class="notification-text">
                          <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                            <span class="visiter-name">{{notification.name}}</span></router-link> があなたの質問に「いいね！」しました。</p></div>
                      </div>
                      <div v-else-if="notification.action === 'follow'">
                        <div><p class="notification-text">
                          <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                            <span class="visiter-name">{{notification.name}}</span></router-link> があなたをフォローしました。</p></div>
                      </div>
                      <div v-else-if="notification.action === 'aclike'">
                        <div><p class="notification-text">
                          <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                            <span class="visiter-name">{{notification.name}}</span></router-link> があなたのコメントに「いいね！」しました。</p></div>
                      </div>
                      <div v-else-if="notification.action === 'qclike'">
                        <div><p class="notification-text">
                          <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                            <span class="visiter-name">{{notification.name}}</span></router-link> があなたのコメントに「いいね！」しました。</p></div>
                      </div>
                      <div v-else-if="notification.action === 'acomment'">
                        <div><p class="notification-text">
                          <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                            <span class="visiter-name">{{notification.name}}</span></router-link> があなたの記事に「コメント」しました。</p>
                        </div>
                      </div>
                      <div v-else-if="notification.action === 'qcomment'">
                        <div><p class="notification-text">
                          <router-link v-bind:to="{ name : 'UserShow', params : { id: notification.userid }}" class="a-tag-n">
                            <span class="visiter-name">{{notification.name}}</span></router-link> があなたの質問に「回答」しました。</p></div>
                      </div>
                    </div>
                  </div>
                  <router-link to="/Notifications" class="a-tag2"><p class="text-center">通知一覧を見る</p></router-link>
                </div>
                <div v-else>
                  <p style="margin:0;"><i class="el-icon-message-solid"></i>お知らせ</p>
                  <p class="text-center">現在通知はございません。</p>
                </div>
                <el-button size="mini" slot="reference" @click="click_notification_btn"><i class="el-icon-message-solid message-solid-i"></i></el-button>
              </el-popover>
            </el-badge>

            <!-- <el-badge v-if="notificationCount === 0" :value="notificationCount" :hidden="true" :max="99" class="badge-all">
              <el-button size="mini"><i class="el-icon-message-solid message-solid-i"></i></el-button>
            </el-badge>
            <el-badge v-else :value="notificationCount" :max="99" class="badge-all">
              <el-button size="mini"><i class="el-icon-message-solid message-solid-i"></i></el-button>
            </el-badge> -->


          </li>
          <li class="a-tag logout-li" @click="myUserPage">
          <!-- <li class="a-tag logout-li"> -->
            <!-- <router-link v-bind:to="{ name : 'UserShow', params : { id: user.id }}" class="a-tag"> -->
              <div class="header-user-icon"></div>
            <!-- </router-link> -->
          </li>
        </ul>
      </nav>
    </div>
  </header>
</template>

<script>
import moment from 'moment'
export default {
  props: {
    loginpage: Boolean,
    isArticle: Boolean,
    isQuestion: Boolean,
    user: Object,
  },
  data(){
    return{
      notificationCount: localStorage.notificationCount,
      notifications: [],
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
    myUserPage: function() {
      this.$router.push({ name : 'UserShow', params : { id: this.user.id }});
      if(this.$router.currentRoute.path !== `/Users/${this.user.id}`){
        this.$router.go({ name : 'UserShow', params : { id: this.user.id }})
      }
    },
    click_notification_btn: function(){
    const jst = this.$cookies.get("JWT");
    this.$axios.get('http://localhost/api/restricted/Notifications',{
      params: {
        // ここにクエリパラメータを指定する
        sender: "header",
      },
      headers: {
        Authorization: `Bearer ${jst}`
      },
    })
      .then(response => {
        this.notifications = response.data.Notifications
        this.notificationCount = response.data.NotificationCount
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
.logout-li{
  margin-left: 20px;
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
  background-color: #ccc;
  width: 40px;
  height: 40px;
  border-radius: 50%;
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
  border-bottom: #ccc 1px solid;
  height: 60px;
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
  padding: 10px;
  color: #000;
}
.notification-text > .a-tag-n {
  text-decoration: none;
  cursor: pointer;
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
</style>
