<script setup lang="ts">
import { useRoute } from "vue-router";
import { ref, onBeforeMount, computed } from "vue";
import Ws from "@/Websocket";
import { useAuthStore } from "@/stores/auth";
import OneGroup from "../components/Onegroup.vue";
import { useGroupsStore } from "@/stores/groups"

const route = useRoute();
const store = useAuthStore();
const gStore = useGroupsStore();
const groupId = Number(route.params.id);

const group = computed(() => gStore.group);
console.log("the group info", group)
onBeforeMount(() => {
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
});

</script>
<template>
  <div>
    <OneGroup :group="group" />
    </div>
</template>

<style scoped></style>

