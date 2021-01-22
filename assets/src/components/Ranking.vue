<template>
  <div v-if="isArticle" class="ranking-all">
    <!-- article -->
    <div class="ranking-period-all">
      <ul class="ranking-period-ul">
        <li v-if="rankingType === 0" class="ranking-title">
            人気ユーザーランキング<i class="el-icon-user-solid"></i>
        </li>
        <li v-else class="ranking-title">
          投稿数ランキング<i class="el-icon-document-checked"></i>
        </li>
        <li class="ranking-period-li" v-bind:class="{ 'now-period': nowPeriod === 'month' }" @click="getRanking('Article','month')">
          月間
        </li>
        <li class="ranking-period-li" v-bind:class="{ 'now-period': nowPeriod === 'all' }" @click="getRanking('Article','all')">
          全て
        </li>
      </ul>
    </div>
    <ul class="ranking-ul" v-loading="rankingLoading">
      <li class="ranking-li" v-for="(u,index) in this.internalRanking" :key="u.id" v-bind:class="{ active: user.id === u.id }">
        <router-link v-bind:to="{ name : 'UserShow', params : { id: u.id }}" class="a-tag">
          <div class="ranking-main">
            <div class="ranking-main">
              <div class="rank-number-div">
                <!--- 順位--->
                <span class="rank-number" v-bind:class="`rank-${index + 1}`">{{ index + 1 }}</span>
              </div>
              <div>
                <!-- アイコン -->
                <!-- <div class="user-icon"></div> -->
                <img class="user-icon" :src="u.imgpath">
              </div>
              <div class="u-div-nameandcomment-width">
                <div class="string-out">
                  <span class="user-name">{{ u.name }}</span>
                </div>
                <div class="string-out">
                  <span class="tag-user-comment">{{u.comment.String}}</span>
                </div>
              </div>
            </div>
            <div class="like-count-div">
              <p class="like-count">{{ u.count }}</p>
              <p class="like-tag">
                <i class="el-icon-star-off" v-if="rankingType === 0"></i>
                <i class="el-icon-document" v-else></i>
              </p>
            </div>
          </div>
        </router-link>
      </li>
    </ul>
  </div>
  <div v-else class="ranking-all">
    <!-- question -->
    <div class="ranking-period-all">
      <ul class="ranking-period-ul">
        <li v-if="rankingType === 0" class="ranking-title">
            やる気ランキング<i class="el-icon-user-solid"></i>
        </li>
        <li v-else class="ranking-title">
          回答数ランキング<i class="el-icon-document-checked"></i>
        </li>
        <li class="ranking-period-li" v-bind:class="{ 'now-period': nowPeriod === 'month' }" @click="getRanking('Question','month')">
          月間
        </li>
        <li class="ranking-period-li" v-bind:class="{ 'now-period': nowPeriod === 'all' }" @click="getRanking('Question','all')">
          全て
        </li>
      </ul>
    </div>
    <ul class="ranking-ul" v-loading="rankingLoading">
      <li class="ranking-li" v-for="(u,index) in this.internalRanking" :key="u.id">
        <router-link v-bind:to="{ name : 'UserShow', params : { id: u.id }}" class="a-tag">
          <div class="ranking-main">
            <div class="ranking-main">
              <div class="rank-number-div">
                <!--- 順位--->
                <span class="rank-number" v-bind:class="`rank-${index + 1}`">{{ index + 1 }}</span>
              </div>
              <div>
                <!-- アイコン -->
                <!-- <div class="user-icon"></div> -->
                <img class="user-icon" :src="u.imgpath">
              </div>
              <div class="u-div-nameandcomment-width">
                <div class="string-out">
                  <span class="user-name">{{ u.name }}</span>
                </div>
                <div class="string-out">
                  <span class="tag-user-comment">{{u.comment.String}}</span>
                </div>
              </div>
            </div>
            <div class="like-count-div">
              <p class="like-count">{{ u.count }}</p>
              <p class="like-tag">
                <i class="el-icon-tickets" v-if="rankingType === 0 "></i>
                <i class="el-icon-chat-dot-round" v-else></i>
                </p>
            </div>
          </div>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script>
  export default {
    props: {
      Ranking: Array,
      rankingType: Number,
      isArticle: Boolean,
      user: {},
    },
    data() {
      return {
        nowPeriod: "month",
        internalRanking: [],
        rankingLoading: false,
      };
    },
    watch: {
      Ranking: { // 外からプロパティの中身が変更になったら実行される
          immediate: true,
          handler(value) {
              this.internalRanking = value;
          }
      }
    },
    methods:{
      getRanking(type,period) {
        this.rankingLoading = true;
        this.$axios.get(`http://localhost/api/restricted/${type}/Ranking`, {
        params: {
          rankingType: this.rankingType,
          period: period,
        },
        headers: {
          Authorization: `Bearer ${this.$cookies.get("JWT")}`
        },
      })
        .then(response => {
          this.internalRanking = response.data.Ranking;
          this.nowPeriod = period
          this.rankingLoading = false;
        })
        .catch(error => {
          if(error.response.status == 401){
            this.$router.push({ path: "/login" });
            this.errorNotify();
            this.rankingLoading = false;
          }
        })
      },
    }
  };
</script>

<style scoped>
.ranking-all{
  margin-left: 20px;
  padding-top: 10px;
  position: sticky;
  top: 0;
}
.ranking-period-all{
  margin-top: 35px;
  padding-top: 10px;
  background: #fff;
  border-radius: 5px 5px 0 0;
}
.ranking-period-ul{
  display: flex;
  margin: 0;
  padding-left:5px;
}
.ranking-period-li{
  padding:6px 6px 0 6px;
  font-size: 0.8em;
  color: #777;
  cursor: pointer;
}
.ranking-title{
  width: 180px;
  padding-left: 5px;
  font-size: 0.9em;
  font-weight: bold;
}
.now-period{
  font-weight: bold;
  color: #333;
}
.ranking-ul{
  background-color: #fff;
  padding: 10px;
  margin-top: 0px;
}
.ranking-main{
  display: flex;
  height: 40px;
}
.rank-number-div{
  width: 23px;
}
.ranking-li{
  width: 250px;
  padding: 7px 0 5px 0;
  border-bottom: #eee solid 1px;
}
.active{
  border: #f0e68c 2.5px solid;
}
.rank-number{
  line-height: 40px;
  margin-right: 10px;
}
.rank-1{
  color: gold;
  font-weight: bold;
  font-size: 1.5em;
}
.rank-2{
  color: silver;
  font-weight: bold;
  font-size: 1.2em;
}
.rank-3{
  color: #cd7f32;
  font-weight: bold;
  font-size: 1.2em;
}
.user-icon{
  background-color: #ccc;
  width: 30px;
  height: 30px;
  object-fit: cover; /* 画像トリミング */
  border-radius: 50%;
}
.user-name{
  width: 150px;
  font-size: 0.9em;
  font-weight: bold;
  margin: 0 0 0 8px;
}
.like-count-div{
  width: 35px;
}
.like-count{
  width: 100%;
  text-align: center;
  font-weight: bold;
  margin: 0;
}
.like-tag{
  font-size: 0.3em;
  color: #999;
  text-align: center;
  margin: 0;
}


.post-count-ranking{
  margin-top: 35px;
}
.u-div-nameandcomment-width{
  width: 158px;
}
.string-out{
  overflow: hidden;
  white-space: nowrap;
	text-overflow: ellipsis;
	-webkit-text-overflow: ellipsis; /* Safari */
	-o-text-overflow: ellipsis; /* Opera */
}
.tag-user-comment{
  font-size: 0.7em;
}
</style>