<template>
  <footer id="footer">
    <div class="footer-all">
      <h1 class="footer-title">KBCBlog</h1>
      <div class="footer-right-div">
        <el-button class="scrollbtn" type="success" icon="el-icon-arrow-up" circle @click="scrollTop"></el-button>
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

<style scoped>
#footer{
  width: 100%;
  height: 200px;
  background-color: #555;
}
.footer-all{
  margin: 80px auto 0 auto;
  width: 1000px;
  display: flex;
}
.footer-title{
  margin: 0;
  padding-top: 30px;
  color: #fff;
  font-size: 50px;
}
.footer-right-div{
  margin: 30px 0 0 65%;
}
.scrollbtn{
  font-size: 40px;
}
</style>
