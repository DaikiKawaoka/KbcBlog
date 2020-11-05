<template>
  <header id="header">
    <div class="body-main header-body-main-div">
      <div class="header-left">
        <h1 class="title">KBC Blog</h1>
        <h3 v-if="!loginpage" class="header-article">記事</h3>
        <h3 v-if="!loginpage" class="header-question">質問</h3>
      </div>
      <nav v-if="!loginpage" id="header-nav">
        <ul>
          <li><router-link tag="li" id="home-nav" to="/" class="a-tag">ホーム</router-link></li>
          <li><router-link tag="li" id="gift-nav" to="/Articles/new" class="a-tag">作成</router-link></li>
          <li class="a-tag" @click="logout">ログアウト</li>
        </ul>
      </nav>
    </div>
  </header>
</template>

<script>

export default {
  props: {
    loginpage: Boolean,
  },
  methods: {
    logout: function() {
      this.$cookies.remove("JWT");
      this.$router.push({ path: "/login" });
    }
  }
}
</script>

<style scoped>
#header{
  height: 60px;
  background: #7dec61;
}
.header-body-main-div{
  display: flex;
}
.header-left{
  margin: 0;
  width: 50%;
  display: flex;
}
.title{
  line-height: 60px;
  margin: 0;
}
.header-article,.header-question{
  margin-left: 30px;
  margin-bottom: 10px;
  padding: 0 12px 7px 12px;
  color: #606266;
  cursor: pointer;
}
.header-article{
  border-bottom: 3px solid #1100ff;
}
#header-nav{
  width: 50%;
}
ul{
  margin: 0;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  list-style: none;
  text-align: right;
}
li{
  display: inline;
}
.a-tag{
  line-height: 60px;
  margin-right: 20px;
  font-size: 15px;
  text-align: right;
  color: #606566;
  text-decoration: none;
  cursor: pointer;
}
</style>
