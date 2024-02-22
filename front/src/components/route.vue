<script setup>
import { ref } from "vue";
import navbar from "./navbar.vue";
const activeId = ref(1);
const activeIndex = ref(0);
const items = [
  {
    text: "浙江",
    children: [
      { text: "杭州", id: 1 },
      { text: "温州", id: 2 },
      { text: "宁波", id: 3, disabled: true },
    ],
  },
  {
    text: "江苏",
    children: [
      { text: "南京", id: 4 },
      { text: "无锡", id: 5 },
      { text: "徐州", id: 6 },
    ],
  },
  { text: "福建", disabled: true },
];
</script>

<template>
  <navbar />
  <div id="app">
    <div class="main">
      <div class="ljName">为您推荐如下路线</div>
      <van-steps direction="vertical" :active="100">
        <van-step>
          <h3>路线1</h3>
          <p>2016-07-12 12:40</p>
        </van-step>
        <van-step>
          <h3>路线2</h3>
          <p>2016-07-11 10:00</p>
        </van-step>
        <van-step>
          <h3>路线3</h3>
          <p>2016-07-11 10:00</p>
        </van-step>
        <van-step>
          <h3>路线4</h3>
          <p>2016-07-10 09:30</p>
        </van-step>
        <van-step>
          <h3>路线5</h3>
          <p>2016-07-11 10:00</p>
        </van-step>
        <van-step>
          <h3>路线6</h3>
          <p>2016-07-10 09:30</p>
        </van-step>
      </van-steps>
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
  // display: flex;
  // justify-content: space-between;
  // flex-wrap: wrap;
  // padding: 1vw 4vw;
  .ljName {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 5vh;
    font-size: 6vw;
  }
}
// .sloganBox {
//   margin: 3vw 0;
//   padding: 1.5vw;
//   width: calc(50vw - 4.5vw - 2vw);
//   height: 30vh;
//   border-radius: $my-border-radius;
//   background: #fff;
// }
</style>
