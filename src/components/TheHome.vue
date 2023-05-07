<template>
    <main>
        <h1>Main page</h1>
    </main>
</template>
<script>
export default {
    name: 'TheHome',

    mounted() {
        let user = localStorage.getItem('user-info')
        if(!user) {
            this.$router.push({name:'SignUp'})
        }
    }
}
</script>