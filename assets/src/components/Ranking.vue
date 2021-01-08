<template>
  <div v-if="isArticle" class="ranking-all">
    <!-- article -->
    <div class="ranking-title-div" v-if="random === 0">
      <h3 class="ranking-title">人気ユーザーランキング<i class="el-icon-user-solid"></i></h3>
    </div>
    <div class="ranking-title-div post-count-ranking" v-else>
      <h3 class="ranking-title">投稿数ランキング<i class="el-icon-document-checked"></i></h3>
    </div>
    <ul class="ranking-ul">
      <li class="ranking-li" v-for="(u,index) in this.Ranking" :key="u.id">
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
                <img class="user-icon" :src="image_path(u)">
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
                <i class="el-icon-star-off" v-if="random === 0"></i>
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
    <div class="ranking-title-div" v-if="random === 0">
      <h3 class="ranking-title">意欲的ユーザーランキング<i class="el-icon-user-solid"></i></h3>
    </div>
    <div class="ranking-title-div post-count-ranking" v-else>
      <h3 class="ranking-title">回答数ランキング<i class="el-icon-document-checked"></i></h3>
    </div>
    <ul class="ranking-ul">
      <li class="ranking-li" v-for="(u,index) in this.Ranking" :key="u.id">
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
                <img class="user-icon" :src="image_path(u)">
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
                <i class="el-icon-tickets" v-if="random === 0 "></i>
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
      random: Number,
      isArticle: Boolean,
    },
    data() {
      return {
      };
    },
    methods:{
      image_path(u){
        if(u.imgdata64.Valid === true){
          return 'data:image/jpeg;base64,' + u.imgdata64.String
          // return require("@/assets/userIcon/" + this.user.imgdata64.String);
        }else{
          if(u.sex === 1){
            return require("@/assets/userIcon/man.png");
          }else if(u.sex === 2){
            return require("@/assets/userIcon/woman.png");
          }
        }
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
.ranking-ul{
  background-color: #fff;
  padding: 10px;
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