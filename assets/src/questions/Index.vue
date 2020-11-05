<template>
  <div id="app">
    <Header :isArticle="false" :isQuestion="true"></Header>
    <div class="index-main body-main">
      <div class="index-menu">
        <div></div>
      </div>
      <div class="question-all-div">
        <div v-for="question in questions" :key="question.id" class="question-show-div">
          <div class="question-user-icon"></div>
          <div class="question-index-body">
            <router-link v-bind:to="{ name : 'QuestionShow', params : { id: question.id }}" class="a-tag">
              <h2 class="question-title-index"> {{question.title}}</h2>
            </router-link>
            <h3 class="question-user-name">{{ question.name }}</h3>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import axios from "axios";
import Header from './../components/Header.vue'

export default {
  name: 'app',
  data(){
    return {
      questions: [],
      user: {},
      type: "question"
    }
  },
  components: {
    Header
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/restricted/Questions',{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.questions = response.data.Questions
        this.user = response.data.user
        console.log(response.data)
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
        }
        console.log(error.response);
      })
  },
}
</script>

<style>
body {
  margin: 0px;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin: 0 auto 0 auto;
  padding-bottom: 30px;
  background-color: #F6F6F4;
}
.body-main{
  width: 1000px;
  margin: 0 auto 0 auto;
}
.index-main{
  display: flex;
  margin-top: 30px;
}
.index-menu{
  width: 300px;
  background-color: #F6F6F4;
}
.question-all-div{
  width: 600px;
  background-color: #fff;
}
.question-show-div{
  padding-left: 20px;
  border: solid 1px #eee;
  display: flex;
}
.question-user-icon{
  margin-top: 25px;
  background-color: #ccc;
  width: 40px;
  height: 40px;
  border-radius: 50%;
}
.question-index-body{
  margin-left: 20px;
}
.question-title-index{
  font-size: 20px;
  width: 500px;
  text-align: left;
}
.question-user-name{
  font-size: 13px;
}
.a-tag{
  color: #000;
  text-decoration: none;
  cursor: pointer;
}
</style>
