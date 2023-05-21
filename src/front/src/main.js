import '@/assets/css/style.css'
import { createApp } from 'vue'
import router from '@/router/routers'
import App from './App.vue'

const app = createApp(App)
/*const cors = require("cors")
const corsOptions = {
   origin:'*', 
   credentials:true,
   optionSuccessStatus:200,
}

app.use(cors(corsOptions))*/
app.use(router)

app.mount('#app')
