import { createApp } from 'vue'
import App from './App.vue'

createApp(App).mount('#app')

/*ChatGPT*/
/*import Vue from 'vue';
import VueRouter from 'vue-router';
import routes from './routes';

Vue.use(VueRouter);

const router = new VueRouter({
    mode: 'history',
    routes: [
        { path: '/sign-up', component: SignUp },
        { path: '/sign-in', component: SignIn },
        { path: '/forgot-password', component: ForgotPassword },
        { path: '/restore-password/:token', component: RestorePassword },
        { path: '/', component: IndexPage },
    ],
})

new Vue({
    router
}).$mount('#app');*/