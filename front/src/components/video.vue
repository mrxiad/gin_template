<script setup>
import { ref, computed,onMounted } from 'vue';
import { useRouter } from 'vue-router';
import navbar from './navbar.vue';
import $ from "jquery"
const router = useRouter();
const urlid = ref(0); // 默认值为0，可以根据实际情况调整

const imgList = ref([]);
const imgList2 = ref([]);

const getImgFile = async () => {
  for (let i = 1; i <= 4; i++) {
    const file = await import(`../assets/fourImg/img${i}.jpg`);
    if (file!=null){
      imgList.value.push(file.default);
    }
  }
  for (let i = 1; i <= 2; i++) {
    const file = await import(`../assets/videoImg/img${i}.jpg`);
    if (file!=null){
      imgList2.value.push(file.default);
    }
  }
};



// 使用计算属性来动态获取当前图片的 URL
const currentImgUrl = computed(() => {
  return imgList2.value[urlid.value] ? imgList2.value[urlid.value] : '';
});


const goPlayPage = (index) => {
  router.push({ 
    name: "article",  
    params:{
      id:index,
    },
   });
};

getImgFile();

let articleTitlesContents = ref([]);
let articleTitle = ref ([])

onMounted(() => {
  // 获取到当前的id，假设id是从URL的最后一部分获取的
  const index = window.location.href.lastIndexOf('/');
  const id = window.location.href.substring(index + 1);
  urlid.value = parseInt(id); // 将id转换为数组索引
  //请求,向后端8080/getTitle发送请求
  $.ajax({
    url: "http://localhost:8080/api/getTitleList",//返回数组四个对象，每个都是{title,content}
    type: "get",
    success: function (data) {
      articleTitlesContents.value = data.data;

      //对articleTitle赋值
      articleTitle.value = articleTitlesContents.value[id].title;
    },
  });
});

const srcUrls = ref([
  "//player.bilibili.com/player.html?aid=776578448&bvid=BV1T14y1P737&cid=932612858&p=1",
  "//player.bilibili.com/player.html?aid=618548088&bvid=BV1Eh4y1A7KT&cid=1271456273&p=1",
]);


const currentVideoUrl= computed(()=>{
  return srcUrls.value[urlid.value];
})



</script>

<template>
  <navbar />
  <div id="app">
    <div class="main">
      <div class="playBox">
        <!-- <iframe src="//player.bilibili.com/player.html?aid=551844951&bvid=BV1Gi4y1y7KZ&cid=514741932&p=1" scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"> </iframe> -->
        <iframe :src="currentVideoUrl" scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"></iframe>
      </div>
      <div class="videoInfo margin">
        <div class="userInfo">
          <div class="userImg">
            <img :src="currentImgUrl" />
          </div>
          <div class="info">
            <div class="nickname"></div>
            <div class="account"></div>
          </div>
        </div>
        <div class="videoName margin">推荐文章</div>
      </div>
      <div class="recommend margin">
        <div class="context">
          <div
            class="articleItem"
            v-for="(Item, index) in imgList"
            :key="index"
            @click="goPlayPage(index)"
          >
            <div class="left">
              <img :src="Item" />
            </div>
            <div class="right">
              <template v-if="articleTitlesContents.length > index">
                {{ articleTitlesContents[index].title }}
              </template>
            </div>
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
.margin {
  margin-top: 3vw;
}
.main {
  .playBox {
    width: 100%;
    height: 30vh;
    background-color: black;
  }
  .playBox iframe {
    width: 100%;
    height: 100%;
  }
  .videoInfo {
    padding: 3vw;
    .userInfo {
      display: flex;
      width: 100%;
      .userImg {
        margin-right: 3vw;
        img {
          width: 9vw;
          height: 9vw;
          border-radius: 50%;
        }
      }
      .info {
        display: flex;
        flex-direction: column;
        justify-content: center;
        .nickname {
          font-size: 3vw;
        }
        .account {
          // color: #f6f6f6f6;
        }
      }
    }
    .videoName {
      font-size: 4vw;
    }
  }
  .recommend {
    width: 100%;
    // border-radius: $my-border-radius;
    .articleItem {
      display: flex;
      padding: 2vw;
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
        // .top {
        font-size: 4vw;
        // }
        // .bottom {
        //   font-size: 3vw;
        // }
      }
    }
  }
}
</style>