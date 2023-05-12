import '@/frontend/assets/css/style.css'
//import '@/assets/css/style.css'
import { createApp } from 'vue'
import router from '@/frontend/routers'
import App from '@/frontend/App.vue'

/*import router from '/routers'*/

createApp(App).use(router).mount('#app')