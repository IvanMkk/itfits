import TheHome from './components/TheHome.vue'
import SignUp from './components/SignUp.vue'
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
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router