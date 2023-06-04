<script setup lang="ts">
import { ref, watch, computed } from "vue";
//import type { Ref } from "vue";
import { useAuthStore } from "@/stores/auth";
import { useRoute } from "vue-router";
import Ws from "@/Websocket";
import { useGroupsStore } from '@/stores/groups';
import SingleGroupInv from "./SingleGroupInv.vue";

const outsider = defineProps(["outsider"]);

const storeG = useGroupsStore();
const route = useRoute();
const toggleShowEvent = ref(false)
//const store = useAuthStore()
//const groupId = Number(route.params.id);

const outsiders = computed(() => storeG.outsiders);

const showEvent = () => {
  toggleShowEvent.value = !toggleShowEvent.value
  isOpen.value = !isOpen.value;
}
const isOpen = ref(false);
//const selectedOption = ref('Select an option');
//
//const options = outsider.outsider.UName;


//function selectOption(option: string): void {
  //selectedOption.value = option;
  //isOpen.value = false;
//}

//const inviteNew = () => { 
//  console.log("the invite is sent to:", outsider.outsider.UId)
//   };
</script>

<template>
  <button @click="showEvent()">Invite new group member</button>
  <div class="overlay" :style="{display: toggleShowEvent ? 'block' : 'none'}">
    <div class="dropdown">
      <ul class="dropdown-menu" :class="{ open: isOpen }">
        <SingleGroupInv v-for="outsider in outsiders" :key="outsider" :outsider="outsider" /> 
      </ul>
    </div>
  </div>
  
</template>

<style scoped>

button {
    display: block;
    margin: 0 auto;
    margin-top: 10px;
    margin-bottom: 15px;
    background-color: #d9d9d9;
    padding: 0.5rem 1rem;
    border-radius: 10rem;
    border: none;
    font-size: 0.6rem;
    letter-spacing: .15rem;
    color: #4a4a4a;
    text-transform: uppercase;
    cursor: pointer; 
}

button:hover {
  background-color: #4a4a4a;
  color: white;
}

.dropdown {
  position: relative;
  width: 100%;
}

.dropdown-menu {
  list-style: none;
  width: 100%;
  padding: 0;
  margin: 0;
  position: absolute;
  top: 100%;
  left: 0;
  background-color: white;
  min-width: 100px;
  max-height: 200px;
  overflow-y: auto; /* Add scrollbar when the height exceeds max-height */
  box-shadow: 0 5px 5px -5px #333;
}

li {
  padding: 10px;
  cursor: pointer;
}

li:hover {
  background-color: #f2f2f2;
}

.open {
  display: block;
}

/* Hide the dropdown menu when it's closed */
.dropdown-menu:not(.open) {
  display: none;
}
</style>
