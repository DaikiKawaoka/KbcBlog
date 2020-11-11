<template>
  <div id="app">
    <Header :user="myUser"></Header>
    <user-form :errors="errors" :user="myUser" :edit="true" @submit="updateUser"></user-form>
  </div>
</template>
<script>
// import axios from "axios";
import UserForm from '../components/UserForm.vue';
import Header from '../components/Header.vue';

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
      user:{},
      errors: {}
    };
  },
  components: {
    UserForm,
    Header,
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
      if (this.myUser.comment.String.length !== 0){
        this.myUser.comment.Valid = true;
      }
      this.$axios
        .patch(`http://localhost/api/restricted/Users/${this.myUser.id}`,this.myUser,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(response => {
          console.log(response);
          console.log("プロフィール編集成功");
          this.$router.push({ name: 'UserShow' , params : { id: this.myUser.id }});
        })
        .catch(error => {
          console.log(error.response.data.ValidationErrors);
          this.errors = error.response.data.ValidationErrors;
        });
    },
  }
};
</script>