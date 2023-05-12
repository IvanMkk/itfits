<template>
    <TheHeader />
    <main>
        <h1>Account</h1>
    </main>
</template>
<script>
import TheHeader from './TheHeader.vue'
export default {
    name: 'UserAccount',

    components: {
        TheHeader
    },

    mounted() {
        let user = localStorage.getItem('user-info')
        if(!user) {
            this.$router.push({name:'SignUp'})
        }
    }
}
</script>