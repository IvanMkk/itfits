import '@/assets/css/style.css'
import { createApp } from 'vue'
import router from '@/router/routers'
import App from './App.vue'

const app = createApp(App)

app.use(router)
app.mount('#app')
