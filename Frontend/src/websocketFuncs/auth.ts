import { useAuthStore } from "../stores/auth";
import router from "../router/index";
import type { follower } from "../interfaces/interfaces"

export const userLogin = (username: string, userId: number, followers: follower[], sessionId: string) => {
  useAuthStore().setUser(username, userId, followers, sessionId);
  router.push("/");
};
export const userSignUp = () => {
  router.push("/login");
};
export const userLogOut = () => {
  useAuthStore().logOut();
  console.log("login v2lja ja pushin uut router")
  router.push("/");
};

export const refreshFollowers = (followers: any) => {
  useAuthStore().refreshFollowers(followers)
}
