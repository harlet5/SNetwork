<script setup lang="ts">
import EmojiPicker from "vue3-emoji-picker";
import "vue3-emoji-picker/css";
import { ref } from "vue";

const showEmoijBox = ref(false);
const emit = defineEmits(["emoij"]);
const emitEmoij = (emoij: any) => {
  emit("emoij", emoij);
  toggleEmoijBox()
};

const toggleEmoijBox = () => {
  showEmoijBox.value = !showEmoijBox.value
}
</script>
<template>
  <div class="overlay">
  <svg class="my-svg-component" @click="toggleEmoijBox">
    <image class="image" width="25" height="25" href="../assets/smile-svgrepo-com.svg" />
  </svg>
  <EmojiPicker :native="true" @select="emitEmoij" v-if="showEmoijBox" class="emoji-picker"/>
</div>
</template>

<style scoped>
.emoji-picker {
  position: absolute;
  bottom: 0;
  right: 0;
  margin-bottom: 60px;
  margin-right: 10px;
}

.my-svg-component {
  height: 30px;
  width: 30px;
  position: absolute;
  bottom: 0;
  right: 0;
  margin-bottom: 12px;
  margin-right: 12px;
}

.image:hover {
  background-color: blue;
  cursor: pointer;
}

.overlay {
  margin-right: 5px;
}
</style>
