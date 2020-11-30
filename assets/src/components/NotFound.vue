<template>
  <div id="login-all">
    <Header :loginpage="loginpage"></Header>
      <h1>NotFound</h1>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import Footer from './../components/Footer.vue';

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
    Footer,
  },
}
</script>

<style>

</style>