<template>
  <div id="app" v-if="user.id !== 920437694 && user.id !== 0 ">
    <Header :isArticle="false" :isQuestion="true" :user="user" :url="url"></Header>
    <Question-form :question="question" :errors="errors" :create="create" @submit="createQuestion" @cancell="goHome"></Question-form>
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
      user: {
        id: 0,
        name: "",
        KBC_mail: "",
      },
      question :{
        userid: 0,
        title: "",
        body: "",
        tag: "",
        category: "",
      },
      errors: [],
      create: true,
      notificationCount: localStorage.notificationCount,
      url: null,
      isEvent: true,
    }
  },
  created () {
    window.addEventListener("beforeunload", this.confirmSave);
    this.url = process.env.VUE_APP_URL
    this.$axios.get(this.url+'api/restricted/Questions/new',{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.user = response.data.user
        if(this.user.id === 920437694){
          this.$router.push("/");
        }
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
    QuestionForm,
    Footer,
  },
  beforeRouteLeave (to, from, next) {
    if(this.isEvent){
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
    }else{
      next();
    }
  },
  destroyed () {
    window.removeEventListener("beforeunload", this.confirmSave);
  },
  methods: {
    createQuestion: function() {
      this.openFullScreen();
      this.question.userid = this.user.id
      this.$axios
        .post(this.url+'api/restricted/Questions', this.question,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.closeFullScreen();
          this.isEvent = false;
          this.$router.push({ path: "/Questions" });
          this.createQuestionAlert();
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

    createQuestionAlert() {
      this.$notify({
        title: 'Success',
        message: '質問を投稿しました。',
        type: 'success'
      });
    },

    openError() {
        this.$message({
          showClose: true,
          message: '入力項目に不備があります。',
          type: 'error'
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