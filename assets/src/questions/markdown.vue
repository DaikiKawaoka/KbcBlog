<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="true" :user="user"></Header>
    <div class="markdown-all-div">
      <div>
        <h4>タイトル</h4>
        <span>{{question.title}}</span>
      </div>
      <div>
        <h4>本文</h4>
        <span>{{compiledMarkdown}}</span>
      </div>
    </div>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import marked from 'marked';

export default {
  name: 'app',
  data(){
    return {
      user: {},
      question : {
        body:"",
      },
    }
  },
  components: {
    Header,
  },
  computed: {
    compiledMarkdown: function () {
      return marked(this.question.body)
    },
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.openFullScreen()
    this.$axios.get(`http://localhost/api/restricted/Questions/${this.$route.params.id}`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },})
      .then(response => {
        this.question = response.data.Question
        this.user = response.data.user
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
  methods: {
    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
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
  },
  watch: {
    notificationCount(newNotificationCount) {
      localStorage.notificationCount = newNotificationCount;
    },
  },

}
</script>

<style scoped>
#app{
  background-color: #fff;
}
.markdown-all-div{
  width: 1000px;
  margin: 0 auto 0 auto;
  background-color: #fff;
}
</style>