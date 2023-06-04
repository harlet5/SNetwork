<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, onBeforeMount, computed } from "vue";
import { useAuthStore } from "@/stores/auth";
import Ws from "@/Websocket";
import {useNotifStore} from "../stores/notif"
import SingleNotification from "./SingleNotification.vue";
const storeNotif = useNotifStore()
const notifs = computed(() => storeNotif.notifications)
const store = useAuthStore();
const test = () => {
//onBeforeMount(() => {
  Ws.send({
      Page: "notifications",
      Data: {
        uid: store.userId,
      },
    }
  );
//});
}

const isOpen = ref(false);

function toggleDropdown(): void {
  isOpen.value = !isOpen.value;
}

</script>
<template>
  <div class="dropdown" @click="test()">
    <div @click="toggleDropdown">notifications <span class="notification-unread-count" v-if="storeNotif.count > 0">{{ storeNotif.count }}</span></div>
    <ul v-show="isOpen">
      <SingleNotification v-for="notif in notifs" :key="notif" :notif="notif" class="notifs" />
    </ul>
  </div>
  </template>
<style scoped>
.dropdown {
  position: relative;
  display: inline-block;
}

ul {
  list-style: none;
  padding: 0;
  margin: 0;
  position: absolute;
  top: 100%;
  left: -250px;
  background-color: gray;
  border: 1px solid #ccc;
  min-width: 375px;
  max-height: 300px;
  overflow-y: auto;
  z-index: 9;
}

.notifs {
  padding: 10px;
  cursor: pointer;
}

.notifs:hover {
  background-color: lightgray;
}

.notification-unread-count {
  padding: 5px;
  background-color: #fa3e3e;
  border-radius: 2px;
  color: white;
  padding: 1px 3px;
  font-size: 10px;
}
</style>

