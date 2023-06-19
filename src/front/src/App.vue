<template>
  <router-view />
</template>

<script>
  import { USER_REQUEST } from "@/store/actions/user";

  export default {
    name: 'App',
    
    created: function() {
      if (this.$store.getters.isAuthenticated) {
        this.$store.dispatch(USER_REQUEST);
      }
    }
  }
</script>
