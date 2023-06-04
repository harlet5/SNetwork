<script setup lang="ts">
import { ref, watch, computed } from "vue";
import type { Ref } from "vue";
import { useAuthStore } from "@/stores/auth";
import { timestamp } from "@vueuse/shared";
import Ws from "@/Websocket";
import type { follower } from "../interfaces/interfaces"
const categorys: any = ref([]);

const images: any = ref([]);
const privacy = ref("Public");
const selectedFollowers: any = ref([]);
const content = ref("");
const contentlen = ref(0);
const regex = /#\w+/g;
const images2: any = ref([]);
const toggleShowPost = ref(false)
const store = useAuthStore()
const errorMessage = ref("")
const emit = defineEmits(["showCreatePost"])
const emitShowPost = () => {
    emit("showCreatePost")
}


const followers: Ref<follower[]> = computed<follower[]>(() => store.followers as follower[])
const showPost = () => {
  toggleShowPost.value = !toggleShowPost.value
}

watch(content, () => {
  categorys.value = content.value.match(regex);
  if (content.value.length > 256) {
    content.value = content.value.slice(0, -1)
  }
  contentlen.value = content.value.length
});

const setPrivacy = (input: string) => {
  privacy.value = input;
  toggleDropDown();
};

const addOrRemoveSelectedFollowerList = (follower: follower) => {
  if (selectedFollowers.value.includes(follower)) {
    selectedFollowers.value.splice(
      selectedFollowers.value.indexOf(follower),
      1
    );
  } else {
    selectedFollowers.value.push(follower);
  }
};

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
const showDropDown: Ref<boolean> = ref(false);
const toggleDropDown = () => {
  console.log(showDropDown.value);
  showDropDown.value = !showDropDown.value;
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
  if (contentlen.value < 8 && images.value.length == 0) {
    return errorMessage.value = "New post needs to be between 8 and 256 characters or include atleast 1 picture"

  } else if (contentlen.value > 256) {
    return errorMessage.value = "New post needs to be between 8 and 256 characters or include atleast 1 picture"

  }
  const d = new Date();
  const time = d.toLocaleDateString().replace(/\//g, ".");
  const newPost = <Post>{
    uid: store.userId,
    content: content.value == null ? "" : content.value,
    privacy: privacy.value,
    whosees: selectedFollowers.value.length == 0 ? [] as string[] : selectedFollowers.value,
    categories: categorys.value === null ? [] as string[] : categorys.value,
    images: images2.value,
    time: time,
    gid: -1,
  }
  Ws.send({
      Page: "newpost",
      Data: 
        newPost
    }
  ); 
  content.value = ''
  privacy.value = "Public"
  images.value = []
  images2.value = []
  selectedFollowers.value = []
  errorMessage.value = ""
  emitShowPost()
};
</script>

<template>
<!--   <button class="add-post-btn" @click="emitShowPost()">Add post</button>
 -->  <div class="overlay">
    <div class="modal">
      <div class="flexcreatepost">
        <div class="createpost">
          <div class="categorys">
            <p class='categories-title'>Categories</p>
            <div class="category" v-for="category in categorys">
              {{ category }}
            </div>
          </div>
          <div class="text-and-count">
          <textarea
            name="content"
            id="content"
            cols="30"
            rows="10"
            class="content"
            placeholder="Type in here..."
            v-model="content"
          ></textarea>
            <div class="count">
              {{ contentlen }} / 256
              <p v-if="errorMessage"> {{ errorMessage }}</p>
            </div>
          </div>
          <div class="dropdown">
            <button @click="toggleDropDown()" class="dropbtn">
              {{ privacy }}
            </button>
            <div
              class="dropdown-content"
              :style="{ display: showDropDown ? 'block' : 'none' }"
            >
              <div @click="setPrivacy('Public')">Public</div>
              <div @click="setPrivacy('Private')">Private</div>
              <div @click="setPrivacy('Almost private')">Almost Private</div>
            </div>
            <div class="followers" v-if="privacy === 'Almost private'">
              <div
                v-if="followers !== null"
                v-for="follower in followers"
                @click="addOrRemoveSelectedFollowerList(follower)"
                :style="{
                  backgroundColor: selectedFollowers.includes(follower)
                    ? '#04aa6d'
                    : 'white',
                  color: selectedFollowers.includes(follower) ?
                  'white' :
                  'black',
                }"
              >
                {{ follower.UName }}
              </div>
              <div class="nofollowers" v-if="followers === null">You have no followers</div>
            </div>
          </div>
          
        </div>
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
</template>

