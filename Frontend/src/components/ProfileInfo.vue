<script setup lang="ts">
import { useProfileStore } from '@/stores/profile';
import { ref, computed, watch } from 'vue';
import { useRoute } from "vue-router";
import Ws from "@/Websocket";
import { useAuthStore } from "@/stores/auth";
import Post from "../components/Post.vue";
import type { profileinf } from '@/interfaces/interfaces';
import type { Ref } from "vue";

const user = useAuthStore();
const route = useRoute();
const uId = ref(route.params.id);
const store: Ref<profileinf> = computed<profileinf>(() => useProfileStore().profile as profileinf)
const previewImage = ref(false);
const filteredPosts = computed(() => store.value.uThreads);
//const profPriv = (store.uPriv === false) ? "Make profile private": "Make profile public";
const profPriv = computed(() => (store.value.uPriv === false) ? "Make profile private": "Make profile public")
const followed = computed(() => (store.value.fStatus === false) ? "Follow" : "Unfollow")
const buttonType = computed(() =>(store.value.uName === useAuthStore().userName.toString()) ? profPriv : followed)
const buttonFunc = () => {
  const flwPage = (store.value.fStatus === false) ? "flwreq" : "flwrem";
  const flwArr = [flwPage, {sid: user.userId, rid: Number(uId.value)}];
  const privArr = ["setprivacy", {uid: user.userId}];
  console.log("buttontype.value: ", buttonType.value.value)

  const data = (buttonType.value.value === "Follow" || buttonType.value.value === "Unfollow") ? flwArr : privArr;
  console.log("data: ", data);
  Ws.send({
      Page: data[0],
      Data: data[1],
    }
  );
};

//v-bind:src="`images/` + store.uPic"
</script>


<template>
  <div class="profile-info">
    <div class="avatar-column">
      <img
      :src="`../../images/${store.uPic}`"
        alt="Avatar"
        class="avatar"
        @click="previewImage = true"
      />
      <div
        class="image-preview-overlay"
        v-if="previewImage"
        @click="previewImage = false"
      >
        <img
        :src="`../../images/${store.uPic}`"
          alt="Image Preview"
          class="image-preview__image"
        />
      </div>
    </div>
    <div class="details-column">
      <div class="detail">
        <span class="detail-label">Username:</span> {{ store.uName }}
      </div>
      <div class="detail">
        <span class="detail-label">First Name:</span> {{ store.uFirst }}
      </div>
      <div class="detail">
        <span class="detail-label">Last Name:</span> {{ store.uLast }}
      </div>
      <div class="detail">
        <span class="detail-label">Age:</span> {{ store.uAge }}
      </div>
      <div class="detail">
        <span class="detail-label">Gender:</span> {{ store.uGender }}
      </div>
      <div class="detail">
        <span class="detail-label">Email:</span> {{ store.uEmail }}
      </div>
      <div class="detail">
        <span class="detail-label">Bio:</span> {{ store.uText }}
      </div>
      <div class="detail">
        <span class="detail-label">Nickname:</span> {{ store.uNick }}
      </div>
      <div class="magicbuttondiv">
        <button class="magicbutton" @click="buttonFunc()">{{ buttonType }}</button>
      </div>
    </div>
  </div>
  <Post v-for="post in filteredPosts" :key="post" :post="post" :style="{width: 40 + '%'}"/>
</template>

<style>
.profile-info {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  margin-top: 50px;
}

.avatar-column {
  max-width: 250px;
  margin-right: 50px;
  margin-top: 50px;
}

.avatar {
  max-width: 100%;
  height: auto;
  cursor: pointer;
}

.details-column {
  display: flex;
  flex-direction: column;
  margin-top: 50px;
  margin-bottom: 50px;
}

.detail {
  margin-bottom: 10px;
}

.detail-label {
  font-weight: bold;
  margin-right: 5px;
  font-size: 0.8rem;
  letter-spacing: .1rem;
  text-transform: uppercase;
}

.image-preview-overlay {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 999;
}

.image-preview__image {
  max-height: 90%;
  max-width: 90%;
}

.magicbutton {
  font-size: 0.7rem;
  width: 100%;
  background-color: #d9d9d9;
  color: #4a4a4a;
  border: none;
  cursor: pointer;
  width: fit-content;
  background-color: #d9d9d9;
  padding: 8px 12px;
  border-radius: 10px;
  margin: 0 auto;
  font-size: 0.7rem;
  letter-spacing: .1rem;
  text-transform: uppercase;
}


.magicbutton:hover {
  background-color: #4a4a4a;
  color: white;
}
  </style>
