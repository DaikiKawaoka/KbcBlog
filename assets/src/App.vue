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
import QuestionIndex from './questions/Index.vue'
import QuestionShow from './questions/Show.vue'
import QuestionEdit from './questions/Edit.vue'
import QuestionNew from './questions/New.vue'
import LoginPage from './sessions/Login.vue'
import UserNew from './users/New'
import UserShow from './users/Show'
import UserEdit from './users/Edit'
import PassEdit from './users/PassEdit'

const router = new VueRouter({
  mode: 'history',
  scrollBehavior() {
    return { x: 0, y: 0 };
  },
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
    {
      path: '/Questions',
      name: 'QuestionIndex',
      component: QuestionIndex
    },
    {
      path: '/Questions/:id(\\d+)',
      name: 'QuestionShow',
      component: QuestionShow
    },
    {
      path: '/Questions/new',
      name: 'QuestionNew',
      component: QuestionNew
    },
    {
      path: '/Questions/:id(\\d+)/edit',
      name: 'QuestionEdit',
      component: QuestionEdit
    },
    {
      path: '/login',
      name: 'LoginPage',
      component: LoginPage
    },
    {
      path: '/Users/new',
      name: 'UserNew',
      component: UserNew
    },
    {
      path: '/Users/:id(\\d+)',
      name: 'UserShow',
      component: UserShow
    },
    {
      path: '/Users/:id(\\d+)/edit',
      name: 'UserEdit',
      component: UserEdit
    },
    {
      path: '/Users/:id(\\d+)/password/edit',
      name: 'PassEdit',
      component: PassEdit
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
