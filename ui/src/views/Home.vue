<template>
  <div class="home">
    <img width="300" src="https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/characters/svg/51.svg">
    <HelloWorld msg="Hello hot reload!"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'Home',
  components: {
    HelloWorld
  }
}
</script>
