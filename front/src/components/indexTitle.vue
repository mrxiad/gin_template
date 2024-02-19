<script setup>
import { ref } from "vue";

const props = defineProps({
  title: String,
});
</script>

<template>
  <div class="padding flexCenter top">
    {{ title }}
  </div>
</template>

<style scoped lang="scss">
.padding {
  padding: 0 5vw;
}

.flexCenter {
  display: flex;
  align-items: center;
}
.top {
  height: 7vh;
  // margin-bottom: 3vw;
  // border-top: 1px solid #f6f6f6f6;
  border-bottom: 1px solid #f6f6f6f6;
  background-color: #fff;
  font-size: 5vw;
  // border-radius: $my-border-radius $my-border-radius 0 0;
}
</style>
