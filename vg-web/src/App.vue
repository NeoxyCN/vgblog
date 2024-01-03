<script setup>
import HelloWorld from './components/HelloWorld.vue'
import TheWelcome from './components/TheWelcome.vue'
import {ref} from 'vue'

const count=ref(0)

function increment(){
  count.value++
}
</script>

<template>
  <button @click="increment">{{ count }}</button>
</template>

<style scoped>
</style>
