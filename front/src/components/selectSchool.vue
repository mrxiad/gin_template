<script setup>
import { ref, onMounted } from 'vue';
import $ from 'jquery';
import navbar from './navbar.vue';

const routeDetails = ref(null); // 使用null作为初始值

onMounted(() => {
  let index = window.location.href.lastIndexOf('/');
  let id = window.location.href.substring(index + 1);
  console.log("当前页面的id为" + id);
  
  $.ajax({
    url: `http://localhost:8080/api/routes/${id}`, // 更新为您的后端API地址
    type: 'GET',
    dataType: 'json',
    success: function(data) {
      routeDetails.value = data; // 直接保存返回的对象
      console.log("接收到的路线数据:", data);
    },
  });
});

</script>

<template>
  <navbar />
  <div id="app">
    <div class="main">
      <div class="ljName">路线</div>
      <div style="width: 100vw;" v-if="routeDetails && routeDetails.Steps">
        <div v-for="(step, index) in routeDetails.Steps" :class="index % 2 == 0 ? 'left' : 'right'" :key="index" style="width: 100vw; display: flex;">
          <div class="lineContext">
            <!-- 条件渲染，根据步骤序号的奇偶性决定线条显示在左侧还是右侧 -->
            <div class="line" :class="(index == routeDetails.Steps.length - 1) ? 'lineAfterTriangle' : 'lineAfterRound'" v-show="index % 2 == 1"></div>
            <div class="context">
                <h3 :class="index % 2 == 0 ? 'right' : 'left'">{{ step.Description }}</h3>
                <p :class="index % 2 == 0 ? 'right' : 'left'">{{ step.Time }}</p>
            </div>
            <div class="line" :class="(index == routeDetails.Steps.length - 1) ? 'lineAfterTriangle' : 'lineAfterRound'" v-show="index % 2 == 0"></div>
          </div>
        </div>
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
