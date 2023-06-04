import { defineStore } from "pinia";

export const useNotifStore = defineStore("notifications", {
  state: () => ({
    notifications: {},
    count: 0,
    }),
  getters: {
 /*    user: (state) => state.userName,
   // userId: (state) => state.userId,
    checkLogin: (state) => state.loggedIn, */
  },
  actions: {
    setNotif(notifications: object) {
      this.notifications = notifications
    },
    addNotifCount(){
      this.count++
    },
    clearNotifCount(){
      this.count = 0
    },
  },
});
