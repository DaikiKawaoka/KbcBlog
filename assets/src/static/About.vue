<template>
  <div id="app" v-if="user">
    <Header :isArticle="false" :isQuestion="false" :user="user" :url="url"></Header>
    <div class="about-all">
      <div class="div-1">
        <h1 class="title">KBC Blog とは</h1>
        <p>河原電子ビジネス専門学校(KBC)の生徒、職員専用の技術ブログサイトです。</p>
      </div>
      <div class="div-2">
        <h1 class="title">目的</h1>
        <ul>
          <li>
            <h3>1 KBC全体の技術力の向上</h3>
          </li>
          <li>
            <h3>2 技術ブログの作成経験を得る</h3>
          </li>
          <li>
            <h3>3 他学年他クラスとの情報交換</h3>
          </li>
        </ul>
      </div>
    </div>
    <Footer :user="user" :url="url"></Footer>
  </div>
</template>

<script>
import Header from './../components/Header.vue'
import Footer from './../components/Footer.vue'

export default {
  name: 'app',
  data(){
    return {
      user: {},
      url: null,
    }
  },
  components: {
    Header,
    Footer,
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    // this.openFullScreen()
    this.url = process.env.VUE_APP_URL
    this.$axios.get(this.url+`api/restricted/About`,{
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },})
      .then(response => {
        this.user = response.data.user
        // this.closeFullScreen()
      })
      .catch(error => {
        if(error.response.status == 401){
          this.user = {}
        }
        // this.closeFullScreen()
      })
  },
  methods: {
  },

}
</script>

<style lang="scss" scoped>
#app{
  background-color: #fff;
}
.about-all{
  width: 650px;
  margin: 0 auto;
  height: 500px;
  p{
    text-align: center;
    font-size: 1.1em;
    // font-weight: bold;
  }
  .title{
    margin-top: 50px;
    text-align: center;
    font-size: 4em;
  }
  h3{
    text-align: center;
  }
}
</style>