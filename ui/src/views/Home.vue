<template>
  <div class="home">
    <img width="300" src="https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/characters/svg/51.svg">
     <div>
         <button class="btn" @click="gcall()">Click Here To Call Go Function!</button>
     </div>
    <HelloWorld msg="Hello hot reload!"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

import g from 'guark'

export default {
  name: 'Home',
  components: {
    HelloWorld
  },

  methods:
  {
    gcall()
    {
        g.call("hello_my_func", {name: "guark working?"}).then(res => {

            console.log("call success", res)

        }).catch(e => {

            console.error("call failed", e)
        })
    },
  }
}
</script>

<style>
.btn {
    padding: 10px 30px;
    background: red;
    color: #FFF;
    border: 2px solid red;
}
</style>

