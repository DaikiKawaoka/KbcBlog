<template>
  <div id="app">
    <Header></Header>
    <user-form :errors="errors" :user="user" @submit="createUser"></user-form>
  </div>
</template>
<script>
import axios from "axios";
import UserForm from '../components/UserForm.vue';
import Header from '../components/Header.vue';

export default {
  components: {
    UserForm,
    Header,
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
       },
       errors: ''
    };
  },
  methods: {
    createUser: function() {
      axios
        .post('http://localhost/api/Users',this.user)
        .then(response => {
          // let e = response.data;
          console.log(response);
          console.log("user作成成功");
          // this.$router.push({ name: 'userShow', params: { id: e.id } });
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