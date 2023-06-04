import { defineStore } from "pinia";
import type { Chat, Message } from "@/interfaces/interfaces";

export const useChatsStore = defineStore("chats", {
  state: () => ({
    chats: [] as Chat[],
    messages: [] as Message[],
    activeChat: -1,
    count: 0,
    lastAct: "",
  }),
  actions: {
    setActiveChat(userid: number) {
      this.activeChat = userid;
    },
    setChats(chats: Chat[]) {
      this.chats = chats;
    },
    setMessages(messages: Message[]) {
      this.messages = messages;
    },
    appendMessages(messages: Message[]) {
      this.messages = [...messages, ...this.messages];
    },
    prependMessage(message: Message) {
      if (this.messages === null) {
        this.messages = new Array();
        this.messages.push(message);
      } else {
        this.messages = [...this.messages, message];
      }
    },
    changeUserActive(userid: number, toggle: boolean) {
      if (this.chats.length === 0) {
        return;
      }
      console.log("chats length: ",this.chats.length)
      this.chats[this.chats.findIndex((chat) => chat.OId === userid)].OActive =
        toggle;
    },
    setCount(count: number) {
      this.count = count;
    },
    setLastAct(act: string) {
      this.lastAct = act;
    },
  },
});
