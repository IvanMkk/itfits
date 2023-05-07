<template>
    <header>
        <nav>
            <a href="#">Log In</a>
            <a v-on:click="signout()" href="#">Sign Out</a>
            <a href="#">Trash</a>
            <a href="#">Wardrobe</a>
        </nav>
    </header>
    <h1>Header</h1>
</template>

<script>
    export default {
        name: 'TheHeader',

        methods: {
            signout() {
                localStorage.clear()
                this.$router.push({name:'TheLogin'})
            }
        }
    }
</script>