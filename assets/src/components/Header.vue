<template>
  <header id="header">
    <div class="header-left">
      <h1 class="title">KBC Blog</h1>
      <h3 class="header-article">記事</h3>
      <h3 class="header-question">質問</h3>
    </div>
    <nav id="header-nav">
      <ul>
        <li><router-link tag="li" id="home-nav" to="/" class="a-tag">ホーム</router-link></li>
        <li><router-link tag="li" id="gift-nav" to="/Articles/new" class="a-tag">作成</router-link></li>
        <li class="a-tag" @click="logout">ログアウト</li>
      </ul>
    </nav>
  </header>
</template>

<script>

export default {
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
  display: flex;
}
.header-left{
  margin: 0;
  padding-left: 20px;
  width: 50%;
  display: flex;
}
.title{
  line-height: 60px;
  margin: 0;
}
.header-article,.header-question{
  padding-left: 30px;
  color: #606266;
  cursor: pointer;
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
  color: #606266;
  text-decoration: none;
  cursor: pointer;
}
</style>
