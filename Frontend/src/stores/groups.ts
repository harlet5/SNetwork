import { defineStore } from "pinia";
import type { group, gchatmsg } from "../interfaces/interfaces";

export const useGroupsStore = defineStore("groups", {
  state: () => ({
    groups: [] as group[],
    uGroups: [] as group[],
    group: {},
    gPosts: {},
    outsiders: {},
    events: {},
    owner: {},
    uStatus: false,
    gChat: [] as gchatmsg[],
    gChatVal: "",
    count: 0,
  }),
  getters: {
    /*    user: (state) => state.userName,
   // userId: (state) => state.userId,
    checkLogin: (state) => state.loggedIn, */
  },
  actions: {
    setGroups(groups: group[]) {
      this.groups = groups;
    },
    setUGroups(uGroups: group[]) {
      this.uGroups = uGroups;
    },
    setGroup(group: object) {
      this.group = group;
    },
    setGPosts(gPosts: object) {
      this.gPosts = gPosts;
    },
    setOutsiders(outsiders: object) {
      this.outsiders = outsiders;
    },
    setEvents(events: object) {
      this.events = events;
    },
    setOwner(owner: object) {
      this.owner = owner;
    },
    setUStatus(uStatus: boolean) {
      this.uStatus = uStatus;
    },
    setGChat(chats: gchatmsg[]) {
      this.gChat = chats;
    },
    setGChatVal(val: string) {
      this.gChatVal = val;
    },
    appendMessages(messages: gchatmsg[]) {
      this.gChat = [...messages, ...this.gChat];
    },
    prependMessage(message: gchatmsg) {
      if (this.gChat === null) {
        this.gChat = new Array();
        this.gChat.push(message);
      } else {
        this.gChat = [...this.gChat, message];
      }
    },
    setCount(count: number) {
      this.count = count;
    },
  },
});

/*

import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {
  const count = ref(0)
  const doubleCount = computed(() => count.value * 2)
  function increment() {
    count.value++
  }

  return { count, doubleCount, increment }
})

*/
