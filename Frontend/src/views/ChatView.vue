<script setup lang="ts">
import Ws from "@/Websocket";
import { useAuthStore } from "../stores/auth";
import Chats from "../components/Chats.vue";
import {
  onBeforeMount,
  computed,
  ref,
  onBeforeUnmount,
  watch,
  onMounted,
} from "vue";
import { useChatsStore } from "@/stores/chats";
import Messages from "../components/Messages.vue";
import type { Ref } from "vue";
import type { Chat } from "../interfaces/interfaces";
import { debounce } from "ts-debounce";
import Emoij from "../components/Emoij.vue";

const store = useAuthStore();
const chatStore = useChatsStore();
const chats: Ref<Chat[]> = computed<Chat[]>(() => chatStore.chats as Chat[]);
const messages = computed(() => chatStore.messages);
const sentMessage = ref("");
const showSendMessage = ref(false);
const selected = computed(() => chatStore.activeChat);
const scrollHeightVal = ref(0);
let otherId = -1;
const showMessageToTrue = (otherId2: number) => {
  showSendMessage.value = true;
  otherId = otherId2;
  scrollHeightVal.value = 0;
  let elem = document.getElementById("chat-messages");
  if (elem) {
    console.log("scrollin alla3")
    elem.scrollTop = elem.scrollHeight;
  }
};

onMounted(() => {
  document
    .getElementById("chat-messages")
    ?.addEventListener("scroll", handleScroll);
});

const handleScroll = debounce((ev: Event) => {
  const target = ev.target as HTMLElement;
  if (target.scrollTop === 0) {
    chatStore.count++;
    chatStore.setLastAct("onechat");
    Ws.send({
      Page: "onechat",
      data: {
        uid: store.userId,
        ouid: selected.value.toString(),
        count: chatStore.count,
      },
    });
  }
}, 500);

onBeforeMount(() => {
  Ws.send({
    Page: "chatroom",
    data: {
      uid: store.userId,
    },
  });
});
onBeforeUnmount(() => {
  chatStore.activeChat = -1;
  chatStore.messages = [];
  document
    .getElementById("chat-messages")
    ?.removeEventListener("scroll", handleScroll);
});
const sendMessage = () => {
  const d = new Date();
  chatStore.setLastAct("newchat");
  Ws.send({
    Page: "newchat",
    data: {
      uid: store.userId,
      ouid: otherId.toString(),
      time: d.toLocaleDateString().replace(/\//g, "."),
      body: sentMessage.value,
      count: 0,
    },
  });
  sentMessage.value = "";
};

watch(messages, () => {
  const elem = document.getElementById("chat-messages");
  if (elem && chatStore.lastAct === "newchat") {
    elem.scrollTop = elem.scrollHeight;
  }
});

const addEmoij = (emoij: any) => {
  sentMessage.value = sentMessage.value + emoij.i;
  const elem = document.getElementById("message-input");
  elem?.focus();
};
</script>

<template>
  <div id="chats-container">
    <div id="chats">
      <Chats
        v-for="chat in chats"
        :key="chat.OId"
        :chat="chat"
        @click="chatStore.activeChat = chat.OId"
        :selected="selected"
        @messagesActivated="showMessageToTrue"
      />
    </div>
    <div class="flex-container">
      <div id="chat-messages">
        <Messages
          v-for="message in messages"
          :key="message.ChId"
          :message="message"
        />
      </div>

      <form
        v-if="showSendMessage"
        @submit.prevent="sendMessage()"
        id="message-form"
        style="display: block"
      >
        <Emoij @emoij="addEmoij" />
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
    </div>
  </div>
</template>

<style scoped>
#chats-container {
  word-wrap: break-word;
  overflow-wrap: break-word;
  height: calc(100vh - 55px);
  width: 100%;
  box-sizing: border-box;
  display: grid;
  grid-template-columns: 300px auto;
  background-color: #292a2d;
  position: fixed;
}

#chats {
  overflow-y: auto;
  border-right: 2px solid black;
}

#chat-messages-placeholder {
  position: relative;
  top: 50%;
  text-align: center;
  font-size: 22px;
}

.flex-container {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  padding: 10px;
  border-left: 2px solid #3c3c3c;
}

#chat-messages {
  overflow-y: auto;
  height: calc(100vh - 125px);
}

#message-form {
  align-self: flex-end;
  width: 100%;
  margin-top: auto;
}

#message-input {
  width: 100%;
  height: 40px;
  padding: 10px;
  background-color: #3b3d42;
  color: #c9c9d8;
}
input {
  padding: 5px;
  border: none;
  border-radius: 5px;
  box-sizing: border-box;
  color: #a9a9b3;
  font-size: 16px;
  background-color: #292a2d;
  outline: none;
}
</style>
