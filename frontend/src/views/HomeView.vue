<script setup lang="ts">
import axios from 'axios';
import TheWelcome from '../components/TheWelcome.vue';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
let apiUrl = (window as any).__ENV__?.API_URL;

const bla = () => {

  if (apiUrl == undefined) {
    apiUrl = "http://localhost:8080";
  }

  axios.get(apiUrl + '/call')
  console.log("bla called");
}
</script>

<template>
  <main>
    <TheWelcome />
    <Button @click="bla">CALLsssssssssssssssssssssss</Button>
  </main>
</template>
