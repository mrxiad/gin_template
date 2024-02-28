<script setup>
import { ref } from "vue";
import navbar from "./navbar.vue";
const count = ref(10)
</script>

<template>
  <navbar />
  <div id="app">
    <div class="main">
      <div class="ljName">为您推荐如下路线</div>

      <div style="width: 100vw;" >
        <!-- 显示到右边的 -->
        <div v-for="item in  count" :class="item % 2==0?'left':'right'" :key="item" style="width: 100vw; display: flex; ">
          <div class="lineContext ">
            <div class="line" :class= "(item==count)? 'lineAfterTriangle' : 'lineAfterRound'" v-show="item % 2 == 1"></div>
            <div class="context">
                <h3 :class="item % 2==0?'right':'left'">路线1路线1路线1路线1路线1路线1路线1路线1路线1路线1路线1</h3>
                <p :class="item % 2==0?'right':'left'">2016-07-12 12:40</p>
            </div>
            <div class="line " :class= "(item==count)? 'lineAfterTriangle' : 'lineAfterRound' " v-show="item % 2 == 0"></div>
          </div>
        </div>
        <!-- 显示到左边的 -->
        <!-- <div style="width: 100vw; display: flex; justify-content: start;">
          <div class="lineContext"  style="display: flex;  width: 51vw; height: 13vh; ">
            <div class="context" style=" width: calc(50vw - 1vw); height: 100%; padding-right: 2vw;">
                <h3 class="right">路线1</h3>
                <p class="right">2016-07-12 12:40</p>
            </div>
            <div class="line" style="width: 1vw;height: 100%; background-color: aqua;"></div>
          </div>
        </div> -->
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
* {
  box-sizing: border-box;
  padding: 0;
  margin: 0;
}
#app {
  height: 100%;
}
.right{
  display: flex;
  justify-content: right;
}
.left{
  display: flex;
  justify-content: left;
}
.start{
  justify-content:start;
}
.end{
  justify-content:end;
}
.main {
  padding-bottom: 7vh;
  .ljName {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 5vh;
    font-size: 6vw;
    margin-bottom: 5vh;
    margin-top: 3vh;
  }
  .lineContext{
    display: flex; 
    width: 51vw; 
    .line{
      display: flex;
      justify-content: center;
      align-items: end;
      width: 1vw;
      height: 100%;
      // 线的颜色
      background-color: aqua;
    }
    .lineAfterTriangle::after{
      content: '';
      transform:rotate(180deg);
      display: inline-block;
      width: 0;
      height: 0;
      border-left: 2vw solid transparent;
      border-right: 2vw solid transparent;
      // 箭头的颜色
      border-bottom:2vw solid aqua; /* 调整颜色和高度以适应需求 */;
      background: #fff;
    }
    .lineAfterRound::after{
      content: '  ';
      display: block;
      position: absolute;
      width: 10px;
      height: 10px;
      border-radius: 50%;
      // 小球的颜色
      background-color: aqua;
      color: #fff;
    }
    .context{
      width: calc(50vw - 1vw); 
      height: 100%; 
      padding: 0 2vw;
    }
  }
}
</style>
