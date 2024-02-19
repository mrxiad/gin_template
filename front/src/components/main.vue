<script setup>
import index from "./index.vue";
import navbar from "./indexTitle.vue";
import { useRouter } from "vue-router";
const router = useRouter();
import { ref, onMounted } from "vue";
let active = ref(1);
const pageSwitch = (index) => {
  // active.value = index;
  if (index == 1) {
    router.push({ path: "/index" });
  }
  if (index == 2) {
    router.push({ path: "/slogan" });
  }
  if (index == 3) {
    router.push({ path: "/user" });
  }
};
onMounted(() => {
  // active.value;
  console.log("路由", router.currentRoute.value.meta.index);
  active.value = router.currentRoute.value.meta.index;
});
router.beforeEach((to, from, next) => {
  // console.log("路由", to.meta.index);
  active.value = to.meta.index;
  // if (to.path == "/login" || to.path == "/register") {
  next();
  // } else {
  //   alert("您还没有登录，请先登录");
  //   next("/login");
  // }
});
</script>


<template>
  <div id="app">
    <!-- <div class="top">
      <navbar/>
    </div> -->

    <div class="context">
      <!-- <index /> -->
      <router-view></router-view>
    </div>
    <div class="bottom">
      <div class="tabbar">
        <div @click="pageSwitch(1)">
          <van-icon
            name="wap-home-o"
            size="6vw"
            :color="active == 1 ? '#fa6400' : '#1989fa'"
          />
        </div>
        <div @click="pageSwitch(2)">
          <van-icon
            name="miniprogram-o"
            size="6vw"
            :color="active == 2 ? '#fa6400' : '#1989fa'"
          />
        </div>
        <div @click="pageSwitch(3)">
          <van-icon
            name="user-circle-o"
            size="6vw"
            :color="active == 3 ? '#fa6400' : '#1989fa'"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
* {
  box-sizing: border-box;
}
#app {
  /* height: 100vh; */
  width: 100vw;
  background: #f6f6f6;
}
.bottom {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  border-top: 2px solid #f6f6f6f6;
}
.bottom .tabbar {
  display: flex;
  justify-content: space-around;
  align-items: center;
  height: 7vh;
  background-color: white;
}
.bottom .tabbar .van-icon {
  position: static !important;
}
.context {
  /* padding: 0 3vw; */
  padding-bottom: 7vh;
  /* height: 100vh; */
}
</style>
