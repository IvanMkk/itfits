<template>
  <h1>Hello there</h1>
  <SignUp />
  <!--SignIn />
  <ForgotPassword />
  <RestorePassword />
  <IndexPage /-->
</template>

<script>
import SignUp from './components/SignUp.vue'
/*import SignIn from './components/SignIn.vue'
import ForgotPassword from './components/ForgotPassword.vue'
import RestorePassword from './components/RestorePassword.vue'
import IndexPage from './components/IndexPage.vue'*/

export default {
  name: 'App',

  components: {
    SignUp/*,
    SignIn,
    ForgotPassword,
    RestorePassword,
    IndexPage*/
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>