<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import navbar from "./navbar.vue";
const router = useRouter();
const imgList = ref([]);
const getImgFile = async () => {
  for (let i = 1; i <= 4; i++) {
    const file = await import(`../assets/videoImg/img${i}.png`);
    imgList.value.push(file.default);
  }
};
const props = defineProps({
  title: String,
});
getImgFile();

const goPlayPage = () => {
  router.push({ path: "/video", query: { id: 1 } });
};

// 新增全屏方法(没用)
const requestFullScreen = () => {
  const element = document.querySelector(".playBox iframe");
  if (element.requestFullscreen) {
    element.requestFullscreen();
  } else if (element.mozRequestFullScreen) {
    element.mozRequestFullScreen();
  } else if (element.webkitRequestFullscreen) {
    element.webkitRequestFullscreen();
  } else if (element.msRequestFullscreen) {
    element.msRequestFullscreen();
  }
};

</script>

<template>
  <navbar />
  <div id="app">
    <div class="main">
      <div class="playBox" @click="requestFullScreen">
        <iframe src="//player.bilibili.com/player.html?aid=551844951&bvid=BV1Gi4y1y7KZ&cid=514741932&p=1" scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"> </iframe>
      </div>
      <div class="videoInfo margin">
        <div class="userInfo">
          <div class="userImg">
            <img src="../assets/videoImg/img1.png" />
          </div>
          <div class="info">
            <div class="nickname">昵称</div>
            <div class="account">日期</div>
          </div>
        </div>
        <div class="videoName margin">视频名称</div>
      </div>
      <div class="recommend margin">
        <div class="context">
          <div
            class="articleItem"
            v-for="(Item, index) in imgList"
            :key="index"
            @click="goPlayPage"
          >
            <div class="left">
              <!-- :src="Item" -->
              <img :src="Item" />
            </div>
            <div class="right">
              公考政策热点公考政策热点公考政策热点公考政策热点公考政策热点公考政策热点公考政策热点
              <!-- <div class="top">公考政策热点</div> -->
              <!-- <div class="bottom">建立高级思维</div> -->
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
