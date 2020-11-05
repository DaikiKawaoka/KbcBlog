<template>
  <header id="header">
    <div class="body-main header-body-main-div">
      <div class="header-left">
        <router-link to="/" class="a-tag"><h1 class="title">KBC Blog</h1></router-link>
        <router-link to="/" class="a-tag2"><h3 v-if="!loginpage" class="header-article" v-bind:class="{ active: isArticle }">記事</h3></router-link>
        <router-link to="/Questions" class="a-tag2"><h3 v-if="!loginpage" class="header-question" v-bind:class="{ active: isQuestion }">質問</h3></router-link>
      </div>
      <nav v-if="!loginpage" id="header-nav">
        <ul>
          <!-- <li><router-link tag="li" id="gift-nav" to="/Articles/new" class="a-tag">作成</router-link></li> -->
          <li>
            <el-dropdown>
              <span class="el-dropdown-link">
                <i class="el-icon-edit toukou-icon"></i>
                投稿<i class="el-icon-caret-bottom el-icon--right"> </i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <router-link to="/Articles/new" class="a-tag2"><el-dropdown-item>記事</el-dropdown-item></router-link>
                <router-link to="/Questions/new" class="a-tag2"><el-dropdown-item>質問</el-dropdown-item></router-link>
              </el-dropdown-menu>
            </el-dropdown>
          </li>
          <li class="a-tag logout-li" @click="logout">ログアウト</li>
        </ul>
      </nav>
    </div>
  </header>
</template>

<script>

export default {
  props: {
    loginpage: Boolean,
    isArticle: Boolean,
    isQuestion: Boolean,
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
  background: #2ee002;
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
  color: #000;
  text-decoration: none;
  cursor: pointer;
}
.header-article,.header-question{
  margin-left: 30px;
  margin-bottom: 10px;
  padding: 0 12px 7px 12px;
  color: #606266;
  cursor: pointer;
}
.active{
  border-bottom: 3px solid #e5ff00;
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
.a-tag{
  line-height: 60px;
  font-size: 15px;
  text-align: right;
  color: #606566;
  text-decoration: none;
  cursor: pointer;
}
.a-tag2{
  text-decoration: none;
  color: #606266;
}
.logout-li{
  margin-left: 20px;
}
.el-dropdown-link {
  cursor: pointer;
  color: #606566;
}
.el-icon-arrow-down {
  font-size: 12px;
}
.toukou-icon{
  font-size: 18px;
}
</style>
