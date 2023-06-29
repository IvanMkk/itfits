<template>
    <header>
        <nav>
            <a href=""><router-link to="/account">Account</router-link></a>
            <a v-on:click="signout()" href="#" class="out">Sign Out</a>
            <a href=""><router-link to="/bin">Trash</router-link></a>
            <a href=""><router-link to="/">Wardrobe</router-link></a>
        </nav>
    </header>
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