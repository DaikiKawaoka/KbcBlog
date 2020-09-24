<template>
  <div id="app">
    <Header></Header>
  </div>
</template>

<script>
import Header from './../components/Header.vue'

export default {
  name: 'app',
  data(){
    return {
      aaa: "",
    }
  },
  components: {
    Header
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/')
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
  width: 1000px;
  margin-left: auto;
  margin-right: auto;
}
</style>
