<template>
  <div class="body-main">
    <el-form :model="user">
      <h2 v-if="edit">プロフィール編集</h2>
      <h2 v-else>ユーザー登録</h2>

      <div v-if="errors.length != 0">
        <ul class="error-ul" v-for="e in errors" :key="e">
            <li class="error-icon-li"><i class="el-icon-warning-outline"></i></li>
            <li><font color="red">{{ e }}</font></li>
        </ul>
      </div>

      <input name="uploadedImage" type="file" ref="file" v-on:change="onFileChange()" v-if="edit">

      <el-form-item label="名前" prop="name"
      :rules="[
          { required: true, message: '入力必須です', trigger: 'blur' },
          { min: 4, max: 20, message: '4~20文字で入力してください', trigger: 'blur' },
      ]">
        <el-input
          v-model="user.name"
          placeholder="name"
          name="user[name]"></el-input>
      </el-form-item>

      <el-form-item label="コメント" v-if="edit">
        <el-input
          v-model="user.comment"
          placeholder="comment">
        </el-input>
      </el-form-item>

      <el-form-item
      label="メールアドレス" prop="KBC_mail"
      :rules="[
          { required: true , pattern: /[\w\-\._]+@(stu.)?kawahara.ac.jp/, message: '河原学園のメールアドレスを入力してください。', trigger: 'blur'},
      ]">
          <el-input v-model="user.KBC_mail"></el-input>
      </el-form-item>

      <el-form-item label="パスワード" prop="password"
      :rules="[
          { required: true, message: '入力必須です', trigger: 'blur' },
          { min: 8, max: 50, message: '8~50文字で入力してください', trigger: 'blur' },
      ]">
          <el-input type="password" v-model="user.password" autocomplete="off"></el-input>
      </el-form-item>

      <el-form-item label="パスワード確認" prop="password"
      :rules="[
          { required: true, message: '入力必須です', trigger: 'blur' },
          { min: 8, max: 50, message: '8~50文字で入力してください', trigger: 'blur' },
      ]">
          <el-input type="password" v-model="user.password_confirmation" autocomplete="off"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="submit" @click="$emit('submit')">登録</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
// import axios from 'axios'

 export default {
   props: {
    user: {},
    errors: Object,
    edit: Boolean,
  },
  data() {
    return{
      uploadedImage: '',
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
i{
  font-size: 25px;
  color:#F56C6C;
}
li{
  list-style: none;
}
.error-icon-li{
  padding-right: 10px;
}
.error-ul{
  display: flex;
}
</style>