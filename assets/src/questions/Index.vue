<template>
  <div id="app" v-if="user">
    <Header :isArticle="false" :isQuestion="true" :user="user" @reset="reset" :url="url"></Header>
    <div class="index-main body-main">
      <Tag @scopetag="scopetag" :tag="tag" :friendsOnly="friendsOnly" @reset="reset" @update:friendsOnly ="friendsOnly=$event"></Tag>
      <div class="question-all-div">
        <div class="question-search-div">
          <el-input placeholder="キーワード検索" v-model="searchText" suffix-icon="el-icon-search" style="width:200px; margin-left: 30px;" @input="scopetag"></el-input>
          <div class="order-div">
            <span class="order-span-left">並び替え</span>
            <span class="order-span" v-bind:class="{'order-span-color': order === 'new' }" @click="questionOrder('new')">新着</span>
            <span class="order-span" v-bind:class="{ 'order-span-color': order === 'like' }" @click="questionOrder('like')">人気</span>
          </div>
        </div>

        <div v-if="questions.length !== 0" v-loading="questionloading"
        element-loading-text="Loading..." element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 1)">
          <div v-for="question in questions" :key="question.id" class="question-show-div">
            <div class="question-user-icon-div">
              <router-link v-bind:to="{ name : 'UserShow', params : { id: question.userid }}" class="a-tag">
                <img class="question-user-icon" :src="question.imgpath">
              </router-link>
            </div>
            <div class="question-index-body">
              <div class="question-tag-div">
                <i class="el-icon-collection-tag tag-icon"></i>
                <span class="question-tag no-magin">{{ question.tag }}</span>
              </div>
              <router-link v-bind:to="{ name : 'QuestionShow', params : { id: question.id }}" class="a-tag">
                <h2 class="question-title-index"> {{question.title}} </h2>
              </router-link>
              <div class="question-index-username-updated">
                <div>
                  <p v-if="question.category === 'Q&A'" class="question-index-category p-QandA"> {{question.category}} </p>
                  <p v-else class="question-index-category p-ikenkoukan"> {{question.category}} </p>
                </div>
                <router-link v-bind:to="{ name : 'UserShow', params : { id: question.userid }}" class="a-tag">
                  <h3 class="question-index-username">by {{ question.name }}</h3>
                </router-link>
                <h3 class="question-index-update"> {{ question.Created | dayjs }}</h3>
                <i class="el-icon-star-on question-star-i"></i>
                <span class="question-likecount-span">{{question.likecount}}</span>
              </div>
            </div>
          </div>


          <infinite-loading @infinite="scrollQuestions" :identifier="infiniteId" spinner="spiral" class="infinite-loading">
            <span slot="no-more" class="thats-all"><i class="el-icon-caret-top"></i> That’s all  <i class="el-icon-caret-top"></i></span>
            <span slot="no-results" class="thats-all"><i class="el-icon-caret-top"></i> That’s all  <i class="el-icon-caret-top"></i></span>
          </infinite-loading>
        </div>

        <div v-else v-loading="questionloading"
        element-loading-text="Loading..." element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 1)">
          <div class="question-show-div not-tag-div" v-if="!friendsOnly">
            <p class="not-tag" v-if="searchText === '' && tag !== '全て' ">タグ『{{ tag }}』の質問はまだありません。</p>
            <p class="not-tag" v-else-if="tag === '全て' && searchText !== '' ">キーワード『{{ searchText }}』の質問はありません。</p>
            <p class="not-tag"  v-else-if="tag !== '全て' && searchText !== '' ">タグ『{{ tag }}』,キーワード『{{ searchText }}』の質問はありません。</p>
            <p class="not-tag" v-else>まだ質問が投稿されていません。</p>
          </div>
          <div class="question-show-div not-tag-div" v-else>
            <p class="not-tag" v-if="searchText === '' && tag !== '全て' ">タグ『{{ tag }}』の質問はまだありません。</p>
            <p class="not-tag" v-else-if="tag === '全て' && searchText !== '' ">キーワード『{{ searchText }}』の質問はありません。</p>
            <p class="not-tag"  v-else-if="tag !== '全て' && searchText !== ''">タグ『{{ tag }}』,キーワード『{{ searchText }}』の質問はありません。</p>
            <p class="not-tag" v-else >フレンドの質問はまだありません。</p>
          </div>
        </div>

      </div>
      <div>
        <Ranking :Ranking="Ranking" :rankingType="rankingType" :url="url"></Ranking>
      </div>
    </div>
    <Footer :user="user" :url="url"></Footer>
  </div>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading'
