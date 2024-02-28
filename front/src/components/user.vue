<script setup>
import { ref } from "vue";
import navbar from "./indexTitle.vue";
</script>

<template>
  <div id="app">
    <div class="main">
      <div class="userInfo">
        <div class="userImg">
          <img src="../assets/videoImg/img1.png" />
        </div>
        <div class="info">
          <div class="nickname">昵称:  小明</div>
          <div class="account">账号:    456</div>
        </div>
      </div>
      <div class="route">
        <div class="routeName">生成路线记录:</div>
        <div class="context">
          <div class="routeContext" v-for="item in 20" :key="item">
            <div class="imgBox">出国</div>
          </div>
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
  height: 100%;

  /* background-color:blue; */
}
.main {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  padding: 3vw;
  width: 100%;
  // background: #f6f6f6;
}
.userInfo {
  width: 100%;
  display: flex;
  border-bottom: 3px solid #f6f6f6;
  padding: 3vw;
  background: #fff;
  border-radius: $my-border-radius;
  margin-bottom: 1vh;
  .userImg {
    margin-right: 5vw;
    img {
      display: block;
      width: 20vw;
      height: 20vw;
      border-radius: 50%;
    }
  }
  .info .nickname,
  .account {
    margin: 2.5vw 0;
  }
}
.route {
  width: 100%;
  border-radius: $my-border-radius;
  background: #fff;
  // .context {
  // overflow-y: auto;
  .routeName,
  .routeContext {
    height: 10vh;
    width: 100%;
    padding: 0 3vw;
    display: flex;
    align-items: center;
    border-bottom: 2px solid #f6f6f6;
  }
  // }
}
</style>
