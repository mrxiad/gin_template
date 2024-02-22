<template>
  <div class="main">
    <div class="context">
        <div class="item" v-for="(Item,index) in 5" :key="index">
            <div class="title">
                {{ index+1 + '. ' }} 题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目
            </div>
            <div class="answer">
                <van-radio-group v-model="checkedList[index]" direction="horizontal">
                    <van-radio name="1">答案 1</van-radio>
                    <van-radio name="2">答案 2</van-radio>
                </van-radio-group>
            </div>
        </div>
        <van-button type="primary" block @click="submit">提交</van-button>
    </div>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import {useRouter} from 'vue-router'
const router = useRouter()
const checkedList = ref([])
const submit = () =>{
    localStorage.setItem('oneLoading', 0)
    router.push({path:'/'})
}
</script>

<style scoped lang="scss">
.main{
    padding: 5vh 10vw;
}
.item{
    height: 10vh;
    margin-bottom: 3vh;
    .title{
        margin-bottom: 1vh;
    }
}
</style>