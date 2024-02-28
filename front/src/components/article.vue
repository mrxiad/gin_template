<script setup>
import $ from "jquery";
import { ref } from "vue";
import navbar from "./navbar.vue";
import { onMounted } from "vue";
let welcomeIsDisplay = ref('')


let articleTitlesContents = ref([]);
let articleContent = ref([])
let articleTitle = ref ([])
onMounted(() => {
  welcomeIsDisplay.value = localStorage.getItem('oneLoading')?localStorage.getItem('oneLoading'):1
  //请求,向后端8080/getTitle发送请求
  $.ajax({
    url: "http://localhost:8080/api/getTitleList",//返回数组四个对象，每个都是{title,content}
    type: "get",
    success: function (data) {
      articleTitlesContents.value = data.data;

      //找到自己url的id
      let index = window.location.href.lastIndexOf('/');
      let id = window.location.href.substring(index+1);
      console.log("当前页面的id为"+id)

      //对articleConent赋值
      articleContent.value = articleTitlesContents.value[id].content;
      articleContent.value = articleContent.value.replace(/\n/g, '<br>');


      //对articleTitle赋值
      articleTitle.value = articleTitlesContents.value[id].title;


      //输出文章标题
      console.log(articleContent.value.substring(0,5))
    },
  });
})

</script>

<template>
  <navbar />
  <div id="app">
    <div class="main">
      <div class="title margin">{{articleTitle}}</div>
      <div class="info margin"></div>
      <div class="context">    
        <div v-html="articleContent"></div>
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
