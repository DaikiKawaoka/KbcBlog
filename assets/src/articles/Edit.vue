<template>
  <div id="app" v-if="user.id !== 920437694 && user.id !== undefined">
    <Header :isArticle="true" :isQuestion="false" :user="user" :url="url"></Header>
    <Article-form :article="article" :user="user" :errors="errors" :create="create" @submit="updateArticle" @cancell="goHome" :tagArray="tagArray"></Article-form>
    <Footer :user="user" :url="url"></Footer>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import Footer from './../components/Footer.vue';
import ArticleForm from './../components/ArticleForm.vue'

export default {
  name: 'app',
  data(){
    return {
      user: {},
      article : {},
      errors: [],
      create: false,
      tagArray: [],
      notificationCount: localStorage.notificationCount,
      url: null,
    }
  },
  components: {
    Header,
    Footer,
    ArticleForm
  },
  beforeRouteLeave (to, from, next) {
    this.$confirm('編集中のものは保存されませんが、よろしいですか？', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
        center: true
      }).then(() => {
        next()
      }).catch(() => {
        next(false)
      });
  },
  destroyed () {
    window.removeEventListener("beforeunload", this.confirmSave);
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.url = process.env.VUE_APP_URL
    window.addEventListener("beforeunload", this.confirmSave);
    this.$axios.get(this.url+`api/restricted/Articles/${this.$route.params.id}/edit`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.user = response.data.user
        if(this.user.id === 920437694){
          this.$router.push("/");
        }
        this.article = response.data.Article
        this.notificationCount = response.data.NotificationCount

        // 文字列のlangsを配列に変換
      const w = this.article.tag;
      this.tagArray = w.split(',');
      if(this.tagArray.length === 0){
        this.tagArray = [];
      }
      if(this.user.id != this.article.userid){
        this.$router.push({ path: "/" });
      }
      })
      .catch(error => {
        if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
        this.errors = error.response.data.ValidationErrors;
      })
  },
  methods: {
    updateArticle: function() {
      this.openFullScreen();
      this.$axios
        .patch(this.url+`api/restricted/Articles/${this.$route.params.id}`, this.article,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.closeFullScreen();
          this.$router.push({ path: `/Articles/${this.article.id}` });
          this.editArticleAlert();
        })
        .catch(error => {
          this.closeFullScreen();
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
          }
          this.errors = error.response.data.ValidationErrors;
          window.scrollTo({
            top: 0,
            behavior: "smooth"
          });
          this.openError()
        });
    },

    confirmSave (event) {
      event.returnValue = "編集中のものは保存されませんが、よろしいですか？";
    },

    openFullScreen() {
      this.loading = this.$loading({
          lock: true,
          text: 'Loading',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.95)'
        });
    },
    closeFullScreen() {
      this.loading.close();
    },

    openError() {
        this.$message({
          showClose: true,
          message: '入力項目に不備があります。',
          type: 'error'
        });
      },

    editArticleAlert() {
      this.$notify({
        title: 'Success',
        message: '記事の内容を変更しました。',
        type: 'success'
      });
    },

    goHome: function(){
      this.$router.push({ path: "/" });
    },
    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
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

</style>
