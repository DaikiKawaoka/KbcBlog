<template>
  <div id="app">
    <Header :user="myUser"></Header>
    <user-form :errors="errors" :user="user" :edit="true" @submit="updateUser"></user-form>
    <Footer></Footer>
  </div>
</template>
<script>
// import axios from "axios";
import UserForm from '../components/UserForm.vue';
import Header from '../components/Header.vue';
import Footer from './../components/Footer.vue';

export default {
  data() {
    return {
      myUser: {
        KBCmail: "",
        id : 0,
        name: "",
        comment: {
          String: "",
          Valid: Boolean
        }
      },
      user:{
        KBCmail: "",
        id : 0,
        name: "",
        comment: {
          String: "",
          Valid: Boolean
        },
        github: {
          String: "",
          Valid: Boolean
        },
        website: {
          String: "",
          Valid: Boolean
        },
        languages: {
          String: "",
          Valid: Boolean
        },
      },
      errors: {},
    };
  },
  components: {
    UserForm,
    Header,
    Footer,
  },
  created () {
    this.$axios.get(`http://localhost/api/restricted/Users/${this.$route.params.id}/edit`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.myUser = response.data.MyUser
        this.user = response.data.User


        if(this.myUser.id !== this.user.id){
          this.$router.push({ path: "/" });
        }
        console.log(response)
      })
      .catch(error => {
        console.log(error)
        if(error.response.status == 401){
            this.$router.push({ path: "/login" });
        }
        if(error.response.status == 400){
          this.$router.push({ path: "/" });
        }
        this.errors = error.response.data.ValidationErrors;
      })
  },
  methods: {
    updateUser: function() {
      // commentが空かチェック
      if (this.user.comment.length !== 0){
        this.user.comment.Valid = true;
      }else{
        this.user.comment.Valid = false;
      }
      // githublinkが空かチェック
      if (this.user.github.String.length !== 0){
        this.user.github.Valid = true;
      }else{
        this.user.github.Valid = false;
      }
      // websitelinkが空かチェック
      if (this.user.website.String.length !== 0){
        this.user.website.Valid = true;
      }else{
        this.user.website.Valid = false;
      }
      // languagesが空かチェック
      if (this.user.languages.String.length !== 0){
        this.user.languages.Valid = true;
      }else{
        this.user.languages.Valid = false;
      }

      this.$axios
        .patch(`http://localhost/api/restricted/Users/${this.myUser.id}`,this.user,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          console.log(response);
          console.log("プロフィール編集成功");
          this.$router.push({ name: 'UserShow' , params : { id: this.user.id }});
        })
        .catch(error => {
          console.log(error.response.data.ValidationErrors);
          this.errors = error.response.data.ValidationErrors;
        });
    },
  }
};
</script>