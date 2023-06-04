<script setup lang="ts">
import { ref, computed } from "vue";
import type { event } from "../interfaces/interfaces";
import type { Ref } from "vue";
import { useAuthStore } from "@/stores/auth";
import { timestamp } from "@vueuse/shared";
import { useRoute } from "vue-router";
import Ws from "@/Websocket";
import { useGroupsStore } from "@/stores/groups";
const gstore = useGroupsStore()
const eName = ref("");
const eDesc = ref("");
const eDate = ref("");
const eTime = ref("");

const toggleShowEvent = ref(false)
const store = useAuthStore()
const route = useRoute();
const groupId = Number(route.params.id);
const eventError = ref("")
const events: Ref<event[]> = computed<event[]>(() => gstore.events as event[])
console.log("events: ", events.value)

const showEvent = () => {
  toggleShowEvent.value = !toggleShowEvent.value
}

interface Event {
  uid: number;
  name: string;
  desc: string;
  time: string;
  gid: number;
}

const createEvent = () => {
  console.log("1")
  const eDateTime = eDate.value + " " + eTime.value
  if (new Date(eDateTime).getTime() < new Date().getTime()) {
    return eventError.value = "New events must happen in future"
  } 
  if (eName.value.length === 0 || eDesc.value.length === 0 || eDateTime.length < 16){
    return eventError.value = "Fill all fields"
  }
  if (events.value !== null && events.value.some((event)=> event.Name === eName.value)){
    return eventError.value = "Event name is already used"
  }
  console.log("2")
  const newEvent:Event = {
    uid: store.userId,
    name: eName.value,
    desc: eDesc.value,
    time: eDateTime,
    gid: groupId,
  }
  console.log(newEvent);
  Ws.send({
      Page: "makeevent",
      Data: 
        newEvent
    }
  ); 
  eName.value = "";
  eDesc.value = "";
  eDate.value = "";
  eTime.value = "";
  showEvent()
};
</script>

<template>
  <button class="create-btn" @click="showEvent()">+ Create new event</button>
  <div class="overlay" :style="{display: toggleShowEvent ? 'block' : 'none'}">
    <div class="modal">
      <div class="flexcreatepost">
        <div class="createpost">
          <input
            name="eName"
            id="eName"
            class="eName"
            placeholder='Name of the event'
            v-model="eName"
          />
          <textarea
            name="eDesc"
            id="eDesc"
            cols="30"
            rows="10"
            class="eDesc"
            placeholder="Event description"
            v-model="eDesc"
          ></textarea>
          <label class="eDate-label" for="eDate">Event time:</label>
          <input 
            type="date"
            id="eDate"
            name="eDate"
            class="eDate"
            v-model="eDate"
            >
          <input 
            type="time"
            id="eTime"
            timeformat="24h"
            name="eTime"
            class="eTime"
            v-model="eTime"
            >
        </div> 
        <div class="err" v-if="eventError">{{ eventError }}</div>
        <button class="submit-btn" @click="createEvent()">Submit</button>
      </div>
    </div>
  </div>
  
</template>

<style scoped>
.err {
  color: red;
}
.create-btn {
    display: block;
    background-color: transparent;
    border-radius: 10rem;
    border: none;
    font-size: 0.8rem;
    letter-spacing: .15rem;
    color: #a9a9b3;
    text-transform: uppercase;
    cursor: pointer; 
}

.create-btn:hover {
  color: white;
}

.eName,
.eDesc,
.eDate,
.eTime {
  background-color: transparent;
  border: 1px solid #626260;
  color: #a9a9b3;
  outline: none;
}

.eName:focus,
.eDesc:focus,
.eDate:focus,
.eTime:focus {
  border: 1px solid white;
}


.eName {
  margin: 10px 0px 5px 0px;
  padding: 5px;
  min-width: 250px;
}

.eDesc {
  resize: none;
  padding: 5px;
  min-width: 250px;
  margin-bottom: 10px;
}

.eDate-label {
  display: block;
  font-size: 0.8rem;
  font-weight: bold;
  color: white;
}

.eTime {
  margin-bottom: 5px;
  margin-left: 10px;
}

.submit-btn {
  margin-bottom: 30px;
  margin-top: 5px;
  background-color: transparent;
  padding: 0.3rem 1rem;
  border-radius: 10rem;
  border: 1px solid #a9a9b3;
  font-size: 0.6rem;
  letter-spacing: .15rem;
  color: #a9a9b3;
  text-transform: uppercase;
  cursor: pointer;
}

.submit-btn:hover {
  color: #292a2d;
  background-color: #a9a9b3;
}

</style>

