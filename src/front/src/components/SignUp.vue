<template>
    <TheHeader />
    <h1>Sign Up</h1>
    <form class="form">
        <div class="form__group">
            <input type="text" v-model="email" id="email" name="email" class="form__control" placeholder="Email">
            <label for="email" class="form__label none">Email</label>
        </div>
        <div class="form__group">
            <input type="password" v-model="password" id="password" name="password" class="form__control" placeholder="Password">
            <label for="password" class="form__label none">Password</label>
        </div>
        <!--button @click.prevent="signUp" class="" id="sign-up">Sign Up</button-->
        <button v-on:click="signUp" class="" id="sign-up">Sign Up</button>
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
            }
        },

        methods: {
            async signUp() {                
                let result = await axios.post(`${process.env.VUE_APP_BASE_URL}/users`, {
                    email: this.email,
                    password: this.password
                })

                alert(result)
                alert(JSON.stringify(result.data))

                if(result.status === 200) {
                    localStorage.setItem("user", JSON.stringify(result.data))
                    this.$router.push({name:'TheHome'})
                }
/*
                result = axios.get('http://0.0.0.0:3000/v1/users', {
                    email: this.email,
                    password: this.password
                })

                if(result.status !== 200) {
                    console.log(JSON.stringify(result.data))
                    
                    return
                }*/

                //localStorage.setItem("user-info", JSON.stringify(result.data))
            }
        },

        mounted() {
            /*const user = JSON.parse(localStorage.getItem('user'));
            
            if (user) {
                this.$router.push('/TheHome');
            }*/
            let user = localStorage.getItem('user')
            if(user) {
                this.$router.push({name:'TheHome'})
            }

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
