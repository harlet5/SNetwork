<script setup lang="ts">
import { useGroupsStore } from '@/stores/groups';
import { useAuthStore } from '@/stores/auth';
import { useRoute } from "vue-router";
import { ref, onBeforeMount, watch, computed } from "vue";
import CreateGroupPost from "./CreateGroupPost.vue";
import Post from "./Post.vue";
import CreateGroupEvent from "./CreateGroupEvent.vue";
import GroupInvite from "./GroupInvite.vue";
import GroupEvents from "./GroupEvents.vue";
import GroupChat from "./GroupChat.vue"
import Ws from "@/Websocket";
import type {user} from "../interfaces/interfaces";
import type {Ref} from "vue";

const store = useGroupsStore();
const group = defineProps(["group"]);
const aStore = useAuthStore();
const route = useRoute();
const groupId = Number(route.params.id);
//import {usePostsStore} from "../stores/posts"
//const store2 = usePostsStore()
const posts = computed(() => store.gPosts);
const events = computed(() => store.events);
const owner: Ref<user> = computed<user>(() => store.owner as user);
const uStatus = computed(() => store.uStatus);
const joinGroup = () => {
  Ws.send({
      Page: "grpreq",
      Data: { 
        sid: aStore.userId, 
        gid: groupId,
      }
    }
  );

}
</script>

<template>
  <div>
  <div class="group">
    <div class="group-header">
      <!--<h3 class="group-user">{{ group.group.Creator }}</h3>-->
    </div>
    <div v-if="uStatus === true" class="group-container">
        <div class="left-sidebar">
          <div class="group-body-div">
            <p class="about-community">About community</p>
            <h2 class="group-title">{{ group.group.Name }}</h2>
            <p class="group-body">{{ group.group.Text }}</p>
          </div>
          <div class="events-section">

            <CreateGroupEvent />
            <GroupEvents v-for="event in events" :key="event" :event="event" />
          </div>
        </div>
        <div class="posts-section">
          <div class="adding-post">
            <CreateGroupPost />
          </div>  
          <Post v-for="post in posts" :key="post" :post="post" :style="{width: 70 + '%'}"/>
        </div>
        <div class="group-chat-section">
          <div class="placeholder">
            <div v-if="owner.UId === aStore.userId">
              <GroupInvite />
            </div>
            <div v-if="owner.UId !== aStore.userId" class="chat-title-div">
              <p class="chat-title">group chat</p>
            </div>
          </div>
          <div class="group-chat-placeholder">
            <GroupChat/>
          </div>
        </div>
    </div>
    <div v-if="uStatus === false" class="join-div">
      <button class="join-group-btn" @click="joinGroup()">Join group</button>
    </div>
    
  </div>
  </div>
</template>

<style scoped>
* {
  margin: 0px;
  padding: 0px;
}

.placeholder {
  height: 50px;
  width: 100%;
}

.group-container {
  display: flex;
  width: 100%;
  height: calc(100vh - 56px);
}

.chat-title-div {
  width: 100%;
  text-align: center;
}

.chat-title {
  display: inline-block;
  margin-top: 15px;
  font-size: 1.2rem;
  letter-spacing: .15rem;
  color: #4a4a4a;
  text-transform: uppercase;
}

.left-sidebar {
  width: 25%;
  box-sizing: border-box;
  background-color: #292a2d;
  display: flex;
  flex-direction: column;
  align-items: center;
  overflow-y: scroll;
}

.group-body-div {
  box-sizing: border-box;
  padding: 5px 15px 15px 15px;
  margin: 50px 0px 40px 0px;
  border: 1px solid #494948;
  background-color: #333;
  border-radius:15px;
  height: auto;
  width: 75%;
  color: white;
}

.about-community {
  font-size: 0.6rem;
  margin-bottom: 5px;
}

.group-title {
  font-size: 1.5rem;
  margin: 0px 0px 10px;
}

.group-body {
  font-size: 0.8rem;  
}

.events-section {
  width: 75%;
}

.posts-section {
  width: 50%;
  overflow-y: auto;
}

.group-chat-section {
  width: 25%;
  box-shadow: -5px 0 5px -5px #333;
  padding-right: 10px;
  height: 100%
}

.join-div {
  width: 100%;
  display: flex;
  align-items: center;
}

.join-group-btn {
  display: inline-block;
  width: fit-content;
  background-color: #d9d9d9;
  padding: 1rem 5rem;
  border-radius: 10rem;
  margin: 0 auto;
  margin-top: 30px;
  font-size: 1rem;
  letter-spacing: .15rem;
  color: #4a4a4a;
  text-transform: uppercase;
  position: relative;
  overflow: hidden;
  cursor: pointer;
}

</style>

