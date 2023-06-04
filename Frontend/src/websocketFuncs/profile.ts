import { useProfileStore } from "../stores/profile";

export const userProfile = (uFirst: string, uLast: string, uAge: string, uGender: string, uEmail: string, uName: string, uTime: string, uPic: string, uText: string, uNick: string, uPriv: boolean, fStatus: boolean, uThreads: object) => {
  useProfileStore().setProfile(uFirst, uLast, uAge, uGender, uEmail, uName, uTime, uPic, uText, uNick, uPriv, fStatus, uThreads);
};

