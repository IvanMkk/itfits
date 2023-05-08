import TheHome from './components/TheHome.vue'
import SignUp from './components/SignUp.vue'
import TheLogin from './components/TheLogin.vue'
import UserAccount from './components/UserAccount.vue'
import RecycleBin from './components/RecycleBin.vue'
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
    },
    {
        name: 'UserAccount',
        component: UserAccount,
        path: '/account'
    },
    {
        name: 'RecycleBin',
        component: RecycleBin,
        path: '/bin'
    }
]

const router = createRouter({
    routes,
    history: createWebHistory()
})

export default router