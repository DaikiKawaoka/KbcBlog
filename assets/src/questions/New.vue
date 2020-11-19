<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="true" :user="user"></Header>
    <Question-form :question="question" :errors="errors" :create="create" @submit="createQuestion" @cancell="goHome"></Question-form>
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
      user: {
        id: 0,
        name: "",
        KBC_mail: "",
      },
      question :{
        userid: 0,
        title: "",
        body: "",
      },
      errors: [],
      create: true,
    }
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/restricted/Questions/new',{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        console.log(response.data)
        this.user.id = response.data.user.id
        this.user.KBC_mail = response.data.user.KBC_mail
        this.user.name = response.data.user.name
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
        }
        console.log(error.response);
      })
  },
  components: {
    Header,
    QuestionForm,
    Footer,
  },
  methods: {
    createQuestion: function() {
      this.question.userid = this.user.id
      this.$axios
        .post('http://localhost/api/restricted/Questions', this.question,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.$router.push({ path: "/Questions" });
          this.createQuestionAlert();
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
          }
          this.errors = error.response.data.ValidationErrors;
        });
    },

    createQuestionAlert() {
      this.$message({
        message: '質問を投稿しました。',
        type: 'success'
      });
    },

    goHome: function(){
      this.$router.push({ path: "/" });
    }
  }
}
</script>

<style>

</style>