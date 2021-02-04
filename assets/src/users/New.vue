<template>
  <div id="app">
    <Header :loginpage="true" :url="url"></Header>
    <user-form :errors="errors" :user="user" @submit="createUser"></user-form>
    <Footer :user="user" :url="url"></Footer>
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
      errors: [],
      url: null,
    };
  },
  methods: {
    createUser: function() {
      this.url = process.env.VUE_APP_URL
      axios
        .post(this.url+'api/Users',this.user)
        .then(() => {
          this.open()
          this.$router.push({ name: 'LoginPage' });
        })
        .catch(error => {
          this.errors = error.response.data.ValidationErrors;
          if(error.response.status == 500){
            if(this.errors === null || this.errors.length === 0){
              this.errors = ["このメールアドレスは既に使われています。"]
            }else{
              this.errors.push("このメールアドレスは既に使われています。")
            }
          }
          window.scrollTo({
            top: 0,
            behavior: "smooth"
          });
        });
    },
    open() {
        this.$notify.success({
          title: 'アカウントを作成できました！',
          message: 'メールアドレスとパスワードを入力してログインしてください。',
          offset: 100
        });
      },
  }
};
</script>