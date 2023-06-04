<script setup lang="ts">
import { ref, watch } from "vue";
import type { Ref } from "vue";
import { useAuthStore } from "@/stores/auth";
import { timestamp } from "@vueuse/shared";
import { useRoute } from "vue-router";
import Ws from "@/Websocket";
import { useGroupsStore } from "@/stores/groups";

const images: any = ref([]);
const privacy = ref("Select Privacy");
const selectedFollowers: any = ref([]);
const content = ref("");
const regex = /#\w+/g;
const images2: any = ref([]);
const toggleShowPost = ref(false)
const store = useAuthStore()
const route = useRoute();
const groupId = Number(route.params.id);

const showPost = () => {
  toggleShowPost.value = !toggleShowPost.value
}

const onFileSelected = (event: any) => {
  if (event.target.files[0].size > 1048576) {
    return;
  }
  const reader = new FileReader();
  reader.addEventListener("load", (e) => {
    images2.value.push(reader.result);
  });
  reader.readAsDataURL(event.target.files[0]);
  images.value.push(URL.createObjectURL(event.target.files[0]));
};

const removeImageFromList = (image: string) => {
  let id = -1;
  images.value = images.value.filter((i: string, index: number) => {
    if (i === image) {
      id = index;
      return false;
    }
    return true;
  });
  images2.value.splice(id, 1);
};

interface Post {
  uid: number;
  content: string;
  privacy: string;
  whosees: Array<string>;
  categories: Array<string>;
  images: Array<string>;
  time: string;
  gid: number;
}

const createPost = () => {
  const d = new Date();
  const time = d.toLocaleDateString().replace(/\//g, ".");
  const newPost = <Post>{
    uid: store.userId,
    content: content.value,
    privacy: "",
    whosees: [],
    categories: [],
    images: images2.value,
    time: time,
    gid: groupId,
  }
  console.log(newPost);
  Ws.send({
      Page: "newpost",
      Data: 
        newPost
    }
  ); 
  Ws.send({
      Page: "onegroup",
      Data: {
        uid: store.userId,
        gid: groupId,
        count: 0,
      },
    }
  );
  useGroupsStore().gChatVal = "onegroup"
  content.value = "";
  images2.value = [];
  showPost();
};
</script>

<template>
  <button class="post-btn" @click="showPost()">Add post</button>
  <div class="overlay" :style="{display: toggleShowPost ? 'block' : 'none'}">
    <div class="modal">
      <div class="flexcreatepost">
        <div class="createpost">
          <textarea
            name="content"
            id="content"
            cols="30"
            rows="10"
            class="content"
            placeholder="Type in here..."
            v-model="content"
          ></textarea>
        </div>
        <div class="xtra-stuff">
        <div class="form">
          <label for="file-upload" class="custom-file-upload"
            >Upload images</label
          >
          <input
            multiple
            id="file-upload"
            type="file"
            accept="image/*"
            @change="onFileSelected"
          />
          <div
            class="image-preview"
            id="image-preview"
            v-for="image in images"
            @click="removeImageFromList(image)"
          >
            <img
              :src="image"
              alt="Image Preview"
              class="image-preview__image"
              style="height: 100%; width: 100%; object-fit: contain"
            />
          </div>
        </div>
        <button class="modal-button" @click="createPost()">Create post</button>
        </div>
      </div>
    </div>
  </div>
  
</template>

<style scoped>
  .post-btn {
    display: block;
    margin: 0 auto;
    margin-top: 50px;
    margin-bottom: 15px;
    background-color: #d9d9d9;
    padding: 0.5rem 2rem;
    border-radius: 10rem;
    font-size: 0.8rem;
    letter-spacing: .15rem;
    color: #4a4a4a;
    text-transform: uppercase;
    position: relative;
    overflow: hidden;
    cursor: pointer;
}

.overlay {
  width: 70%;
  margin: 0 auto;
  box-sizing: border-box;
  margin-bottom: 30px;
}

.modal {
  background-color: white;
  height: 150px;
  min-width: auto;
  margin: 0 auto;
  border-radius: 10px;
  border: 1px solid #f1f1f1;
  padding: 3px;
  display: flex;
  flex-direction: column;
  box-shadow: rgb(204, 219, 232) 3px 3px 6px 0px inset, rgba(255, 255, 255, 0.5) -3px -3px 6px 1px inset;
}

.flexcreatepost {
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 100%;
}

.createpost {
  display: flex;
}

.content {
  resize: none;
  height: 110px;
  outline: none;
  border: none;
  line-height: 20px;
  box-sizing: border-box;
  padding: 10px;
  width: 100%;
  background-color: transparent;
}

.xtra-stuff {
  display: flex;
  justify-content: space-between;
}

.form {
  display: flex;
}

.custom-file-upload {
  border: 1px solid #ccc;
  display: inline-block;
  padding: 6px 10px;
  margin-left: 10px;
  cursor: pointer;
  height: 10px;
  font-size: 0.6rem;
}

.image-preview {
  width: 30px;
  height: 30px;
  border: 2px solid #dddddd;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  color: #cccccc;
}


input[type="file"] {
  display: none;
}

.modal-button {
  font-size: 0.6rem;
  background-color: #d9d9d9;
  color: #4a4a4a;
  border: none;
  border-radius: 15px;
  cursor: pointer;
  padding: 6px 12px;
  height: 30px
}

.modal-button:hover {
  background-color: #4a4a4a;
  color: white;
}

</style>

