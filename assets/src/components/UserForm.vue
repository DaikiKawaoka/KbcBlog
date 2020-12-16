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

      <!-- <input name="uploadedImage" type="file" ref="file" v-on:change="onFileChange()" v-if="edit"> -->

      <el-form-item label="名前" prop="name"
      :rules="[
          { required: true, message: '入力必須です', trigger: 'blur' },
          { min: 4, max: 10, message: '4~10文字で入力してください', trigger: 'blur' },
      ]">
        <el-input
          v-model="user.name"
          placeholder="name"
          maxlength="10"
          show-word-limit
          name="user[name]"></el-input>
      </el-form-item>

      <el-form-item label="コメント" v-if="edit"
      :rules="[
          { max: 150, message: '最大150文字です', trigger: 'blur' },
      ]">
        <el-input
          v-model="user.comment.String"
          maxlength="150"
          show-word-limit
          placeholder="comment">
        </el-input>
      </el-form-item>

  <el-form-item label="好きな言語 TOP 3" v-if="edit">
    <div class="select-div">
      <el-select
        style="width: 320px;"
        v-model="langarray"
        v-on:change="arrayChangeString"
        multiple
        size="large"
        :multiple-limit=3
        :collapse-tags="false"
        no-match-text="一致する言語がありません"
        filterable
        :allow-create="false"
        :default-first-option="true"
        placeholder="好きな言語を最大3つまで選択してください">
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value">
        </el-option>
      </el-select>
    </div>
  </el-form-item>



      <el-form-item label="GithubアカウントのURL" prop="github.String" v-if="edit"
      :rules="[
          { pattern: /https\:\/\/github.com\/[\w]+/, message: '自分のGithubアカウントのURLを入力してください。', trigger: 'blur'},
      ]">
        <el-input
          v-model="user.github.String"
          show-word-limit
          placeholder="例 https://github.com/DaikiKawaoka">
        </el-input>
      </el-form-item>

      <el-form-item label="WebサイトのURL" prop="website.String" v-if="edit"
      :rules="[
          { pattern: /https?:\/\/[\w/:%#\$&\?\(\)~\.=\+\-]+/, message: 'リンクを入力してください。', trigger: 'blur'},
      ]">
        <el-input
          v-model="user.website.String"
          show-word-limit
          placeholder="例  http://ec2-13-230-88-76.ap-northeast-1.compute.amazonaws.com/portfolio/">
        </el-input>
      </el-form-item>

      <el-form-item v-if="!edit"
      label="メールアドレス" prop="KBC_mail"
      :rules="[
          { required: true , pattern: /[\w\-\._]+@(stu.)?kawahara.ac.jp/, message: '河原学園のメールアドレスを入力してください。', trigger: 'blur'},
      ]">
          <el-input v-model="user.KBC_mail"></el-input>
      </el-form-item>

      <el-form-item label="パスワード" prop="password" v-if="!edit"
      :rules="[
          { required: true, message: '入力必須です', trigger: 'blur' },
          { min: 8, max: 50, message: '8~50文字で入力してください', trigger: 'blur' },
      ]">
          <el-input type="password" v-model="user.password" autocomplete="off"></el-input>
      </el-form-item>

      <el-form-item label="パスワード確認" prop="password" v-if="!edit"
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
    user: Object,
    errors: Object,
    edit: Boolean,
  },
  data() {
    return{
      langarray: [],
      options: [{
          value: 'HTML',
          label: 'HTML'
        }, {
          value: 'CSS',
          label: 'CSS'
        }, {
          value: 'Java',
          label: 'Java'
        }, {
          value: 'Python',
          label: 'Python'
        }, {
          value: 'JavaScript',
          label: 'JavaScript'
        }, {
          value: 'C',
          label: 'C'
        }, {
          value: 'C++',
          label: 'C++'
        }, {
          value: 'C#',
          label: 'C#'
        }, {
          value: 'PHP',
          label: 'PHP'
        }, {
          value: 'Ruby',
          label: 'Ruby'
        }, {
          value: 'Rust',
          label: 'Rust'
        }, {
          value: 'Go',
          label: 'Go'
        }, {
          value: 'TypeScript',
          label: 'TypeScript'
        }, {
          value: 'R',
          label: 'R'
        }, {
          value: 'Perl',
          label: 'Perl'
        }, {
          value: 'Kotlin',
          label: 'Kotlin'
        }, {
          value: 'Swift',
          label: 'Swift'
        }],
      uploadedImage: '',
    }
    },
    beforeUpdate() {
      // 文字列のlangsを配列に変換
        this.langarray = this.user.languages.String.split(',');
        if(this.langarray[0].length == 0){
          this.langarray = [];
        }
    },
    methods: {
      arrayChangeString(){
        this.user.languages.String = this.langarray.join(',');
        // console.log(this.user.languages);
      },
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
.select-div{
  margin-top: 40px;
  margin-right: 50px;
}
</style>