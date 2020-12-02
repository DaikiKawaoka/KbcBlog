<template>
  <div class="tag-all">
    <el-form-item label="タグ">
    <div class="select-div">
      <el-select
        style="width: 400px;"
        v-model="tagArray"
        v-on:change="arrayChangeString"
        multiple
        size="large"
        :multiple-limit=4
        :collapse-tags="false"
        no-match-text="一致するタグがありません"
        filterable
        :allow-create="false"
        :default-first-option="true"
        placeholder="関連するタグを最大5つまで選択してください。">
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value">
        </el-option>
      </el-select>
    </div>
  </el-form-item>

  </div>
</template>

<script>
  export default {
    props: {
      article: Object,
      tagArray: Array,
    },
    data() {
      return {
        tagArray2: [],
        options: [
          {
            value: 'プログラミング',
            label: 'プログラミング'
          }, {
            value: 'プログラミング言語',
            label: 'プログラミング言語',
          }, {
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
            value: 'Node.js',
            label: 'Node.js'
          }, {
            value: 'C/C++',
            label: 'C/C++'
          }, {
            value: 'C#',
            label: 'C#'
          }, {
            value: 'SQL',
            label: 'SQL'
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
          },{
            value: 'VB',
            label: 'VB'
          },{
            value: 'VBA',
            label: 'VBA'
          },{
            value: 'COBOL',
            label: 'COBOL'
          }, {
            value: 'フレームワーク',
            label: 'フレームワーク',
          } ,{
            value: 'Ruby on Rails',
            label: 'Ruby on Rails'
          }, {
            value: 'Vue.js',
            label: 'Vue.js'
          }, {
            value: 'React',
            label: 'React'
          }, {
            value: 'AngularJS',
            label: 'AngularJS'
          }, {
            value: 'Spring Framework',
            label: 'Spring Framework'
          }, {
            value: 'Play Framework',
            label: 'Play Framework'
          },{
            value: 'Bootstrap',
            label: 'Bootstrap'
          }, {
            value: 'CakePHP',
            label: 'CakePHP'
          }, {
            value: '9Laravel',
            label: 'Laravel'
          }, {
            value: 'Django',
            label: 'Django'
          }, {
            value: 'OS',
            label: 'OS',
          }, {
            value: 'Linux',
            label: 'Linux'
          }, {
            value: 'Windows',
            label: 'Windows'
          }, {
            value: 'macOS',
            label: 'macOS'
          }, {
            value: 'iOS',
            label: 'iOS'
          }, {
            value: 'Android',
            label: 'Android'
          }, {
            value: 'Webサーバ',
            label: 'Webサーバ',
          }, {
            value: 'Apache',
            label: 'Apache'
          }, {
            value: 'Nginx',
            label: 'Nginx'
          }, {
            value: 'データベース',
            label: 'データベース',
          }, {
            value: 'MySQL',
            label: 'MySQL'
          }, {
            value: 'PostgreSQL',
            label: 'PostgreSQL'
          }, {
            value: 'Oracle Database',
            label: 'Oracle Database'
          }, {
            value: 'SQLite',
            label: 'SQLite'
          }, {
            value: 'MongoDB',
            label: 'MongoDB'
          }, {
            value: '技術',
            label: '技術',
          }, {
            value: 'AWS',
            label: 'AWS'
          }, {
            value: 'Docker',
            label: 'Docker'
          }, {
            value: 'kubernetes',
            label: 'kubernetes'
          }, {
            value: 'Git/GitHub',
            label: 'Git/GitHub'
          }, {
            value: 'WordPress',
            label: 'WordPress'
          }, {
            value: 'Firebase',
            label: 'Firebase'
          }, {
            value: '国家試験',
            label: '国家試験',
          }, {
            value: 'ITパスポート試験',
            label: 'ITパスポート試験'
          },{
            value: '基本情報技術者試験',
            label: '基本情報技術者試験'
          }, {
            value: '応用情報技術者試験',
            label: '応用情報技術者試験'
          }, {
            value: '情報処理安全確保支援士試験',
            label: '情報処理安全確保支援士試験'
          }, {
            value: 'データベーススペシャリスト試験',
            label: 'データベーススペシャリスト試験'
          }, {
            value: 'ネットワークスペシャリスト試験',
            label: 'ネットワークスペシャリスト試験'
          }, {
            value: 'エンベデッドシステムスペシャリスト試験',
            label: 'エンベデッドシステムスペシャリスト試験'
          }, {
            value: 'KBC',
            label: 'KBC',
          }, {
            value: 'ITイノベーション科',
            label: 'ITイノベーション科'
          }, {
            value: 'ITエンジニア科',
            label: 'ITエンジニア科'
          }, {
            value: 'ゲームクリエイター科',
            label: 'ゲームクリエイター科'
          }, {
            value: '情報ビジネス科',
            label: '情報ビジネス科'
          }, {
            value: 'その他',
            label: 'その他',
          }, {
            value: '初心者',
            label: '初心者'
          }, {
            value: '就活',
            label: '就活'
          }]
      };
    },
    mounted(){
      this.tagArray2 = this.tagArray
    },
    beforeUpdate(){
      this.tagArray2 = this.article.tag.split(',');
      if(this.tagArray2[0].length == 0){
        this.tagArray2 = [];
      }
      console.log(this.tagArray2)
      console.log(this.article)
  },
    methods: {
      arrayChangeString(){
        this.article.tag = this.tagArray2.join(',');
        console.log(this.article.tag)
      },
    },
  };
</script>

<style scoped>
.tag-all{
  width: 210px;
  background-color: #F6F6F4;
  /* background-color: #15202b; */
}
.tag-header{
  margin-top: 20px;
}
.tag-header-text{
  color: #999;
}
ul{
  padding: 0;
}
</style>