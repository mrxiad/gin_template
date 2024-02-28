<template>
  <div class="main">
    <div class="context">
        <!-- 年级 -->
        <div class="item">
            <div class="title">1. 年级</div>
            <div class="answer">
                <van-radio-group v-model="checkedList[0]" direction="horizontal">
                    <van-radio name="大一">大一</van-radio>
                    <van-radio name="大二">大二</van-radio>
                    <van-radio name="大三">大三</van-radio>
                    <van-radio name="大四">大四</van-radio>
                </van-radio-group>
            </div>
        </div>
        <!-- 专业 -->
        <div class="item">
            <div class="title">2. 专业</div>
            <div class="answer">
                <van-radio-group v-model="checkedList[1]" direction="horizontal">
                    <van-radio name="经济学">经济学</van-radio>
                    <van-radio name="工商管理">工商管理</van-radio>
                    <van-radio name="会计学">会计学</van-radio>
                    <van-radio name="法学">法学</van-radio>
                </van-radio-group>
            </div>
        </div>
        <!-- 绩点 -->
        <div class="item">
            <div class="title">3. 绩点</div>
            <div class="answer">
                <van-radio-group v-model="checkedList[2]" direction="horizontal">
                    <van-radio name="2-2.5">2-2.5</van-radio>
                    <van-radio name="2.5-3">2.5-3</van-radio>
                    <van-radio name="3-3.5">3-3.5</van-radio>
                    <van-radio name="3.5-4">3.5-4</van-radio>
                </van-radio-group>
            </div>
        </div>
        <!-- 实习经历 -->
        <div class="item">
            <div class="title">4. 实习经历</div>
            <div class="answer">
                <van-radio-group v-model="checkedList[3]" direction="horizontal">
                    <van-radio name="暂无">暂无</van-radio>
                    <van-radio name="一段">一段</van-radio>
                    <van-radio name="两段">两段</van-radio>
                    <van-radio name="三段">三段</van-radio>
                    <van-radio name="四段及以上">四段及以上</van-radio>
                </van-radio-group>
            </div>
        </div>
        <!-- 竞赛经历 -->
        <div class="item">
            <div class="title">5. 竞赛经历（竞赛相关的创新创业学分）</div>
            <div class="answer">
                <input type="text" v-model="checkedList[4]" placeholder="填空____（分）"/>
            </div>
        </div>
        <!-- 学生干部经历 -->
        <div class="item">
            <div class="title">6. 学生干部经历</div>
            <div class="answer">
                <van-radio-group v-model="checkedList[5]" direction="horizontal">
                    <van-radio name="校级部长">校级部长</van-radio>
                    <van-radio name="校级部员">校级部员</van-radio>
                    <van-radio name="院级部长">院级部长</van-radio>
                    <van-radio name="院级部员">院级部员</van-radio>
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
const checkedList = ref(Array(6).fill('')) // 更新为六项
const submit = () => {
    localStorage.setItem('oneLoading', 0)
    router.push({path:'/'})
}
</script>

<style scoped lang="scss">
.main{
    padding: 5vh 10vw;
}
.item{
    height: auto; /* 更新为auto以适应内容 */
    margin-bottom: 3vh;
    .title{
        margin-bottom: 1vh;
    }
    input[type="text"] {
        width: 100%; /* 确保输入框宽度合适 */
        padding: 0.5em;
        margin-top: 0.5em; /* 为输入框添加上边距 */
    }
}
</style>
