<template>
  <div id="login-all">
    <Header></Header>
    <el-card class="box-card login">

      <div slot="header" class="clearfix">
        <span>Login</span>
        <el-button style="float: right; padding: 3px 0" type="text">Forget password</el-button>
      </div>

      <el-form label-width="80px">
        <el-form-item label="KBC_mail">
          <el-input type="email" v-model="KBC_mail"></el-input>
        </el-form-item>
        <el-form-item label="Password">
          <el-input type="password" v-model="password"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button style="float: right" type="primary" @click="loginUser">Login</el-button>
        </el-form-item>
      </el-form>

    </el-card>
  </div>
</template>

<script>
// import axios from "axios";
import Header from './../components/Header.vue'

export default {
  data() {
    return {
      KBC_mail: '',
      password: '',
      errors: ''
    };
  },
  components: {
    Header,
  },
  methods: {
    loginUser: function() {
      axios
        .post('/api/v1/sessions',{password: this.password, email: this.email})
        .then(response => {
          let e = response.data;
          this.$router.push({ name: 'staticHome'});
        })
        .catch(error => {
          console.error(error);
          if (error.response.data && error.response.data.errors) {
            this.errors = error.response.data.errors;
          }
        });
    }
  }
}
</script>

<style>
#login-all{

}
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