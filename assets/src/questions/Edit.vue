<template>
  <div id="app" v-if="user.id !== 920437694 && user.id !== undefined">
    <Header :isArticle="false" :isQuestion="true" :user="user" :url="url"></Header>
    <Question-form :question="question" :user="user" :errors="errors" :create="create" @submit="updateQuestion" @cancell="goHome"></Question-form>
    <Footer :user="user" :url="url"></Footer>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import Footer from './../components/Footer.vue';
import QuestionForm from './../components/QuestionForm.vue'

export default {
  name: 'app',
  data(){
    return {
      user: {},
      question : {},
      errors: [],
      create: false,
      notificationCount: localStorage.notificationCount,
      url: null,
    }
  },
  components: {
    Header,
    Footer,
    QuestionForm
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    window.addEventListener("beforeunload", this.confirmSave);
    this.url = process.env.VUE_APP_URL
    this.$axios.get(this.url+`api/restricted/Questions/${this.$route.params.id}/edit`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.user = response.data.user
        if(this.user.id === 920437694){
          this.$router.push("/");
        }
        this.question = response.data.Question
        this.notificationCount = response.data.NotificationCount
        if(this.user.id != this.question.userid){
          this.$router.push({ path: "/Questions" });
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
  methods: {
    updateQuestion: function() {
      this.openFullScreen();
      this.$axios
        .patch(this.url+`api/restricted/Questions/${this.$route.params.id}`, this.question,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.closeFullScreen();
          this.$router.push({ path: `/Questions/${this.question.id}` });
          this.editQuestionAlert();
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

    editQuestionAlert() {
      this.$notify({
        title: 'Success',
        message: '質問の内容を変更しました。',
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

    confirmSave (event) {
      event.returnValue = "編集中のものは保存されませんが、よろしいですか？";
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
