<template>
  <div class="tag-all">
    <div class="tag-header">
      <span class="tag-header-text">絞り込み</span>
    </div>
    <ul class="tag-ul">
      <li>フレンドのみ</li>
    </ul>
    <span>タグ</span>
    <div>
      <el-cascader :options="options" :show-all-levels="false" filterable></el-cascader>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        options: [{
          label: 'プログラミング',
          children: [{
            label: '言語',
            children: [{
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
        }]
          }, {
            label: 'フレームワーク',
            children: [{
              value: 'top nav',
              label: 'Ruby on Rails'
            }, {
              value: 'top nav',
              label: 'Vue.js'
            }, {
              value: 'top nav',
              label: 'React'
            }, {
              value: 'top nav',
              label: 'AngularJS'
            }, {
              value: 'top nav',
              label: 'Spring Framework'
            }, {
              value: 'top nav',
              label: 'Play Framework'
            },{
              value: 'top nav',
              label: 'Bootstrap'
            }, {
              value: 'top nav',
              label: 'CakePHP'
            }, {
              value: 'top nav',
              label: 'Laravel'
            }, {
              value: 'top nav',
              label: 'Django'
            },]
          }, {
            label: 'OS',
            children: [{
              value: 'side nav',
              label: 'Linux'
            }, {
              value: 'top nav',
              label: 'Windows'
            }, {
              value: 'top nav',
              label: 'macOS'
            }, {
              value: 'top nav',
              label: 'iOS'
            }, {
              value: 'top nav',
              label: 'Android'
            }]
          }, {
            label: 'Webサーバ',
            children: [{
              value: 'side nav',
              label: 'Apache'
            }, {
              value: 'top ',
              label: 'Nginx'
            }]
          } , {
            label: 'データベース',
            children: [{
              value: 'side nav',
              label: 'MySQL'
            }, {
              value: 'top nav',
              label: 'PostgreSQL'
            }, {
              value: 'top nav',
              label: 'Oracle Database'
            }, {
              value: 'top nav',
              label: 'SQLite'
            }, {
              value: 'top nav',
              label: 'MongoDB'
            }]
          }, {
            label: '技術',
            children: [{
              value: 'side nav',
              label: 'AWS'
            }, {
              value: 'top nav',
              label: 'Docker'
            }, {
              value: 'top nav',
              label: 'kubernetes'
            }, {
              value: 'top nav',
              label: 'Git/GitHub'
            }, {
              value: 'top nav',
              label: 'WordPress'
            }, {
              value: 'top nav',
              label: 'Firebase'
            }]
          }]
        }, {
          label: '国家試験',
          children: [
            {
              value: 'layout',
              label: 'ITパスポート試験'
            },{
              value: 'layout',
              label: '基本情報技術者試験'
            }, {
              value: 'color',
              label: '応用情報技術者試験'
            }, {
              value: 'typography',
              label: '情報処理安全確保支援士試験'
            }, {
              value: 'icon',
              label: 'データベーススペシャリスト試験'
            }, {
              value: 'button',
              label: 'ネットワークスペシャリスト試験'
            }, {
              value: 'button',
              label: 'エンベデッドシステムスペシャリスト試験'
            }]
          }, {
          label: 'KBC',
          children: [{
            value: 'axure',
            label: 'ITイノベーション科'
          }, {
            value: 'sketch',
            label: 'ITエンジニア科'
          }, {
            value: 'docs',
            label: 'ゲームクリエイター科'
          }, {
            value: 'docs',
            label: '情報ビジネス科'
          }]
        }, {
          label: 'その他',
          children: [{
            value: 'axure',
            label: '初心者'
          }, {
            value: 'docs',
            label: '就活'
          }, {
            value: 'docs',
            label: 'KBC'
          }]
        }]
      };
    }
  };
</script>

<style scoped>
.tag-all{
  width: 210px;
  background-color: #F6F6F4;
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