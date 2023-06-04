<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref } from "vue";

const comment = defineProps(["comment"]);
const router = useRouter();
const navigateToProfile = (id: number) => {
  router.replace(`/user/${id}`);
};

const toggleShowPopUp = ref("");

const showPopUp = (pic: string) => {
  if (toggleShowPopUp.value === "") {
    toggleShowPopUp.value = pic;
  } else {
    toggleShowPopUp.value = "";
  }
};
</script>

<template>
  <div class="one-comment">
    <div class="user-content">
      <div
        class="comment-author"
        @click="navigateToProfile(comment.comment.CoUId)"
      >
        {{ comment.comment.CoUName }}
      </div>
      <div class="grid-comment-content">{{ comment.comment.CoBody }}</div>
      <div class="all-images">
        <div class="post-images" v-for="pic in comment.comment.CoPics">
          <img :src="pic" alt="Image 1" id="proov" @click="showPopUp(pic)" />

          <div class="overlay" v-if="toggleShowPopUp === pic">
            <span class="close" @click="showPopUp(pic)">X</span>
            <img :src="pic" class="modal-content" id="img01" />
          </div>
        </div>
      </div>
    </div>
    <div class="date-commented">{{ comment.comment.CoTime }}</div>
  </div>
</template>

<style scoped>
.one-comment {
  margin-bottom: 15px;
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

.user-content {
  background-color: #ececec;
  padding: 10px;
  border-radius: 10px;
  border-radius: 0px 20px 0px 20px;
}

.comment-author {
  font-weight: bold;
  font-size: 0.8rem;
}

.date-commented {
  font-size: 0.7rem;
  text-align: right;
}
</style>
