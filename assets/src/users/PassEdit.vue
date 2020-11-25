<template>
  <div id="app">
    <Header :user="myUser"></Header>
    <div class="passedit-all-div">
      <h1>パスワード変更</h1>
      <div class="passedit-main-div">
        <el-form>
          <el-form-item label="現在のパスワード">
              <el-input type="password" v-model="passwords.currentPassword" autocomplete="off" show-password></el-input>
          </el-form-item>

          <el-form-item label="新しいパスワード">
              <el-input type="password" v-model="passwords.newPassword" autocomplete="off" show-password></el-input>
          </el-form-item>

          <el-form-item label="新しいパスワードを確認">
              <el-input type="password" v-model="passwords.passwordConfirmation" autocomplete="off" show-password></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="passwordEdit">パスワード変更</el-button>
          </el-form-item>

        </el-form>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>
<script>
// import axios from "axios";
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
      user:{},
      passwords:{
        currentPassword:"",
        newPassword:"",
        passwordConfirmation:"",
      },
      errors: {},
    };
  },
  components: {
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
    passwordEdit: function() {
      this.$axios
        .patch(`http://localhost/api/restricted/Users/${this.myUser.id}/password/edit`,this.passwords,{
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
        });
    },

    editUserAlert() {
      this.$notify({
        title: 'Success',
        message: 'パスワードを変更しました。',
        type: 'success'
      });
    },

    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },
  }
};
</script>

<style scoped>
h1{
  text-align: center;
}
.passedit-all-div{
  width: 1000px;
  margin: 0 auto 100px auto;
}
.passedit-main-div{
  width: 500px;
  margin: 0 auto 0 auto;
}

</style>