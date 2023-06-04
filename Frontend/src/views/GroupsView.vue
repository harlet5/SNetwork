<script setup lang="ts">
import { useRoute } from "vue-router";
import { ref, onBeforeMount, computed } from "vue";
import Ws from "@/Websocket";
import { useAuthStore } from "@/stores/auth";
import { useGroupsStore } from "@/stores/groups"
import AllGroups from "../components/AllGroups.vue";
import InGroups from "../components/InGroups.vue";
import CreateGroup from "../components/CreateGroup.vue";
import type {group} from "../interfaces/interfaces"
import type { Ref } from "vue"

const route = useRoute();
const groupsStore = useGroupsStore(); 
const groups: Ref<group[]> = computed<group[]>(() => groupsStore.groups.filter((group) => !groupsStore.uGroups.some((uGroup)=> group.Id === uGroup.Id)));
const uGroups = computed(() => groupsStore.uGroups);
const store = useAuthStore();
//const comments = computed(() => store.comments);
//const post = computed(() => store.post);

//const commentcontent = ref();

onBeforeMount(() => {
  Ws.send({
      Page: "groups",
      Data: {
        uid: store.userId,
      },
    }
  );
});
</script>
<template>
  <div>
    <CreateGroup />

    <h3 class="heading">My groups:</h3>
    <InGroups v-for="uGroup in uGroups" :key="uGroup.Id" :uGroup="uGroup" v-if="Object.keys(uGroups).length > 0" />
    <p class="no-groups" v-else>You haven't joined any groups yet!</p>

    <h3 class="heading">All groups:</h3>
    <AllGroups v-for="group in groups" :key="group.Id" :group="group" v-if="Object.keys(groups).length > 0" />
    <p class="no-groups" v-else>There are no groups yet!</p>

  </div>
</template>

<style scoped>
.heading {
  text-align: center;
}

.no-groups {
  text-align: center;
}
</style>
