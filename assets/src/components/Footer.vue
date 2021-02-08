<template>
  <footer id="footer">
    <div class="footer-all">
      <div class="footer-body">
        <h1 class="footer-title">KBCBlog</h1>
        <div class="footer-right-div">
          <div class="menu">
            <router-link v-if="user !== null && user !== undefined " to="/About" class="a-tag"><p class="about-css">About</p></router-link>
            <p class="about-css" v-else @click="guestLogin">About</p>
            <!-- <router-link to="/Help" class="a-tag"><p class="about-css">Help</p></router-link> -->
          </div>
          <div class="btn">
            <el-button class="scrollbtn" type="success" icon="el-icon-arrow-up" circle @click="scrollTop"></el-button>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
import axios from "axios";
export default {
  props: {
    user: null,
    url: null,
  },
  methods: {
    guestLogin: function() {
      axios
        .get(this.url+'api/guestLogin')
        .then(response => {
          let token = response.data.token;
          this.$cookies.set('JWT',token,"10000h");
          this.$router.push({ path: "/About" });
        })
        .catch(error => {
          // if (error.response.data && error.response.data.errors) {
            this.errors = error.response.data;
          // }
        });
    },
    scrollTop: function(){
      window.scrollTo({
        top: 0,
        behavior: "smooth"
      });
    },
  }
}
</script>

<style lang="scss" scoped>
#footer{
  width: 100%;
  height: 150px;
  background-color: #555;
  .footer-all{
    width: 1000px;
    height: 150px;
    background: #555;
    margin: 0 auto 0 auto;
    .footer-body{
      display: flex;
    }
  }
}

.footer-title{
  margin: 0;
  padding-top: 30px;
  color: #fff;
  font-size: 50px;
  width: 50%;
}
.footer-right-div{
  // margin: 30px 0 0 50%;
  display: flex;
  width: 50%;
  margin-top: 30px;
  .menu{
    margin-left: 300px;
    .about-css{
      margin: 0 0 10px 0;
      font-size: 1.2em;
      color: #ddd;
      cursor: pointer;
      margin-right: 50px;
    }
  }
  .btn{
    // margin-left: 100px;
    .scrollbtn{
      font-size: 40px;
    }
  }
}

</style>
