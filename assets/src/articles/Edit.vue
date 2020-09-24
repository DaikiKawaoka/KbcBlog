<template>
  <div id="app">

  </div>
</template>

<script>

export default {
  name: 'app',
  data(){
    return {
      aaa: "",
    }
  },
  components: {
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/Articles/${this.$route.params.id}/edit')
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
}
</style>
