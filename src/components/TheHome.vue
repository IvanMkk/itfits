<template>
    <TheHeader />
    <main>
        <h1>Main page</h1>
    </main>
</template>
<script>
import TheHeader from './TheHeader.vue'
export default {
    name: 'TheHome',

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