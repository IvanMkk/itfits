import TheHome from './components/TheHome.vue'
import SignUp from './components/SignUp.vue'
import TheLogin from './components/TheLogin.vue'
import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        name: 'TheHome',
        component: TheHome,
        path: '/'
    },
    {
        name: 'SignUp',
        component: SignUp,
        path: '/signup'
    },
    {
        name: 'TheLogin',
        component: TheLogin,
        path: '/login'
    }
]

const router = createRouter({
    routes,
    history: createWebHistory()
})

export default router