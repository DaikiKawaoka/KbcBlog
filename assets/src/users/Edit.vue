<template>
  <div id="app" v-if="user.id !== 920437694 && user.id !== 0 ">
    <Header :user="myUser" :url="url"></Header>
    <user-form :errors="errors" :user="user" :edit="true" @submit="updateUser"></user-form>
    <Footer :user="myUser" :url="url"></Footer>
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
      url: null,
      isEvent: true,
    };
  },
  components: {
    UserForm,
    Header,
    Footer,
  },
  created () {
    window.addEventListener("beforeunload", this.confirmSave);
    this.url = process.env.VUE_APP_URL
    this.$axios.get(this.url+`api/restricted/Users/${this.$route.params.id}/edit`,{
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
  beforeRouteLeave (to, from, next) {
    if(this.isEvent){
      this.$confirm('編集中のものは保存されませんが、よろしいですか？', 'Warning', {
            confirmButtonText: 'OK',
            cancelButtonText: 'Cancel',
            type: 'warning',
            center: true
          }).then(() => {
            next()
          }).catch(() => {
            next(false)
          });
    }else{
      next();
    }
  },
  destroyed () {
    window.removeEventListener("beforeunload", this.confirmSave);
  },
  methods: {
    updateUser: function() {
      this.openFullScreen();
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
        .patch(this.url+`api/restricted/Users/${this.myUser.id}`,this.user,{
          headers: {
            Authorization: `Bearer ${this.$cookies.get("JWT")}`
          },
        })
        .then(() => {
          this.closeFullScreen();
          this.isEvent = false;
          this.$router.push({ name: 'UserShow' , params : { id: this.user.id }});
          this.editUserAlert();
        })
        .catch(error => {
          this.closeFullScreen();
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
    openFullScreen() {
      this.loading = this.$loading({
          lock: true,
          text: 'Loading',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.95)'
        });
    },
    closeFullScreen() {
      this.loading.close();
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

    confirmSave (event) {
      event.returnValue = "編集中のものは保存されませんが、よろしいですか？";
    },
  },
  watch: {
    notificationCount(newNotificationCount) {
      localStorage.notificationCount = newNotificationCount;
    },
  },
};
</script>