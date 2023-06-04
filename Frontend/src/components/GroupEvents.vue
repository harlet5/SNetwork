<script setup lang="ts">
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import Ws from "@/Websocket";
const store = useAuthStore();
const router = useRouter();
const event = defineProps(["event"]);
const navigateToProfile = (id: number) => {
  router.replace(`/user/${id}`);
};
const eventstatus = (eId: number, gId: number, stts: string) => {
  Ws.send({
      Page: "setevent",
      Data: { 
        uid: store.userId,
        eid: eId, 
        gid: gId,
        status: stts,
      }
    }
  );
}
</script>
<template>
  <div class="event">
    <div class="event-header">
      <div class="event-name">{{ event.event.Name }}</div>
      <div class="author-name" @click="navigateToProfile(event.event.Creator)">Organiser: {{ event.event.UName }}</div>
      <div class="datetime">{{event.event.Time}}</div>
    </div>
    <div class="event-content">
      <p class="event-body">{{ event.event.Text }}</p>
      <div class="going-counter">
        <span class="yes-count">{{ event.event.YesCount }} going</span>
        <span class="no-count">{{ event.event.NoCount }} not going</span>
        <div class="buttons">
          <button class="going-btn"  @click="eventstatus(event.event.Id, event.event.Gid, 'yes')">Going</button>
          <button class="not-going-btn"  @click="eventstatus(event.event.Id, event.event.Gid, 'no')">Not going</button>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.event {
  padding: 15px;
  border: 1px solid #494948;
  background-color: #333;
  font-size: 0.8rem;
  max-width: 250px;
  margin: 10px 10px 10px 0px;
  border-radius: 3px;
  color: white;
}

.event-name {
  font-weight: bold;
  font-size: 1.2rem;
  margin: 5px;
  text-align: center;
}

.event-body {
  font-size: 0.9rem;
}

.going-counter {
  display: flex;
  justify-content: space-between;
}

.buttons {
  display: inherit;
  gap: 10px;
}

.going-btn,
.not-going-btn {
  border: none;
  padding: 2px 5px 2px 5px;
  border-radius: 3px;
  background-color: transparent;
  cursor: pointer;
}

.going-btn {
  background-color: #4caf50;
  color: white;
}

.not-going-btn {
  background-color: #575756;
  color: white;
}

</style>
