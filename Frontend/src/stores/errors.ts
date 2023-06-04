import { defineStore } from "pinia";

export const useErrorsStore = defineStore("errors", {
  state: () => ({
    loginError: "",
    registerErrorEmail: "",
    registerErrorUsername: "",
    groupError: "",
  }),
  getters: {
    /*    user: (state) => state.userName,
   // userId: (state) => state.userId,
    checkLogin: (state) => state.loggedIn, */
  },
  actions: {
    setLoginError(msg: string) {
      this.loginError = msg;
    },
    setRegisterErrorEmail(msg: string) {
      this.registerErrorEmail = msg;
    },
    setRegisterErrorUsername(msg: string) {
      this.registerErrorUsername = msg;
    },
    setGroupError(msg: string) {
      this.groupError = msg;
    },
  },
});
