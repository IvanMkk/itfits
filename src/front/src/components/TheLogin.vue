<template>
    <TheHeader />
    <h1>Login</h1>
    <form class="form">
        <div class="form__group">
            <input type="text" v-model="email" id="email" name="email" class="form__control" placeholder="Email">
            <label for="email" class="form__label none">Email</label>
        </div>
        <div class="form__group">
            <input type="password" v-model="password" id="password" name="password" class="form__control" placeholder="Password">
            <label for="password" class="form__label none">Password</label>
        </div>
        <button type="submit" @click.prevent="login">Login</button>
        <!--button v-on:click="login" class="" id="login">Login</button-->
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
                password: '' //embisda
            }
        },

        methods: {
            async login() {
                if(this.email != "" && this.password != "") {
                    await axios.post(`${process.env.VUE_APP_BASE_URL}/users`, {
                        email: this.email,
                        password: this.password
                    })
                    .then(response => {
                        console.log(response)
                        console.log(response.data)
                        const user = response.data.find(user => user.email === this.email && user.password === this.password)
                        if(user.email === this.email && user.password === this.password) {
                            console.log(user.email)
                            console.log(user.password)
                            localStorage.setItem("user-info", JSON.stringify(user.data))
                            //this.$router.push({name:'TheHome'})
                        } else {
                            console.log('User does not exist')
                        }
                        //const user = response.data.find(user => user.email === this.email && user.password === this.password)
                    })
                    .catch(error => {
                        console.log(error)
                        alert("User not found, please register")
                        //this.$router.push({name:'SignUp'})
                    })
                }
            },                
        },

        mounted() {
            let user = localStorage.getItem('user')
            if(user) {
                //this.$router.push({name:'TheHome'})
            }
        }
    }
</script>
