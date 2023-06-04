<script setup lang="ts">
import { RouterLink } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { ref } from "vue";
import Notifications from "./Notifications.vue";
import { useUnseenCountStore } from "@/stores/unseencount";
import { useNotifStore } from "@/stores/notif";

const store = useAuthStore();
const toggleMenu = ref(false);
</script>

<template>
  <div>
    <div class="navbar">
      <div class="title">Social Network</div>
      <a href="#" class="toggle-button">
        <span class="bar"></span>
        <span class="bar"></span>
        <span class="bar"></span>
      </a>
      <div class="navbar-links">
        <ul>
          <RouterLink class="tag" to="/"><li>home</li></RouterLink>
          <RouterLink class="tag" v-if="store.loggedIn" to="/posts"
            ><li>posts</li></RouterLink
          >
          <RouterLink class="tag" v-if="!store.loggedIn" to="/register"
            ><li>register</li></RouterLink
          >
          <RouterLink class="tag" v-if="!store.loggedIn" to="/login"
            ><li>login</li></RouterLink
          >
          <RouterLink class="tag" v-if="store.loggedIn" to="/logout"
            ><li>logout</li></RouterLink
          >
          <RouterLink class="tag" v-if="store.loggedIn" to="/chat"
            ><li>
              chat
              <span
                class="chat-unread-count"
                v-if="useUnseenCountStore().unseenchatcount > 0"
                >{{ useUnseenCountStore().unseenchatcount }}</span
              >
            </li>
          </RouterLink>
          <RouterLink class="tag" v-if="store.loggedIn" to="/profile"
            ><li>profile</li></RouterLink
          >
          <RouterLink class="tag" v-if="store.loggedIn" to="/groups"
            ><li>groups</li></RouterLink
          >
          <Notifications
            class="tag"
            v-if="store.loggedIn"
            @click="useNotifStore().clearNotifCount"
            ><li>notifications</li></Notifications
          >
        </ul>
      </div>
    </div>
  </div>
</template>

<style scoped>
* {
  box-sizing: border-box;
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #333;
  color: white;
  height: 56px;
  padding: 0px;
  font-family: "Open sans";
}
.title {
  font-size: 1.5rem;
  margin: 0.5rem;
}
.navbar-links ul {
  margin: 0;
  padding: 0;
  display: flex;
}

.navbar-links li {
  list-style: none;
}
.tag {
  text-decoration: none;
  color: white;
  padding: 1rem;
  display: block;
}
.tag:hover {
  background-color: #555;
}
.toggle-button {
  position: absolute;
  top: 0.75rem;
  right: 1rem;
  display: none;
  flex-direction: column;
  justify-content: space-between;
  width: 30px;
  height: 21px;
}

.toggle-button .bar {
  height: 3px;
  width: 100%;
  background-color: white;
  border-radius: 10px;
}

@media (max-width: 600px) {
  .toggle-button {
    display: flex;
  }
  .navbar-links {
    display: none;
    width: 100%;
  }
  .navbar {
    flex-direction: column;
    align-items: flex-start;
  }
  .navbar-links ul {
    flex-direction: column;
    width: 100%;
  }
  .navbar-links li {
    text-align: center;
  }

  .tag {
    padding: 0.5rem 1rem;
  }
  .navbar-links.active {
    display: flex;
  }
}

.chat-unread-count {
  padding: 5px;
  background-color: #fa3e3e;
  border-radius: 2px;
  color: white;
  padding: 1px 3px;
  font-size: 10px;
}

.groups-unread-count {
  padding: 5px;
  background-color: #fa3e3e;
  border-radius: 2px;
  color: white;
  padding: 1px 3px;
  font-size: 10px;
}
</style>
