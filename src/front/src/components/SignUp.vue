<template>
    <TheHeader />
    <h1>Sign Up</h1>
    <form class="form" @submit.prevent="signUp">
        <div class="form__group">
            <input type="text" v-model="email" id="email" name="email" class="form__control" placeholder="Email">
            <label for="email" class="form__label none">Email</label>
            <p v-if="error === true" class="message-error">This user already exist</p>
        </div>
        <div class="form__group">
            <input type="password" v-model="password" id="password" name="password" class="form__control" placeholder="Password">
            <label for="password" class="form__label none">Password</label>
        </div>
        <button class="" id="sign-up">Sign Up</button>
        <p class="message-error">{{ message }}</p>
    </form>
    <a href=""><router-link to="/login">Login</router-link></a>
</template>

<script>
    import TheHeader from './TheHeader.vue'
    import axios from 'axios'
    export default {
        name: 'SignUp',

        components: {
            TheHeader
        },

        data() {
            return {
                email: '',
                password: '',
                error: false,
                message: ''
            }
        },

        methods: {
            async signUp() {                
                if(this.email != "" && this.password != "") {
                    await axios.post(`${process.env.VUE_APP_BASE_URL}/users`, {
                        email: this.email,
                        password: this.password
                    })
                    .then(response => {
                        localStorage.setItem('token', response.data.token)
                        this.$router.push({name:'TheLogin'})
                    })
                    .catch(error => {
                        if(error.response.status == 400) {
                            return this.error = true
                        }
                    })
                } else {
                    return this.message = 'Fill in the email and password'
                }
            }
        },

        mounted() {
            let token = localStorage.getItem('token')

            if(token) {
                this.$router.push({name:'TheHome'})
            }
        }
    }
</script>
