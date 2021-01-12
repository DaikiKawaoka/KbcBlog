<template>
  <div id="app">
    <Header :loginpage="true"></Header>
    <user-form :errors="errors" :user="user" @submit="createUser"></user-form>
    <Footer></Footer>
  </div>
</template>
<script>
import axios from "axios";
import UserForm from '../components/UserForm.vue';
import Header from '../components/Header.vue';
import Footer from './../components/Footer.vue';

export default {
  components: {
    UserForm,
    Header,
    Footer,
  },
  created () {
    // this.login_user();
  },
  data() {
    return {
      user: {
        name: '',
        KBC_mail: '',
        password: '',
        password_confirmation: '',
        sex: '',
      },
      errors: []
    };
  },
  methods: {
    createUser: function() {
      axios
        .post('http://localhost/api/Users',this.user)
        .then(() => {
          console.log("user作成成功");
          this.$router.push({ name: 'LoginPage' });
        })
        .catch(error => {
          this.errors = error.response.data.ValidationErrors;
          window.scrollTo({
            top: 0,
            behavior: "smooth"
          });
        });
    },
  }
};
</script>