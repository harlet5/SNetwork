<script setup lang="ts">
import Ws from "@/Websocket";
import { useAuthStore } from "@/stores/auth";
import { useChatsStore } from "@/stores/chats";
const emit = defineEmits(["messagesActivated"]);
const emitSelectedOption = (otherId: number) => {
  emit("messagesActivated", otherId);
};
const store = useAuthStore();
const chatStore = useChatsStore();

const {chat, selected } = defineProps(["chat", "selected"]);
const getMessages = (id: number) => {
chatStore.count = 0
chatStore.setLastAct("onechat")
  Ws.send({
      Page: "onechat",
      data: {
        uid: store.userId,
        ouid: id.toString(),
        count: chatStore.count,
      },
    }
  );
};
</script>
<template>
  <body>
    <div
      @click="
        getMessages(chat.OId);
        emitSelectedOption(chat.OId);
      "
      :class="{ chat: true, online: chat.OActive, active:  chat.OId === selected }"
      :id="`chat-${chat.OId}`"
    >
      <div>
        <img :src="`../../images/${chat.OProf}`" alt="hi" />
      </div>
      <div class="chat-info">
        <div class="chat-info-upper">
          <div :class="{isOnline: chat.OActive}"></div>
          <p class="chat-username">{{ chat.OName }}</p>
        </div>
        <div class="chat-last-message"></div>
        <div
          class="typing-indicator"
          :id="`chat-${chat.OId}-typing-indicator`"
        >
          <span></span><span></span><span></span>
        </div>
      </div>
      <div
        class="unread-messages-count active"
        :id="`chat-${chat.OId}-unread-messages-count`"
        v-if="chat.Unread > 0"
      >
        {{ chat.Unread }}
      </div>
    </div>
  </body>
</template>

<style scoped>

body {
  word-wrap: break-word;
  overflow-wrap: break-word;
}

.chats-title {
  font-size: 18px;
  font-weight: 600;
  padding: 10px;
  margin: 0px;
}

.chat {
  display: grid;
  grid-template-columns: 50px 175px 25px;
  grid-gap: 10px;
  padding: 10px;
  background-color: #292a2d;
  color: #c9c9d8;
  cursor: pointer;
}

.chat-info {
  display: flex;
  flex-direction: column;
}

.chat.active {
  background-color: #3b3d42;
  color: #c9c9d8;
}

.chat:hover {
    background-color: #3b3d42;
}

.chat-info-upper {
  display: inherit;
}

.chat-info p {
  word-wrap: normal;
  max-width: 188px;
  overflow: hidden;
  text-overflow: ellipsis;
}

.isOnline {
  display: inline-block;
  height: 10px;
  width: 10px;
  border-radius: 10px;
  background-color: green;
  margin-right: 5px;
  margin-top: 7px;
  background-color: limegreen;
}

.unread-messages-count {
  align-self: center;
  width: 25px;
  height: 25px;
  font-size: 13px;
  text-align: center;
  line-height: 25px;
  border-radius: 100%;
  vertical-align: middle;
  background-color: #c9c9d8;
  color: #292a2d;
}

.chat-unread-messages-count.active {
  background-color: #3b3d42;
  color: #c9c9d8;
}

.chat img {
  border-radius: 50%;
  max-width: 50px;
  margin: 0;
  padding: 0;
}

.chat p {
  padding: 0;
  margin: 0;
}

#chat-messages-placeholder {
  position: relative;
  top: 50%;
  text-align: center;
  font-size: 22px;
}

input,
select,
textarea {
  padding: 5px;
  border: none;
  border-radius: 5px;
  box-sizing: border-box;
  color: #a9a9b3;
  font-size: 16px;
  background-color: #292a2d;
  outline: none;
}

.typing-indicator {
  display: none;
  will-change: transform;
  padding: 5px 0;
  animation: 2s infinite ease-out;
}

.typing-indicator p {
  /* display: inline; */
  margin: 0;
  padding: 0;
}

.typing-indicator::before,
.typing-indicator::after {
  height: 15px;
  width: 15px;
  border-radius: 50%;
}

.typing-indicator span {
  height: 9px;
  width: 9px;
  float: left;
  margin: 0 1px;
  background-color: #c9c9d8;
  border-radius: 50%;
  opacity: 0.4;
}

.chat.active .typing-indicator span {
  background-color: #3b3d42;
}

.typing-indicator span:nth-of-type(1) {
  animation: 1s blink infinite 0.3333s;
}

.typing-indicator span:nth-of-type(2) {
  animation: 1s blink infinite 0.6666s;
}

.typing-indicator span:nth-of-type(3) {
  animation: 1s blink infinite 0.9999s;
}

@keyframes blink {
  50% {
    opacity: 1;
  }
}
</style>
