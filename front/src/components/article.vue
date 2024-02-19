<script setup>
import { ref } from "vue";
import navbar from "./navbar.vue";
</script>

<template>
  <navbar />
  <div id="app">
    <div class="main">
      <div class="title margin">标题</div>
      <div class="info margin">信息</div>
      <div class="context">
        我的年快过完了，从初一到初七，从内蒙到济南 回家第一天拍的牛马群
        走的那天窗户上按完手印发现外面挺好看
        去机场的路上的雕塑，朋友说寓意：立马滚蛋 飞机降落前拍到的北京落脚点
        在济南第一天晚上济南下起鹅毛大雪（P了） 逛街遇到济南棣棣
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.margin {
  margin-bottom: 2vw;
}
.main {
  padding: 3vw;
  .title {
    font-size: 6vw;
  }
  .info {
    font-size: 3vw;
  }
  .context {
    line-height: 7vw;
  }
}
</style>
