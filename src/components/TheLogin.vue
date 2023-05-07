<template>
    <h1>Login</h1>
    <div class="form__group">
        <input type="text" v-model="email" id="email" name="email" class="form__control" placeholder="Email">
        <label for="email" class="form__label none">Email</label>
    </div>
    <div class="form__group">
        <input type="password" v-model="password" id="password" name="password" class="form__control" placeholder="Password">
        <label for="password" class="form__label none">Password</label>
    </div>
    <button v-on:click="login" class="" id="login">Login</button>
    <a href=""><router-link to="/signup">Sign Up</router-link></a>
</template>

<script>
    import axios from 'axios'
    export default {
        name: 'TheLogin',

        data() {
            return {
                email: '',
                password: ''
            }
        },

        methods: {
            async login() {
                let result = await axios.get(`http://localhost:3000/users?email=${this.email}&password=${this.password}`)
                console.log(result)
            
                if(result.status == 200 && result.data.length > 0) {
                    localStorage.setItem("user-info", JSON.stringify(result.data[0]))
                    this.$router.push({name:'TheHome'})
                }
            }
        },

        mounted() {
            let user = localStorage.getItem('user-info')
            if(user) {
                this.$router.push({name:'TheHome'})
            }
        }
    }
</script>