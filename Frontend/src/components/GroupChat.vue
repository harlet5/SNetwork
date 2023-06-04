<script setup lang="ts">
import { useGroupsStore } from "@/stores/groups";
import { useAuthStore } from "@/stores/auth";
import Ws from "@/Websocket";
import { computed, onBeforeUnmount, ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import type { Ref } from "vue";
import { debounce } from "ts-debounce";
import Emoij from "../components/Emoij.vue";


const store = useGroupsStore();
const chats: Ref<Chat[]> = computed<Chat[]>(() => store.gChat as Chat[]);
const sentMessage = ref("");
const route = useRoute();
const GId = Number(route.params.id);

interface Chat {
  ChId: number,
  ChUId: number,
  ChGId: number,
  ChUName: string,
  ChBody: string,
  ChTime: string,
}
const sendMessage = () => {
  const d = new Date();
  Ws.send({
      Page: "newgroupchat",
      data: {
        uid: useAuthStore().userId,
        ouid: GId,
        time: d.toLocaleDateString().replace(/\//g, "."),
        body: sentMessage.value,
        count: 0,
      },
    }
  );
  useGroupsStore().gChatVal = "newgroupchat"
  sentMessage.value = "";
};
onBeforeUnmount(() => {
    store.setGChat([])
    document
    .getElementById("chat-messages")
    ?.removeEventListener("scroll", handleScroll);
    store.setCount(0)
})
const computedClasses = (uid: number) => {
  if (uid == useAuthStore().userId) {
    return 'message sended-message'
  } else {
    return 'message received-message'
  }
}

onMounted(() => {
  document
    .getElementById("chat-messages")
    ?.addEventListener("scroll", handleScroll);
});

const handleScroll = debounce((ev: Event) => {
  const target = ev.target as HTMLElement;
  if (target.scrollTop === 0) {
    store.count++
    store.setGChatVal("scroll")
    Ws.send({
      Page: "onegroup",
      Data: {
        uid: useAuthStore().userId,
        gid: GId,
        count: store.count,
      },
    }
  );
  }
}, 500);
const addEmoij = (emoij: any) => {
  sentMessage.value = sentMessage.value + emoij.i;
  const elem = document.getElementById("message-input")
  console.log(elem)
  elem?.focus()
};
</script>

<template>

<div class="all-messages" id="chat-messages">
  <div class="messages" v-for="chat in chats">    
    <div :id="`message-${chat.ChId}`" :class="computedClasses(chat.ChUId)">
      <p class="message-name">{{ chat.ChUName }}</p>
      <p class="message-text">{{chat.ChBody}}</p>
      <div class="message-info" :id="`message-${chat.ChId}-info`">
        <p class="message-date">{{chat.ChTime}}</p>
      </div>
    </div>
  </div>
</div>
<form
      @submit.prevent="sendMessage()"
      id="message-form"
      style="display: block"
    >
      <input
        required
        id="message-input"
        v-model="sentMessage"
        type="text"
        size="64"
        placeholder="Send message"
        autocomplete="off"
        autofocus="false"
      />
</form>
<Emoij class="emoji" @emoij="addEmoij"/>

</template>

<style scoped>

* {
  margin: 0;
  padding: 0;
}

.all-messages {
  padding-left: 10px;
  height: calc(100vh - 170px);
  overflow-y: auto;
}

.message {
  width: fit-content;
  min-width: 180px;
  max-width: 180px;
  margin: 2px 10px 10px 0px;
  padding: 10px;
  border-radius: 15px;
}

.message-name {
  font-size: 0.7rem;
}

.message-text {
  margin: 0;
  margin-bottom: 5px;
  font-size: 0.9em;
  font-weight: bold;
}

.message-info {
  display: grid;
}

.message-info p {
  margin: 0;
  font-size: 0.7rem;
  margin-left: auto;
  margin-right: 0;
}

.received-message {
  background-color: #d9d9d9;
}

.sended-message {
  background-color: #969594;
  margin-left: auto;
  margin-right: 0;
}

form {
  margin-left: 10px;
  position: absolute;
  bottom: 0;
  margin-bottom: 10px;
  margin-right: 10px;
  background-color: #d9d9d9;
}

#message-input {
  width: 93%;
  height: 40px;
  padding: 10px;
  background-color: #d9d9d9;
}

input {
  padding: 5px;
  border: none;
  border-radius: 5px;
  box-sizing: border-box;
  font-size: 16px;
  outline: none;
}

</style>
