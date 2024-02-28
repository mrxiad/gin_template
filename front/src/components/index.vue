<script setup>
import $ from "jquery";
import { onMounted, ref } from "vue";
import indexTitle from "./indexTitle.vue";
import { useRouter } from "vue-router";
import welcome from './welcome.vue'
const router = useRouter();
const welcomeIsDisplay = ref('')
const imgList = ref([]);
const articleTitles = ref([]);
const getImgFile = async () => {
  for (let i = 1; i <= 4; i++) {
    const file = await import(`../assets/videoImg/img${i}.png`);
    imgList.value.push(file.default);
  }
};
getImgFile();
const goPlayPage = () => {
  router.push({ path: "/video" });
};
const goArticlePage = (index) => {
  if (index === undefined) {
    console.error('Article ID is required');
    return; // 提前退出，不执行导航
  }
  console.log(index);
  router.push({ 
    name: 'article',
    params: { id: index },
  });
};
onMounted(() => {
  welcomeIsDisplay.value = localStorage.getItem('oneLoading')?localStorage.getItem('oneLoading'):1

  //请求,向后端8080/getTitle发送请求
  $.ajax({
    url: "http://localhost:8080/api/getTitleList",//返回数组四个对象，每个都是{title,content}
    type: "get",
    success: function (data) {
      articleTitles.value = data.data;
      console.log(articleTitles.value);
    },
  });
})


</script>

<template>
  <div id="app">
    <welcome v-if="welcomeIsDisplay == 1"></welcome>
    <div class="main">
      <div class="videoBox">
        <indexTitle :title="'视频专区'"></indexTitle>
        <div class="videoBoxMain">
          <div
            v-for="(Item, index) in imgList"
            :key="index"
            @click="goPlayPage"
          >
            <img :src="Item" style="width: 40vw; height: 12vh" />
          </div>
        </div>
      </div>
      <div class="article">
        <indexTitle :title="'文章专区'"></indexTitle>
        <div class="context">
          <div
            class="articleItem"
            v-for="(Item, index) in imgList"
            :key="index"
            @click="goArticlePage(index)"
          >
            <div class="left">
              <img :src="Item" />
            </div>
            <div class="right">
              <div class="top">{{articleTitles[index].title}}</div>
              <div class="bottom">建立高级思维</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
#app {
  height: 100%;
  /* background-color:blue; */
}
.main {
  overflow: auto;
  padding: 0 2vw;
}

.videoBoxMain {
  display: flex;
  overflow: hidden;
  overflow-x: scroll;
  margin-bottom: 3vw;
  background: #fff;
  border-radius: $my-border-radius;
  div {
    width: 40vw;
    height: 12vh;
    margin: 0 0.5vw;
  }
}
// .videoBoxMain div {
//   width: 40vw;
//   height: 15vh;
//   margin: 0 0.5vw;
// }
.article {
  margin: 3vw 0;
  overflow: hidden;
  .context {
    // padding: 0 3vw;
    border-radius: $my-border-radius;
    .articleItem {
      display: flex;
      padding: 3vw;
      border-top: 1px solid #f6f6f6f6;
      border-bottom: 1px solid #f6f6f6f6;
      background: #fff;
      .left img {
        width: 30vw;
        height: 10vh;
        border-radius: $my-border-radius;
      }
      .right {
        padding-left: 1vw;
        .top {
          font-size: 5vw;
        }
        .bottom {
          font-size: 3vw;
        }
      }
    }
  }
}
// .articleItem {
//   display: flex;
//   padding: 3vw;
//   border-top: 1px solid #f6f6f6f6;
//   border-bottom: 1px solid #f6f6f6f6;
//   background: #fff;
//   .left img {
//     width: 30vw;
//     height: 10vh;
//     border-radius: $my-border-radius;
//   }
//   .right {
//     padding-left: 1vw;
//     .top {
//       font-size: 5vw;
//     }
//     .bottom {
//       font-size: 3vw;
//     }
//   }
// }
// .articleItem .left img {
//   width: 30vw;
//   height: 10vh;
//   border-radius: $my-border-radius;
// }
// .articleItem .right {
//   padding-left: 1vw;
// }
// .articleItem .right .top {
//   font-size: 5vw;
// }
// .articleItem .right .bottom {
//   font-size: 3vw;
// }
</style>
