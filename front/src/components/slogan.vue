<script setup>
import { ref } from "vue";
// import navbar from "./indexTitle.vue";
import { useRouter } from "vue-router";
const router = useRouter();
const imgList = ref([]);
const getImgFile = async () => {
  for (let i = 1; i <= 4; i++) {
    const file = await import(`../assets/videoImg/img${i}.png`);
    imgList.value.push(file.default);
  }
};
getImgFile();
const selectSchool = () => {
  router.push({ path: "/selectSchool" });
};
</script>

<template>
  <div id="app">
    <div class="main">
      <div
        class="sloganBox"
        v-for="item in 4"
        :key="item"
        @click="selectSchool"
      >
        <div>考研</div>
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
  padding: 1vw 4vw;
}
.sloganBox {
  margin: 3vw 0;
  padding: 1.5vw;
  width: calc(50vw - 4.5vw - 2vw);
  height: 30vh;
  border-radius: $my-border-radius;
  background: #fff;
}
</style>
