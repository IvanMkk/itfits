<template>
<TheHeader />
<TheContent />
<TheFooter />
</template>

<script>
import TheHeader from './components/TheHeader.vue'
import TheContent from './components/TheContent.vue'
import TheFooter from './components/TheFooter.vue'

export default {
  name: 'App',
  components: {
    TheHeader,
    TheContent,
    TheFooter
  }
}
</script>

<style lang="less">
  @import '@/assets/css/clean.css';
  @import '@/assets/css/style.less';
</style>
