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
                email: '', //embisda
                password: '', //embisda
                message: ''
            }
        },

        methods: {
            async login() {
                if(this.email != "" && this.password != "") {
                    /*const response = await axios.get(`${process.env.VUE_APP_BASE_URL}/users?email=${this.email}&password=${this.password}`)*/
                    const response = await axios.post(`${process.env.VUE_APP_BASE_URL}/users/auth`, {
                        email: this.email,
                        password: this.password
                       /* "data": {
                            email: this.email,
                            password: this.password
                        }*/
                    })
                    .then(response => {
                        if (response.response.status == 200 /*&& response.response.data.length > 0*/) {
                            localStorage.setItem('token', response.data.token)
                            localStorage.setItem('userInfo', JSON.stringify(response.response.data))
                            this.$router.push({name:'TheHome'})
                        }

                        //console.log(response)
                        //localStorage.setItem('token', response.data.token)
                        //this.$router.push({name:'TheHome'})
                    })
                    .catch(error => {
                        if(error.response.status == 400) {
                            return this.message = 'Sever received your request but the content was not valid'
                        }
                    })
                    console.log(response)
                } else {
                    return this.message = 'Fill in the email and password'
                }
            },                
        },

        mounted() {
            /*let user = localStorage.getItem('user')
            if(user) {
                this.$router.push({name:'TheHome'})
            }*/
        }
    }
</script>
