
<template>
  <v-typical
    class="blink"
    :steps="['Welcome to the Social Network', 700, 'Welcome to the Future', 700, 'Welcome home!', 700, 'Prepare to be immersed 🤖', 1000, 'Navigate on using the buttons above ☝️', 1000]"
    :loop=1
    :wrapper="'h2'"
  ></v-typical>
</template>

<script lang="ts">
import VTypical from "vue-typical";

export default {
  components: {
    VTypical,
  },
};
</script>


<style scoped>
h2 {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-size: x-large;
}
</style>
