<template>
 
  <div class="main">
    <div class="contextBox" >
        <!-- <div class="top"> -->
            <div class="close" @click.stop="close">关闭</div>
        <!-- </div> -->
        <!-- <div class="bottom"> -->
            <div class="left vertical-text">欢迎使用您的 </div>
            <div class="right vertical-text">人生指南针</div>
        <!-- </div> -->
    </div>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import {useRouter} from 'vue-router'
const router = useRouter()
const close = () => {
    router.push({path: '/topic'})
}
const buttonClicked = () =>{}

</script>

<style scoped lang="scss">
.main{
    position: fixed;
    z-index: 99999;
    display: flex;
    justify-content: center;
    align-content: center;
    width: 100vw;
    height: 100vh;
    background-color: rgba(255,255,255,.9);
    .contextBox{
        width: 80vw;
        height: 75vh;
        position: relative;
        border: 1px solid black;
        margin-top: 12.5vh;
        display: flex;
        justify-content: center;
        // align-items: center;
        background-color: rgba(255,255,255);
        .close{
            width: 10vw;
            height: 8vw;
            position:absolute;
            top: 2vw;
            right: 2vw;
            font-size: 5vw;
        }
        .left,.right{
            font-size: 8vw;
            padding-top: 35%;
        }
        .left{
            
        }
        .right{
           margin-top: 40%;
        }
    }
}
.vertical-text {
  writing-mode: vertical-rl; /* 文字从上到下，从右到左 */
//   transform: rotate(180deg); /* 可选，使文字从上到下，从左到右 */
}
</style>