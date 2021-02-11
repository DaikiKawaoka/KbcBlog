<template>
  <div id="app" v-if="user.id !== 920437694 && user.id !== undefined ">
    <Header :user="myUser" :url="url"></Header>
    <div class="passedit-all-div">
      <h1>パスワード変更</h1>
      <div class="passedit-main-div">

        <div v-if="errors.length != 0">
          <ul class="error-ul" v-for="e in errors" :key="e">
              <li class="error-icon-li"><i class="el-icon-warning-outline"></i></li>
              <li><font color="red">{{ e }}</font></li>
          </ul>
        </div>

        <el-form :model="passwords" ref="passForm">
          <el-form-item label="現在のパスワード" prop="currentPassword"
              :rules="[
              { required: true, message: '入力必須です', trigger: 'blur' },
              { min: 8, max: 50, message: '8~50文字で入力してください', trigger: 'blur' },
              { pattern: /^(?=.*?[a-z])(?=.*?\d)[a-z\d]{8,100}$/i, message: '半角英字と半角数字それぞれ1文字以上含めてください。', trigger: 'blur' },
          ]">
              <el-input type="password" v-model="passwords.currentPassword" autocomplete="off" show-password></el-input>
          </el-form-item>

          <el-form-item label="新しいパスワード" prop="newPassword"
          :rules="[
              { required: true, message: '入力必須です', trigger: 'blur' },
              { min: 8, max: 50, message: '8~50文字で入力してください', trigger: 'blur' },
              { pattern: /^(?=.*?[a-z])(?=.*?\d)[a-z\d]{8,100}$/i, message: '半角英字と半角数字それぞれ1文字以上含めてください。', trigger: 'blur' },
          ]">
              <el-input type="password" v-model="passwords.newPassword" autocomplete="off" show-password></el-input>
          </el-form-item>

          <el-form-item label="新しいパスワードを確認" prop="passwordConfirmation"
          :rules="[
              { required: true, message: '入力必須です', trigger: 'blur' },
              { min: 8, max: 50, message: '8~50文字で入力してください', trigger: 'blur' },
              { pattern: /^(?=.*?[a-z])(?=.*?\d)[a-z\d]{8,100}$/i, message: '半角英字と半角数字それぞれ1文字以上含めてください。', trigger: 'blur' },
          ]">
              <el-input type="password" v-model="passwords.passwordConfirmation" autocomplete="off" show-password></el-input>
          </el-form-item>

          <br>
          <el-form-item>
            <el-button type="primary" @click="passForm('passForm')">パスワード変更</el-button>
          </el-form-item>

        </el-form>
      </div>
    </div>
    <Footer :user="myUser" :url="url"></Footer>
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
      notificationCount: localStorage.notificationCount,
      url: null,
    };
  },
  components: {
    Header,
    Footer,
  },
  created () {
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
        if(this.myUser.id !== this.user.id){
          this.$router.push({ path: "/" });
        }
          this.notificationCount = response.data.NotificationCount
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
    passForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.passwordEdit();
          } else {
            console.log("aaa")
            this.openError()
            return false;
          }
        });
      },
      openError() {
        this.$message({
          showClose: true,
          message: '入力項目に不備があります。',
          type: 'error'
        });
      },
    passwordEdit: function() {
      this.openFullScreen();
      if(this.passwords.newPassword === this.passwords.passwordConfirmation){
        this.$axios
          .patch(this.url+`api/restricted/Users/${this.myUser.id}/password/edit`,this.passwords,{
            headers: {
              Authorization: `Bearer ${this.$cookies.get("JWT")}`
            },
          })
          .then(() => {
            this.closeFullScreen();
            this.$router.push({ name: 'UserShow' , params : { id: this.user.id }});
            this.editUserAlert();
          })
          .catch(error => {
            this.closeFullScreen();
            this.errors = error.response.data.ValidationErrors;
          });
        }else{
          this.closeFullScreen();
          this.errors = ["パスワードと確認パスワードが異なります。"];
        }
    },

    editUserAlert() {
      this.$notify({
        title: 'Success',
        message: 'パスワードを変更しました。',
        type: 'success'
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

<style scoped>
h1{
  text-align: center;
}
.passedit-all-div{
  width: 1000px;
  padding: 15px 0 40px 0;
  background-color: #F6F6F4;
  margin: 0 auto;
}
.passedit-main-div{
  width: 500px;
  height: 450px;
  margin: 0 auto 0 auto;
}
i{
  font-size: 25px;
  color:#F56C6C;
}
li{
  list-style: none;
}
.error-icon-li{
  padding-right: 10px;
}
.error-ul{
  display: flex;
}

</style>