<template>
  <div id="app" v-if="user.id !== 920437694 && user.id !== 0">
    <Header :isArticle="true" :isQuestion="false" :user="user" :url="url"></Header>
    <Article-form :article="article" :errors="errors" :create="create" @submit="createArticle" @cancell="goHome"></Article-form>
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
      user: {
        id: 0,
        name: "",
        KBC_mail: "",
      },
      article :{
        userid: 0,
        title: "",
        body: "",
        tag: "",
      },
      errors: [],
      create: true,
      notificationCount: localStorage.notificationCount,
      url: null,
    }
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.url = process.env.VUE_APP_URL
    this.$axios.get(this.url+'api/restricted/Articles/new',{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.user = response.data.user
        if(this.user.id === 920437694){
          this.$router.push("/");
        }
        this.article.userid = this.user.id
        this.notificationCount = response.data.NotificationCount
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
  },
  components: {
    Header,
    Footer,
    ArticleForm
  },
  methods: {
    createArticle: function() {
      this.$axios
        .post(this.url+'api/restricted/Articles', this.article,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.$router.push({ path: "/" });
          this.createArticleAlert();
        })
        .catch(error => {
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

    openError() {
        this.$message({
          showClose: true,
          message: '入力項目に不備があります。',
          type: 'error'
        });
      },

    createArticleAlert() {
      this.$notify({
        title: 'Success',
        message: '記事を投稿しました。',
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
