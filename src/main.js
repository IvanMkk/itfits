import '@/assets/css/style.css'
import { createApp } from 'vue'
import router from './routers'
import App from './App.vue'

/*import router from '/routers'*/

createApp(App).use(router).mount('#app')