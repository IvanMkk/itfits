import '@/assets/css/style.css'
import { createApp } from 'vue'
import router from '@/router/routers'
import store from "@store/store"
import App from './App.vue'

const app = createApp(App)

app.use(router)
app.use(store)
app.mount('#app')
