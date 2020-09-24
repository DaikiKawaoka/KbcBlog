<template>
  <el-form :model="article">

    <h2 v-if="create">技術記事作成</h2>
    <h2 v-else>技術記事編集</h2>

    <div v-if="errors.length != 0">
      <ul v-for="e in errors" :key="e">
        <li><font color="red">{{ e }}</font></li>
      </ul>
    </div>

    <el-form-item label="タイトル">
      <el-input
        type="text"
        placeholder="Please input"
        v-model="article.title"
        maxlength="100"
        show-word-limit
      >
      </el-input>
    </el-form-item>

    <el-row>
      <el-col :span="4"><div class="grid-content bg-purple"></div></el-col>
      <el-col :span="4"><div class="grid-content bg-purple-light"></div></el-col>
      <el-col :span="4"><div class="grid-content bg-purple"></div></el-col>
      <el-col :span="4"><div class="grid-content bg-purple-light"></div></el-col>
      <el-col :span="4"><div class="grid-content bg-purple"></div></el-col>
      <el-col :span="4">
        <el-button v-if="isActive" type="primary" icon="el-icon-open" circle v-on:click="active"></el-button>
        <el-button v-else icon="el-icon-turn-off" circle v-on:click="active"></el-button>
      </el-col>
    </el-row>


    <el-form-item label="本文">
      <el-input
        type="textarea"
        :rows="20"
        placeholder="Please input"
        v-model="article.body">
      </el-input>
    </el-form-item>

    <el-row>
      <el-col :span="12"><el-button type="danger">キャンセル</el-button></el-col>
      <el-col :span="12"><el-button type="success" @click="$emit('submit')" native-type="submit">セーブ</el-button></el-col>
    </el-row>
  </el-form>
</template>

<script>

 export default {
   props: {
    article: Object,
    errors: Array,
    create: Boolean,
  },
  data(){
    return{
      isActive: false
    }
  },
  methods: {
    active: function () {
      this.isActive = !this.isActive
    }
  }
 }
</script>

<style scoped>
.grid-content {
    border-radius: 4px;
    min-height: 36px;
  }
</style>