<script setup lang="ts">
import { RouterView } from "vue-router";
import Nav from "./components/Nav.vue";
import Ws from "./Websocket";

Ws.connect();
</script>
<template>
  <body>
    <Nav />
    <main>
      <RouterView />
    </main>
  </body>
</template>

<style scoped>
</style>
