<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="false" :user="user"></Header>
    <Footer></Footer>
  </div>
</template>

<script>
// import axios from "axios";
import Header from './../components/Header.vue'
import Footer from './../components/Footer.vue'

export default {
  name: 'app',
  data(){
    return {
      user: {},
      notificationCount: localStorage.notificationCount,
    }
  },
  components: {
    Header,
    Footer,
  },
  // filters: {
  //   moment: function (date) {
  //     return moment(date).fromNow();
  //   },
  //   moment2: function (date) {
  //     // locale関数を使って言語を設定すると、日本語で出力される
  //     return moment(date).utc().format('YYYY/MM/DD');
  //   }
  // },
  watch: {
    searchText(newSearchText) {
      localStorage.searchText = newSearchText;
    },
    order(newOrder) {
      localStorage.order = newOrder;
    },
    tag(newTag) {
      localStorage.tag = newTag;
    },
    friendsOnly(newFriendsOnly) {
      localStorage.friendsOnly = newFriendsOnly;
    },
    notificationCount(newNotificationCount) {
      localStorage.notificationCount = newNotificationCount;
    },
  },

  created () {
    const jst = this.$cookies.get("JWT");
    this.setUp()
    this.$axios.get('http://localhost/api/restricted/Articles',{
      // params: {
      //   // ここにクエリパラメータを指定する
      //   // cursor: this.cursor,
      // },
      headers: {
        Authorization: `Bearer ${jst}`
      },
    })
      .then(response => {
        this.user = response.data.user
        // this.cursor = response.data.Cursor
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
  methods:{

    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },
  },
}
</script>

<style>
</style>
