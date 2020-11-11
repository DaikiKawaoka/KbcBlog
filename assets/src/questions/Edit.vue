<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="true" :user="user"></Header>
    <Question-form :question="question" :user="user" :errors="errors" :create="create" @submit="updateQuestion" @cancell="goHome"></Question-form>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import QuestionForm from './../components/QuestionForm.vue'

export default {
  name: 'app',
  data(){
    return {
      user: {},
      question : {},
      errors: [],
      create: false,
    }
  },
  components: {
    Header,
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
        this.question = response.data.Question
        if(this.user.id != this.question.userid){
          this.$router.push({ path: "/Questions" });
        }
        console.log(response)
      })
      .catch(error => {
        console.log(error)
        if(error.response.status == 401){
            this.$router.push({ path: "/login" });
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
        .then(response => {
          this.$router.push({ path: `/Questions/${this.question.id}` });
          console.log(response)
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
          }
          this.errors = error.response.data.ValidationErrors;
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
