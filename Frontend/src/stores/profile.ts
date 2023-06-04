import { defineStore } from "pinia";
import { useLocalStorage } from "@vueuse/core";
import type { Ref } from "vue";
import type { profileinf } from "@/interfaces/interfaces";

const ls = <T>(id: string, defaultValue: T): Ref<T> =>
  useLocalStorage(id, defaultValue);

export const useProfileStore = defineStore("profile", {
  state: () => ({
    profile: {
      uFirst: "",
      uLast: "",
      uAge: "",
      uGender: "",
      uEmail: "",
      uName: "",
      uTime: "",
      uPic: "",
      uText: "",
      uNick: "",
      uPriv: false,
      fStatus: false,
      uThreads: {},
    } as profileinf,
  }),
  getters: {
    /*    user: (state) => state.userName,
   // userId: (state) => state.userId,
    checkLogin: (state) => state.loggedIn, */
  },
  actions: {
    setProfile(
      uFirst: string,
      uLast: string,
      uAge: string,
      uGender: string,
      uEmail: string,
      uName: string,
      uTime: string,
      uPic: string,
      uText: string,
      uNick: string,
      uPriv: boolean,
      fStatus: boolean,
      uThreads: object
    ) {
      (this.profile.uFirst = uFirst),
        (this.profile.uLast = uLast),
        (this.profile.uAge = uAge),
        (this.profile.uGender = uGender),
        (this.profile.uEmail = uEmail),
        (this.profile.uName = uName),
        (this.profile.uTime = uTime),
        (this.profile.uPic = uPic),
        (this.profile.uText = uText),
        (this.profile.uNick = uNick),
        (this.profile.uPriv = uPriv),
        (this.profile.fStatus = fStatus),
        (this.profile.uThreads = uThreads);
    },
  },
});
