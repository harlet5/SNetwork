<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, onBeforeMount, computed } from "vue";
import Ws from "@/Websocket";
import { useAuthStore } from "@/stores/auth";

const notif = defineProps(["notif"]);
const router = useRouter();
const store = useAuthStore();

const flwObj = {
  sid: notif.notif.Sender,
  rid: notif.notif.Reciver,
}
const grpInvObj = {
  rid: notif.notif.Reciver,
  sid: notif.notif.Sender,
  gid: notif.notif.Gid,
}

const grpReqObj = {
    sid: notif.notif.Sender,
    gid: notif.notif.Gid,
  }
const notifSend = (url: string, data: object) => {
  Ws.send({
      Page: url,
      Data: data,
    }
  );
}

const redirToGroup = (id: number) => {
  router.replace(`/groups/${id}`);
}

const redirToUsr = (id: number) => {
  router.replace(`/user/${id}`)
}
//const navigateToPost = (id: number) => {
//};
//const navigateToProfile = (id: number) => {
  //router.replace(`/user/${id}`);
//};
</script>
<template>
  <div class="notification">
    <div v-if="notif.notif.Ntype === 'invite'">
      <div v-if="notif.notif.Status === 'accepted'">
        <div v-if="store.userId === notif.notif.Sender" @click="redirToGroup(notif.notif.Gid)">{{ notif.notif.Rname }} has accepted your invite to {{ notif.notif.Gname }}</div>
        <div v-else @click="redirToGroup(notif.notif.Gid)">You accepted {{ notif.notif.Sname }}'s invite to join {{ notif.notif.Gname }}</div>
      </div>
      <div v-if="notif.notif.Status === 'declined'">
        <div v-if="store.userId === notif.notif.Sender" @click="redirToGroup(notif.notif.Gid)">{{ notif.notif.Rname }} has declined your invite to join {{ notif.notif.Gname }}</div>
        <div v-else @click="redirToGroup(notif.notif.Gid)">You declined {{ notif.notif.Sname }}'s invite to join {{ notif.notif.Gname }}</div>
      </div>
      <div v-if="notif.notif.Status === 'unseen'">
        <div v-if="store.userId === notif.notif.Sender" @click="redirToUsr(notif.notif.Reciver)">You sent a group invite to {{ notif.notif.Rname }}</div>
        <div v-else><p @click="redirToGroup(notif.notif.Gid)">{{ notif.notif.Sname }} has sent you an invite to {{ notif.notif.Gname }}</p>
        <div class='buttondiv'><button class='accept' @click="notifSend('grpinvacc', grpInvObj)">Accept</button>
        <button class='decline' @click="notifSend('grpinvdec', grpInvObj)">Decline</button></div></div>
      </div>
    </div>
    <div v-if="notif.notif.Ntype === 'joinreq'">
      <div v-if="notif.notif.Status === 'accepted'">
        <div v-if="store.userId === notif.notif.Sender" @click="redirToGroup(notif.notif.Gid)">{{ notif.notif.Rname }} has accepted your request to join {{ notif.notif.Gname }}</div>
        <div v-else @click="redirToGroup(notif.notif.Gid)">You accepted {{ notif.notif.Sname }}'s request to join {{ notif.notif.Gname }}</div>
      </div>
      <div v-if="notif.notif.Status === 'declined'">
        <div v-if="store.userId === notif.notif.Sender" @click="redirToGroup(notif.notif.Gid)">{{ notif.notif.Rname }} has declined your request to join {{ notif.notif.Gname }}</div>
        <div v-else @click="redirToGroup(notif.notif.Gid)">You declined {{ notif.notif.Sname }}'s request to join {{ notif.notif.Gname }}</div>
      </div>
      <div v-if="notif.notif.Status === 'unseen'">
        <div v-if="store.userId === notif.notif.Sender" @click="redirToGroup(notif.notif.Gid)">You sent a request to join {{ notif.notif.Gname }} to {{ notif.notif.Rname }}</div>
        <div v-else><p @click="redirToGroup(notif.notif.Gid)">{{ notif.notif.Sname }} wants to join {{ notif.notif.Gname }}</p>
        <div class='buttondiv'><button class='accept' @click="notifSend('grpreqacc', grpReqObj)">Accept</button>
        <button class='decline' @click="notifSend('grpreqdec', grpReqObj)">Decline</button></div></div>
      </div>
    </div>
    <div v-if="notif.notif.Ntype === 'follow'">
      <div v-if="notif.notif.Status === 'accepted'">
        <div v-if="store.userId === notif.notif.Sender" @click="redirToUsr(notif.notif.Reciver)">{{ notif.notif.Rname }} has accepted your follow request</div>
        <div v-else @click="redirToUsr(notif.notif.Sender)">You accepted {{ notif.notif.Sname }}'s follow request</div>
      </div>
      <div v-if="notif.notif.Status === 'declined'">
        <div v-if="store.userId === notif.notif.Sender" @click="redirToUsr(notif.notif.Reciver)">{{ notif.notif.Rname }} has declined your follow request</div>
        <div v-else @click="redirToUsr(notif.notif.Sender)">You declined {{ notif.notif.Sname }}'s follow request</div>
      </div>
      <div v-if="notif.notif.Status === 'unseen'">
        <div v-if="store.userId === notif.notif.Sender" @click="redirToUsr(notif.notif.Reciver)">You sent a follow request to {{ notif.notif.Rname }}</div>
        <div v-else><p @click="redirToUsr(notif.notif.Sender)">{{ notif.notif.Sname }} has sent you a follow request</p>
        <div class='buttondiv'><button class='accept' @click="notifSend('flwacc', flwObj)">Accept</button>
        <button class='decline' @click="notifSend('flwdec', flwObj)">Decline</button></div></div>
      </div>
    </div>
    <div v-if="notif.notif.Ntype === 'event'"> 
      <div v-if="notif.notif.Status === 'unseen'">
        <div v-if="store.userId === notif.notif.Sender" @click="redirToGroup(notif.notif.Gid)">You created a new event at {{ notif.notif.Gname }}</div>
        <div v-else @click="redirToGroup(notif.notif.Gid)">{{ notif.notif.Sname }} created a new event at {{ notif.notif.Gname }}</div>
      </div>
    </div>
  </div>

 </template>
<style scoped>

.notification {
  padding: 10px;
  cursor: pointer;
}

.notification:hover {
  background-color: lightgray;
}
.accept {
    background-color: green;
  }
.decline {
    background-color: red;
  }
</style>
