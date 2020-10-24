<template>
  <el-form :model="user">
    <h2 v-if="edit">プロフィール編集</h2>
    <h2 v-else>ユーザー登録</h2>
    <div v-if="errors.length != 0">
      <ul v-for="e in errors" :key="e">
        <li><font color="red">{{ e }}</font></li>
      </ul>
    </div>

    <input name="uploadedImage" type="file" ref="file" v-on:change="onFileChange()" v-if="edit">

    <el-form-item label="名前">
      <el-input
        v-model="user.name"
        placeholder="name"
        name="user[name]"></el-input>
    </el-form-item>
    <el-form-item label="ユーザーネーム">
      <el-input
        v-model="user.user_name"
        placeholder="user_name"
        name="user[user_name]"></el-input>
    </el-form-item>

    <el-form-item label="コメント" v-if="edit">
      <el-input
        v-model="user.comment"
        placeholder="comment">
      </el-input>
    </el-form-item>

    <el-form-item
    label="メールアドレス" prop="email"
    :rules="[
        { required: true, message: 'タイトルは必ず入力してください。' },
        { max: 30, message: '30文字以内で入力してください。' },
        { type: 'email', message: 'メール形式で入力してください。' }
    ]">
        <el-input type="text" v-model="user.email" autocomplete="off"></el-input>
    </el-form-item>

    <el-form-item label="パスワード" prop="password">
        <el-input type="password" v-model="user.password" autocomplete="off"></el-input>
    </el-form-item>

    <el-form-item label="パスワード確認" prop="password">
        <el-input type="password" v-model="user.password_confirmation" autocomplete="off"></el-input>
    </el-form-item>

    <el-form-item>
      <el-button type="submit" @click="$emit('submit')">登録</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
// import axios from 'axios'

 export default {
   props: {
    user: {},
    // errors: '',
    // edit: false,
  },
  data() {
    return{
      uploadedImage: '',
      errors: '',
      edit: false,
    }
    },
    methods: {
      // onImageUploaded(e) {
      //   // event(=e)から画像データを取得する
      //   // this.user.image = e.target.files[0]
      //   // console.log(this.user.image);
      //   // console.log(this.$parent.user.image);
      // },
      onFileChange() {
      // let file = event.target.files[0] || event.dataTransfer.files
      // let reader = new FileReader()
      // reader.onload = () => {
      //   this.uploadedImage = event.target.result
      //   this.user.image = this.uploadedImage
      // }
      // reader.readAsDataURL(file)
    },
    }
 }
</script>

<style scoped>

</style>