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
            /*const user = JSON.parse(localStorage.getItem('user'));
            
            if (user) {
                this.$router.push('/TheHome');
            }*/
            /*
            let user = localStorage.getItem('user')
            if(user) {
                this.$router.push({name:'TheHome'})
            }*/

            /*const btnAddGarment = document.querySelector('#btn-new')
            const radioContainer = document.querySelector('.radio-container')
            const radioDetails = document.querySelectorAll('.radio-details')*/
            /*const inputBrand = document.querySelector('#brand')
            const inputType = document.querySelector('#type')
            const btnAdd = document.querySelector('#btn-add')
            const btnCncl = document.querySelector('#btn-cancel')*/
            /*const divItem = Array.from(document.querySelectorAll('.item p'));
            const divSizes = Array.from(document.querySelectorAll('.sizes p'));*/

            /*window.addEventListener('DOMContentLoaded', () => {
                const forms = document.querySelectorAll('.form')

                for (const form of forms) {
                    const groups = form.querySelectorAll('.form__group')

                    for (const group of groups) {
                        const control = group.querySelector('.form__control')
                        const labels = group.querySelector('.form__label')

                        if (control.value.length > 0) {
                            group.classList.add('form__group--active')
                            labels.classList.remove('none')
                        }

                        control.addEventListener('input', () => {
                            if (control.value.length > 0) {
                                group.classList.add('form__group--active')
                                labels.classList.remove('none')
                            } else {
                                group.classList.remove('form__group--active')
                                labels.classList.add('none')
                            }

                            if (inputBrand.value.length > 0 && inputType.value.length > 0) {
                                btnAdd.classList.remove('none')
                                btnCncl.classList.add('button-small')
                            } else {
                                btnAdd.classList.add('none')
                                btnCncl.classList.remove('button-small')
                            }
                        })
                    }
                }
            })*/
        }
    }
</script>
