import '@/assets/css/style.css'
import { createApp } from 'vue'
import router from '@/router/routers'
import App from './App.vue'

createApp(App).use(router).mount('#app')