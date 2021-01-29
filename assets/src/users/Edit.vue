<template>
  <div id="app" v-if="user.id !== 920437694 && user.id !== 0 ">
    <Header :user="myUser"></Header>
    <user-form :errors="errors" :user="user" :edit="true" @submit="updateUser"></user-form>
    <Footer :user="myUser"></Footer>
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
      errors: [],
      notificationCount: localStorage.notificationCount,
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
        if(this.user.id === 920437694){
          this.$router.push("/");
        }
        this.notificationCount = response.data.NotificationCount
        if(this.myUser.id !== this.user.id){
          this.$router.push({ path: "/" });
        }
      })
      .catch(error => {
        if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
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
        .then(() => {
          this.$router.push({ name: 'UserShow' , params : { id: this.user.id }});
          this.editUserAlert();
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

    editUserAlert() {
      this.$notify({
        title: 'Success',
        message: 'プロフィールを変更しました。',
        type: 'success'
      });
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
};
</script>