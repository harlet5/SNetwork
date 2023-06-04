<script setup lang="ts">
import Ws from "@/Websocket";
import { ref, onBeforeMount, computed } from "vue";
import Post from "../components/Post.vue";
import { usePostsStore } from "../stores/posts";
import CreatePost from "../components/CreatePost.vue";
import { useAuthStore } from "@/stores/auth";
const store = usePostsStore();
const posts: any = computed(() => store.posts);
const search = ref("");
const filtersApplied: any = ref([]);
const showCreatePost = ref(false);
const filteredPosts = computed(() => {
  if (filtersApplied.value.length == 0) {
    return posts.value;
  } else {
    return posts.value.filter((post: any) => {
      let output = false;
      for (let i = 0; i < filtersApplied.value.length; i++) {
        if (post.TCats !== null) {
          output = post.TCats.some(
            (cat: any) =>
              cat.CName.toLowerCase() === filtersApplied.value[i].toLowerCase()
          );
        }

        if (!output) {
          return output;
        }
      }
      return output;
    });
  }
});
const toggleCreatePost = () => {
  showCreatePost.value = !showCreatePost.value;
  console.log("toggle tootab: ", showCreatePost.value)
};
onBeforeMount(() => {
  Ws.send({
    Page: "threads",
    Data: { uname: useAuthStore().userName },
  });
});

const addFilter = () => {
  filtersApplied.value.push(search.value.toLowerCase());
  search.value = "";
};

const removeFilter = (filter: string) => {
  filtersApplied.value = filtersApplied.value.filter(
    (fil: string) => fil !== filter
  );
};
</script>
<template>
  <header>
    <p class="filter-title">Filter:</p>
    <form @submit.prevent="addFilter()">
      <input v-model.trim="search" type="text" placeholder="Insert a category" />
    </form>
    <div
      class="filters"
      v-for="filter in filtersApplied"
      @click="removeFilter(filter)"
    >
      <p class="x">
        x <span class="no-x"> {{ filter }} </span>
      </p>
    </div>
  </header>

  <div class="button-div">
    <button class="add-post-btn" @click="toggleCreatePost()">Add post</button>
  </div>
  <CreatePost class="createpost" @showCreatePost="toggleCreatePost" v-if="showCreatePost"/>

  <div class="grid-posts">
    <div class="grid-post">
      <Post
        v-for="post in filteredPosts"
        :key="post"
        :post="post"
        :style="{ width: 60 + '%' }"
      />
    </div>
  </div>
</template>

<style scoped>
.createpost {
width: 100%;
height: 100%;
}

.button-div {
  width: 100%;
  display: flex;
}

button {
    display: inline-block;
    width: fit-content;
    background-color: #d9d9d9;
    padding: 1rem 5rem;
    border-radius: 10rem;
    margin: 0 auto;
    margin-bottom: 30px;
    font-size: 1rem;
    letter-spacing: .15rem;
    color: #4a4a4a;
    text-transform: uppercase;
    position: relative;
    overflow: hidden;
    cursor: pointer;
}
header {
  margin: 10px;
  display: flex;
  align-items: center;
}

.filter-title {
  font-weight: bold;
  margin-right: 10px;
}

header input {
  border: none;
  background-color: rgba(128, 128, 128, 0.1);
  padding: 10px;
  border-radius: 5px;
  outline: none;
}

.filters {
  cursor: pointer;
  margin-left: 15px;
}

.grid-posts {
  display: flex;
  flex-direction: column;
}

.x {
  color: red;
  background-color: rgba(128, 128, 128, 0.1);
  padding: 4px;
  border-radius: 5px;
  font-weight: bold;
}

.x:hover {
  background-color: #a9a9b3;
}

.no-x {
  color: black;
  font-weight: normal;
}
</style>
