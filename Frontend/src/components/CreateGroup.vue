<script setup lang="ts">
import { ref, watch } from "vue";
import type { Ref } from "vue";
import { useAuthStore } from "@/stores/auth";
import { timestamp } from "@vueuse/shared";
import Ws from "@/Websocket";
import { useGroupsStore } from "@/stores/groups";

const groupname = ref("");
const groupdesc = ref("");
const toggleShowGroup = ref(false)
const store = useAuthStore()
const groupError = ref("")

const showGroup = () => {
  toggleShowGroup.value = !toggleShowGroup.value
}

interface Group {
  uid: number;
  gname: string;
  gtext: string;
}

const createGroup = () => {
  if (groupname.value.length < 3) {
    return groupError.value = "Group name must be atleast 3 characters long"
  } else if (groupdesc.value.length < 5) {
    return groupError.value = "Group description must be atleast 5 characters long"
  } else if(useGroupsStore().groups.some((group) => group.Name === groupname.value)){
    return groupError.value = "Group name already exists"
  }
  const newGroup = <Group>{
    uid: store.userId,
    gname: groupname.value,
    gtext: groupdesc.value,
  };
  Ws.send({
      Page: "newgroup",
      Data: newGroup,
    }
  );
  groupError.value = ""
  groupname.value = ""
  groupdesc.value = ""
  toggleShowGroup.value = false;
};
</script>

<template>
  <button class="createButton" @click="showGroup()">Create new group</button>
  <div class="overlay" v-show="toggleShowGroup">
    <div class="modal">
      <button class="close-button" @click="showGroup()">X</button>
      
      <div class="flexcreategroup">
        <div class="create-group-container">
          <h3 class="label">Group name:</h3>
          <textarea
            name="groupname"
            id="groupname"
            cols="30"
            rows="1"
            class="groupname"
            v-model="groupname"
          ></textarea>
          <h3 class="label">Group description:</h3>
          <textarea
            name="groupdesc"
            id="groupdesc"
            cols="30"
            rows="10"
            class="groupdesc"
            v-model="groupdesc"
          ></textarea>
          <div v-if="groupError.length > 0">{{groupError}}</div>
        </div>
        <button class="modal-button" @click="createGroup()">Create new group</button>
      </div>
    </div>
  </div>
</template>

<style scoped>

.createButton {
  max-width: 50ch;
  width: 80%;
  background-color: #f0f0f0;
  border-radius: 4px;
  padding: 10px;
  margin-bottom: 30px;
  margin-top: 15px;
  cursor: pointer;
  transition: background-color 0.2s;
  box-sizing: border-box;
  text-align: center;
  font-size: 18px;
  font-weight: bold;
  display: block;
  margin-left: auto;
  margin-right: auto;
  border: none;
  color: inherit;
}

.createButton:hover {
  background-color: #e0e0e0;
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
.label {
  font-weight: bold;
  font-size: 16px;
  margin-bottom: 3px;
}

.create-group-container {
  display: flex;
  flex-direction: column;
  width: 80%;
}

.close-button {
  position: absolute;
  top: 10px;
  right: 10px;
  background-color: transparent;
  border: none;
  font-size: 18px;
  font-weight: bold;
  cursor: pointer;
}

.close-button:hover {
  color: red;
}

.modal {
  position: relative;
  background-color: white;
  border-radius: 10px;
  padding: 20px;
  max-width: 500px;
  width: 90%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.groupname,
.groupdesc {
  width: 100%;
  resize: none;
  padding: 5px;
  margin: 5px 0;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 14px;
}

.modal-button {
  background-color: #4caf50;
  border-radius: 3px;
  cursor: pointer;
  color: white;
  border: none;
}
</style>