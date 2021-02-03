<template>
  <div id="app">
    <Header :loginpage="loginpage"></Header>
    <div class="body-main-login">

      <el-card class="box-card login">

        <div slot="header" class="clearfix">
          <span>ログイン</span>

          <div v-if="errors.length != 0" class="err-div-login">
            <ul class="error-ul" v-for="e in errors" :key="e">
              <span class="error-span">
                <li class="error-icon-li"><i class="el-icon-warning-outline err-i"></i></li>
                <li><font color="red" class="err-text">{{ e }}</font></li>
              </span>
            </ul>
          </div>

          <el-button style="float: right; padding: 3px 0" type="text" @click="signup">新規作成</el-button>
          <el-button style="float: right; padding: 3px 0; margin-right: 10px;" type="text" @click="guestLogin">ゲストログイン</el-button>
        </div>

        <el-form label-width="80px" ref="loginForm" :model="loginUser">
          <el-form-item label="KBC_mail" prop="KBC_mail"
          :rules="[
              { required: true , pattern: /[\w\-\._]+@(stu.)?kawahara.ac.jp/, message: '河原学園のメールアドレスを入力してください。', trigger: 'blur'},
          ]">
            <el-input type="email" placeholder="河原学園のメールアドレス" v-model="loginUser.KBC_mail"></el-input>
          </el-form-item>
          <el-form-item label="Password" prop="password"
      :rules="[
          { required: true, message: '入力必須です', trigger: 'blur' },
          { min: 8, max: 50, message: '8~50文字で入力してください', trigger: 'blur' },
      ]">
      <!-- { pattern: /^(?=.*?[a-z])(?=.*?\d)[a-z\d]{8,100}$/i, message: '半角英字と半角数字それぞれ1文字以上含めてください。', trigger: 'blur' }, -->
            <el-input type="password" placeholder="パスワード" v-model="loginUser.password"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button style="float: right" type="primary" @click="loginForm('loginForm')">ログイン</el-button>
          </el-form-item>
        </el-form>

      </el-card>
    </div>
    <Footer :user="null"></Footer>
  </div>
</template>

<script>
import axios from "axios";
import Header from './../components/Header.vue';
import Footer from './../components/Footer.vue';

export default {
  data() {
    return {
      loginUser:{
        KBC_mail: '',
        password: '',
      },
      errors: '',
      loginpage: true
    };
  },
  components: {
    Header,
    Footer,
  },
  created(){
    document.title = `KBC Blog`;
  },
  methods: {
    signup: function(){
      this.$router.push("/Users/new");
    },
    login: function() {
      axios
        .post('http://localhost/api/Login',{password: this.loginUser.password, KBC_mail: this.loginUser.KBC_mail})
        .then(response => {
          let token = response.data.token;
          this.$cookies.set('JWT',token,"10000h");
          this.$router.push({ path: "/" });
        })
        .catch(error => {
          // if (error.response.data && error.response.data.errors) {
            this.errors = error.response.data;
          // }
        });
    },
    guestLogin: function() {
      axios
        .get('http://localhost/api/guestLogin')
        .then(response => {
          let token = response.data.token;
          this.$cookies.set('JWT',token,"10000h");
          this.$router.push({ path: "/" });
        })
        .catch(error => {
          // if (error.response.data && error.response.data.errors) {
            this.errors = error.response.data;
          // }
        });
    },
    loginForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.login()
          } else {
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
  }
}
</script>

<style lang="scss" scoped>
.body-main-login{
  height: 450px;
  position:relative;
  .box-card {
  width: 480px;
}
  .login {
    position: relative;
    top: 100px;
    right: 0px;
    bottom: 0px;
    left: 0px;
    margin: auto;
  }
  .err-div-login{
    position:absolute;
    top: 200px;
    left: -20px;
    z-index: 2;
  }
  .err-i{
    font-size: 18px;
    color:#F56C6C;
  }
  li{
    list-style: none;
  }
  .error-span{
    display: flex;
    text-align: center;
  }
  .err-text{
    font-size: 0.9em;
  }
}
</style>