import Header from './../components/Header.vue'
import Footer from './../components/Footer.vue'
import Tag from './../components/Tag.vue'
import Ranking from './../components/Ranking.vue'
import dayjs from 'dayjs'
import timezone from 'dayjs/plugin/timezone'
import utc from 'dayjs/plugin/utc'

export default {
  name: 'app',
  data(){
    return {
      questions: [],
      user: null,
      Ranking: [],
      searchText:"",
      tag: String,
      order: String,
      friendsOnly: Boolean,
      cursor: 0,
      page: 1,
      infiniteId: +new Date(),
      notificationCount: localStorage.notificationCount,
      rankingType: null,
      questionloading: false,
      url: null,
    }
  },
  components: {
    Header,
    Footer,
    Tag,
    Ranking,
    InfiniteLoading,
  },
  filters: {
    dayjs: function (date) {
      dayjs.extend(utc)
      dayjs.extend(timezone)
      return dayjs(date).tz('UTC').format('YYYY/MM/DD')
    }
  },
  watch: {
    searchText(newSearchText) {
      localStorage.searchText = newSearchText;
    },
    order(newOrder) {
      localStorage.order = newOrder;
    },
    tag(newTag) {
      localStorage.tag = newTag;
    },
    friendsOnly(newFriendsOnly) {
      localStorage.friendsOnly = newFriendsOnly;
    },
    notificationCount(newNotificationCount) {
      localStorage.notificationCount = newNotificationCount;
    },
  },
  created () {
    this.url = process.env.VUE_APP_URL
    document.title = `KBC Blog`;
    this.openFullScreen()
    const jst = this.$cookies.get("JWT");
    if(jst === null){
      this.$router.push({ path: "/login" });
    }
    this.setUp()
    this.$axios.get(this.url+'api/restricted/Questions',{
      params: {
        // ここにクエリパラメータを指定する
        friendsOnly: this.friendsOnly,
        searchText: this.searchText,
        order: this.order,
        tag: this.tag,
        cursor: this.cursor,
        rankingType: this.rankingType,
      },
      headers: {
        Authorization: `Bearer ${jst}`
      },
    })
      .then(response => {
        this.questions = response.data.Questions
        this.Ranking = response.data.Ranking
        this.user = response.data.user
        this.cursor = response.data.Cursor
        this.notificationCount = response.data.NotificationCount
        this.closeFullScreen()
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
        this.closeFullScreen()
      })
  },

  methods:{

    openFullScreen() {
      this.loading = this.$loading({
          lock: true,
          text: 'Loading',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.98)'
        });
    },
    closeFullScreen() {
      this.loading.close();
    },

    errorNotify() {
      this.$notify.error({
        title: 'Error',
        message: 'あなたのセッションはタイムアウトしました。再度ログインしてください。'
      });
    },

    questionOrder(c) {
      this.questionloading = true
      this.changeType()
      this.$axios.get(this.url+'api/restricted/Questions/scope', {
      params: {
        friendsOnly: this.friendsOnly,
        searchText: this.searchText,
        order: c,
        tag: this.tag,
        cursor: this.cursor,
      },
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        if(c === "new"){
          this.order = "new"
        }else{
          this.order = "like"
        }
        this.questions = response.data.Questions
        this.cursor = response.data.Cursor
        this.questionloading = false
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
        this.questionloading = false
      })
    },

    scopetag() {
      this.changeType()
      this.$axios.get(this.url+'api/restricted/Questions/scope', {
      params: {
        friendsOnly: this.friendsOnly,
        searchText: this.searchText,
        order: this.order,
        tag: this.tag,
        cursor: this.cursor,
      },
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        this.questions = response.data.Questions
        this.cursor = response.data.Cursor
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
    },

    scrollQuestions($state) {
      this.$axios.get(this.url+'api/restricted/Questions/scope', {
      params: {
        friendsOnly: this.friendsOnly,
        searchText: this.searchText,
        order: this.order,
        tag: this.tag,
        cursor: this.cursor,
      },
      headers: {
        Authorization: `Bearer ${this.$cookies.get("JWT")}`
      },
    })
      .then(response => {
        if (response.data.Questions.length) {
          this.questions = this.questions.concat(response.data.Questions);
          this.cursor = response.data.Cursor
          this.page += 1;
          $state.loaded();
        } else {
          $state.complete();
        }
      })
      .catch(error => {
        if(error.response.status == 401){
          this.$router.push({ path: "/login" });
          this.errorNotify();
        }
      })
    },
    changeType() {
      this.page = 1;
      this.infiniteId += 1;
      this.questions = [];
      this.cursor = 0;
    },

    setUp() {
      if (localStorage.searchText) {
        this.searchText = localStorage.searchText;
      }
      if (localStorage.tag) {
        this.tag = localStorage.tag;
      }else{
        this.tag = '全て'
      }
      if (localStorage.order) {
        this.order = localStorage.order;
      }else{
        this.order = 'new'
      }
      if (localStorage.friendsOnly) {
        if(localStorage.friendsOnly === "true"){
          this.friendsOnly = true
        }else{
          this.friendsOnly = false
        }
      }else{
        this.friendsOnly = false
      }

      this.rankingType = Math.floor(Math.random() * Math.floor(2));
    },
    reset(){
      this.searchText = ''
      this.tag = '全て'
      this.order = 'new'
      this.friendsOnly = false
      this.scopetag()
    }
  }
}
</script>

<style>

.question-all-div{
  width: 480px;
  background-color: #F6F6F4;
  /* background-color: #15202b; */
}
.question-show-div{
  background-color: #FFF;
  padding-left: 10px;
  border: solid 1px #eee;
  border-radius: 6px;
  margin-bottom: 5px;
  display: flex;
}
.question-show-div:hover {
  transition: 0.10s ;
  border: solid 1.3px #409eff;
  /* box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
  transform: translateY(-0.1875em); */
}
.question-user-icon-div{
  margin-top: 25px;
}
.question-user-icon{
  object-fit: cover; /* 画像トリミング */
  width: 30px;
  height: 30px;
  border-radius: 50%;
}
.question-index-body{
  margin-left: 10px;
}
.question-index-username-updated,.question-search-div{
  display: flex;
}
.question-index-category{
  padding: 1.5px;
  margin-right: 10px;
  border-radius: 2px;
  color: #FFF;
  font-weight: bold;
}
.p-QandA{
  background: #dc143c;
  font-size: 0.9em;
}
.p-ikenkoukan{
  background: #4169e1;
  font-size: 0.8em;
}
.question-title-index{
  font-size: 1.1em;
  width: 440px;
  text-align: left;
  margin: 0;
}
.question-index-username,.question-index-update{
  font-size: 13px;
  color: #999;
}
.question-index-update{
  margin-left: 20px;
}
.question-search-div{
  background-color: #F6F6F4;
  /* background-color: #15202b; */
  padding-bottom: 8px;
  /* margin-left: 20px; */
  padding-top: 10px;
  position: sticky;
  top: 0;
}

.question-star-i{
  color: orange;
  margin-top: 13px;
  margin-left: 20px;
  font-size: 1.3em;
}
.question-likecount-span{
  margin-top: 15px;
  font-size: 0.8em;
  color: #777;
}

.question-tag-div{
  display: flex;
}

.question-tag{
  font-size: 0.6em;
  color: #999;
  width: 440px;
}
.infinite-loading{
  height: 60px;
  line-height: 60px;
}
.thats-all{
  color: #777;
}
</style>
