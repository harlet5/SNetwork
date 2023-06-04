<script setup lang="ts">
import { useRoute } from "vue-router";
import { ref, onBeforeMount, computed } from "vue";
import Ws from "@/Websocket";
import { useAuthStore } from "@/stores/auth";

const outsider = defineProps(["outsider"]);
const route = useRoute();
const store = useAuthStore();
const groupId = Number(route.params.id);
//console.log("the thingy: ", outsider)
const invite = (uid: number) => {
 console.log("the uid:", uid)
 Ws.send({
      Page: "grpinv",
      Data: { 
        sid: store.userId,
        rid: uid, 
        gid: groupId,
      }
    }
  );
}
</script>
<template>
  <div class="outsider">
      <div>
        <div class='buttondiv'>
          <span class="username">{{ outsider.outsider.UName }}</span>
          <button class='accept' @click="invite(outsider.outsider.UId)">Invite</button>
        </div>
      </div>
  </div>

 </template>
<style scoped>

.notification {
  padding: 10px;
  cursor: pointer;
}

.buttondiv {
  display: flex;
  justify-content: space-between;
  padding: 8px;
  margin: 0px 30px;
  border-bottom: 1px solid #d9d9d9
}

.notification:hover {
  background-color: lightgray;
}

.accept {
  background-color: #d9d9d9;
  border-radius: 3px;
  letter-spacing: .1rem;
  cursor: pointer;
  color: green;
  border: none;
  text-transform: uppercase;
  font-size: 0.6rem;
  letter-spacing: .15rem;
}

.accept:hover {
  background-color: green;
  color: #d9d9d9
}

.decline {
    background-color: red;
  }
</style>
