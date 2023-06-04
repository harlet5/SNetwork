<script setup lang="ts">
import { useRoute } from "vue-router";
import { ref, onBeforeMount, computed, watch } from "vue";
import Comment from "../components/Comment.vue";
import Ws from "@/Websocket";
import { usePostsStore } from "@/stores/posts";
import { useAuthStore } from "@/stores/auth";
import Onepost from "../components/Onepost.vue";
const route = useRoute();
const postId = route.params.id as string;
const store = usePostsStore();
const comments = computed(() => store.comments);
const post = computed(() => store.post);
watch(comments, () => {
  const elem = document.getElementById("comment-section");
  if (elem) {
    elem.scrollTop = elem.scrollHeight;
  }
});

const commentcontent = ref("");
onBeforeMount(() => {
  Ws.send({
    Page: "onethread",
    Data: {
      tid: postId,
    },
  });
});
const createComment = () => {
  if (commentcontent.value.length === 0 && images2.value.length === 0) {
    return;
  }
  const d = new Date();
  Ws.send({
    Page: "newcomm",
    Data: {
      tid: postId,
      uid: useAuthStore().userId as number,
      body: commentcontent.value == null ? "" : commentcontent.value,
      time: d.toLocaleDateString().replace(/\//g, "."),
      images: images2.value,
    },
  });
  Ws.send({
    Page: "onethread",
    Data: {
      tid: postId,
    },
  });
  commentcontent.value = "";
  images.value = [];
  images2.value = [];
};
const images: any = ref([]);
const images2: any = ref([]);
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

</script>
<template>
  <link
    href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css"
    rel="stylesheet"
  />
  <link
    href="https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"
    rel="stylesheet"
  />

  <div class="one-post">
    <Onepost :post="post" />
    <div class="comment-section-and-form">
      <p class="comment-section-title">Comments</p>
      <div class="comment-section" id="comment-section">
        <Comment
          v-for="comment in comments"
          :key="comment"
          :comment="comment"
        />
      </div>
      <div class="grid-form">
        <form @submit.prevent="createComment()">
          <div class="form-div">
            <input
              class="comment-input"
              v-model="commentcontent"
              placeholder="Write a comment here..."
            />
            <div class="comment-icons">
              <div class="form">
                <label class="image-label">
                  <i class="fa fa-image"></i>
                  <input
                    multiple
                    type="file"
                    accept="image/*"
                    @change="onFileSelected"
                    style="display: none"
                    name="image"
                  />
                </label>
              </div>
              <button class="comment-button gg-log-in"></button>
            </div>
          </div>

          <div class="images">
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
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import url("https://unpkg.com/css.gg@2.0.0/icons/css/log-in.css");

.one-post {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 55%;
  height: calc(100vh - 55px);
  margin: 0 auto;
}

.comment-section-and-form {
  width: 100%;
}

.comment-section {
  box-sizing: border-box;
  width: 100%;
  padding: 10px 15px 10px 15px;
  height: auto;
  max-height: 45vh;
  overflow-y: auto;
}

.comment-section-title {
  font-weight: bold;
  padding-left: 14px;
}

.grid-form {
  width: 100%;
  padding-top: 15px;
  box-shadow: 0 -5px 5px -5px #333;
  margin-top: 3px;
}

.form-div {
  width: 100%;
  border-radius: 20px;
  background-color: #ececec;
  padding: 5px 15px 5px 15px;
  display: flex;
}

#file-upload {
  display: none;
}

.comment-icons {
  display: flex;
  gap: 20px;
  align-items: center;
  justify-content: center;
}

.comment-input {
  background-color: transparent;
  border: none;
  outline: none;
  padding-right: 10px;
  font-size: 0.8rem;
  width: 100%;
}

.image-label {
  margin: 0;
}

.images {
  display: flex;
  gap: 3px;
}

.comment-button {
  margin-left: auto;
  background: transparent;
  border: none;
  outline: none;
}

.gg-log-in {
  height: 11px;
}

.gg-log-in::before {
  width: 15px;
  height: 15px;
  background-color: transparent;
}

.comment-button:hover,
.fa-image:hover {
  color: blue;
  cursor: pointer;
}

.form {
  display: flex;
}

.image-preview {
  width: 150px;
  height: 100px;
  border: 2px solid #dddddd;
  margin-top: 15px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  color: #cccccc;
}
</style>
