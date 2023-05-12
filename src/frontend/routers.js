import TheHome from '@/frontend/components/TheHome.vue'
import SignUp from '@/frontend/components/SignUp.vue'
import TheLogin from '@/frontend/components/TheLogin.vue'
import UserAccount from '@/frontend/components/UserAccount.vue'
import RecycleBin from '@/frontend/components/RecycleBin.vue'
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