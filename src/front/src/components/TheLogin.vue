<template>
    <TheHeader />
    <h1>Login</h1>
    <form class="form" @submit.prevent="login">
        <div class="form__group">
            <input type="text" v-model="email" id="email" name="email" class="form__control" placeholder="Email">
            <label for="email" class="form__label none">Email</label>
        </div>
        <div class="form__group">
            <input type="password" v-model="password" id="password" name="password" class="form__control" placeholder="Password">
            <label for="password" class="form__label none">Password</label>
        </div>
        <button class="" id="login">Login</button>
        <p class="message-error">{{ message }}</p>
    </form>
    <a href=""><router-link to="/signup">Sign Up</router-link></a>
</template>

<script>
    import TheHeader from './TheHeader.vue'
    import axios from 'axios'
    export default {
        name: 'TheLogin',

        components: {
            TheHeader
        },

        data() {
            return {
                email: '',
                password: '',
                message: ''
            }
        },

        methods: {
            async login() {
                if(this.email != "" && this.password != "") {
                    await axios.post(`${process.env.VUE_APP_BASE_URL}/users/auth`, {
                        email: this.email,
                        password: this.password
                    })
                    .then(response => {
                        if (response.status == 200) {
                            localStorage.setItem('token', response.data.token.auth_token)
                            localStorage.setItem('user', this.email)
                            this.$router.push({name:'TheHome'})
                        }
                    })
                    .catch(error => {
                        console.log(error)
                        if(error.response.status == 400) {
                            return this.message = this.email + ' user does not exist or wrong password'
                        }
                    })
                } else {
                    return this.message = 'Fill in the email and password'
                }
            },                
        },

        mounted() {
            let token = localStorage.getItem('token')

            if(token) {
                this.$router.push({name:'TheHome'})
            }
        }
    }
</script>
