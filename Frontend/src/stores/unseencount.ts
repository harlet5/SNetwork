import { defineStore } from "pinia";

export const useUnseenCountStore = defineStore("unseen", {
  state: () => ({
    unseenchatcount: 0,
  }),
  getters: {
    /*    user: (state) => state.userName,
     // userId: (state) => state.userId,
      checkLogin: (state) => state.loggedIn, */
  },
  actions: {
    setUnseenChatCount(count: number) {
      this.unseenchatcount = count;
    },
  },
});
