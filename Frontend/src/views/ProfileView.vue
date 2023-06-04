<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import { ref, onBeforeMount, watch } from "vue";
import Ws from "@/Websocket";
import { useAuthStore } from "@/stores/auth";
import { useProfileStore } from "@/stores/profile";
import ProfileInfo from "../components/ProfileInfo.vue";

const router = useRouter();
const route = useRoute();
const uId = ref(route.params.id);
console.log("the uid", uId.value);
const store = useAuthStore();

if (Number(uId) === store.userId) {
  router.replace(`/profile`);
}
//const comments = computed(() => store.comments);
//const post = computed(() => store.post);

//const commentcontent = ref();

onBeforeMount(() => {
  console.log("onbefouremount profile: ", uId.value,typeof uId.value, useAuthStore().userId )
  if (typeof uId.value === undefined || typeof uId.value === null){
    console.log("esimene ver")
    Ws.send({
    Page: "front",
    Data: {
      cur: useAuthStore().userId,
      uid: useAuthStore().userId,      
    },
  });
  } else if (uId.value !== undefined && uId.value !== null){
    console.log("teine ver")

    Ws.send({
    Page: "front",
    Data: {
      cur: useAuthStore().userId,
      uid: Number(uId.value),      
    },
  });
  }

});

watch(
  () => route.params.search,
  (search) => {
    console.log("search: ", search);
    Ws.send({
      Page: "front",
      Data: {
        cur: useAuthStore().userId,
        uid: useAuthStore().userId,
      },
    });
  },
  { deep: true, immediate: true }
);
</script>
<template>
  <div>
    <ProfileInfo />
  </div>
</template>

<style scoped></style>
