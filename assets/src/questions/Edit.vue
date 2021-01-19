<template>
  <div id="app" v-if="user.id !== 920437694 && user.id !== undefined">
    <Header :isArticle="false" :isQuestion="true" :user="user"></Header>
    <Question-form :question="question" :user="user" :errors="errors" :create="create" @submit="updateQuestion" @cancell="goHome"></Question-form>
    <Footer></Footer>
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
    }
  },
  components: {
    Header,
    Footer,
    QuestionForm
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get(`http://localhost/api/restricted/Questions/${this.$route.params.id}/edit`,{
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
  methods: {
    updateQuestion: function() {
      this.$axios
        .patch(`http://localhost/api/restricted/Questions/${this.$route.params.id}`, this.question,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.$router.push({ path: `/Questions/${this.question.id}` });
          this.editQuestionAlert();
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
