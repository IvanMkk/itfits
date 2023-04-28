<template>
    <h1>Sign Up</h1>
    <!--form action="" class="form"-->
        <div class="form__group">
            <input type="text" v-model="email" id="email" name="email" class="form__control" placeholder="Email">
            <label for="email" class="form__label none">Email</label>
        </div>
        <div class="form__group">
            <input type="password" v-model="password" id="password" name="password" class="form__control" placeholder="Password">
            <label for="password" class="form__label none">Password</label>
        </div>
        <button v-on:click="signUp" class="" id="sign-up">Sign Up</button>
    <!--/form-->
    <!--div>
        <form @‌submit.prevent="onSubmit">
            <label>Email:</label>
            <input type="email" v-model="email">
            <label>Password:</label>
            <input type="password" v-model="password">
            <button type="submit">Sign Up</button>
            <router-link to="/sign-in">Already have an account?</router-link>
        </form>
    </div-->
</template>

<script>
import axios from 'axios'
export default {
    name: 'SignUp',

    data() {
        return {
            email: '',
            password: '',
        }
    },

    methods: {
        async signUp() {
            let result= await axios.post("http://localhost:3000/users", {
                email:this.email,
                password:this.password
            })

            console.log(result)

            if(result.status == 201) {
                alert("sign up done")
                localStorage.setItem("user-info", JSON.stringify(result.data))
            }
        }
    }

    /*methods: {
        onSubmit() {
            this.$store.dispatch('signUp', { email: this.email, password: this.password })
            this.$router.push('/')
        },
    },*/
}
</script>