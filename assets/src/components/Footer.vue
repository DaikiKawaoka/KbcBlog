<template>
  <footer id="footer">
    <div class="footer-all">
      <h1 class="footer-title">KBCBlog</h1>
      <div class="footer-right-div">
        <div class="menu">
          <router-link to="/About" class="a-tag"><p class="about-css">About</p></router-link>
          <!-- <router-link to="/Help" class="a-tag"><p class="about-css">Help</p></router-link> -->
        </div>
        <div class="btn">
          <el-button class="scrollbtn" type="success" icon="el-icon-arrow-up" circle @click="scrollTop"></el-button>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>

export default {
  methods: {
    scrollTop: function(){
      window.scrollTo({
        top: 0,
        behavior: "smooth"
      });
    }
  }
}
</script>

<style lang="scss" scoped>
#footer{
  width: 100%;
  height: 200px;
  background-color: #555;
  .footer-all{
  margin: 80px auto 0 auto;
  width: 1000px;
  display: flex;
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
    margin-left: 200px;
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
