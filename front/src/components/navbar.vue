<template>
  <!-- <div class="main"> -->
  <!--        -->
  <van-nav-bar
    left-text="返回"
    :title="title"
    left-arrow
    @click-left="onClickLeft"
  />
  <!-- </div> -->
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();
const props = defineProps({
  title: String,
});
const onClickLeft = () => {
  router.go(-1);
};
</script>

<style lang="scss" scoped></style>
