<template>
  <div id="login-all">
    <Header :loginpage="loginpage"></Header>
    <div class="body-main">
      <div v-if="errors.length != 0">
        <ul class="error-ul" v-for="e in errors" :key="e">
            <li class="error-icon-li"><i class="el-icon-warning-outline"></i></li>
            <li><font color="red">{{ e }}</font></li>
        </ul>
      </div>

      <el-card class="box-card login">

        <div slot="header" class="clearfix">
          <span>Login</span>
          <el-button style="float: right; padding: 3px 0" type="text" @click="signup">Signup</el-button>
        </div>

        <el-form label-width="80px">
          <el-form-item label="KBC_mail">
            <el-input type="email" placeholder="河原学園のメールアドレス" v-model="KBC_mail"></el-input>
          </el-form-item>
          <el-form-item label="Password">
            <el-input type="password" placeholder="パスワード" v-model="password"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button style="float: right" type="primary" @click="login">Login</el-button>
          </el-form-item>
        </el-form>

      </el-card>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Header from './../components/Header.vue'

export default {
  data() {
    return {
      KBC_mail: '',
      password: '',
      errors: '',
      loginpage: true
    };
  },
  components: {
    Header,
  },
  methods: {
    signup: function(){
      this.$router.push("/Users/new");
    },
    login: function() {
      axios
        .post('http://localhost/api/Login',{password: this.password, KBC_mail: this.KBC_mail})
        .then(response => {
          let token = response.data.token;
          this.$cookies.set('JWT',token,"1h");
          this.$router.push({ path: "/" });
        })
        .catch(error => {
          // if (error.response.data && error.response.data.errors) {
            this.errors = error.response.data;
          // }
        });
    }
  }
}
</script>

<style>
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
</style>