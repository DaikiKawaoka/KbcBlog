<template>
  <header id="header">
    <h1>KBC Blog</h1>
    <nav id="header-nav">
      <ul>
        <li><router-link tag="li" id="home-nav" to="/" class="a-tag">ホーム</router-link></li>
        <li><router-link tag="li" id="gift-nav" to="/Articles/new" class="a-tag">作成</router-link></li>
      </ul>
    </nav>
  </header>
</template>

<script>
export default {
  props: {
  }
}
</script>

<style scoped>
#header{
  height: 60px;
  background: #ffb;
  display: flex;
}
h1{
  margin: 0;
  line-height: 60px;
}
ul{
  margin: 0;
  width: 500px;
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
