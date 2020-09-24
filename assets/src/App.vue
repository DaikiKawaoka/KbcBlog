<template>
  <div>
    <router-view></router-view>
  </div>
</template>

<script>
import Vue from 'vue'
import VueRouter from 'vue-router'
import ArticleIndex from './articles/Index.vue'
import ArticleShow from './articles/Show.vue'
import ArticleEdit from './articles/Edit.vue'
import ArticleNew from './articles/New.vue'

const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'ArticleIndex',
      component: ArticleIndex
    },
    {
      path: '/Articles/:id(\\d+)',
      name: 'ArticleShow',
      component: ArticleShow
    },
    {
      path: '/Articles/new',
      name: 'ArticleNew',
      component: ArticleNew
    },
    {
      path: '/Articles/:id(\\d+)/edit',
      name: 'ArticleEdit',
      component: ArticleEdit
    },
  ]
})

Vue.use(VueRouter)

export default {
  router
}
</script>

<style scoped>
</style>
