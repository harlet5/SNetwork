<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
const store = useAuthStore()
const message = defineProps(["message"]);
const computedClasses = () => {
  if (message.message.ChOId == store.userId) {
    return 'message received-message'
  } else {
    return 'message sended-message'
  }
}

/*
{ "message": { "ChId": 66, "ChUId": 8, "ChOId": 4, "ChUName": "123", "ChOName": "kevinn", "ChBody": "t", "ChTime": "2.20.2023", "ChStatus": "false" } }

*/
</script>

<template>
  <div :id="`message-${message.message.ChId}`" :class="computedClasses()">
    <p class="message-text">{{message.message.ChBody}}</p>
    <div class="message-info" :id="`message-${message.message.ChId}-info`">
      <p class="message-date">{{message.message.ChTime}}</p>
      <p class="message-status" :id="`message-${message.message.ChId}-status`">✓✓</p>
    </div>
  </div>
</template>

<style scoped>
.message {
  width: max-content;
  min-width: 180px;
  max-width: 500px;
  margin: 10px 10px 10px 10px;
  padding: 10px;
  border-radius: 15px;
  background-color: #3b3d42;
  color: #c9c9d8;
}

.message-text {
  margin: 0;
  margin-bottom: 5px;
}

.message-info {
  display: grid;
  grid-template-columns: auto 20px;
}

.message-info p {
  margin: 0;
  font-size: 0.8rem;
  margin-left: auto;
  margin-right: 0;
}

.message-status {
  letter-spacing: -5px;
}

.sended-message {
  margin-left: auto;
  margin-right: 0;
}

.received-message {
  background-color: #4c4f56;
}

</style>
