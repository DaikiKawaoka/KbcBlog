<template>
  <div id="app">
    <div class="markdown-all-div">
      <h2>タイトル</h2>
      <div>
        <span>{{article.title}}</span>
      </div>
      <h2>本文</h2>
      <div>
        <h3>src</h3>
        <span v-text="compiledMarkdown"></span>
      </div>
      <div>
        <h3>markdown</h3>
        <span>{{article.body}}</span>
      </div>
    </div>
  </div>
</template>

<script>
// import Header from './../components/Header.vue'
import marked from 'marked';

export default {
  name: 'app',
  data(){
    return {
      article : {
        body: "",
      },
      url: null,
    }
  },
  components: {
    // Header,
  },
  computed: {
    compiledMarkdown: function () {
      return marked(this.article.body)
    },
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.url = process.env.VUE_APP_URL
    this.openFullScreen()
    this.$axios.get(this.url+`api/restricted/Articles/${this.$route.params.id}/markdown`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },})
      .then(response => {
        this.article = response.data.Article
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