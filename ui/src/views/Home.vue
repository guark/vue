<template>
  <div class="home">
    <img width="300" src="https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/characters/svg/51.svg">
     <div>
         <button class="btn" @click="gcall()">Click Here To Call Go Function!</button> |
         <button class="btn" @click="clipboardWrite()">Call Plugin (Write "guark" to clipboard)!</button>
         <button class="btn" @click="warning()">Call Plugin (Dialog warning)!</button>
         <button class="btn" @click="info()">Call Plugin (Dialog info)!</button>
         <button class="btn" @click="notify()">Call Plugin (Notify)!</button>
         <button class="btn" @click="select()">Call Plugin (Dialog Select file)!</button>
     </div>
    <HelloWorld msg="Hello World Guark Vue Template!"/>
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

    clipboardWrite()
    {
      g.call("clipboard.write", {text: "guark"}).then(res =>
      {

        console.log("clipboard write success")

      }).catch(e => {

        console.error("write to clipboard error!")

      })
    },


    warning()
    {
      g.call("dialog.warning", {text: "this is a warning from js api! is this working?"})
    },

    info()
    {
      g.call("dialog.info", {text: "this is a info from js api! is this working?"})
    },


    notify()
    {
      g.call("notify.send", {text: "this is a notify test call from ui javascript api."})
    },

    select()
    {
      g.call("dialog.file", {title: "Select File For Guark APP!"}).then(res => {
        console.log("you selected", res)
      })
    }
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

