<template>
  <div id="app">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js Show"/>
    <p>{{aaa}}</p>
  </div>
</template>

<script>
import HelloWorld from './../components/HelloWorld.vue'

export default {
  name: 'app',
  data(){
    return {
      aaa: "",
    }
  },
  components: {
    HelloWorld
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/Articles/${this.$route.params.id}')
      .then(response => {
        this.aaa = response.data
        console.log(response)
      })
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
