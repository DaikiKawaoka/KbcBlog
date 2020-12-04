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
      <el-cascader @change="event" v-model="scopetag" :options="options" :show-all-levels="false" filterable :props="{ checkStrictly: true,expandTrigger:'hover' }" ></el-cascader>
    </div>
  </div>
</template>

<script>
  export default {
    props: {
      tag: String,
    },
    data() {
      return {
        value: '全て',
        scopetag: "全て",
        options: [{
          value: '全て',
          label: '全て',
        }, {
          value: ' プログラミング ',
          label: 'プログラミング',
          children: [{
            value: ' 言語 ',
            label: '言語',
            children: [{
          value: ' HTML ',
          label: 'HTML'
        }, {
          value: ' CSS ',
          label: 'CSS'
        }, {
          value: ' Java ',
          label: 'Java'
        }, {
          value: ' Python ',
          label: 'Python'
        }, {
          value: ' JavaScript ',
          label: 'JavaScript'
        }, {
          value: ' Node.js ',
          label: 'Node.js'
        }, {
          value: ' C/C++ ',
          label: 'C/C++'
        }, {
          value: ' C# ',
          label: 'C#'
        }, {
          value: ' SQL ',
          label: 'SQL'
        }, {
          value: ' PHP ',
          label: 'PHP'
        }, {
          value: ' Ruby ',
          label: 'Ruby'
        }, {
          value: ' Go ',
          label: 'Go'
        }, {
          value: ' TypeScript ',
          label: 'TypeScript'
        }, {
          value: ' Kotlin ',
          label: 'Kotlin'
        }, {
          value: ' Swift ',
          label: 'Swift'
        }]
          }, {
            value: ' フレームワーク ',
            label: 'フレームワーク',
            children: [{
              value: ' Rails ',
              label: 'Rails'
            }, {
              value: ' Vue.js ',
              label: 'Vue.js'
            }, {
              value: ' React ',
              label: 'React'
            }, {
              value: ' AngularJS ',
              label: 'AngularJS'
            }, {
              value: ' Spring Framework ',
              label: 'Spring Framework'
            }, {
              value: ' Play Framework ',
              label: 'Play Framework'
            },{
              value: ' Bootstrap ',
              label: 'Bootstrap'
            }, {
              value: ' CakePHP ',
              label: 'CakePHP'
            }, {
              value: ' Laravel ',
              label: 'Laravel'
            }, {
              value: ' Django ',
              label: 'Django'
            },]
          }, {
            value: ' OS ',
            label: 'OS',
            children: [{
              value: ' Linux ',
              label: 'Linux'
            }, {
              value: ' Windows ',
              label: 'Windows'
            }, {
              value: ' macOS ',
              label: 'macOS'
            }, {
              value: ' iOS ',
              label: 'iOS'
            }, {
              value: ' Android ',
              label: 'Android'
            }]
          }, {
            value: ' Webサーバ ',
            label: 'Webサーバ',
            children: [{
              value: ' Apache ',
              label: 'Apache'
            }, {
              value: 'Nginx',
              label: 'Nginx'
            }]
          } , {
            value: ' データベース ',
            label: 'データベース',
            children: [{
              value: ' MySQL ',
              label: 'MySQL'
            }, {
              value: ' PostgreSQL ',
              label: 'PostgreSQL'
            }, {
              value: ' Oracle Database ',
              label: 'Oracle Database'
            }, {
              value: ' SQLite ',
              label: 'SQLite'
            }, {
              value: ' MongoDB ',
              label: 'MongoDB'
            }]
          }, {
            value: ' ネットワーク ',
            label: 'ネットワーク',
          }, {
            value: ' セキュリティ ',
            label: 'セキュリティ',
          }, {
            value: ' 技術 ',
            label: '技術',
            children: [{
              value: ' AWS ',
              label: 'AWS'
            }, {
              value: ' Docker ',
              label: 'Docker'
            }, {
              value: ' kubernetes ',
              label: 'kubernetes'
            }, {
              value: ' Git/GitHub ',
              label: 'Git/GitHub'
            }, {
              value: ' WordPress ',
              label: 'WordPress'
            }, {
              value: ' Firebase ',
              label: 'Firebase'
            }]
          }]
        }, {
          value: ' 国家試験 ',
          label: '国家試験',
          children: [
          {
            value: ' ITパスポート試験 ',
            label: 'ITパスポート試験'
          },{
            value: ' 基本情報技術者試験 ',
            label: '基本情報技術者試験'
          }, {
            value: ' 応用情報技術者試験 ',
            label: '応用情報技術者試験'
          }, {
            value: ' 情報処理安全確保支援士試験 ',
            label: '情報処理安全確保支援士試験'
          }]
        }, {
          value: ' KBC ',
          label: 'KBC',
          children: [{
            value: ' ITイノベーション科 ',
            label: 'ITイノベーション科'
          }, {
            value: ' ITエンジニア科 ',
            label: 'ITエンジニア科'
          }, {
            value: ' ゲームクリエイター科 ',
            label: 'ゲームクリエイター科'
          }, {
            value: ' 情報ビジネス科 ',
            label: '情報ビジネス科'
          }]
        }, {
          value: ' その他 ',
          label: 'その他',
          children: [{
            value: ' 初心者 ',
            label: '初心者'
          }, {
            value: ' 就活 ',
            label: '就活'
          }]
        }]
      };
    },
    methods: {
      event() {
        this.$parent.tag = this.scopetag[this.scopetag.length - 1]
        console.log(this.scopetag)
        this.$emit('scopetag', this.$parent.tag);
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