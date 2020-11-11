<template>
  <div id="app">
    <Header :user="myUser"></Header>
    <user-form :errors="errors" :user="user" :edit="true" @submit="updateUser"></user-form>
  </div>
</template>
<script>
import axios from "axios";
import UserForm from '../components/UserForm.vue';
import Header from '../components/Header.vue';

export default {
  data() {
    return {
      myUser: null,
      user: null,
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
        if(this.myUser.id != this.user.id){
          this.$router.push({ path: "/" });
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
    updateUser: function() {
      axios
        .post('http://localhost/api/Users',this.user)
        .then(response => {
          // let e = response.data;
          console.log(response);
          console.log("user作成成功");
          this.$router.push({ name: 'LoginPage' });
        })
        .catch(error => {
          console.log(error.response.data.ValidationErrors);
          this.errors = error.response.data.ValidationErrors;
        });
    },
    // login_user: function() {
    //   axios
    //     .get('/api/v1/sessions.json')
    //     .then(response => {
    //       if (response.status !== 201){
    //         this.$router.push({ name: 'staticHome'})
    //       }
    //     })
    // },
  }
};
</script>