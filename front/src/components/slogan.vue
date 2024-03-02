<script setup>
import { ref } from "vue";
// import navbar from "./indexTitle.vue";
import { useRouter } from "vue-router";

const router = useRouter();

const selectSchool = (index) => {
  router.push({ 
    name: "selectSchool",
    params:{
      id:index,
    },
   });
};

// 新增：定义四个不同目标的描述
const slogans = ref([
  { text: "保研：超越限界，探索未来这段路" },
  { text: "就业：拓展视野，再书职场新篇章" },
  { text: "出国：跨越边界，再读世界这本书" },
  { text: "考研：游历考海，开启学业新征程" }
]);

</script>

<template>
  <div id="app">
    <div class="main">
      <div
        class="sloganBox"
        v-for="(slogan, index) in slogans"
        :key="index"
        @click="selectSchool(index)"
      >
        <div class="text">{{ slogan.text }}</div>
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