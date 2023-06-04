import { defineStore } from "pinia";

export const usePostsStore = defineStore("posts", {
  state: () => ({
    posts: {},
    comments: {},
    post: {},
  }),
  getters: {
 /*    user: (state) => state.userName,
   // userId: (state) => state.userId,
    checkLogin: (state) => state.loggedIn, */
  },
  actions: {
    setPosts(posts: object) {
      this.posts = posts
    },
    setComments(comments: object){
      this.comments = comments
    },
    setPost(post: object) {
      this.post = post
    }
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
