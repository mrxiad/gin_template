import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/Main/index.vue'
// createRouter 创建路由实例，===> new VueRouter()
// history 是路由模式，hash模式，history模式
// createWebHistory() 是开启history模块   http://xxx/user
// createWebHashHistory() 是开启hash模式    http://xxx/#/user

// vite 的配置 import.meta.env.BASE_URL 是路由的基准地址，默认是 ’/‘
// https://vitejs.dev/guide/build.html#public-base-path
// 如果将来你部署的域名路径是：http://xxx/my-path/user
// vite.config.ts  添加配置  base: my-path，路由这就会加上 my-path 前缀了
// import index from "./components/index.vue";
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: '/',
      component: import('../components/main.vue'),
      redirect:'/index',
      children:[  
        {
          path: '/index',
          name: 'index',
          component: import('../components/index.vue'),
          meta:{
            index:1,
          },
          children:[]
        },{
          path: '/slogan',
          name: 'slogan',
          component: import('../components/slogan.vue'),
          meta:{
            index:2,
          },
          children:[]
        },{
          path: '/user',
          name: 'user',
          component: import('../components/user.vue'),
          meta:{
            index:3,
          },
          children:[]
        }
      ]
    },
    {
      path:'/video',
      name:'video',
      component: import('../components/video.vue'),
    },
    {
      path:'/selectSchool',
      name:'selectSchool',
      component: import('../components/selectSchool.vue'),
    },
    {
      path:'/article',
      name:'article',
      component: import('../components/article.vue'),
    },
    {
      path:'/topic',
      name:'topic',
      component: import('../components/topic.vue'),
    },
    {
      path:'/welcome',
      name:'welcome',
      component: import('../components/welcome.vue'),
    }

  ]
})

export default router
