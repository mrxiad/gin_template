import { createApp } from 'vue'
import App from './App.vue'
const app = createApp(App)
import vant from 'vant'
import 'vant/lib/index.css'  //引入vant所有组件样式
import router from './router/index'
import "@/styles/index.scss"
app.use(vant)
app.use(router)
app.mount('#app') // 挂载一定要放到最后