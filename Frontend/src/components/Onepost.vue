<script setup lang="ts">
import { usePostsStore } from "@/stores/posts";
import { useRouter } from "vue-router";
import { ref } from "vue";

const store = usePostsStore();
const post = defineProps(["post"]);
const router = useRouter();
const navigateToProfile = (id: number) => {
  router.replace(`/user/${id}`);
};
const toggleShowPopUp = ref("");

const showPopUp = (pic: string) => {
  if (toggleShowPopUp.value === ""){
    toggleShowPopUp.value = pic;
  } else {
    toggleShowPopUp.value = "";
  }
  console.log(toggleShowPopUp.value);
};
</script>

<template>
  <div class="grid-info">
    <div class="header">
      <img
        class="author-profile-picture"
        :src="`../../images/${post.post.TProf}`"
        alt="Author Profile Picture"
      />
      <div class="post-author" @click="navigateToProfile(post.post.TUId)">
        {{ post.post.TUName }}
      </div>
      <div class="post-timing">{{ post.post.TTime }}</div>
    </div>
    <div class="content">
      <p>{{ post.post.TBody }}</p>
      <div class="all-images">
        <div v-for="pic in post.post.TPics" class="post-images">
          <img :src="pic" alt="Image 1" id="proov" @click="showPopUp(pic)" />

          <div class="overlay" v-if="toggleShowPopUp === pic">
            <span class="close" @click="showPopUp(pic)">X</span>
            <img :src="pic" class="modal-content" id="img01" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.grid-info {
  min-width: 250px;
  width: 100%;
  min-height: 120px;
  margin: 0 auto;
  margin-top: 40px;
  margin-bottom: 10px;
  padding: 10px;
  border-radius: 10px;
  display: flex;
  box-sizing: border-box;
}

.all-images {
  display: flex;
  flex-wrap: wrap;
  margin-bottom: 10px;
}

.post-images #proov {
  max-width: 100px;
  max-height: 100px;
  object-fit: cover;
  margin-right: 10px;
  margin-bottom: 10px;
}

#proov:hover {
  opacity: 0.7;
  cursor: pointer;
}

.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal {
  position: relative;
  z-index: 1;
  padding-top: 100px;
  left: 0;
  top: 0;
  width: 50%;
  height: 50%;
  overflow: auto;
  background-color: rgb(0, 0, 0);
  background-color: rgba(0, 0, 0, 0.9);
}

.modal-content {
  margin: auto;
  display: block;
  width: 80%;
  max-width: 700px;
}

.modal-content,
#caption {
  animation-name: zoom;
  animation-duration: 0.6s;
}

@keyframes zoom {
  from {
    transform: scale(0);
  }
  to {
    transform: scale(1);
  }
}

.close {
  position: absolute;
  top: 15px;
  right: 35px;
  color: white;
  font-size: 40px;
  font-weight: bold;
  transition: 0.3s;
}

.close:hover,
.close:focus {
  color: red;
  text-decoration: none;
  cursor: pointer;
}

@media only screen and (max-width: 700px) {
  .modal-content {
    width: 100%;
  }
}

.header {
  display: inherit;
  align-items: center;
  margin-right: 25px;
  flex-direction: column;
  flex-grow: 1;
}

.author-profile-picture {
  width: 50px;
  height: 50px;
  object-fit: cover;
  border-radius: 50%;
  border: 1px solid black;
  font-size: 10px;
  text-align: center;
}

.post-author {
  font-weight: bold;
  font-size: 1rem;
  margin-bottom: 7px;
  text-align: center;
}

.post-timing {
  font-size: 0.7rem;
  text-align: center;
}

.content {
  position: relative;
  display: flex;
  flex-direction: column;
  flex-grow: 8;
  height: auto;
  min-height: 100px;
}
</style>