<style scoped>

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

.overlay {
  width: 100%;
  height: calc(100vh - 125px);
  background-color: rgba(0, 0, 0, 0.77);
  z-index: 10;
  padding: 20px;
  box-shadow: black 0px 30px 60px -12px inset, black 0px 18px 36px -18px inset;
  box-sizing: border-box;
}
.modal {
  width: 55%;
  background-color: white;
  min-width: 570px;
  margin: 0 auto;
  border-radius: 10px;
  padding: 30px;
  display: flex;
  flex-direction: column;
  box-shadow: rgb(204, 219, 232) 3px 3px 6px 0px inset, rgba(255, 255, 255, 0.5) -3px -3px 6px 1px inset;
}

.flexcreatepost {
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 100%;
  width: 100%;
}

.createpost {
  display: flex;
  margin-bottom: 10px;
}

.categorys {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  min-width: 110px;
  border-right: 1px solid gray;
  height: 300px;
  overflow-y: scroll;
}

.categories-title{
  margin: 0 auto;
  margin-top: 15px;
  color: black;
  text-transform: uppercase;
  letter-spacing: .1rem;
  font-size: 0.9rem;
}

.text-and-count {
  flex-grow: 3;
  display: flex;
  flex-direction: column;
}

.content {
  resize: none;
  height: 265px;
  outline: none;
  border: none;
  letter-spacing: .15rem;
  padding: 15px 15px 0px 15px;
  line-height: 20px
}

.count {
  padding-left: 15px;
  color: #8e8e8e;
  font-size: 0.8rem;
}

.dropdown {
  position: relative;
  display: inline-block;
  flex-grow: 0.5;
  display: flex;
  flex-direction: column;
  height: 300px;
}

.dropbtn {
  background-color: #04aa6d;
  color: white;
  padding: 1rem;
  margin-bottom: 5px;
  font-size: 0.6rem;
  font-weight: bold;
  border: none;
  width: 150px;
}

.dropdown-content {
  position: relative;
  background-color: #f1f1f1;
  min-width: 150px;
  box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
  z-index: 1;
}

.dropdown-content div {
  color: black;
  padding: 10px 16px;
  text-decoration: none;
  display: block;
  cursor: pointer;
  font-size: 0.8rem;
}

.dropdown-content div:hover {
  background-color: #ddd;
}

.dropbtn:hover {
  background-color: #3e8e41;
}

.form {
  display: flex;
}

.custom-file-upload {
  border: 1px solid #ccc;
  display: inline-block;
  padding: 6px 12px;
  cursor: pointer;
  height: 20px;
}

.image-preview {
  width: 150px;
  height: 100px;
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

.followers {
  display: flex;
  flex-direction: column;
  margin-left: 5px;
  margin-top: 10px;
  height: 100%;
  overflow-y: auto;  
}

.followers div {
  cursor: pointer;
  box-sizing: border-box;
  padding: 5px;
}

.finishpost {
  width: 200px;
}

.modal-button {
  font-size: 0.7rem;
  width: 100%;
  background-color: #d9d9d9;
  color: #4a4a4a;
  border: none;
  cursor: pointer;
  margin-top: 30px;
  margin-bottom: 10px;
}

.modal-button:hover {
  background-color: #4a4a4a;
  color: white;
}

</style>
