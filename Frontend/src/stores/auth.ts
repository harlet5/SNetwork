import { defineStore } from "pinia";
import { useLocalStorage } from "@vueuse/core";
import type { Ref } from "vue";
import type { follower } from "../interfaces/interfaces";
const ls = <T>(id: string, defaultValue: T): Ref<T> =>
  useLocalStorage(id, defaultValue);

export const useAuthStore = defineStore("auth", {
  state: () => ({
    userName: ls("userName", ""),
    loggedIn: ls("loggedIn", false),
    userId: ls("userId", -1),
    followers: ls("followers", [] as follower[]),
    sessionId: ls("sessionId", ""),
  }),
  getters: {
    /*    user: (state) => state.userName,
   // userId: (state) => state.userId,
    checkLogin: (state) => state.loggedIn, */
  },
  actions: {
    setUser(
      userName: string,
      userId2: number,
      followers: follower[],
      sessionId: string
    ) {
      (this.userName = userName),
        (this.loggedIn = true),
        (this.userId = userId2),
        (this.followers = followers),
        (this.sessionId = sessionId);
    },
    refreshFollowers(followers: follower[]) {
      this.followers = followers;
    },
    logOut() {
      (this.userName = ""),
        (this.loggedIn = false),
        (this.userId = -1),
        (this.followers = []),
        (this.sessionId = "");
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
